package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/types"
	jsc "github.com/skydivervrn/jetbrains-space-api-client-go"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &spaceProjectRepositoryResource{}
	_ resource.ResourceWithConfigure = &spaceProjectRepositoryResource{}
)

// spaceProjectRepositoryResource is the resource implementation.
type spaceProjectRepositoryResource struct {
	client *jsc.Client
}
type spaceProjectRepositoryModel struct {
	ID                       types.String `tfsdk:"id"`
	ProjectId                types.String `tfsdk:"project_id"`
	Name                     types.String `tfsdk:"name"`
	Description              types.String `tfsdk:"description"`
	LatestRepositoryActivity types.String `tfsdk:"latest_repository_activity"`
	DefaultBranchName        types.String `tfsdk:"default_branch_name"`
}

// NewSpaceProjectRepositoryResource is a helper function to simplify the provider implementation.
func NewSpaceProjectRepositoryResource() resource.Resource {
	return &spaceProjectRepositoryResource{}
}

// Configure adds the provider configured client to the resource.
func (r *spaceProjectRepositoryResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*jsc.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *jsc.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

// Metadata returns the resource type name.
func (r *spaceProjectRepositoryResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project_repository"
}

// Schema defines the schema for the resource.
func (r *spaceProjectRepositoryResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"project_id": schema.StringAttribute{
				Required:    true,
				Description: "Id of parent Space project",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "id of git repo",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of a repository",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "Project description",
			},
			"default_branch_name": schema.StringAttribute{
				Optional:    true,
				Description: "Default branch name",
			},
			"latest_repository_activity": schema.StringAttribute{
				Computed:    true,
				Description: "Latest activity in repository",
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *spaceProjectRepositoryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state spaceProjectRepositoryModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	createRepositoryData := jsc.CreateRepositoryData{
		Name:          state.Name.ValueString(),
		Description:   state.Description.ValueString(),
		DefaultBranch: state.DefaultBranchName.ValueString(),
		ProjectId:     state.ProjectId.ValueString(),
	}
	repository, err := r.client.CreateRepository(createRepositoryData)
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to Create Space repository name: %s", state.Name.ValueString()),
			err.Error(),
		)
		return
	}
	state.ID = types.StringValue(repository.ID)
	state.Name = types.StringValue(repository.Name)
	state.Description = types.StringValue(repository.Description)
	if state.DefaultBranchName.ValueString() == "" {
		state.DefaultBranchName = types.StringValue("master")
	}

	//Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *spaceProjectRepositoryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state spaceProjectRepositoryModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	repository, err := r.client.GetRepository(state.Name.ValueString(), state.ProjectId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to Read Space repository name: %s", state.Name.ValueString()),
			err.Error(),
		)
		return
	}
	state.Name = types.StringValue(repository.Name)
	state.Description = types.StringValue(repository.Description)
	state.LatestRepositoryActivity = types.StringValue(repository.LatestActivity)
	state.DefaultBranchName = types.StringValue(strings.Split(repository.DefaultBranch.Head, "/")[2])

	//Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *spaceProjectRepositoryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

}

// Delete deletes the resource and removes the Terraform state on success.
func (r *spaceProjectRepositoryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state spaceProjectRepositoryModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := r.client.DeleteRepository(state.ProjectId.ValueString(), state.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to Delete Space repository name: %s", state.Name.ValueString()),
			err.Error(),
		)
		return
	}
}

