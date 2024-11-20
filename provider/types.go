package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type zonesModel struct {
	Zone       types.String `tfsdk:"zone"`       //required
	Type       types.String `tfsdk:"type"`       //required
	MTU        types.Int64  `tfsdk:"mtu"`        //optional
	Nodes      types.Set    `tfsdk:"nodes"`      //optional
	IPAM       types.String `tfsdk:"ipam"`       //optional
	DNS        types.String `tfsdk:"dns"`        //optional
	ReverseDNS types.String `tfsdk:"reversedns"` //optional
	DNSZone    types.String `tfsdk:"dnszone"`    //optional
	VLAN       types.Object `tfsdk:"vlan"`       //optional
	QinQ       types.Object `tfsdk:"qinq"`       //optional
	VXLAN      types.Object `tfsdk:"vxlan"`      //optional
	EVPN       types.Object `tfsdk:"evpn"`       //optional
}

type VLANConfigModel struct {
	Bridge types.String `tfsdk:"bridge"` //required
}

type QinQConfigModel struct {
	Bridge       types.String `tfsdk:"bridge"`       //required
	Tag          types.Int64  `tfsdk:"tag"`          //required
	VLANProtocol types.String `tfsdk:"vlanprotocol"` //optional
}

type VXLANConfigModel struct {
	Peer types.Set `tfsdk:"peer"` //required
}

type EVPNConfigModel struct {
	Controller              types.String `tfsdk:"controller"`              //required
	VRFVXLAN                types.Int64  `tfsdk:"vrf_vxlan"`               //required
	MAC                     types.String `tfsdk:"mac"`                     //optional
	ExitNodes               types.Set    `tfsdk:"exitnodes"`               //optional
	PrimaryExitNode         types.String `tfsdk:"primaryexitnode"`         //optional
	ExitNodesLocalRouting   types.Bool   `tfsdk:"exitnodeslocalrouting"`   //optional
	AdvertiseSubnets        types.Bool   `tfsdk:"advertisesubnets"`        //optional
	DisableARPNdSuppression types.Bool   `tfsdk:"disablearpndsuppression"` //optional
	RouteTargetImport       types.String `tfsdk:"rtimport"`                //optional
}
