package provider

import (
	"context"
	"fmt"

	"github.com/Dodai-Dodai/terraform-provider-proxmox-sdn/client"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

//type proxmoxSDNZoneResourceModel = zonesModel

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

func (r *proxmoxSDNZoneResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.SSHProxmoxClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got %T", req.ProviderData),
		)
		return
	}
	r.client = client
}

func (r *proxmoxSDNZoneResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zone"
}

func (r *proxmoxSDNZoneResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone": schema.StringAttribute{
				Description: "The name of the zone",
				Required:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of the zone",
				Required:    true,
			},
			"mtu": schema.Int64Attribute{
				Description: "The MTU Num of the Zone",
				Optional:    true,
				Computed:    true,
			},
			"nodes": schema.SetAttribute{
				Description: "Set of nodes in the zone",
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
			},
			"ipam": schema.StringAttribute{
				Description: "The IPAM of the zone",
				Optional:    true,
				Computed:    true,
			},
			"dns": schema.StringAttribute{
				Description: "The DNS of the zone",
				Optional:    true,
				Computed:    true,
			},
			"reversedns": schema.StringAttribute{
				Description: "The reverse dns of the zone",
				Optional:    true,
				Computed:    true,
			},
			"dnszone": schema.StringAttribute{
				Description: "The DNS zone of the zone",
				Optional:    true,
				Computed:    true,
			},

			// VLAN Config
			"vlan": schema.SingleNestedAttribute{
				Description: "The VLAN configuration of the zone",
				Optional:    true,
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"bridge": schema.StringAttribute{
						Description: "The bridge of the VLAN",
						Required:    true,
					},
				},
			},

			// QinQ Config
			"qinq": schema.SingleNestedAttribute{
				Description: "The QinQ configuration of the zone",
				Optional:    true,
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"bridge": schema.StringAttribute{
						Description: "The bridge of the QinQ",
						Required:    true,
					},
					"tag": schema.Int64Attribute{
						Description: "The tag of the QinQ",
						Required:    true,
					},
					"vlanprotocol": schema.StringAttribute{
						Description: "The VLAN protocol of the QinQ",
						Optional:    true,
						Computed:    true,
					},
				},
			},

			// VXLAN Config
			"vxlan": schema.SingleNestedAttribute{
				Description: "The VXLAN configuration of the zone",
				Optional:    true,
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"peer": schema.SetAttribute{
						Description: "The peer of the VXLAN",
						ElementType: types.StringType,
						Required:    true,
					},
				},
			},

			// EVPN Config
			"evpn": schema.SingleNestedAttribute{
				Description: "The EVPN configuration of the zone",
				Optional:    true,
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"controller": schema.StringAttribute{
						Description: "The controller of the EVPN",
						Required:    true,
					},
					"vrf_vxlan": schema.Int64Attribute{
						Description: "The VRF VXLAN of the EVPN",
						Required:    true,
					},
					"mac": schema.StringAttribute{
						Description: "The MAC of the EVPN",
						Optional:    true,
						Computed:    true,
					},
					"exitnodes": schema.SetAttribute{
						Description: "The exit nodes of the EVPN",
						ElementType: types.StringType,
						Optional:    true,
						Computed:    true,
					},
					"primaryexitnode": schema.StringAttribute{
						Description: "The primary exit node of the EVPN",
						Optional:    true,
						Computed:    true,
					},
					"exitnodeslocalrouting": schema.BoolAttribute{
						Description: "The exit nodes local routing of the EVPN",
						Optional:    true,
						Computed:    true,
					},
					"advertisesubnets": schema.BoolAttribute{
						Description: "The advertise subnets of the EVPN",
						Optional:    true,
						Computed:    true,
					},
					"disablearpndsuppression": schema.BoolAttribute{
						Description: "The disable arp nd suppression of the EVPN",
						Optional:    true,
						Computed:    true,
					},
					"rtimport": schema.StringAttribute{
						Description: "The route import of the EVPN",
						Optional:    true,
						Computed:    true,
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
	if resp.Diagnostics.HasError() {
		return
	}

	// planの構造体をclientの構造体に変換
	zone, diagsConvert := convertZonesModeltoClientSDNZone(ctx, plan)
	resp.Diagnostics.Append(diagsConvert...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.CreateSDNZone(*zone)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to create SDN zone",
			err.Error(),
		)
		return
	}

	createdZone, err := r.client.GetSDNZone(zone.Zone)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to get created SDN zone",
			err.Error(),
		)
		return
	}

	state, diagsState := convertSDNZoneToZonesModel(ctx, *createdZone)
	resp.Diagnostics.Append(diagsState...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func convertZonesModeltoClientSDNZone(ctx context.Context, model zonesModel) (*client.SDNZone, diag.Diagnostics) {
	var diags diag.Diagnostics

	zone := &client.SDNZone{
		Zone: model.Zone.ValueString(),
		Type: model.Type.ValueString(),
	}

	if !model.MTU.IsNull() {
		mtu := model.MTU.ValueInt64()
		zone.MTU = &mtu
	}

	if !model.Nodes.IsNull() {
		var nodes []string
		diagNodes := model.Nodes.ElementsAs(ctx, &nodes, false)
		diags.Append(diagNodes...)
		if diagNodes.HasError() {
			return nil, diags
		}
		zone.Nodes = nodes
	}

	if !model.IPAM.IsNull() {
		ipam := model.IPAM.ValueString()
		zone.IPAM = &ipam
	}

	if !model.DNS.IsNull() {
		dns := model.DNS.ValueString()
		zone.DNS = &dns
	}

	if !model.ReverseDNS.IsNull() {
		reverseDNS := model.ReverseDNS.ValueString()
		zone.ReverseDNS = &reverseDNS
	}

	if !model.DNSZone.IsNull() {
		dnsZone := model.DNSZone.ValueString()
		zone.DNSZone = &dnsZone
	}

	if model.VLAN != nil {
		zone.VLAN = &client.VLANConfig{
			Bridge: model.VLAN.Bridge.ValueString(),
		}
	}

	if model.QinQ != nil {
		zone.QinQ = &client.QinQConfig{
			Bridge: model.QinQ.Bridge.ValueString(),
			Tag:    model.QinQ.Tag.ValueInt64(),
		}
		if !model.QinQ.VLANProtocol.IsNull() {
			vlanProtocol := model.QinQ.VLANProtocol.ValueString()
			zone.QinQ.VLANProtocol = &vlanProtocol
		}
	}

	if model.VXLAN != nil {
		var peers []string
		diagsPeers := model.VXLAN.Peer.ElementsAs(ctx, &peers, false)
		diags.Append(diagsPeers...)
		if diagsPeers.HasError() {
			return nil, diags
		}
		zone.VXLAN = &client.VXLANConfig{
			Peer: peers,
		}
	}

	if model.EVPN != nil {
		evpnConfig := &client.EVPNConfig{
			Controller: model.EVPN.Controller.ValueString(),
			VRFVXLAN:   model.EVPN.VRFVXLAN.ValueInt64(),
		}

		if !model.EVPN.MAC.IsNull() {
			mac := model.EVPN.MAC.ValueString()
			evpnConfig.MAC = &mac
		}

		if !model.EVPN.ExitNodes.IsNull() {
			var exitNodes []string
			diagsExitNodes := model.EVPN.ExitNodes.ElementsAs(ctx, &exitNodes, false)
			diags.Append(diagsExitNodes...)
			if diagsExitNodes.HasError() {
				return nil, diags
			}
			evpnConfig.ExitNodes = exitNodes
		}

		if !model.EVPN.PrimaryExitNode.IsNull() {
			primaryExitNode := model.EVPN.PrimaryExitNode.ValueString()
			evpnConfig.PrimaryExitNode = &primaryExitNode
		}

		if !model.EVPN.ExitNodesLocalRouting.IsNull() {
			exitNodesLocalRouting := model.EVPN.ExitNodesLocalRouting.ValueBool()
			evpnConfig.ExitNodesLocalRouting = &exitNodesLocalRouting
		}

		if !model.EVPN.AdvertiseSubnets.IsNull() {
			advertiseSubnets := model.EVPN.AdvertiseSubnets.ValueBool()
			evpnConfig.AdvertiseSubnets = &advertiseSubnets
		}

		if !model.EVPN.DisableARPNdSuppression.IsNull() {
			disableARPNdSuppression := model.EVPN.DisableARPNdSuppression.ValueBool()
			evpnConfig.DisableARPNdSuppression = &disableARPNdSuppression
		}

		if !model.EVPN.RouteTargetImport.IsNull() {
			routeTargetImport := model.EVPN.RouteTargetImport.ValueString()
			evpnConfig.RouteTargetImport = &routeTargetImport
		}

		zone.EVPN = evpnConfig
	}
	return zone, diags
}

// Read refreshes the Terraform state with the latest data.
func (r *proxmoxSDNZoneResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state zonesModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	zone, err := r.client.GetSDNZone(state.Zone.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to get SDN zone",
			err.Error(),
		)
		return
	}

	updatedState, diagsState := convertSDNZoneToZonesModel(ctx, *zone)
	resp.Diagnostics.Append(diagsState...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, updatedState)
	resp.Diagnostics.Append(diags...)
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *proxmoxSDNZoneResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan zonesModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	zone, diagsConvert := convertZonesModeltoClientSDNZone(ctx, plan)
	resp.Diagnostics.Append(diagsConvert...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.UpdateSDNZone(*zone)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to update SDN zone",
			err.Error(),
		)
		return
	}

	updatedZone, err := r.client.GetSDNZone(zone.Zone)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to get updated SDN zone",
			err.Error(),
		)
		return
	}

	updatedSteta, diagsState := convertSDNZoneToZonesModel(ctx, *updatedZone)
	resp.Diagnostics.Append(diagsState...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, updatedSteta)
	resp.Diagnostics.Append(diags...)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *proxmoxSDNZoneResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state zonesModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteSDNZone(state.Zone.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to delete SDN zone",
			err.Error(),
		)
		return
	}

}
