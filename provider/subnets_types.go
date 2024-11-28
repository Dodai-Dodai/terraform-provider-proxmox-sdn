package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type dhcpRangeModel struct {
	StartAddress types.String `tfsdk:"start_address"`
	EndAddress   types.String `tfsdk:"end_address"`
}

type subnetsModel struct {
	Subnet        types.String     `tfsdk:"subnet"`          // required field
	Type          types.String     `tfsdk:"type"`            // required but always set 'subnet'
	Vnet          types.String     `tfsdk:"vnet"`            // required field
	DhcpDnsServer types.String     `tfsdk:"dhcp_dns_server"` // optional
	DhcpRange     []dhcpRangeModel `tfsdk:"dhcp_range"`      // optional
	DnsZonePrefix types.String     `tfsdk:"dns_zone_prefix"` // optional
	Gateway       types.String     `tfsdk:"gateway"`         // optional
	Snat          types.Bool       `tfsdk:"snat"`            // optional
	Zone          types.String     `tfsdk:"zone"`            // optional
}
