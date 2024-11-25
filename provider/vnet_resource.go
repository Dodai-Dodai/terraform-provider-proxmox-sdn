package provider

import (
	"context"
	"fmt"

	"github.com/Dodai-Dodai/terraform-provider-proxmox-sdn/client"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var (
	_ resource.Resource              = &proxmoxSDNVnetResource{}
	_ resource.ResourceWithConfigure = &proxmoxSDNVnetResource{}
)

func NewProxmoxSDNVnetResource() resource.Resource {
	return &proxmoxSDNVnetResource{}
}

type proxmoxSDNVnetResource struct {
	client *client.SSHProxmoxClient
}

func (r *proxmoxSDNVnetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.SSHProxmoxClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected provider data type",
			fmt.Sprintf("Expected *client.SSHProxmoxClient, got %T", req.ProviderData),
		)
		return
	}
	r.client = client
}

func (r *proxmoxSDNVnetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vnet"
}

func (r *proxmoxSDNVnetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"vnet": schema.StringAttribute{
				Description: "vnet name",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"zone": schema.StringAttribute{
				Description: "zone name",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"type": schema.StringAttribute{
				Description: "vnet type",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"tag": schema.Int64Attribute{
				Description: "vnet tag",
				Optional:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"vlanaware": schema.BoolAttribute{
				Description: "vnet vlanaware",
				Optional:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *proxmoxSDNVnetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan vnetsModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	zone, diags := convertVnetsModeltoClientVnet(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	err := r.client.CreateVnet(*zone)
	if err != nil {
		resp.Diagnostics.AddError(
			"CreateVnet failed",
			err.Error(),
		)
		return
	}

	state, diags := convertSDNVnettoVnetsModel(ctx, *zone)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func convertVnetsModeltoClientVnet(ctx context.Context, model vnetsModel) (*client.SDNVnet, diag.Diagnostics) {
	var diags diag.Diagnostics

	var vnet client.SDNVnet

	vnet.Vnet = model.Vnet.ValueString()
	vnet.Zone = model.Zone.ValueString()
	vnet.Type = model.Type.ValueString()

	if !model.Tag.IsNull() && !model.Tag.IsUnknown() {
		tag := model.Tag.ValueInt64()
		vnet.Tag = &tag
	}

	vnet.Vlanaware = BoolToIntBoolPointer(model.Vlanaware)

	return &vnet, diags
}

func (r *proxmoxSDNVnetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state vnetsModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	vnet, err := r.client.GetVnet(state.Vnet.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to get vnet", err.Error())
		return
	}

	updatedState, diagsState := convertSDNVnettoVnetsModel(ctx, *vnet)
	resp.Diagnostics.Append(diagsState...)
	if diagsState.HasError() {
		return
	}

	diags = resp.State.Set(ctx, updatedState)
	resp.Diagnostics.Append(diags...)
}

func (r *proxmoxSDNVnetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan vnetsModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	zone, diags := convertVnetsModeltoClientVnet(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	err := r.client.UpdateVnet(*zone)
	if err != nil {
		resp.Diagnostics.AddError(
			"UpdateVnet failed",
			err.Error(),
		)
		return
	}

	state, diags := convertSDNVnettoVnetsModel(ctx, *zone)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *proxmoxSDNVnetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state vnetsModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	err := r.client.DeleteVnet(state.Vnet.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("DeleteVnet failed", err.Error())
		return
	}
}
