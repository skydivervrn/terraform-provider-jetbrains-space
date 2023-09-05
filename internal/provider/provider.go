// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	jsc "github.com/skydivervrn/jetbrains-space-api-client-go"
	"os"
)

// Ensure SpaceProvider satisfies various provider interfaces.
var _ provider.Provider = &SpaceProvider{}

// SpaceProvider defines the provider implementation.
type SpaceProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// SpaceProviderModel describes the provider data model.
type SpaceProviderModel struct {
	URL   types.String `tfsdk:"url"`
	TOKEN types.String `tfsdk:"token"`
}

func (p *SpaceProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "jetbrains-space"
	resp.Version = p.version
}

func (p *SpaceProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"url": schema.StringAttribute{
				Description: "Jetbrains Space URL. If not set, provider will try to read environment variable TERRAFORM_PROVIDER_JETBRAINS_SPACE_URL",
				Optional:    true,
			},
			"token": schema.StringAttribute{
				Description: "Jetbrains Space token. If not set, provider will try to read environment variable TERRAFORM_PROVIDER_JETBRAINS_SPACE_TOKEN",
				Optional:    true,
				Sensitive:   true,
			},
		},
	}
}

func (p *SpaceProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data SpaceProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	url := os.Getenv("TERRAFORM_PROVIDER_JETBRAINS_SPACE_URL")
	token := os.Getenv("TERRAFORM_PROVIDER_JETBRAINS_SPACE_TOKEN")

	// Example client configuration for data sources and resources
	client, err := jsc.NewClient(url, token)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Jetbrains Space API Client",
			"An unexpected error occurred when creating the Jetbrains API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Jetbrains Space Client Error: "+err.Error(),
		)
		return
	}
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *SpaceProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *SpaceProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewSpaceProjectDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &SpaceProvider{
			version: version,
		}
	}
}
