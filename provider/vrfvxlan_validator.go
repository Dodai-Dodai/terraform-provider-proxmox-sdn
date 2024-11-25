// vrfvxlan_validator.go
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.Int64 = VrfvxlanValidator{}

type VrfvxlanValidator struct {
	TypeAttributeName string
}

func (v VrfvxlanValidator) Description(_ context.Context) string {
	return fmt.Sprintf("Requires 'vrfvxlan' to be set when '%s' is 'evpn', otherwise 'vrfvxlan' must be null", v.TypeAttributeName)
}

func (v VrfvxlanValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v VrfvxlanValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	typeAttrPath := path.Root(v.TypeAttributeName)
	var typeAttr types.String

	diags := req.Config.GetAttribute(ctx, typeAttrPath, &typeAttr)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	if typeAttr.IsUnknown() || typeAttr.IsNull() {
		return
	}

	typeValue := typeAttr.ValueString()

	if typeValue == "evpn" {
		// 値は設定されている必要がある
		if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Missing Required Attribute",
				fmt.Sprintf("'vrfvxlan' must be set when '%s' is 'evpn'", v.TypeAttributeName),
			)
			return
		}
	} else {
		// evpnでない場合、vrfvxlanは設定されていてはいけない
		if !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid Attribute",
				fmt.Sprintf("'vrfvxlan' must be null when '%s' is not 'evpn'", v.TypeAttributeName),
			)
			return
		}
	}
}
