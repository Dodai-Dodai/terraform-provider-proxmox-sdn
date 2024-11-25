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
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &proxmoxSDNZoneResource{}
	_ resource.ResourceWithConfigure = &proxmoxSDNZoneResource{}
)

func NewProxmoxSDNZoneResource() resource.Resource {
	return &proxmoxSDNZoneResource{}
}

type proxmoxSDNZoneResource struct {
	client *client.SSHProxmoxClient
}

func (r *proxmoxSDNZoneResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *proxmoxSDNZoneResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zone"
}

func (r *proxmoxSDNZoneResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone": schema.StringAttribute{
				Description: "The name of the zone",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"type": schema.StringAttribute{
				Description: "The type of the zone",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"mtu": schema.Int64Attribute{
				Description: "The MTU Num of the Zone",
				Optional:    true,
				//Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"nodes": schema.SetAttribute{
				Description: "Set of nodes in the zone",
				ElementType: types.StringType,
				Optional:    true,
				//Computed:    true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.RequiresReplace(),
				},
			},
			"ipam": schema.StringAttribute{
				Description: "The IPAM of the zone",
				Optional:    true,
				//Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"dns": schema.StringAttribute{
				Description: "The DNS of the zone",
				Optional:    true,
				//Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"reversedns": schema.StringAttribute{
				Description: "The reverse dns of the zone",
				Optional:    true,
				//Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"dnszone": schema.StringAttribute{
				Description: "The DnsZone of the zone",
				Optional:    true,
				//Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"bridge": schema.StringAttribute{
				Description: "The bridge of the zone",
				Optional:    true,
				//Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					BridgeValidator{
						TypeAttributeName: "type",
					},
				},
			},
			"tag": schema.Int64Attribute{
				Description: "The tag of the zone",
				Optional:    true,
				//Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Validators: []validator.Int64{
					TagValidator{
						TypeAttributeName: "type",
					},
				},
			},
			"vlanprotocol": schema.StringAttribute{
				Description: "The VLAN Protocol of the zone",
				Optional:    true,
				//Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				// PLANがQinQかどうかを確認して、それ以外の場合は無効にするvalidatorを設定、加えて文字列が正しいかを確認
				Validators: []validator.String{
					VLANProtocolValidator{
						TypeAttributeName: "type",
					},
				},
			},

			"peers": schema.SetAttribute{
				Description: "Set of peers in the zone",
				ElementType: types.StringType,
				Optional:    true,
				//Computed:    true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.RequiresReplace(),
				},
				// PLANがVXLANのとき必須にして、それ以外の場合は無効にするvalidatorを設定
				Validators: []validator.Set{
					PeersValidator{
						TypeAttributeName: "type",
					},
				},
			},

			"controller": schema.StringAttribute{
				Description: "The controller of the zone",
				Optional:    true,
				//Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					ControllerValidator{
						TypeAttributeName: "type",
					},
				},
			},

			"vrf_vxlan": schema.Int64Attribute{
				Description: "The VRFVXLAN of the zone",
				Optional:    true,
				//Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Validators: []validator.Int64{
					VrfvxlanValidator{
						TypeAttributeName: "type",
					},
				},
			},

			"mac": schema.StringAttribute{
				Description: "The MAC of the zone",
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					MACValidator{
						TypeAttributeName: "type",
					},
				},
			},

			"exitnodes": schema.SetAttribute{
				Description: "Set of exit nodes in the zone",
				ElementType: types.StringType,
				Optional:    true,
				//Computed:    true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.RequiresReplace(),
				},
				Validators: []validator.Set{
					ExitNodesValidator{
						TypeAttributeName: "type",
					},
				},
			},

			"primaryexitnode": schema.StringAttribute{
				Description: "The primary exit node of the zone",
				Optional:    true,
				//Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					PrimaryExitNodeValidator{
						TypeAttributeName: "type",
					},
				},
			},

			"exitnodeslocalrouting": schema.BoolAttribute{
				Description: "The exit nodes local routing of the zone",
				Optional:    true,
				//Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Validators: []validator.Bool{
					ExitnodesLocalRoutingValidator{
						TypeAttributeName: "type",
					},
				},
			},

			"advertisesubnets": schema.BoolAttribute{
				Description: "The advertise subnets of the zone",
				Optional:    true,
				//Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Validators: []validator.Bool{
					AdvertiseSubnetsValidator{
						TypeAttributeName: "type",
					},
				},
			},

			"disablearpndsuppression": schema.BoolAttribute{
				Description: "The disable arp nd suppression of the zone",
				Optional:    true,
				//Computed:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Validators: []validator.Bool{
					DisableARPNdSuppressionValidator{
						TypeAttributeName: "type",
					},
				},
			},

			"rtimport": schema.StringAttribute{
				Description: "The route target import of the zone",
				Optional:    true,
				//Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					RouteTargetImportValidator{
						TypeAttributeName: "type",
					},
				},
			},
		},
	}
}

func (r *proxmoxSDNZoneResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan zonesModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	zone, diags := convertZonesModeltoClientSDNZone(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	err := r.client.CreateSDNZone(*zone)
	if err != nil {
		resp.Diagnostics.AddError("Failed to create SDN Zone", err.Error())
		return
	}

	createdZone, err := r.client.GetSDNZone(zone.Zone)
	if err != nil {
		resp.Diagnostics.AddError("Failed to get SDN Zone", err.Error())
		return
	}

	state, diags := convertSDNZoneToZonesModel(ctx, *createdZone)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)

}

func BoolToIntBoolPointer(b types.Bool) *client.IntBool {
	if b.IsNull() || b.IsUnknown() {
		return nil
	}
	value := b.ValueBool()
	intBoolValue := client.IntBool(value)
	return &intBoolValue
}

