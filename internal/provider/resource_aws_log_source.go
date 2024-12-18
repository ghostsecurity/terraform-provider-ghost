package provider

import (
	"context"
	"errors"
	"fmt"

	"github.com/ghostsecurity/terraform-provider-ghost/internal/api"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AWSLogSourceResource struct {
	client *api.GhostClient
}

type AWSLogSourceResourceModel struct {
	LogForwarderID types.String `tfsdk:"log_forwarder_id"`
	S3BucketName   types.String `tfsdk:"s3_bucket_name"`
	RoleARN        types.String `tfsdk:"role_arn"`
	SQSARN         types.String `tfsdk:"sqs_arn"`
	Region         types.String `tfsdk:"region"`
	AccountID      types.String `tfsdk:"account_id"`
}

func NewAWSLogSourceResource() resource.Resource {
	return &AWSLogSourceResource{}
}

func (r *AWSLogSourceResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "ghost_aws_log_source"
}

func (r *AWSLogSourceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "A log source configures a log forwarder to ingest files from an AWS S3 bucket. Only one source can be created per log forwarder.",
		Attributes: map[string]schema.Attribute{
			"s3_bucket_name": schema.StringAttribute{
				MarkdownDescription: "S3 bucket name ",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"log_forwarder_id": schema.StringAttribute{
				MarkdownDescription: "Log forwarder UUID.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"account_id": schema.StringAttribute{
				MarkdownDescription: "AWS account ID the s3 bucket belongs to.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"region": schema.StringAttribute{
				MarkdownDescription: "AWS region the SQS queue is deployed in.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"sqs_arn": schema.StringAttribute{
				MarkdownDescription: "ARN for the SQS queue to receive S3 object notifications from.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"role_arn": schema.StringAttribute{
				MarkdownDescription: "ARN for the AWS IAM role that allows reading from the bucket and SQS queue. The role must be configured to allow a google federated identity to assume the role based on the log forwarder subject_id.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

// Configure fetches the Ghost API client from the provider to be used to make authenticated
// requests to the API in the CRUD handlers.
func (r *AWSLogSourceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*api.GhostClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected ProviderData type",
			fmt.Sprintf("Expected *api.Client, got: %T.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *AWSLogSourceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AWSLogSourceResourceModel

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	id, err := uuid.Parse(data.LogForwarderID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Log forwarder ID is not a UUID",
			err.Error(),
		)
		return
	}

	source := &api.ForwarderSource{
		BucketName: data.S3BucketName.ValueString(),
		RoleARN:    data.RoleARN.ValueString(),
		SQSARN:     data.SQSARN.ValueString(),
		Region:     data.Region.ValueString(),
		AccountID:  data.AccountID.ValueString(),
	}

	// Make the API request to create the log forwarder
	_, err = r.client.UpdateLogForwarderSource(id, source)
	if err != nil {
		details := err.Error()
		if errors.Is(err, api.ErrNotFound) {
			details = "Log forwarder not found"
		}

		resp.Diagnostics.AddError(
			"Error creating aws log source",
			details,
		)
		return
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *AWSLogSourceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AWSLogSourceResourceModel

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	id, err := uuid.Parse(data.LogForwarderID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"ID is not a UUID",
			err.Error(),
		)
		return
	}

	// Make the API request to fetch the log forwarder
	getResp, err := r.client.GetLogForwarder(id)
	if err != nil {
		// If the log forwarder no longer exists, remove the source from state.
		if errors.Is(err, api.ErrNotFound) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Error fetching aws log source",
			err.Error(),
		)
		return
	}

	// If the source does not exist for the log forwarder remove it from state.
	if getResp.Source == nil {
		resp.State.RemoveResource(ctx)
		return
	}

	// Set the forwarder details in the state.
	data.S3BucketName = types.StringValue(getResp.Source.BucketName)
	data.RoleARN = types.StringValue(getResp.Source.RoleARN)
	data.SQSARN = types.StringValue(getResp.Source.SQSARN)
	data.Region = types.StringValue(getResp.Source.Region)
	data.AccountID = types.StringValue(getResp.Source.AccountID)

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update is not supported and all attribute changes should require a replace
func (r *AWSLogSourceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *AWSLogSourceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AWSLogSourceResourceModel

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	id, err := uuid.Parse(data.LogForwarderID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"ID is not a UUID",
			err.Error(),
		)
		return
	}

	// Make the API request to remove the source from the log forwarder
	if _, err := r.client.UpdateLogForwarderSource(id, nil); err != nil {
		// If the log forwarder has already been deleted do not error.
		if !errors.Is(err, api.ErrNotFound) {
			resp.Diagnostics.AddError(
				"Error deleting aws log source",
				err.Error(),
			)
			return
		}
	}

	resp.State.RemoveResource(ctx)
}
