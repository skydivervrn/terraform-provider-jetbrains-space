package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	jsc "github.com/skydivervrn/jetbrains-space-api-client-go"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &spaceProjectDataSource{}
	_ datasource.DataSourceWithConfigure = &spaceProjectDataSource{}
)

// spaceProjectDataSource is the data source implementation.
type spaceProjectDataSource struct {
	client *jsc.Client
}

type spaceProjectModel struct {
	ID                       types.String `tfsdk:"id"`
	Key                      types.String `tfsdk:"key"`
	Name                     types.String `tfsdk:"name"`
	Private                  types.Bool   `tfsdk:"private"`
	Description              types.String `tfsdk:"description"`
	Icon                     types.String `tfsdk:"icon"`
	LatestRepositoryActivity types.String `tfsdk:"latest_repository_activity"`
	CreatedAtIso             types.String `tfsdk:"created_at_iso"`
	CreatedAtTimestamp       types.Int64  `tfsdk:"created_at_timestamp"`
	Archived                 types.Bool   `tfsdk:"archived"`
}

// NewSpaceProjectDataSource is a helper function to simplify the provider implementation.
func NewSpaceProjectDataSource() datasource.DataSource {
	return &spaceProjectDataSource{}
}

// Metadata returns the data source type name.
func (d *spaceProjectDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

// Schema defines the schema for the data source.
func (d *spaceProjectDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:    true,
				Description: "The uniq ID of existed Space project to get data from",
			},
			"key": schema.StringAttribute{
				Computed:    true,
				Description: "Magic field. Project name written in uppercase",
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: "Name of a project",
			},
			"private": schema.BoolAttribute{
				Computed:    true,
				Description: "Defines if project private or public",
			},
			"description": schema.StringAttribute{
				Computed:    true,
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

// Read refreshes the Terraform state with the latest data.
func (d *spaceProjectDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state spaceProjectModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	project, err := d.client.GetProject(state.ID.ValueString())
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

func (d *spaceProjectDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

	d.client = client
}

//func dataSourceSpaceProject() *schema.Resource {
//	return &schema.Resource{
//		// This description is used by the documentation generator and the language server.
//		Description: "Space project data source in the Terraform provider Space.",
//
//		ReadContext: dataSourceSpaceProjectRead,
//
//		Schema: map[string]*schema.Schema{
//			"id": &schema.Schema{
//				Type:        schema.TypeString,
//				Required:    true,
//				Description: "Unique id of Space project.",
//			},
//			"key": &schema.Schema{
//				Type:     schema.TypeString,
//				Computed: true,
//			},
//			"name": &schema.Schema{
//				Type:     schema.TypeString,
//				Computed: true,
//			},
//			"private": &schema.Schema{
//				Type:     schema.TypeBool,
//				Computed: true,
//			},
//			"description": &schema.Schema{
//				Type:     schema.TypeString,
//				Computed: true,
//			},
//			"icon": &schema.Schema{
//				Type:     schema.TypeInt,
//				Computed: true,
//			},
//			"latest_repository_activity": &schema.Schema{
//				Type:     schema.TypeString,
//				Computed: true,
//			},
//			"created_at_iso": &schema.Schema{
//				Type:     schema.TypeString,
//				Computed: true,
//			},
//			"created_at_timestamp": &schema.Schema{
//				Type:     schema.TypeInt,
//				Computed: true,
//			},
//			"archived": &schema.Schema{
//				Type:     schema.TypeBool,
//				Computed: true,
//			},
//		},
//	}
//}
//
//func dataSourceSpaceProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
//
//	// Warning or errors can be collected in a slice type
//	var diags diag.Diagnostics
//	id := d.Get("id").(string)
//
//	client := m.(*sc.Client)
//	projectData, err := client.GetProject(id)
//	if err != nil {
//		return diag.FromErr(err)
//	}
//	resourceSpaceProjectSetAllValues(ctx, d, projectData)
//
//	// always run
//	d.SetId(id)
//
//	return diags
//}