//
//
//func resourceSpaceRepository() *schema.Resource {
//	return &schema.Resource{
//		// This description is used by the documentation generator and the language server.
//		Description:   "This resource will create repository in project for JetBrains Space",
//		CreateContext: resourceSpaceRepositoryCreate,
//		ReadContext:   resourceSpaceRepositoryRead,
//		//UpdateContext: resourceSpaceRepositoryUpdate,
//		DeleteContext: resourceSpaceRepositoryDelete,
//		Importer: &schema.ResourceImporter{
//			StateContext: schema.ImportStatePassthroughContext,
//		},
//		Schema: map[string]*schema.Schema{
//			"project_id": {
//				Type:        schema.TypeString,
//				Required:    true,
//				ForceNew:    true,
//				Description: "Id of parent Space project",
//			},
//			"id": {
//				Type:     schema.TypeString,
//				Computed: true,
//			},
//			"name": {
//				Type:        schema.TypeString,
//				Required:    true,
//				ForceNew:    true,
//				Description: "Repository name",
//			},
//			"description": {
//				Type:     schema.TypeString,
//				Computed: true,
//				Optional: true,
//			},
//			"latest_repository_activity": {
//				Type:     schema.TypeString,
//				Computed: true,
//			},
//			"default_branch_name": {
//				Type:     schema.TypeString,
//				ForceNew: true,
//				Computed: true,
//				Optional: true,
//			},
//		},
//	}
//}
//
//func resourceSpaceRepositorySetAllValues(ctx context.Context, d *schema.ResourceData, repository sc.Repository) diag.Diagnostics {
//	var diags diag.Diagnostics
//	if err := d.Set("name", repository.Name); err != nil {
//		return diag.FromErr(err)
//	}
//	if err := d.Set("description", repository.Description); err != nil {
//		return diag.FromErr(err)
//	}
//	if err := d.Set("latest_repository_activity", repository.LatestActivity); err != nil {
//		return diag.FromErr(err)
//	}
//	if err := d.Set("default_branch_name", repository.DefaultBranch); err != nil {
//		return diag.FromErr(err)
//	}
//	return diags
//}
//
//func resourceSpaceRepositoryCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
//	// use the meta value to retrieve your client from the provider configure method
//	client := meta.(*sc.Client)
//	data := sc.CreateRepositoryData{}
//	var diags diag.Diagnostics
//	projectId := d.Get("project_id").(string)
//	repositoryName := d.Get("name").(string)
//	data.Description = d.Get("description").(string)
//	data.DefaultBranch = d.Get("default_branch_name").(string)
//
//	repository, err := client.CreateRepository(repositoryName, projectId, data)
//	if err != nil {
//		return diag.FromErr(err)
//	}
//	resourceSpaceRepositorySetAllValues(ctx, d, repository)
//	d.SetId(repository.ID)
//
//	// write logs using the tflog package
//	// see https://pkg.go.dev/github.com/hashicorp/terraform-plugin-log/tflog
//	// for more information
//	tflog.Trace(ctx, "created a resource")
//	resourceSpaceRepositoryRead(ctx, d, meta)
//	return diags
//}
//
//func resourceSpaceRepositoryRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
//	projectId := d.Get("project_id").(string)
//	var diags diag.Diagnostics
//	repositoryName := d.Get("name").(string)
//	client := meta.(*sc.Client)
//	repositoryData, err := client.GetRepository(repositoryName, projectId)
//	if err != nil {
//		return diag.FromErr(err)
//	}
//	resourceSpaceRepositorySetAllValues(ctx, d, repositoryData)
//
//	id := d.Id()
//	// always run
//	d.SetId(id)
//
//	return diags
//}
//
////func resourceSpaceRepositoryUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
////	id := d.Id()
////	var diags diag.Diagnostics
////	name := d.Get("name").(string)
////	client := meta.(*sc.Client)
////	repositoryData, err := client.UpdateRepository(id, name)
////	if err != nil {
////		return diag.FromErr(err)
////	}
////	if err := d.Set("name", repositoryData.Name); err != nil {
////		return diag.FromErr(err)
////	}
////	resourceSpaceRepositoryRead(ctx, d, meta)
////	return diags
////}
//
//func resourceSpaceRepositoryDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
//	projectId := d.Get("project_id").(string)
//	repoName := d.Get("name").(string)
//	var diags diag.Diagnostics
//	client := meta.(*sc.Client)
//	err := client.DeleteRepository(projectId, repoName)
//	if err != nil {
//		return diag.FromErr(err)
//	}
//	return diags
//}
