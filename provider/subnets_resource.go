package provider

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Dodai-Dodai/terraform-provider-proxmox-sdn/client"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var (
	_ resource.Resource              = &proxmoxSDNSubnetResource{}
	_ resource.ResourceWithConfigure = &proxmoxSDNSubnetResource{}
)

func NewProxmoxSDNSubnetsResource() resource.Resource {
	return &proxmoxSDNSubnetResource{}
}

type proxmoxSDNSubnetResource struct {
	client *client.SSHProxmoxClient
}

func (r *proxmoxSDNSubnetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *proxmoxSDNSubnetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_subnets"
}

func (r *proxmoxSDNSubnetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"subnet": schema.StringAttribute{
				Description: "subnet name",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"cidr": schema.StringAttribute{
				Description: "cidr",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				// subnetと同じ値でなければならない制約

			},
			"type": schema.StringAttribute{
				Description: "subnet type",
				Computed:    true,
			},
			"vnet": schema.StringAttribute{
				Description: "vnet name",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"dhcp_dns_server": schema.StringAttribute{
				Description: "dhcp dns server",
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"dhcp_range": schema.ListNestedAttribute{
				Description: "dhcp range",
				Optional:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"start_address": schema.StringAttribute{
							Description: "start address",
							Required:    true,
						},
						"end_address": schema.StringAttribute{
							Description: "end address",
							Required:    true,
						},
					},
				},
			},
			"dns_zone_prefix": schema.StringAttribute{
				Description: "dns zone prefix",
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"gateway": schema.StringAttribute{
				Description: "gateway",
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"snat": schema.BoolAttribute{
				Description: "snat",
				Optional:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
			"zone": schema.StringAttribute{
				Description: "zone name",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *proxmoxSDNSubnetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state subnetsModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	subnet, err := r.client.GetSubnet(state.Vnet.ValueString(), state.Zone.ValueString(), state.Subnet.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading subnet", err.Error())
		return
	}

	updatedState, diags := convertSDNSubnetstoSubnetsModel(ctx, *subnet)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	diags = resp.State.Set(ctx, updatedState)
	resp.Diagnostics.Append(diags...)
}

func convertSubnetsModeltoClientSubnet(ctx context.Context, model subnetsModel) (*client.SDNSubnets, diag.Diagnostics) {
	var diags diag.Diagnostics
	var subnet client.SDNSubnets

	subnet.Subnet = model.Subnet.ValueString()
	subnet.Cidr = model.Cidr.ValueString()
	subnet.Vnet = model.Vnet.ValueString()
	subnet.Type = model.Type.ValueString()

	if !model.DhcpDnsServer.IsNull() && !model.DhcpDnsServer.IsUnknown() {
		dhcpDnsServer := model.DhcpDnsServer.ValueString()
		subnet.DhcpDnsServer = &dhcpDnsServer
	}

	//DhcpRange
	if len(model.DhcpRange) > 0 {
		var dhcpRanges []client.DhcpRange
		for _, dhcpRangeModel := range model.DhcpRange {
			dhcpRange := client.DhcpRange{
				StartAddress: dhcpRangeModel.StartAddress.ValueString(),
				EndAddress:   dhcpRangeModel.EndAddress.ValueString(),
			}
			dhcpRanges = append(dhcpRanges, dhcpRange)
		}
		subnet.DhcpRange = dhcpRanges
	}

	if !model.DnsZonePrefix.IsNull() && !model.DnsZonePrefix.IsUnknown() {
		dnsZonePrefix := model.DnsZonePrefix.ValueString()
		subnet.DnsZonePrefix = &dnsZonePrefix
	}

	if !model.Gateway.IsNull() && !model.Gateway.IsUnknown() {
		gateway := model.Gateway.ValueString()
		subnet.Gateway = &gateway
	}

	// Snat intboolの処理
	subnet.Snat = BoolToIntBoolPointer(model.Snat)

	if !model.Zone.IsNull() && !model.Zone.IsUnknown() {
		zone := model.Zone.ValueString()
		subnet.Zone = &zone
	}
	return &subnet, diags
}

func (r *proxmoxSDNSubnetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan subnetsModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	subnet, diags := convertSubnetsModeltoClientSubnet(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	err := r.client.CreateSubnet(*subnet)
	if err != nil {
		resp.Diagnostics.AddError("CreateSubnet failed", err.Error())
		return
	}

	// 20秒待つ
	// これはProxmoxのAPIの仕様で、サブネットの作成後にすぐに取得すると、作成されたサブネットが取得できないことがあるため
	// この問題はProxmoxのAPIの問題であり、Terraformの問題ではない
	time.Sleep(20 * time.Second)

	// VNet内のすべてのサブネットを取得
	subnets, err := r.client.GetSubnets(subnet.Vnet)
	if err != nil {
		resp.Diagnostics.AddError("Failed to get subnets", err.Error())
		return
	}

	// subnetsをデバッグ用に出力
	for _, s := range subnets {
		fmt.Printf("Subnet: %v", s)
		// cidrをデバッグ用に出力
		fmt.Printf("Cidr: %v\n", s.Cidr)
	}

	// CIDR表記で一致するサブネットを探す
	var createdSubnet *client.SDNSubnets
	for _, s := range subnets {
		if s.Cidr == subnet.Cidr {
			createdSubnet = &s
			break
		}
	}

	if createdSubnet == nil {
		resp.Diagnostics.AddError("Failed to find created subnet", "Created subnet not found in subnets")
		return
	}

	// 状態を設定
	state, diags := convertSDNSubnetstoSubnetsModel(ctx, *createdSubnet)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *proxmoxSDNSubnetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan subnetsModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	subnet, diags := convertSubnetsModeltoClientSubnet(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	err := r.client.UpdateSubnet(*subnet)
	if err != nil {
		resp.Diagnostics.AddError(
			"UpdateSubnet failed",
			err.Error(),
		)
		return
	}

	state, diags := convertSDNSubnetstoSubnetsModel(ctx, *subnet)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *proxmoxSDNSubnetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state subnetsModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	// スラッシュをハイフンに置換
	subnetID := strings.ReplaceAll(state.Subnet.ValueString(), "/", "-")

	err := r.client.DeleteSubnet(state.Vnet.ValueString(), state.Zone.ValueString(), subnetID)
	if err != nil {
		resp.Diagnostics.AddError("DeleteSubnet failed", err.Error())
		return
	}
}
