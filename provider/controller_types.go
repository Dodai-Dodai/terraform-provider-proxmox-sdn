package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type controllerModel struct {
	Controller types.String `tfsdk:"controller"`                 //required
	Type       types.String `tfsdk:"type"`                       //required
	ASN        types.Int64  `tfsdk:"asn"`                        //optional
	BGPMAR     types.Bool   `tfsdk:"bgp_multipath_aspath_relax"` //optional
	EBGP       types.Bool   `tfsdk:"ebgp"`                       //optional
	EBGPMH     types.Int64  `tfsdk:"ebgp_multihop"`              //optional
	ISISDomain types.String `tfsdk:"isis_domain"`                //optional
	ISISIfaces types.String `tfsdk:"isis_ifaces"`                //optional
	ISISNet    types.String `tfsdk:"isis_net"`                   //optional
	Loopback   types.String `tfsdk:"loopback"`                   //optional
	Node       types.String `tfsdk:"node"`                       //optional
	Peers      types.Set    `tfsdk:"peers"`                      //optional
}