func convertZonesModeltoClientSDNZone(ctx context.Context, model zonesModel) (*client.SDNZone, diag.Diagnostics) {
	var diags diag.Diagnostics

	var zone client.SDNZone

	zone.Zone = model.Zone.ValueString()
	zone.Type = model.Type.ValueString()

	if !model.MTU.IsUnknown() && !model.MTU.IsNull() {
		mtu := model.MTU.ValueInt64()
		zone.MTU = &mtu
	}

	if !model.IPAM.IsUnknown() && !model.IPAM.IsNull() {
		ipam := model.IPAM.ValueString()
		zone.IPAM = &ipam
	}

	if !model.DNS.IsUnknown() && !model.DNS.IsNull() {
		dns := model.DNS.ValueString()
		zone.DNS = &dns
	}

	if !model.ReverseDNS.IsUnknown() && !model.ReverseDNS.IsNull() {
		reversedns := model.ReverseDNS.ValueString()
		zone.ReverseDNS = &reversedns
	}

	if !model.DNSZone.IsUnknown() && !model.DNSZone.IsNull() {
		dnszone := model.DNSZone.ValueString()
		zone.DNSZone = &dnszone
	}

	if !model.Bridge.IsUnknown() && !model.Bridge.IsNull() {
		bridge := model.Bridge.ValueString()
		zone.Bridge = &bridge
	}

	if !model.Tag.IsUnknown() && !model.Tag.IsNull() {
		tag := model.Tag.ValueInt64()
		zone.Tag = &tag
	}

	if !model.VLANProtocol.IsUnknown() && !model.VLANProtocol.IsNull() {
		vlanprotocol := model.VLANProtocol.ValueString()
		zone.VLANProtocol = &vlanprotocol
	}

	if !model.Controller.IsUnknown() && !model.Controller.IsNull() {
		controller := model.Controller.ValueString()
		zone.Controller = &controller
	}

	if !model.VRFVXLAN.IsUnknown() && !model.VRFVXLAN.IsNull() {
		vrfvxlan := model.VRFVXLAN.ValueInt64()
		zone.VRFVXLAN = &vrfvxlan
	}

	if !model.MAC.IsUnknown() && !model.MAC.IsNull() {
		mac := model.MAC.ValueString()
		zone.MAC = &mac
	}

	if !model.PrimaryExitNode.IsUnknown() && !model.PrimaryExitNode.IsNull() {
		primaryexitnode := model.PrimaryExitNode.ValueString()
		zone.PrimaryExitNode = &primaryexitnode
	}

	zone.ExitNodesLocalRouting = BoolToIntBoolPointer(model.ExitNodesLocalRouting)
	zone.AdvertiseSubnets = BoolToIntBoolPointer(model.AdvertiseSubnets)
	zone.DisableARPNdSuppression = BoolToIntBoolPointer(model.DisableARPNdSuppression)

	if !model.RouteTargetImport.IsUnknown() && !model.RouteTargetImport.IsNull() {
		rtimport := model.RouteTargetImport.ValueString()
		zone.RouteTargetImport = &rtimport
	}

	// nodes, peers, exitnodesの変換
	if !model.Nodes.IsUnknown() && !model.Nodes.IsNull() {
		var nodes []string
		diagNodes := model.Nodes.ElementsAs(ctx, &nodes, false)
		diags.Append(diagNodes...)
		if diagNodes.HasError() {
			return &zone, diags
		}
		zone.Nodes = nodes
	}

	if !model.Peers.IsUnknown() && !model.Peers.IsNull() {
		var peers []string
		diagPeers := model.Peers.ElementsAs(ctx, &peers, false)
		diags.Append(diagPeers...)
		if diagPeers.HasError() {
			return &zone, diags
		}
		zone.Peers = peers
	}

	if !model.ExitNodes.IsUnknown() && !model.ExitNodes.IsNull() {
		var exitnodes []string
		diagExitnodes := model.ExitNodes.ElementsAs(ctx, &exitnodes, false)
		diags.Append(diagExitnodes...)
		if diagExitnodes.HasError() {
			return &zone, diags
		}
		zone.ExitNodes = exitnodes
	}
	return &zone, diags
}

func (r *proxmoxSDNZoneResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state zonesModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	zone, err := r.client.GetSDNZone(state.Zone.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to get SDN Zone", err.Error())
		return
	}

	updatedState, diagsState := convertSDNZoneToZonesModel(ctx, *zone)
	resp.Diagnostics.Append(diagsState...)
	if diagsState.HasError() {
		return
	}

	diags = resp.State.Set(ctx, updatedState)
	resp.Diagnostics.Append(diags...)
}

func (r *proxmoxSDNZoneResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// プランからデータを取得
	var plan zonesModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// 現在の状態を取得
	var state zonesModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// プランからSDNZoneを生成
	zone, diags := convertZonesModeltoClientSDNZone(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// SDNゾーンを更新
	err := r.client.UpdateSDNZone(*zone)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating SDN Zone",
			fmt.Sprintf("Could not update SDN Zone %s: %s", zone.Zone, err.Error()),
		)
		return
	}

	// 更新されたゾーンを取得
	updatedZone, err := r.client.GetSDNZone(zone.Zone)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading SDN Zone",
			fmt.Sprintf("Could not read SDN Zone %s: %s", zone.Zone, err.Error()),
		)
		return
	}

	// 状態を更新
	updatedState, diags := convertSDNZoneToZonesModel(ctx, *updatedZone)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, updatedState)
	resp.Diagnostics.Append(diags...)
}

func (r *proxmoxSDNZoneResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state zonesModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	err := r.client.DeleteSDNZone(state.Zone.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to delete SDN Zone", err.Error())
		return
	}
}
