package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type zonesModel struct {
	Zone                    types.String `tfsdk:"zone"`                    //required
	Type                    types.String `tfsdk:"type"`                    //required
	MTU                     types.Int64  `tfsdk:"mtu"`                     //optional
	Nodes                   types.Set    `tfsdk:"nodes"`                   //optional
	IPAM                    types.String `tfsdk:"ipam"`                    //optional
	DNS                     types.String `tfsdk:"dns"`                     //optional
	ReverseDNS              types.String `tfsdk:"reversedns"`              //optional
	DNSZone                 types.String `tfsdk:"dnszone"`                 //optional
	Bridge                  types.String `tfsdk:"bridge"`                  //optional
	Tag                     types.Int64  `tfsdk:"tag"`                     //optional
	VLANProtocol            types.String `tfsdk:"vlanprotocol"`            //optional
	Peers                   types.Set    `tfsdk:"peers"`                   //optional
	Controller              types.String `tfsdk:"controller"`              //optional
	VRFVXLAN                types.Int64  `tfsdk:"vrf_vxlan"`               //optional
	MAC                     types.String `tfsdk:"mac"`                     //optional
	ExitNodes               types.Set    `tfsdk:"exitnodes"`               //optional
	PrimaryExitNode         types.String `tfsdk:"primaryexitnode"`         //optional
	ExitNodesLocalRouting   types.Bool   `tfsdk:"exitnodeslocalrouting"`   //optional
	AdvertiseSubnets        types.Bool   `tfsdk:"advertisesubnets"`        //optional
	DisableARPNdSuppression types.Bool   `tfsdk:"disablearpndsuppression"` //optional
	RouteTargetImport       types.String `tfsdk:"rtimport"`                //optional
}
