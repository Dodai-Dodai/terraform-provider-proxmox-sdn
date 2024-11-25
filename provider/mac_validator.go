// MAC_validator.go
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.String = MACValidator{}

type MACValidator struct {
	TypeAttributeName string
}

func (v MACValidator) Description(_ context.Context) string {
	return fmt.Sprintf("'MAC''%s' is 'evpn'; otherwise, 'MAC' must be null", v.TypeAttributeName)
}

func (v MACValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v MACValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
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
		// evpnの場合、値は設定されていなくてもいい
		if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
			return
		}
	} else {
		// evpnでない場合、MACは設定されていてはいけない
		if !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid Attribute",
				fmt.Sprintf("'MAC' must be null when '%s' is not 'evpn'", v.TypeAttributeName),
			)
			return
		}
	}
}
