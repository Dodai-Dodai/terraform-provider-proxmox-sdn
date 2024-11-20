package provider

import (
	"context"
	"fmt"

	"github.com/Dodai-Dodai/terraform-provider-proxmox-sdn/client"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &proxmoxSDNZoneDataSource{}
	_ datasource.DataSourceWithConfigure = &proxmoxSDNZoneDataSource{}
)

func NewProxmoxSDNZoneDataSource() datasource.DataSource {
	return &proxmoxSDNZoneDataSource{}
}

type proxmoxSDNZoneDataSource struct {
	client *client.SSHProxmoxClient
}

func (d *proxmoxSDNZoneDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zone"
}

func (d *proxmoxSDNZoneDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zones": schema.ListNestedAttribute{
				Description: "List of SDN zones",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
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
							Description: "The DnsZone of the zone",
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
									Description: "The bridge of VLAN zone",
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
									Description: "The bridge of QinQ zone",
									Required:    true,
								},
								"tag": schema.Int64Attribute{
									Description: "The tag num of QinQ zone",
									Required:    true,
								},
								"vlanprotocol": schema.StringAttribute{
									Description: "VLAN Protocol of the zone",
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
									Description: "Set of Peers in the vxlan zone",
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
									Description: "The controller of EVPN zone",
									Required:    true,
								},
								"vrf_vxlan": schema.Int64Attribute{
									Description: "The VRFVXLAN of EVPN zone",
									Required:    true,
								},
								"mac": schema.StringAttribute{
									Description: "The MAC of EVPN zone",
									Optional:    true,
									Computed:    true,
								},
								"exitnodes": schema.SetAttribute{
									Description: "Set of ExitNodes in the EVPN zone",
									ElementType: types.StringType,
									Optional:    true,
									Computed:    true,
								},
								"primaryexitnode": schema.StringAttribute{
									Description: "The PrimaryExitNode of EVPN zone",
									Optional:    true,
									Computed:    true,
								},
								"exitnodeslocalrouting": schema.BoolAttribute{
									Description: "The ExitNodesLocalRouting of EVPN zone",
									Optional:    true,
									Computed:    true,
								},
								"advertisesubnets": schema.BoolAttribute{
									Description: "The AdvertiseSubnets of EVPN zone",
									Optional:    true,
									Computed:    true,
								},
								"disablearpndsuppression": schema.BoolAttribute{
									Description: "The DisableARPNdSuppression of EVPN zone",
									Optional:    true,
									Computed:    true,
								},
								"rtimport": schema.StringAttribute{
									Description: "The RouteTargetImport of EVPN zone",
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *proxmoxSDNZoneDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.SSHProxmoxClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.SSHProxmoxClient, got %T", req.ProviderData),
		)
		return
	}
	d.client = client
}

