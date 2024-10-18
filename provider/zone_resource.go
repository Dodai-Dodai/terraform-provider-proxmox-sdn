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
				Description: "The advertise subnets status of the zone",
				Computed:    true,
			},
			"bridge": schema.StringAttribute{
				Description: "The bridge of the zone",
				Computed:    true,
			},
			"bridge_disable_mac_learning": schema.BoolAttribute{
				Description: "The bridge disable MAC learning status of the zone",
				Computed:    true,
			},
			"controller": schema.StringAttribute{
				Description: "The controller of the zone",
				Computed:    true,
			},
			"disable_arp_discovery": schema.BoolAttribute{
				Description: "The disable ARP discovery status of the zone",
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
				Description: "The exit nodes local routing status of the zone",
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
func setStringIfNotNull(dest *string, src types.String) {
	if !src.IsNull() {
		*dest = src.ValueString()
	}
}

// Int64型のマッピング
func setIntIfNotNull(dest *int, src types.Int64) {
	if !src.IsNull() {
		*dest = int(src.ValueInt64())
	}
}

// Bool型のマッピング
func setBoolIfNotNull(dest *bool, src types.Bool) {
	if !src.IsNull() {
		*dest = src.ValueBool()
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

	// 全ての属性をplanに設定
	plan.Zone = types.StringValue(createdZone.Zone)
	plan.Type = types.StringValue(createdZone.Type)
	plan.DHCP = types.StringValue(createdZone.DHCP)
	plan.DNS = types.StringValue(createdZone.DNS)
	plan.DNSZone = types.StringValue(createdZone.DNSZone)
	plan.Digest = types.StringValue(createdZone.Digest)
	plan.IPAM = types.StringValue(createdZone.IPAM)
	plan.MTU = types.Int64Value(int64(createdZone.MTU))
	plan.Nodes = types.StringValue(createdZone.Nodes)
	plan.Pending = types.BoolValue(createdZone.Pending)
	plan.ReverseDNS = types.StringValue(createdZone.ReverseDNS)
	plan.State = types.StringValue(createdZone.State)
	plan.AdvertiseSubnets = types.BoolValue(createdZone.AdvertiseSubnets)
	plan.Bridge = types.StringValue(createdZone.Bridge)
	plan.BridgeDisableMACLearning = types.BoolValue(createdZone.BridgeDisableMACLearning)
	plan.Controller = types.StringValue(createdZone.Controller)
	plan.DisableARPDiscovery = types.BoolValue(createdZone.DisableARPDiscovery)
	plan.DPID = types.Int64Value(int64(createdZone.DPID))
	plan.ExitNodes = types.StringValue(createdZone.ExitNodes)
	plan.ExitNodesLocalRouting = types.BoolValue(createdZone.ExitNodesLocalRouting)
	plan.MAC = types.StringValue(createdZone.MAC)
	plan.Peers = types.StringValue(createdZone.Peers)
	plan.RouteTargetImport = types.StringValue(createdZone.RouteTargetImport)
	// if createdZone.Tag != nil {
	// 	plan.Tag = types.Int64Value(int64(*createdZone.Tag))
	// }
	plan.VLANProtocol = types.StringValue(createdZone.VLANProtocol)
	plan.VRFVXLAN = types.Int64Value(int64(createdZone.VRFVXLAN))
	plan.VXLANPort = types.Int64Value(int64(createdZone.VXLANPort))

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

	state.Digest = types.StringValue(zone.Digest)
	state.Pending = types.BoolValue(zone.Pending)
	state.State = types.StringValue(zone.State)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *proxmoxSDNZoneResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *proxmoxSDNZoneResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
