package provider

import (
	"context"
	"fmt"

	client2 "github.com/Dodai-Dodai/terraform-provider-proxmox-sdn/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
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
	client *client2.SSHProxmoxClient
}

func (d *proxmoxSDNZoneDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zone"
}

func (d *proxmoxSDNZoneDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Proxmox SDN Zone Data Source",
		Attributes: map[string]schema.Attribute{
			"zones": schema.ListNestedAttribute{
				Description: "The list of zones",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"zone": schema.StringAttribute{
							Description: "The ID of the zone",
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: "The type of the zone",
							Computed:    true,
						},
						"dhcp": schema.StringAttribute{
							Description: "The DHCP configuration of the zone",
							Computed:    true,
						},
						"dns": schema.StringAttribute{
							Description: "The DNS configuration of the zone",
							Computed:    true,
						},
						"dns_zone": schema.StringAttribute{
							Description: "The DNS zone of the zone",
							Computed:    true,
						},
						"digest": schema.StringAttribute{
							Description: "The digest of the zone",
							Computed:    true,
						},
						"ipam": schema.StringAttribute{
							Description: "The IPAM configuration of the zone",
							Computed:    true,
						},
						"mtu": schema.Int64Attribute{
							Description: "The MTU of the zone",
							Computed:    true,
						},
						"nodes": schema.StringAttribute{
							Description: "The nodes of the zone",
							Computed:    true,
						},
						"pending": schema.BoolAttribute{
							Description: "The pending status of the zone",
							Computed:    true,
						},
						"reverse_dns": schema.StringAttribute{
							Description: "The reverse DNS configuration of the zone",
							Computed:    true,
						},
						"state": schema.StringAttribute{
							Description: "The state of the zone",
							Computed:    true,
						},
						"advertise_subnets": schema.BoolAttribute{
							Description: "The advertise subnets configuration of the zone",
							Computed:    true,
						},
						"bridge": schema.StringAttribute{
							Description: "The bridge of the zone",
							Computed:    true,
						},
						"bridge_disable_mac_learning": schema.BoolAttribute{
							Description: "The bridge disable MAC learning configuration of the zone",
							Computed:    true,
						},
						"controller": schema.StringAttribute{
							Description: "The controller of the zone",
							Computed:    true,
						},
						"disable_arp_discovery": schema.BoolAttribute{
							Description: "The disable ARP discovery configuration of the zone",
							Computed:    true,
						},
						"dp_id": schema.Int64Attribute{
							Description: "The DP ID of the zone",
							Computed:    true,
						},
						"exit_nodes": schema.StringAttribute{
							Description: "The exit nodes of the zone",
							Computed:    true,
						},
						"exit_nodes_local_routing": schema.BoolAttribute{
							Description: "The exit nodes local routing configuration of the zone",
							Computed:    true,
						},
						"mac": schema.StringAttribute{
							Description: "The MAC of the zone",
							Computed:    true,
						},
						"peers": schema.StringAttribute{
							Description: "The peers of the zone",
							Computed:    true,
						},
						"route_target_import": schema.StringAttribute{
							Description: "The route target import of the zone",
							Computed:    true,
						},
						// "tag": schema.Int64Attribute{
						// 	Description: "The tag of the zone",
						// 	Computed:    true,
						// },
						"vlan_protocol": schema.StringAttribute{
							Description: "The VLAN protocol of the zone",
							Computed:    true,
						},
						"vrf_vxlan": schema.Int64Attribute{
							Description: "The VRF VXLAN of the zone",
							Computed:    true,
						},
						"vxlan_port": schema.Int64Attribute{
							Description: "The VXLAN port of the zone",
							Computed:    true,
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

	client, ok := req.ProviderData.(*client2.SSHProxmoxClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client2.Client, got %T", req.ProviderData),
		)
		return
	}
	d.client = client
}

func convertSDNZoneToZonesModel(zone client2.SDNZone) zonesModel {
	derefString := func(s *string) string {
		if s == nil {
			return ""
		}
		return *s
	}

	derefInt := func(i *int) int {
		if i == nil {
			return 0
		}
		return *i
	}

	derefBool := func(b *bool) bool {
		if b == nil {
			return false
		}
		return *b
	}
	zoneState := zonesModel{
		Zone:                     types.StringValue(zone.Zone),
		Type:                     types.StringValue(zone.Type),
		DHCP:                     types.StringValue(derefString(zone.DHCP)),
		DNS:                      types.StringValue(derefString(zone.DNS)),
		DNSZone:                  types.StringValue(derefString(zone.DNSZone)),
		Digest:                   types.StringValue(derefString(zone.Digest)),
		IPAM:                     types.StringValue(derefString(zone.IPAM)),
		MTU:                      types.Int64Value(int64(derefInt(zone.MTU))),
		Nodes:                    types.StringValue(derefString(zone.Nodes)),
		Pending:                  types.BoolValue(derefBool(zone.Pending)),
		ReverseDNS:               types.StringValue(derefString(zone.ReverseDNS)),
		State:                    types.StringValue(derefString(zone.State)),
		AdvertiseSubnets:         types.BoolValue(derefBool(zone.AdvertiseSubnets)),
		Bridge:                   types.StringValue(derefString(zone.Bridge)),
		BridgeDisableMACLearning: types.BoolValue(derefBool(zone.BridgeDisableMACLearning)),
		Controller:               types.StringValue(derefString(zone.Controller)),
		DisableARPDiscovery:      types.BoolValue(derefBool(zone.DisableARPDiscovery)),
		DPID:                     types.Int64Value(int64(derefInt(zone.DPID))),
		ExitNodes:                types.StringValue(derefString(zone.ExitNodes)),
		ExitNodesLocalRouting:    types.BoolValue(derefBool(zone.ExitNodesLocalRouting)),
		MAC:                      types.StringValue(derefString(zone.MAC)),
		Peers:                    types.StringValue(derefString(zone.Peers)),
		RouteTargetImport:        types.StringValue(derefString(zone.RouteTargetImport)),
		// Tag:                      types.Int64Null(),
		VLANProtocol: types.StringValue(derefString(zone.VLANProtocol)),
		VRFVXLAN:     types.Int64Value(int64(derefInt(zone.VRFVXLAN))),
		VXLANPort:    types.Int64Value(int64(derefInt(zone.VXLANPort))),
	}

	// if zone.Tag != nil {
	// 	zoneState.Tag = types.Int64Value(int64(*zone.Tag))
	// }

	return zoneState
}

type proxmoxSDNZoneDataSourceModel struct {
	Zones []zonesModel `tfsdk:"zones"`
}

// proxmoxSDNZoneDataSourceModel maps the data source schema data.
type zonesModel struct {
	Zone                     types.String `tfsdk:"zone"`
	Type                     types.String `tfsdk:"type"`
	DHCP                     types.String `tfsdk:"dhcp"`
	DNS                      types.String `tfsdk:"dns"`
	DNSZone                  types.String `tfsdk:"dns_zone"`
	Digest                   types.String `tfsdk:"digest"`
	IPAM                     types.String `tfsdk:"ipam"`
	MTU                      types.Int64  `tfsdk:"mtu"`
	Nodes                    types.String `tfsdk:"nodes"`
	Pending                  types.Bool   `tfsdk:"pending"`
	ReverseDNS               types.String `tfsdk:"reverse_dns"`
	State                    types.String `tfsdk:"state"`
	AdvertiseSubnets         types.Bool   `tfsdk:"advertise_subnets"`
	Bridge                   types.String `tfsdk:"bridge"`
	BridgeDisableMACLearning types.Bool   `tfsdk:"bridge_disable_mac_learning"`
	Controller               types.String `tfsdk:"controller"`
	DisableARPDiscovery      types.Bool   `tfsdk:"disable_arp_discovery"`
	DPID                     types.Int64  `tfsdk:"dp_id"`
	ExitNodes                types.String `tfsdk:"exit_nodes"`
	ExitNodesLocalRouting    types.Bool   `tfsdk:"exit_nodes_local_routing"`
	MAC                      types.String `tfsdk:"mac"`
	Peers                    types.String `tfsdk:"peers"`
	RouteTargetImport        types.String `tfsdk:"route_target_import"`
	// Tag                      types.Int64  `tfsdk:"tag"`
	VLANProtocol types.String `tfsdk:"vlan_protocol"`
	VRFVXLAN     types.Int64  `tfsdk:"vrf_vxlan"`
	VXLANPort    types.Int64  `tfsdk:"vxlan_port"`
}

func (d *proxmoxSDNZoneDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state proxmoxSDNZoneDataSourceModel

	zones, err := d.client.GetSDNZones()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get zones", err.Error())
		return
	}

	for _, zone := range zones {
		zoneState := convertSDNZoneToZonesModel(zone)
		state.Zones = append(state.Zones, zoneState)
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}
