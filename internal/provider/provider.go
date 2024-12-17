package provider

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/ghostsecurity/terraform-provider-ghost/internal/api"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// GhostProvider implementa a terraform provider for interacting with the
// Ghost Security API.
type GhostProvider struct {
	version string
}

type GhostProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
}

func (p *GhostProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "ghost"
	resp.Version = p.version
}

func (p *GhostProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "Ghost API endpoint. For example https://api.ghostsecurity.com",
				Optional:            true,
			},
		},
	}
}

func (p *GhostProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data GhostProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiKey := os.Getenv("GHOST_API_KEY")
	if apiKey == "" {
		resp.Diagnostics.AddError("Missing API key", "GHOST_API_KEY must be set")
		return
	}

	endpoint := data.Endpoint.ValueString()
	if endpoint == "" {
		endpoint = "https://api.ghostsecurity.com"
	}

	apiURL, err := url.Parse(endpoint)
	if err != nil {
		resp.Diagnostics.AddError("Invalid Endpoint",
			fmt.Sprintf("The provided endpoint %q is not a valid URL", data.Endpoint.ValueString()))
		return
	}

	client, err := api.NewClient(apiURL.String(), apiKey)
	if err != nil {
		resp.Diagnostics.AddError("Error initializing API client", err.Error())
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *GhostProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewLogForwarderResource,
	}
}

func (p *GhostProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *GhostProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &GhostProvider{
			version: version,
		}
	}
}
