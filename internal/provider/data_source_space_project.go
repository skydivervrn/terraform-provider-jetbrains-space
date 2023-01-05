package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sc "github.com/skydivervrn/terraform-provider-space/space-api-client"
)

func dataSourceSpaceProject() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Space project data source in the Terraform provider Space.",

		ReadContext: dataSourceSpaceProjectRead,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique id of Space project.",
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"private": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"icon": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"latest_repository_activity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at_iso": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at_timestamp": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"archived": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceSpaceProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	id := d.Get("id").(string)

	client := m.(*sc.Client)
	projectData, err := client.GetProject(id)
	if err != nil {
		return diag.FromErr(err)
	}
	resourceSpaceProjectSetAllValues(ctx, d, projectData)

	// always run
	d.SetId(id)

	return diags
}
