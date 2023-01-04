package provider

import (
	"context"
	sc "github.com/skydivervrn/space-api-client"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSpaceProject() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "This resource will create project for JetBrains Space https://jetbrains.space",

		CreateContext: resourceSpaceProjectCreate,
		ReadContext:   resourceSpaceProjectRead,
		UpdateContext: resourceSpaceProjectUpdate,
		DeleteContext: resourceSpaceProjectDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique key which couldn't be changed in future.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Space project name.",
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

func resourceSpaceProjectSetAllValues(ctx context.Context, d *schema.ResourceData, project sc.Project) diag.Diagnostics {
	var diags diag.Diagnostics
	if err := d.Set("key", project.Key.Key); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", project.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("private", project.Private); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", project.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("icon", project.Icon); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("latest_repository_activity", project.LatestRepositoryActivity); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_at_iso", project.CreatedAt.Iso); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_at_timestamp", project.CreatedAt.Timestamp); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("archived", project.Archived); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceSpaceProjectCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*sc.Client)
	var diags diag.Diagnostics

	project, err := client.CreateProject(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	resourceSpaceProjectSetAllValues(ctx, d, project)
	d.SetId(project.ID)

	// write logs using the tflog package
	// see https://pkg.go.dev/github.com/hashicorp/terraform-plugin-log/tflog
	// for more information
	tflog.Trace(ctx, "created a resource")
	resourceSpaceProjectRead(ctx, d, meta)
	return diags
}

func resourceSpaceProjectRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	id := d.Id()
	var diags diag.Diagnostics

	client := meta.(*sc.Client)
	projectData, err := client.GetProject(id)
	if err != nil {
		return diag.FromErr(err)
	}
	resourceSpaceProjectSetAllValues(ctx, d, projectData)

	// always run
	d.SetId(id)

	return diags
}

func resourceSpaceProjectUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	id := d.Id()
	var diags diag.Diagnostics
	name := d.Get("name").(string)
	client := meta.(*sc.Client)
	projectData, err := client.UpdateProject(id, name)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", projectData.Name); err != nil {
		return diag.FromErr(err)
	}
	resourceSpaceProjectRead(ctx, d, meta)
	return diags
}

func resourceSpaceProjectDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	id := d.Id()
	var diags diag.Diagnostics
	client := meta.(*sc.Client)
	err := client.DeleteProject(id)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
