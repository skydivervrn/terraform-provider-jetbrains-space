package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/types"
	jsc "github.com/skydivervrn/jetbrains-space-api-client-go"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &spaceProjectResource{}
	_ resource.ResourceWithConfigure = &spaceProjectResource{}
)

// spaceProjectResource is the resource implementation.
type spaceProjectResource struct {
	client *jsc.Client
}

// NewSpaceProjectResource is a helper function to simplify the provider implementation.
func NewSpaceProjectResource() resource.Resource {
	return &spaceProjectResource{}
}

// Configure adds the provider configured client to the resource.
func (r *spaceProjectResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *spaceProjectResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

// Schema defines the schema for the resource.
func (r *spaceProjectResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The uniq ID of existed Space project to get data from",
			},
			"key": schema.StringAttribute{
				Computed:    true,
				Description: "Magic field. Project name written in uppercase",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of a project",
			},
			"private": schema.BoolAttribute{
				Optional:    true,
				Description: "Defines if project private or public",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "Project description",
			},
			"icon": schema.StringAttribute{
				Computed:    true,
				Description: "TBD",
			},
			"latest_repository_activity": schema.StringAttribute{
				Computed:    true,
				Description: "Latest activity in project repos",
			},
			"created_at_iso": schema.StringAttribute{
				Computed:    true,
				Description: "TBD",
			},
			"created_at_timestamp": schema.Int64Attribute{
				Computed:    true,
				Description: "TBD",
			},
			"archived": schema.BoolAttribute{
				Computed:    true,
				Description: "TBD",
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *spaceProjectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state spaceProjectModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	createProjectData := jsc.CreateProjectData{
		Name:        state.Name.ValueString(),
		Description: state.Description.ValueString(),
		Private:     state.Private.ValueBool(),
	}
	project, err := r.client.CreateProject(createProjectData)
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to Create Space project name: %s", state.Name.ValueString()),
			err.Error(),
		)
		return
	}
	state.ID = types.StringValue(project.ID)
	state.Key = types.StringValue(project.Key.Key)
	state.Name = types.StringValue(project.Name)
	state.Private = types.BoolValue(project.Private)
	state.Description = types.StringValue(project.Description)
	state.Icon = types.StringValue(project.Icon)
	state.LatestRepositoryActivity = types.StringValue(project.LatestRepositoryActivity)
	state.CreatedAtIso = types.StringValue(project.CreatedAt.Iso)
	state.CreatedAtTimestamp = types.Int64Value(project.CreatedAt.Timestamp)
	state.Archived = types.BoolValue(project.Archived)

	//Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *spaceProjectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state spaceProjectModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	project, err := r.client.GetProject(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to Read Space project id: %s", state.ID.ValueString()),
			err.Error(),
		)
		return
	}
	state.Key = types.StringValue(project.Key.Key)
	state.Name = types.StringValue(project.Name)
	state.Private = types.BoolValue(project.Private)
	state.Description = types.StringValue(project.Description)
	state.Icon = types.StringValue(project.Icon)
	state.LatestRepositoryActivity = types.StringValue(project.LatestRepositoryActivity)
	state.CreatedAtIso = types.StringValue(project.CreatedAt.Iso)
	state.CreatedAtTimestamp = types.Int64Value(project.CreatedAt.Timestamp)
	state.Archived = types.BoolValue(project.Archived)

	//Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *spaceProjectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state spaceProjectModel
	var config spaceProjectModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	updateProjectData := jsc.UpdateProjectData{
		Id:          state.ID.ValueString(),
		Name:        config.Name.ValueString(),
		Description: config.Description.ValueString(),
		Private:     config.Private.ValueBool(),
	}
	project, err := r.client.UpdateProject(updateProjectData)
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to Update Space project name: %s", state.Name.ValueString()),
			err.Error(),
		)
		return
	}
	state.ID = types.StringValue(project.ID)
	state.Key = types.StringValue(project.Key.Key)
	state.Name = types.StringValue(project.Name)
	state.Private = types.BoolValue(project.Private)
	state.Description = types.StringValue(project.Description)
	state.Icon = types.StringValue(project.Icon)
	state.LatestRepositoryActivity = types.StringValue(project.LatestRepositoryActivity)
	state.CreatedAtIso = types.StringValue(project.CreatedAt.Iso)
	state.CreatedAtTimestamp = types.Int64Value(project.CreatedAt.Timestamp)
	state.Archived = types.BoolValue(project.Archived)

	//Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *spaceProjectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state spaceProjectModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := r.client.DeleteProject(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Unable to Delete Space project name: %s", state.Name.ValueString()),
			err.Error(),
		)
		return
	}
}