func convertSDNZoneToZonesModel(ctx context.Context, zone client.SDNZone) (zonesModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	nodeset, diagNodes := types.SetValueFrom(ctx, types.StringType, zone.Nodes)
	if diagNodes.HasError() {
		diags.Append(diagNodes...)
		return zonesModel{}, diags
	}
	diags.Append(diagNodes...)

	zoneModel := zonesModel{
		Zone:       types.StringValue(zone.Zone),
		Type:       types.StringValue(zone.Type),
		MTU:        types.Int64PointerValue(zone.MTU),
		Nodes:      nodeset,
		IPAM:       types.StringPointerValue(zone.IPAM),
		DNS:        types.StringPointerValue(zone.DNS),
		ReverseDNS: types.StringPointerValue(zone.ReverseDNS),
		DNSZone:    types.StringPointerValue(zone.DNSZone),
	}

	if zone.VLAN != nil {
		vlanModel := VLANConfigModel{
			Bridge: types.StringValue(zone.VLAN.Bridge),
		}

		vlanObject, diag := types.ObjectValueFrom(ctx, map[string]attr.Type{
			"bridge": types.StringType,
		}, vlanModel)
		if diag.HasError() {
			diags.Append(diag...)
			return zoneModel, diags
		}

		zoneModel.VLAN = vlanObject
	} else {
		// VLANがnilの場合、nullのObjectを作成
		zoneModel.VLAN = types.ObjectNull(map[string]attr.Type{
			"bridge": types.StringType,
		})
	}

	if zone.QinQ != nil {
		qinqModel := QinQConfigModel{
			Bridge: types.StringValue(zone.QinQ.Bridge),
			Tag:    types.Int64Value(zone.QinQ.Tag),
		}
		if zone.QinQ.VLANProtocol != nil {
			qinqModel.VLANProtocol = types.StringPointerValue(zone.QinQ.VLANProtocol)
		}

		qinqObject, diag := types.ObjectValueFrom(ctx, map[string]attr.Type{
			"bridge":       types.StringType,
			"tag":          types.Int64Type,
			"vlanprotocol": types.StringType,
		}, qinqModel)
		if diag.HasError() {
			diags.Append(diag...)
			return zoneModel, diags
		}

		zoneModel.QinQ = qinqObject
	} else {
		// QinQがnilの場合、nullのObjectを作成
		zoneModel.QinQ = types.ObjectNull(map[string]attr.Type{
			"bridge":       types.StringType,
			"tag":          types.Int64Type,
			"vlanprotocol": types.StringType,
		})
	}

	if zone.VXLAN != nil {
		peerSet, diagPeers := types.SetValueFrom(ctx, types.StringType, zone.VXLAN.Peer)
		if diagPeers.HasError() {
			diags.Append(diagPeers...)
			return zoneModel, diags
		}

		vxlanModel := VXLANConfigModel{
			Peer: peerSet,
		}

		vxlanObject, diag := types.ObjectValueFrom(ctx, map[string]attr.Type{
			"peer": types.SetType{ElemType: types.StringType},
		}, vxlanModel)
		if diag.HasError() {
			diags.Append(diag...)
			return zoneModel, diags
		}

		zoneModel.VXLAN = vxlanObject
	} else {
		// VXLANがnilの場合、nullのObjectを作成
		zoneModel.VXLAN = types.ObjectNull(map[string]attr.Type{
			"peer": types.SetType{ElemType: types.StringType},
		})
	}

	if zone.EVPN != nil {
		var exitNodesSet types.Set
		if zone.EVPN.ExitNodes != nil {
			var diagExitNodes diag.Diagnostics
			exitNodesSet, diagExitNodes = types.SetValueFrom(ctx, types.StringType, zone.EVPN.ExitNodes)
			if diagExitNodes.HasError() {
				diags.Append(diagExitNodes...)
				return zoneModel, diags
			}
		} else {
			// ExitNodesがnilの場合、nullのSetを作成
			exitNodesSet = types.SetNull(types.StringType)
		}

		evpnModel := EVPNConfigModel{
			Controller: types.StringValue(zone.EVPN.Controller),
			VRFVXLAN:   types.Int64Value(zone.EVPN.VRFVXLAN),
			ExitNodes:  exitNodesSet,
		}
		if zone.EVPN.MAC != nil {
			evpnModel.MAC = types.StringPointerValue(zone.EVPN.MAC)
		}
		if zone.EVPN.PrimaryExitNode != nil {
			evpnModel.PrimaryExitNode = types.StringPointerValue(zone.EVPN.PrimaryExitNode)
		}
		if zone.EVPN.ExitNodesLocalRouting != nil {
			evpnModel.ExitNodesLocalRouting = types.BoolPointerValue(zone.EVPN.ExitNodesLocalRouting)
		}
		if zone.EVPN.AdvertiseSubnets != nil {
			evpnModel.AdvertiseSubnets = types.BoolPointerValue(zone.EVPN.AdvertiseSubnets)
		}
		if zone.EVPN.DisableARPNdSuppression != nil {
			evpnModel.DisableARPNdSuppression = types.BoolPointerValue(zone.EVPN.DisableARPNdSuppression)
		}
		if zone.EVPN.RouteTargetImport != nil {
			evpnModel.RouteTargetImport = types.StringPointerValue(zone.EVPN.RouteTargetImport)
		}

		evpnObject, diag := types.ObjectValueFrom(ctx, map[string]attr.Type{
			"controller":              types.StringType,
			"vrf_vxlan":               types.Int64Type,
			"mac":                     types.StringType,
			"exitnodes":               types.SetType{ElemType: types.StringType},
			"primaryexitnode":         types.StringType,
			"exitnodeslocalrouting":   types.BoolType,
			"advertisesubnets":        types.BoolType,
			"disablearpndsuppression": types.BoolType,
			"rtimport":                types.StringType,
		}, evpnModel)
		if diag.HasError() {
			diags.Append(diag...)
			return zoneModel, diags
		}

		zoneModel.EVPN = evpnObject
	} else {
		// EVPNがnilの場合、nullのObjectを作成
		zoneModel.EVPN = types.ObjectNull(map[string]attr.Type{
			"controller":              types.StringType,
			"vrf_vxlan":               types.Int64Type,
			"mac":                     types.StringType,
			"exitnodes":               types.SetType{ElemType: types.StringType},
			"primaryexitnode":         types.StringType,
			"exitnodeslocalrouting":   types.BoolType,
			"advertisesubnets":        types.BoolType,
			"disablearpndsuppression": types.BoolType,
			"rtimport":                types.StringType,
		})
	}

	return zoneModel, diags

}

func (d *proxmoxSDNZoneDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state struct {
		Zones []zonesModel `tfsdk:"zones"`
	}

	zones, err := d.client.GetSDNZones()
	if err != nil {
		resp.Diagnostics.AddError("Error getting zones", err.Error())
		return
	}

	var diags diag.Diagnostics

	for _, zone := range zones {
		zoneModel, diagZone := convertSDNZoneToZonesModel(ctx, zone)
		if diagZone.HasError() {
			resp.Diagnostics.Append(diagZone...)
			continue
		}
		diags.Append(diagZone...)
		state.Zones = append(state.Zones, zoneModel)
	}

	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	diags.Append(resp.State.Set(ctx, state)...)
	resp.Diagnostics.Append(diags...)
}
