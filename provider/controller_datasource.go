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
	_ datasource.DataSource              = &proxmoxSDNControllerDatasource{}
	_ datasource.DataSourceWithConfigure = &proxmoxSDNControllerDatasource{}
)

func NewProxmoxSDNControllerDatasource() datasource.DataSource {
	return &proxmoxSDNControllerDatasource{}
}

type proxmoxSDNControllerDatasource struct {
	client *client.SSHProxmoxClient
}

func (d *proxmoxSDNControllerDatasource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_controller"
}

// Schema
func (d *proxmoxSDNControllerDatasource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"controllers": schema.ListNestedAttribute{
				Description: "List of controller",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"controller": schema.StringAttribute{
							Description: "controller name",
							Required:    true,
						},
						"type": schema.StringAttribute{
							Description: "controller type",
							Required:    true,
						},
						"asn": schema.Int64Attribute{
							Description: "controller asn",
							Optional:    true,
						},
						"bgp_multipath_aspath_relax": schema.BoolAttribute{
							Description: "controller bgp-multipath-aspath-relax",
							Optional:    true,
						},
						"ebgp": schema.BoolAttribute{
							Description: "controller ebgp",
							Optional:    true,
						},
						"ebgp_multihop": schema.Int64Attribute{
							Description: "controller ebgp-multihop",
							Optional:    true,
						},
						"isis_domain": schema.StringAttribute{
							Description: "controller isis-domain",
							Optional:    true,
						},
						"isis_ifaces": schema.StringAttribute{
							Description: "controller isis-ifaces",
							Optional:    true,
						},
						"isis_net": schema.StringAttribute{
							Description: "controller isis-net",
							Optional:    true,
						},
						"loopback": schema.StringAttribute{
							Description: "controller loopback",
							Optional:    true,
						},
						"node": schema.StringAttribute{
							Description: "controller node",
							Optional:    true,
						},
						"peers": schema.SetAttribute{
							Description: "controller peers",
							Optional:    true,
							ElementType: types.StringType,
						},
					},
				},
			},
		},
	}
}

func (d *proxmoxSDNControllerDatasource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func convertSDNControllertoControllersModel(ctx context.Context, controller client.SDNController) (controllerModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// set型に変換
	peers, diagPerrs := types.SetValueFrom(ctx, types.StringType, controller.Peers)
	if diagPerrs.HasError() {
		diags = append(diags, diagPerrs...)
		return controllerModel{}, diags
	}
	diags.Append(diagPerrs...)

	// bool型に変換
	bgpMAR := IntBoolPointerToBoolPointer(controller.BgpMultipathAsPathRelax)
	ebgp := IntBoolPointerToBoolPointer(controller.Ebgp)

	controllersModel := controllerModel{
		Controller: types.StringValue(controller.Controller),
		Type:       types.StringValue(controller.Type),
		ASN:        types.Int64PointerValue(controller.ASN),
		BGPMAR:     types.BoolPointerValue(bgpMAR),
		EBGP:       types.BoolPointerValue(ebgp),
		EBGPMH:     types.Int64PointerValue(controller.EbgpMultihop),
		ISISDomain: types.StringPointerValue(controller.ISISDomain),
		ISISIfaces: types.StringPointerValue(controller.ISISIfaces),
		ISISNet:    types.StringPointerValue(controller.ISISNet),
		Loopback:   types.StringPointerValue(controller.Loopback),
		Node:       types.StringPointerValue(controller.Node),
		Peers:      peers,
	}

	return controllersModel, diags
}

// Read
func (d *proxmoxSDNControllerDatasource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state struct {
		Controllers []controllerModel `tfsdk:"controllers"`
	}

	controllers, err := d.client.GetSDNControllers()
	if err != nil {
		resp.Diagnostics.AddError("Failed to get controllers", err.Error())
		return
	}

	var diags diag.Diagnostics

	for _, controller := range controllers {
		controllerModel, d := convertSDNControllertoControllersModel(ctx, controller)
		if d.HasError() {
			resp.Diagnostics.Append(d...)
			continue
		}
		diags.Append(d...)
		state.Controllers = append(state.Controllers, controllerModel)
	}

	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	diags.Append(resp.State.Set(ctx, state)...)
	resp.Diagnostics.Append(diags...)
}
