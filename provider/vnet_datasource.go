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
	_ datasource.DataSource              = &proxmoxSDNVnetDatasource{}
	_ datasource.DataSourceWithConfigure = &proxmoxSDNVnetDatasource{}
)

func NewProxmoxSDNVnetDatasource() datasource.DataSource {
	return &proxmoxSDNVnetDatasource{}
}

type proxmoxSDNVnetDatasource struct {
	client *client.SSHProxmoxClient
}

func (d *proxmoxSDNVnetDatasource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vnet"
}

func (d *proxmoxSDNVnetDatasource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"vnets": schema.ListNestedAttribute{
				Description: "List of vnet",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vnet": schema.StringAttribute{
							Description: "vnet name",
							Required:    true,
						},
						"zone": schema.StringAttribute{
							Description: "zone name",
							Required:    true,
						},
						"type": schema.StringAttribute{
							Description: "vnet type",
							Optional:    true,
							Computed:    true,
						},
						"tag": schema.Int64Attribute{
							Description: "vnet tag",
							Optional:    true,
						},
						"vlanaware": schema.BoolAttribute{
							Description: "vnet vlanaware",
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func (d *proxmoxSDNVnetDatasource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func convertSDNVnettoVnetsModel(ctx context.Context, vnet client.SDNVnet) (vnetsModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// boolとIntBoolの変換
	vlanaware := IntBoolPointerToBoolPointer(vnet.Vlanaware)

	vnetsModel := vnetsModel{
		Vnet:      types.StringValue(vnet.Vnet),
		Zone:      types.StringValue(vnet.Zone),
		Type:      types.StringValue(vnet.Type),
		Tag:       types.Int64PointerValue(vnet.Tag),
		Vlanaware: types.BoolPointerValue(vlanaware),
	}

	return vnetsModel, diags
}

func (d *proxmoxSDNVnetDatasource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state struct {
		Vnets []vnetsModel `tfsdk:"vnets"`
	}

	vnets, err := d.client.GetVnets()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get vnets", err.Error())
		return
	}

	var diags diag.Diagnostics

	for _, vnet := range vnets {
		vnetModel, d := convertSDNVnettoVnetsModel(ctx, vnet)
		if d.HasError() {
			resp.Diagnostics.Append(d...)
			continue
		}
		diags.Append(d...)
		state.Vnets = append(state.Vnets, vnetModel)
	}

	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	diags.Append(resp.State.Set(ctx, state)...)
	resp.Diagnostics.Append(diags...)
}
