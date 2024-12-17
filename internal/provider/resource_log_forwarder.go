package provider

import (
	"context"
	"errors"
	"fmt"

	"github.com/ghostsecurity/terraform-provider-ghost/internal/api"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LogForwarderResource struct {
	client *api.GhostClient
}

type LogForwarderResourceModel struct {
	ID        types.String `tfsdk:"id"`
	SubjectID types.String `tfsdk:"subject_id"`
	Name      types.String `tfsdk:"name"`
}

func NewLogForwarderResource() resource.Resource {
	return &LogForwarderResource{}
}

func (r *LogForwarderResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "ghost_log_forwarder"
}

func (r *LogForwarderResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "A log forwarder receives logs from a Cloud Provider account to be processed in the Ghost platform.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Unique name for the log forwarder.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "Log forwarder UUID.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"subject_id": schema.StringAttribute{
				MarkdownDescription: "OAuth subject_id that must be used when creating a role in a cloud provider account to allow the Ghost platform to ingest logs.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

// Configure fetches the Ghost API client from the provider to be used to make authenticated
// requests to the API in the CRUD handlers.
func (r *LogForwarderResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *LogForwarderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LogForwarderResourceModel

	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	name := data.Name.ValueString()

	// Make the API request to create the log forwarder
	createResp, err := r.client.CreateLogForwarder(ctx, name)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating log forwarder",
			err.Error(),
		)
		return
	}

	// Set the forwarder ID and subject ID in the state
	// when forwarder was created successfully.
	data.ID = types.StringValue(createResp.ID)
	data.SubjectID = types.StringValue(createResp.SubjectID)

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *LogForwarderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LogForwarderResourceModel

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := data.ID.ValueString()

	// Make the API request to fetch the log forwarder
	getResp, err := r.client.GetLogForwarder(ctx, id)
	if err != nil {
		// If the log forwarder no longer exists, remove it from state.
		if errors.Is(err, api.ErrNotFound) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Error fetching log forwarder",
			err.Error(),
		)
		return
	}

	// Set the forwarder details in the state.
	data.ID = types.StringValue(getResp.ID)
	data.SubjectID = types.StringValue(getResp.SubjectID)
	data.Name = types.StringValue(getResp.Name)

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update is not supported and all attribute changes should require a replace
func (r *LogForwarderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *LogForwarderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LogForwarderResourceModel

	diags := req.State.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := data.ID.ValueString()

	// Make the API request to fetch the log forwarder
	if err := r.client.DeleteLogForwarder(ctx, id); err != nil {
		// If the log forwarder has already been deleted do not error.
		if !errors.Is(err, api.ErrNotFound) {
			resp.Diagnostics.AddError(
				"Error deleting log forwarder",
				err.Error(),
			)
			return
		}
	}

	resp.State.RemoveResource(ctx)
}
