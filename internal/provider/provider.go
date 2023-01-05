package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sc "github.com/skydivervrn/terraform-provider-space/space-api-client"
)

const (
	spaceProjectResourceName           = "space_project"
	spaceRepositoryResourceName        = "space_git_repository"
	spaceProjectDataSourceResourceName = "space_project_data_source"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"url": &schema.Schema{
					Type:        schema.TypeString,
					Optional:    true,
					Description: "URL for space instance. Example: https://test.jetbrains.space",
					DefaultFunc: schema.EnvDefaultFunc("SPACE_URL", nil),
				},
				"token": &schema.Schema{
					Type:        schema.TypeString,
					Optional:    true,
					Sensitive:   true,
					Description: "Token for access Space instance. https://www.jetbrains.com/help/space/personal-tokens.html",
					DefaultFunc: schema.EnvDefaultFunc("SPACE_TOKEN", nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				spaceProjectDataSourceResourceName: dataSourceSpaceProject(),
			},
			ResourcesMap: map[string]*schema.Resource{
				spaceProjectResourceName:    resourceSpaceProject(),
				spaceRepositoryResourceName: resourceSpaceRepository(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
		var diags diag.Diagnostics
		url := d.Get("url").(string)
		token := d.Get("token").(string)
		client, err := sc.NewClient(url, token)
		if err != nil {
			return nil, diag.FromErr(err)
		}
		return client, diags
	}
}
