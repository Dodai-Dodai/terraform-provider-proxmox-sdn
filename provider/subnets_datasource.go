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
	_ datasource.DataSource              = &proxmoxSDNSubnetDatasource{}
	_ datasource.DataSourceWithConfigure = &proxmoxSDNSubnetDatasource{}
)

func NewProxmoxSDNSubnetsDatasource() datasource.DataSource {
	return &proxmoxSDNSubnetDatasource{}
}

type proxmoxSDNSubnetDatasource struct {
	client *client.SSHProxmoxClient
}

func (d *proxmoxSDNSubnetDatasource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_subnets"
}

func (d *proxmoxSDNSubnetDatasource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"subnets": schema.ListNestedAttribute{
				Description: "List of subnets",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"subnet": schema.StringAttribute{
							Description: "subnet name",
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: "subnet type",
							Optional:    true,
							Computed:    true,
						},
						"vnet": schema.StringAttribute{
							Description: "vnet name",
							Optional:    true,
						},
						"dhcp_dns_server": schema.StringAttribute{
							Description: "dhcp dns server",
							Optional:    true,
						},
						"dhcp_range": schema.ListNestedAttribute{
							Description: "dhcp range",
							Optional:    true,
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
						},
						"gateway": schema.StringAttribute{
							Description: "gateway",
							Optional:    true,
						},
						"snat": schema.BoolAttribute{
							Description: "snat",
							Optional:    true,
						},
						"zone": schema.StringAttribute{
							Description: "zone",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *proxmoxSDNSubnetDatasource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.SSHProxmoxClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.SSHProxmoxClient, got %T", req.ProviderData),
		)
	}
	d.client = client
}

func convertSDNSubnetstoSubnetsModel(ctx context.Context, subnet client.SDNSubnets) (subnetsModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// DhcpRangeの変換
	if len(subnet.DhcpRange) > 0 {
		var dhcpRanges []client.DhcpRange
		for _, dr := range subnet.DhcpRange {
			dhcpRange := client.DhcpRange{
				StartAddress: dr.StartAddress,
				EndAddress:   dr.EndAddress,
			}
			dhcpRanges = append(dhcpRanges, dhcpRange)
		}
		subnet.DhcpRange = dhcpRanges
	}

	// boolとIntBoolの変換
	snat := IntBoolPointerToBoolPointer(subnet.Snat)

	// Convert []client.DhcpRange to []dhcpRangeModel
	var dhcpRanges []dhcpRangeModel
	for _, dr := range subnet.DhcpRange {
		dhcpRange := dhcpRangeModel{
			StartAddress: types.StringValue(dr.StartAddress),
			EndAddress:   types.StringValue(dr.EndAddress),
		}
		dhcpRanges = append(dhcpRanges, dhcpRange)
	}

	// SDNSubnetsをSubnetsModelに変換
	subnets := subnetsModel{
		Subnet:        types.StringValue(subnet.Subnet),
		Type:          types.StringValue(subnet.Type),
		Vnet:          types.StringValue(subnet.Vnet),
		DhcpDnsServer: types.StringPointerValue(subnet.DhcpDnsServer),
		DhcpRange:     dhcpRanges,
		DnsZonePrefix: types.StringPointerValue(subnet.DnsZonePrefix),
		Gateway:       types.StringPointerValue(subnet.Gateway),
		Snat:          types.BoolPointerValue(snat),
		Zone:          types.StringPointerValue(subnet.Zone),
	}
	return subnets, diags
}

func (d *proxmoxSDNSubnetDatasource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// vnetの一覧を取得 -> vnetごとにsubnetを取得の流れ
	type sdnSubnetsDataSourceModel struct {
		Subnets []subnetsModel `tfsdk:"subnets"`
	}
	var state sdnSubnetsDataSourceModel
	// vnetsを取得
	vnets, err := d.client.GetVnets()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get vnets", err.Error())
		return
	}

	var diags diag.Diagnostics
	var allSubnets []subnetsModel

	// vnetごとにsubnetを取得
	for _, vnet := range vnets {
		subnets, err := d.client.GetSubnets(vnet.Vnet)
		if err != nil {
			resp.Diagnostics.AddError("Failed to get subnets", err.Error())
			return
		}

		for _, subnet := range subnets {
			subnetsModel, diags := convertSDNSubnetstoSubnetsModel(ctx, subnet)
			if diags.HasError() {
				resp.Diagnostics.Append(diags...)
				return
			}
			allSubnets = append(allSubnets, subnetsModel)
		}
		state.Subnets = allSubnets
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
