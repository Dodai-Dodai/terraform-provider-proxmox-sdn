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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &proxmoxSDNControllerResource{}
	_ resource.ResourceWithConfigure = &proxmoxSDNControllerResource{}
)

func NewProxmoxSDNControllerResource() resource.Resource {
	return &proxmoxSDNControllerResource{}
}

type proxmoxSDNControllerResource struct {
	client *client.SSHProxmoxClient
}

func (r *proxmoxSDNControllerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *proxmoxSDNControllerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_controller"
}

func (r *proxmoxSDNControllerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"controller": schema.StringAttribute{
				Description: "controller name",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"type": schema.StringAttribute{
				Description: "controller type",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"asn": schema.Int64Attribute{
				Description: "controller asn",
				Optional:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				// typeがevpnとbgpのとき必須

				// 値は0 - 4294967296

			},
			"bgp_multipath_aspath_relax": schema.BoolAttribute{
				Description: "controller bgp-multipath-aspath-relax",
				Optional:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				// typeがbgpの場合のみ有効

			},
			"ebgp": schema.BoolAttribute{
				Description: "controller ebgp",
				Optional:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				// typeがbgpの場合のみ有効

			},
			"ebgp_multihop": schema.Int64Attribute{
				Description: "controller ebgp-multihop",
				Optional:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				// typeがbgpの場合のみ有効
				DeprecationMessage: "ebgp_multihop is deprecated for isis controller",

				// 値は1 - 100

			},
			"isis_domain": schema.StringAttribute{
				Description: "controller isis-domain",
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				// typeがisisの場合のとき必須
				DeprecationMessage: "isis_domain is deprecated for evpn and bgp controller",
			},
			"isis_ifaces": schema.StringAttribute{
				Description: "controller isis-ifaces",
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				// typeがisisの場合のとき必須

			},
			"isis_net": schema.StringAttribute{
				Description: "controller isis-net",
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				// typeがisisの場合のとき必須

			},
			"loopback": schema.StringAttribute{
				Description: "controller loopback",
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				// typeがbgpとisisの場合のみ有効

			},
			"node": schema.StringAttribute{
				Description: "controller node",
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				// typeがbgpとisisの場合のとき必須

			},
			"peers": schema.SetAttribute{
				Description: "controller peers",
				Optional:    true,
				ElementType: types.StringType,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.RequiresReplace(),
				},
				// typeがevpnとbgpの場合のとき必須

			},
		},
	}
}

func convertControllerModeltoClientController(ctx context.Context, model controllerModel) (*client.SDNController, diag.Diagnostics) {
	var diags diag.Diagnostics

	var controller client.SDNController

	controller.Controller = model.Controller.ValueString()
	controller.Type = model.Type.ValueString()

	if !model.ASN.IsNull() && !model.ASN.IsUnknown() {
		asn := model.ASN.ValueInt64()
		controller.ASN = &asn
	}

	controller.BgpMultipathAsPathRelax = BoolToIntBoolPointer(model.BGPMAR)
	controller.Ebgp = BoolToIntBoolPointer(model.EBGP)

	if !model.EBGPMH.IsNull() && !model.EBGPMH.IsUnknown() {
		ebgpMH := model.EBGPMH.ValueInt64()
		controller.EbgpMultihop = &ebgpMH
	}

	if !model.ISISDomain.IsNull() && !model.ISISDomain.IsUnknown() {
		isisDomain := model.ISISDomain.ValueString()
		controller.ISISDomain = &isisDomain
	}

	if !model.ISISIfaces.IsNull() && !model.ISISIfaces.IsUnknown() {
		isisIfaces := model.ISISIfaces.ValueString()
		controller.ISISIfaces = &isisIfaces
	}

	if !model.ISISNet.IsNull() && !model.ISISNet.IsUnknown() {
		isisNet := model.ISISNet.ValueString()
		controller.ISISNet = &isisNet
	}

	if !model.Loopback.IsNull() && !model.Loopback.IsUnknown() {
		loopback := model.Loopback.ValueString()
		controller.Loopback = &loopback
	}

	if !model.Node.IsNull() && !model.Node.IsUnknown() {
		node := model.Node.ValueString()
		controller.Node = &node
	}

	if !model.Peers.IsNull() && !model.Peers.IsUnknown() {
		var peers []string
		diagPeers := model.Peers.ElementsAs(ctx, &peers, false)
		diags.Append(diagPeers...)
		if diagPeers.HasError() {
			return &controller, diags
		}
		controller.Peers = peers
	}

	return &controller, diags

}

func (r *proxmoxSDNControllerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan controllerModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	controller, diags := convertControllerModeltoClientController(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	err := r.client.CreateSDNController(*controller)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to create controller",
			err.Error(),
		)
		return
	}

	state, diags := convertSDNControllertoControllersModel(ctx, *controller)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *proxmoxSDNControllerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var stete controllerModel

	diags := req.State.Get(ctx, &stete)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	controller, err := r.client.GetSDNController(stete.Controller.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to get controller", err.Error())
		return
	}

	updatedState, diagsState := convertSDNControllertoControllersModel(ctx, *controller)
	resp.Diagnostics.Append(diagsState...)
	if diagsState.HasError() {
		return
	}

	diags = resp.State.Set(ctx, updatedState)
	resp.Diagnostics.Append(diags...)
}

func (r *proxmoxSDNControllerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan controllerModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	controller, diags := convertControllerModeltoClientController(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	err := r.client.UpdateSDNController(*controller)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to update controller",
			err.Error(),
		)
		return
	}

	state, diags := convertSDNControllertoControllersModel(ctx, *controller)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *proxmoxSDNControllerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state controllerModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	err := r.client.DeleteSDNController(state.Controller.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to delete controller", err.Error())
		return
	}
}
