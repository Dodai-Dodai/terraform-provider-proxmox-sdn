package provider

import (
	"context"
	"fmt"

	"github.com/Dodai-Dodai/terraform-provider-proxmox-sdn/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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

						"bridge": schema.StringAttribute{
							Description: "The bridge of the zone",
							Optional:    true,
							Computed:    true,
							// PLANがVLANかQinQのとき必須にして、それ以外の場合は無効にするvalidatorを設定
							Validators: []validator.String{
								BridgeValidator{
									TypeAttributeName: "type",
								},
							},
						},

						"tag": schema.Int64Attribute{
							Description: "The tag of the zone",
							Optional:    true,
							Computed:    true,
							// PLANがQinQのとき必須にして、それ以外の場合は無効にするvalidatorを設定
							Validators: []validator.Int64{
								TagValidator{
									TypeAttributeName: "type",
								},
							},
						},

						"vlanprotocol": schema.StringAttribute{
							Description: "The VLAN Protocol of the zone",
							Optional:    true,
							Computed:    true,
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
							Computed:    true,
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
							Computed:    true,
							Validators: []validator.String{
								ControllerValidator{
									TypeAttributeName: "type",
								},
							},
						},

						"vrf_vxlan": schema.Int64Attribute{
							Description: "The VRFVXLAN of the zone",
							Optional:    true,
							Computed:    true,
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
							Computed:    true,
							Validators: []validator.Set{
								ExitNodesValidator{
									TypeAttributeName: "type",
								},
							},
						},

						"primaryexitnode": schema.StringAttribute{
							Description: "The primary exit node of the zone",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{
								PrimaryExitNodeValidator{
									TypeAttributeName: "type",
								},
							},
						},

						"exitnodeslocalrouting": schema.BoolAttribute{
							Description: "The exit nodes local routing of the zone",
							Optional:    true,
							Computed:    true,
							Validators: []validator.Bool{
								ExitnodesLocalRoutingValidator{
									TypeAttributeName: "type",
								},
							},
						},

						"advertisesubnets": schema.BoolAttribute{
							Description: "The advertise subnets of the zone",
							Optional:    true,
							Computed:    true,
							Validators: []validator.Bool{
								AdvertiseSubnetsValidator{
									TypeAttributeName: "type",
								},
							},
						},

						"disablearpndsuppression": schema.BoolAttribute{
							Description: "The disable arp nd suppression of the zone",
							Optional:    true,
							Computed:    true,
							Validators: []validator.Bool{
								DisableARPNdSuppressionValidator{
									TypeAttributeName: "type",
								},
							},
						},

						"rtimport": schema.StringAttribute{
							Description: "The route target import of the zone",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{
								RouteTargetImportValidator{
									TypeAttributeName: "type",
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

func IntBoolPointerToBoolPointer(b *client.IntBool) *bool {
	if b == nil {
		return nil
	}
	boolValue := bool(*b)
	return &boolValue
}

func convertSDNZoneToZonesModel(ctx context.Context, zone client.SDNZone) (zonesModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	nodeset, diagNodes := types.SetValueFrom(ctx, types.StringType, zone.Nodes)
	if diagNodes.HasError() {
		diags.Append(diagNodes...)
		return zonesModel{}, diags
	}
	diags.Append(diagNodes...)

	peersset, diagPeers := types.SetValueFrom(ctx, types.StringType, zone.Peers)
	if diagPeers.HasError() {
		diags.Append(diagPeers...)
		return zonesModel{}, diags
	}
	diags.Append(diagPeers...)

	exitnodesset, diagExitnodes := types.SetValueFrom(ctx, types.StringType, zone.ExitNodes)
	if diagExitnodes.HasError() {
		diags.Append(diagExitnodes...)
		return zonesModel{}, diags
	}
	diags.Append(diagExitnodes...)

	exitNodesLocalRouting := IntBoolPointerToBoolPointer(zone.ExitNodesLocalRouting)
	advertiseSubnets := IntBoolPointerToBoolPointer(zone.AdvertiseSubnets)
	disableARPNdSuppression := IntBoolPointerToBoolPointer(zone.DisableARPNdSuppression)

	zoneModel := zonesModel{
		Zone:                    types.StringValue(zone.Zone),
		Type:                    types.StringValue(zone.Type),
		MTU:                     types.Int64PointerValue(zone.MTU),
		Nodes:                   nodeset,
		IPAM:                    types.StringPointerValue(zone.IPAM),
		DNS:                     types.StringPointerValue(zone.DNS),
		ReverseDNS:              types.StringPointerValue(zone.ReverseDNS),
		DNSZone:                 types.StringPointerValue(zone.DNSZone),
		Bridge:                  types.StringPointerValue(zone.Bridge),
		Tag:                     types.Int64PointerValue(zone.Tag),
		VLANProtocol:            types.StringPointerValue(zone.VLANProtocol),
		Peers:                   peersset,
		Controller:              types.StringPointerValue(zone.Controller),
		VRFVXLAN:                types.Int64PointerValue(zone.VRFVXLAN),
		MAC:                     types.StringPointerValue(zone.MAC),
		ExitNodes:               exitnodesset,
		PrimaryExitNode:         types.StringPointerValue(zone.PrimaryExitNode),
		ExitNodesLocalRouting:   types.BoolPointerValue(exitNodesLocalRouting),
		AdvertiseSubnets:        types.BoolPointerValue(advertiseSubnets),
		DisableARPNdSuppression: types.BoolPointerValue(disableARPNdSuppression),
		RouteTargetImport:       types.StringPointerValue(zone.RouteTargetImport),
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
