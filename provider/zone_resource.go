package provider

import (
	"context"
	"fmt"

	client2 "github.com/Dodai-Dodai/terraform-provider-proxmox-sdn/client"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-framework/resource"
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
	client *client2.SSHProxmoxClient
}

func (r *proxmoxSDNZoneResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = client
}

func (r *proxmoxSDNZoneResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zone"
}

func (r *proxmoxSDNZoneResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone": schema.StringAttribute{
				Description: "The ID of the zone",
				Required:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of the zone",
				Required:    true,
			},
			"dhcp": schema.StringAttribute{
				Description: "The DHCP configuration of the zone",
				Optional:    true,
				Computed:    true,
			},
			"dns": schema.StringAttribute{
				Description: "The DNS configuration of the zone",
				Optional:    true,
				Computed:    true,
			},
			"dns_zone": schema.StringAttribute{
				Description: "The DNS zone of the zone",
				Optional:    true,
				Computed:    true,
			},
			"digest": schema.StringAttribute{
				Description: "The digest of the zone",
				Computed:    true,
			},
			"ipam": schema.StringAttribute{
				Description: "The IPAM configuration of the zone",
				Optional:    true,
				Computed:    true,
			},
			"mtu": schema.Int64Attribute{
				Description: "The MTU of the zone",
				Optional:    true,
				Computed:    true,
			},
			"nodes": schema.StringAttribute{
				Description: "The nodes of the zone",
				Optional:    true,
				Computed:    true,
			},
			"pending": schema.BoolAttribute{
				Description: "The pending status of the zone",
				Computed:    true,
			},
			"reverse_dns": schema.StringAttribute{
				Description: "The reverse DNS configuration of the zone",
				Optional:    true,
				Computed:    true,
			},
			"state": schema.StringAttribute{
				Description: "The state of the zone",
				Computed:    true,
			},
			"advertise_subnets": schema.BoolAttribute{
				Description: "The advertise subnets status of the zone",
				Optional:    true,
				Computed:    true,
			},
			"bridge": schema.StringAttribute{
				Description: "The bridge of the zone",
				Optional:    true,
				Computed:    true,
			},
			"bridge_disable_mac_learning": schema.BoolAttribute{
				Description: "The bridge disable MAC learning status of the zone",
				Optional:    true,
				Computed:    true,
			},
			"controller": schema.StringAttribute{
				Description: "The controller of the zone",
				Optional:    true,
				Computed:    true,
			},
			"disable_arp_discovery": schema.BoolAttribute{
				Description: "The disable ARP discovery status of the zone",
				Optional:    true,
				Computed:    true,
			},
			"dp_id": schema.Int64Attribute{
				Description: "The DP ID of the zone",
				Optional:    true,
				Computed:    true,
			},
			"exit_nodes": schema.StringAttribute{
				Description: "The exit nodes of the zone",
				Optional:    true,
				Computed:    true,
			},
			"exit_nodes_local_routing": schema.BoolAttribute{
				Description: "The exit nodes local routing status of the zone",
				Optional:    true,
				Computed:    true,
			},
			"mac": schema.StringAttribute{
				Description: "The MAC of the zone",
				Optional:    true,
				Computed:    true,
			},
			"peers": schema.StringAttribute{
				Description: "The peers of the zone",
				Optional:    true,
				Computed:    true,
			},
			"route_target_import": schema.StringAttribute{
				Description: "The route target import of the zone",
				Optional:    true,
				Computed:    true,
			},
			// "tag": schema.Int64Attribute{
			// 	Description: "The tag of the zone",
			// 	Optional:    true,
			// },
			"vlan_protocol": schema.StringAttribute{
				Description: "The VLAN protocol of the zone",
				Optional:    true,
				Computed:    true,
			},
			"vrf_vxlan": schema.Int64Attribute{
				Description: "The VRF VXLAN of the zone",
				Optional:    true,
				Computed:    true,
			},
			"vxlan_port": schema.Int64Attribute{
				Description: "The VXLAN port of the zone",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

type proxmoxSDNZoneResourceModel struct {
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

// ヘルパー関数の定義

// String型のマッピング
func setStringIfNotNull(dest **string, src types.String) {
	if !src.IsNull() {
		value := src.ValueString()
		*dest = &value
	}
}

// Int64型のマッピング
func setIntIfNotNull(dest **int64, src types.Int64) {
	if !src.IsNull() {
		value := src.ValueInt64()
		*dest = &value
	}
}

// Bool型のマッピング
func setBoolIfNotNull(dest **bool, src types.Bool) {
	if !src.IsNull() {
		value := src.ValueBool()
		*dest = &value
	}
}

func (r *proxmoxSDNZoneResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan proxmoxSDNZoneResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	zone := client2.SDNZone{
		Zone: plan.Zone.ValueString(),
		Type: plan.Type.ValueString(),
	}

	setStringIfNotNull(&zone.DHCP, plan.DHCP)
	setStringIfNotNull(&zone.DNS, plan.DNS)
	setStringIfNotNull(&zone.DNSZone, plan.DNSZone)
	setStringIfNotNull(&zone.Digest, plan.Digest)
	setStringIfNotNull(&zone.IPAM, plan.IPAM)
	setIntIfNotNull(&zone.MTU, plan.MTU)
	setStringIfNotNull(&zone.Nodes, plan.Nodes)
	setBoolIfNotNull(&zone.Pending, plan.Pending)
	setStringIfNotNull(&zone.ReverseDNS, plan.ReverseDNS)
	setStringIfNotNull(&zone.State, plan.State)
	setBoolIfNotNull(&zone.AdvertiseSubnets, plan.AdvertiseSubnets)
	setStringIfNotNull(&zone.Bridge, plan.Bridge)
	setBoolIfNotNull(&zone.BridgeDisableMACLearning, plan.BridgeDisableMACLearning)
	setStringIfNotNull(&zone.Controller, plan.Controller)
	setBoolIfNotNull(&zone.DisableARPDiscovery, plan.DisableARPDiscovery)
	setIntIfNotNull(&zone.DPID, plan.DPID)
	setStringIfNotNull(&zone.ExitNodes, plan.ExitNodes)
	setBoolIfNotNull(&zone.ExitNodesLocalRouting, plan.ExitNodesLocalRouting)
	setStringIfNotNull(&zone.MAC, plan.MAC)
	setStringIfNotNull(&zone.Peers, plan.Peers)
	setStringIfNotNull(&zone.RouteTargetImport, plan.RouteTargetImport)
	// if !plan.Tag.IsNull() {
	// 	tag := int(plan.Tag.ValueInt64())
	// 	zone.Tag = &tag
	// }
	setStringIfNotNull(&zone.VLANProtocol, plan.VLANProtocol)
	setIntIfNotNull(&zone.VRFVXLAN, plan.VRFVXLAN)
	setIntIfNotNull(&zone.VXLANPort, plan.VXLANPort)

	err := r.client.CreateSDNZone(zone)
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

	// ヘルパー関数を定義
	derefString := func(s *string) types.String {
		if s == nil {
			return types.StringNull()
		}
		return types.StringValue(*s)
	}

	derefInt := func(i *int64) types.Int64 {
		if i == nil {
			return types.Int64Null()
		}
		return types.Int64Value(int64(*i))
	}

	derefBool := func(b *bool) types.Bool {
		if b == nil {
			return types.BoolNull()
		}
		return types.BoolValue(*b)
	}

	// planにすべての属性を設定
	plan.Zone = types.StringValue(createdZone.Zone)
	plan.Type = types.StringValue(createdZone.Type)
	plan.DHCP = derefString(createdZone.DHCP)
	plan.DNS = derefString(createdZone.DNS)
	plan.DNSZone = derefString(createdZone.DNSZone)
	plan.Digest = derefString(createdZone.Digest)
	plan.IPAM = derefString(createdZone.IPAM)
	plan.MTU = derefInt(createdZone.MTU)
	plan.Nodes = derefString(createdZone.Nodes)
	plan.Pending = derefBool(createdZone.Pending)
	plan.ReverseDNS = derefString(createdZone.ReverseDNS)
	plan.State = derefString(createdZone.State)
	plan.AdvertiseSubnets = derefBool(createdZone.AdvertiseSubnets)
	plan.Bridge = derefString(createdZone.Bridge)
	plan.BridgeDisableMACLearning = derefBool(createdZone.BridgeDisableMACLearning)
	plan.Controller = derefString(createdZone.Controller)
	plan.DisableARPDiscovery = derefBool(createdZone.DisableARPDiscovery)
	plan.DPID = derefInt(createdZone.DPID)
	plan.ExitNodes = derefString(createdZone.ExitNodes)
	plan.ExitNodesLocalRouting = derefBool(createdZone.ExitNodesLocalRouting)
	plan.MAC = derefString(createdZone.MAC)
	plan.Peers = derefString(createdZone.Peers)
	plan.RouteTargetImport = derefString(createdZone.RouteTargetImport)
	plan.VLANProtocol = derefString(createdZone.VLANProtocol)
	plan.VRFVXLAN = derefInt(createdZone.VRFVXLAN)
	plan.VXLANPort = derefInt(createdZone.VXLANPort)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *proxmoxSDNZoneResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state proxmoxSDNZoneResourceModel
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

	// helper functions
	derefString := func(s *string) types.String {
		if s == nil {
			return types.StringNull()
		}
		return types.StringValue(*s)
	}

	derefInt := func(i *int64) types.Int64 {
		if i == nil {
			return types.Int64Null()
		}
		return types.Int64Value(int64(*i))
	}

	derefBool := func(b *bool) types.Bool {
		if b == nil {
			return types.BoolNull()
		}
		return types.BoolValue(*b)
	}

	state.Zone = types.StringValue(zone.Zone)
	state.Type = types.StringValue(zone.Type)
	state.DHCP = derefString(zone.DHCP)
	state.DNS = derefString(zone.DNS)
	state.DNSZone = derefString(zone.DNSZone)
	state.Digest = derefString(zone.Digest)
	state.IPAM = derefString(zone.IPAM)
	state.MTU = derefInt(zone.MTU)
	state.Nodes = derefString(zone.Nodes)
	state.Pending = derefBool(zone.Pending)
	state.ReverseDNS = derefString(zone.ReverseDNS)
	state.State = derefString(zone.State)
	state.AdvertiseSubnets = derefBool(zone.AdvertiseSubnets)
	state.Bridge = derefString(zone.Bridge)
	state.BridgeDisableMACLearning = derefBool(zone.BridgeDisableMACLearning)
	state.Controller = derefString(zone.Controller)
	state.DisableARPDiscovery = derefBool(zone.DisableARPDiscovery)
	state.DPID = derefInt(zone.DPID)
	state.ExitNodes = derefString(zone.ExitNodes)
	state.ExitNodesLocalRouting = derefBool(zone.ExitNodesLocalRouting)
	state.MAC = derefString(zone.MAC)
	state.Peers = derefString(zone.Peers)
	state.RouteTargetImport = derefString(zone.RouteTargetImport)
	state.VLANProtocol = derefString(zone.VLANProtocol)
	state.VRFVXLAN = derefInt(zone.VRFVXLAN)
	state.VXLANPort = derefInt(zone.VXLANPort)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *proxmoxSDNZoneResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan proxmoxSDNZoneResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	zone := client2.SDNZone{
		Zone: plan.Zone.ValueString(),
		Type: plan.Type.ValueString(),
	}

	setStringIfNotNull(&zone.DHCP, plan.DHCP)
	setStringIfNotNull(&zone.DNS, plan.DNS)
	setStringIfNotNull(&zone.DNSZone, plan.DNSZone)
	setStringIfNotNull(&zone.Digest, plan.Digest)
	setStringIfNotNull(&zone.IPAM, plan.IPAM)
	setIntIfNotNull(&zone.MTU, plan.MTU)
	setStringIfNotNull(&zone.Nodes, plan.Nodes)
	setBoolIfNotNull(&zone.Pending, plan.Pending)
	setStringIfNotNull(&zone.ReverseDNS, plan.ReverseDNS)
	setStringIfNotNull(&zone.State, plan.State)
	setBoolIfNotNull(&zone.AdvertiseSubnets, plan.AdvertiseSubnets)
	setStringIfNotNull(&zone.Bridge, plan.Bridge)
	setBoolIfNotNull(&zone.BridgeDisableMACLearning, plan.BridgeDisableMACLearning)
	setStringIfNotNull(&zone.Controller, plan.Controller)
	setBoolIfNotNull(&zone.DisableARPDiscovery, plan.DisableARPDiscovery)
	setIntIfNotNull(&zone.DPID, plan.DPID)
	setStringIfNotNull(&zone.ExitNodes, plan.ExitNodes)
	setBoolIfNotNull(&zone.ExitNodesLocalRouting, plan.ExitNodesLocalRouting)
	setStringIfNotNull(&zone.MAC, plan.MAC)
	setStringIfNotNull(&zone.Peers, plan.Peers)
	setStringIfNotNull(&zone.RouteTargetImport, plan.RouteTargetImport)
	// if !plan.Tag.IsNull() {
	// 	tag := int(plan.Tag.ValueInt64())
	// 	zone.Tag = &tag
	// }
	setStringIfNotNull(&zone.VLANProtocol, plan.VLANProtocol)
	setIntIfNotNull(&zone.VRFVXLAN, plan.VRFVXLAN)
	setIntIfNotNull(&zone.VXLANPort, plan.VXLANPort)

	err := r.client.UpdateSDNZone(zone)
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

	// helper functions
	derefString := func(s *string) types.String {
		if s == nil {
			return types.StringNull()
		}
		return types.StringValue(*s)
	}

	derefInt := func(i *int64) types.Int64 {
		if i == nil {
			return types.Int64Null()
		}
		return types.Int64Value(int64(*i))
	}

	derefBool := func(b *bool) types.Bool {
		if b == nil {
			return types.BoolNull()
		}
		return types.BoolValue(*b)
	}

	// planにすべての属性を設定
	plan.Zone = types.StringValue(updatedZone.Zone)
	plan.Type = types.StringValue(updatedZone.Type)
	plan.DHCP = derefString(updatedZone.DHCP)
	plan.DNS = derefString(updatedZone.DNS)
	plan.DNSZone = derefString(updatedZone.DNSZone)
	plan.Digest = derefString(updatedZone.Digest)
	plan.IPAM = derefString(updatedZone.IPAM)
	plan.MTU = derefInt(updatedZone.MTU)
	plan.Nodes = derefString(updatedZone.Nodes)
	plan.Pending = derefBool(updatedZone.Pending)
	plan.ReverseDNS = derefString(updatedZone.ReverseDNS)
	plan.State = derefString(updatedZone.State)
	plan.AdvertiseSubnets = derefBool(updatedZone.AdvertiseSubnets)
	plan.Bridge = derefString(updatedZone.Bridge)
	plan.BridgeDisableMACLearning = derefBool(updatedZone.BridgeDisableMACLearning)
	plan.Controller = derefString(updatedZone.Controller)
	plan.DisableARPDiscovery = derefBool(updatedZone.DisableARPDiscovery)
	plan.DPID = derefInt(updatedZone.DPID)
	plan.ExitNodes = derefString(updatedZone.ExitNodes)
	plan.ExitNodesLocalRouting = derefBool(updatedZone.ExitNodesLocalRouting)
	plan.MAC = derefString(updatedZone.MAC)
	plan.Peers = derefString(updatedZone.Peers)
	plan.RouteTargetImport = derefString(updatedZone.RouteTargetImport)
	plan.VLANProtocol = derefString(updatedZone.VLANProtocol)
	plan.VRFVXLAN = derefInt(updatedZone.VRFVXLAN)
	plan.VXLANPort = derefInt(updatedZone.VXLANPort)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *proxmoxSDNZoneResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state proxmoxSDNZoneResourceModel
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
