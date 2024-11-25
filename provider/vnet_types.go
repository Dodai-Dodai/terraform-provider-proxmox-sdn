package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type vnetsModel struct {
	Vnet      types.String `tfsdk:"vnet"`      //required
	Zone      types.String `tfsdk:"zone"`      //required
	Type      types.String `tfsdk:"type"`      //optional but always set 'vnet'
	Tag       types.Int64  `tfsdk:"tag"`       //optional
	Vlanaware types.Bool   `tfsdk:"vlanaware"` //optional
}
