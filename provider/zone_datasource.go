package provider

import (
	"context"
	"fmt"

	"github.com/Dodai-Dodai/terraform-provider-proxmox-sdn/client"
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
		zoneModel.VLAN = &VLANConfig{
			Bridge: types.StringValue(zone.VLAN.Bridge),
		}
	}

	if zone.QinQ != nil {
		zoneModel.QinQ = &QinQConfig{
			Bridge:       types.StringValue(zone.QinQ.Bridge),
			Tag:          types.Int64Value(zone.QinQ.Tag),
			VLANProtocol: types.StringPointerValue(zone.QinQ.VLANProtocol),
		}
	}

	if zone.VXLAN != nil {
		peerSet, diagPeers := types.SetValueFrom(ctx, types.StringType, zone.VXLAN.Peer)
		if diagPeers.HasError() {
			diags.Append(diagPeers...)
			return zoneModel, diags
		}
		diags.Append(diagPeers...)

		zoneModel.VXLAN = &VXLANConfig{
			Peer: peerSet,
		}
	}

	if zone.EVPN != nil {
		exitNodesSet, diagExitNodes := types.SetValueFrom(ctx, types.StringType, zone.EVPN.ExitNodes)
		if diagExitNodes.HasError() {
			diags.Append(diagExitNodes...)
			return zoneModel, diags
		}
		diags.Append(diagExitNodes...)

		zoneModel.EVPN = &EVPNConfig{
			Controller:              types.StringValue(zone.EVPN.Controller),
			VRFVXLAN:                types.Int64Value(zone.EVPN.VRFVXLAN),
			MAC:                     types.StringPointerValue(zone.EVPN.MAC),
			ExitNodes:               exitNodesSet,
			PrimaryExitNode:         types.StringPointerValue(zone.EVPN.PrimaryExitNode),
			ExitNodesLocalRouting:   types.BoolPointerValue(zone.EVPN.ExitNodesLocalRouting),
			AdvertiseSubnets:        types.BoolPointerValue(zone.EVPN.AdvertiseSubnets),
			DisableARPNdSuppression: types.BoolPointerValue(zone.EVPN.DisableARPNdSuppression),
			RouteTargetImport:       types.StringPointerValue(zone.EVPN.RouteTargetImport),
		}
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
