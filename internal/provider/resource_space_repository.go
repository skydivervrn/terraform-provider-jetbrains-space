package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	//sc "github.com/skydivervrn/space-api-client"
	sc "github.com/skydivervrn/terraform-provider-space/space-api-client"
)

func resourceSpaceRepository() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description:   "This resource will create repository in project for JetBrains Space",
		CreateContext: resourceSpaceRepositoryCreate,
		ReadContext:   resourceSpaceRepositoryRead,
		//UpdateContext: resourceSpaceRepositoryUpdate,
		DeleteContext: resourceSpaceRepositoryDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Id of parent Space project",
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Repository name",
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"latest_repository_activity": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_branch_name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func resourceSpaceRepositorySetAllValues(ctx context.Context, d *schema.ResourceData, repository sc.Repository) diag.Diagnostics {
	var diags diag.Diagnostics
	if err := d.Set("name", repository.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", repository.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("latest_repository_activity", repository.LatestActivity); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("default_branch_name", repository.DefaultBranch); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceSpaceRepositoryCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*sc.Client)
	data := sc.CreateRepositoryData{}
	var diags diag.Diagnostics
	projectId := d.Get("project_id").(string)
	repositoryName := d.Get("name").(string)
	data.Description = d.Get("description").(string)
	data.DefaultBranch = d.Get("default_branch_name").(string)

	repository, err := client.CreateRepository(repositoryName, projectId, data)
	if err != nil {
		return diag.FromErr(err)
	}
	resourceSpaceRepositorySetAllValues(ctx, d, repository)
	d.SetId(repository.ID)

	// write logs using the tflog package
	// see https://pkg.go.dev/github.com/hashicorp/terraform-plugin-log/tflog
	// for more information
	tflog.Trace(ctx, "created a resource")
	resourceSpaceRepositoryRead(ctx, d, meta)
	return diags
}

func resourceSpaceRepositoryRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	projectId := d.Get("project_id").(string)
	var diags diag.Diagnostics
	repositoryName := d.Get("name").(string)
	client := meta.(*sc.Client)
	repositoryData, err := client.GetRepository(repositoryName, projectId)
	if err != nil {
		return diag.FromErr(err)
	}
	resourceSpaceRepositorySetAllValues(ctx, d, repositoryData)

	id := d.Id()
	// always run
	d.SetId(id)

	return diags
}

//func resourceSpaceRepositoryUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
//	id := d.Id()
//	var diags diag.Diagnostics
//	name := d.Get("name").(string)
//	client := meta.(*sc.Client)
//	repositoryData, err := client.UpdateRepository(id, name)
//	if err != nil {
//		return diag.FromErr(err)
//	}
//	if err := d.Set("name", repositoryData.Name); err != nil {
//		return diag.FromErr(err)
//	}
//	resourceSpaceRepositoryRead(ctx, d, meta)
//	return diags
//}

func resourceSpaceRepositoryDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	projectId := d.Get("project_id").(string)
	repoName := d.Get("name").(string)
	var diags diag.Diagnostics
	client := meta.(*sc.Client)
	err := client.DeleteRepository(projectId, repoName)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
