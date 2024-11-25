// bridge_validator.go
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.String = BridgeValidator{}

type BridgeValidator struct {
	TypeAttributeName string
}

func (v BridgeValidator) Description(_ context.Context) string {
	return fmt.Sprintf("Requires 'bridge' to be set when '%s' is 'vlan' or 'qinq', otherwise 'bridge' must be null", v.TypeAttributeName)
}

func (v BridgeValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v BridgeValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
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

	if typeValue == "vlan" || typeValue == "qinq" {
		// 値が設定されなければならない
		if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Missing Required Attribute",
				fmt.Sprintf("'bridge' must be set when '%s' is 'vlan' or 'qinq'", v.TypeAttributeName),
			)
			return
		}
	} else {
		// 値が設定されていてはいけない
		if !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid Attribute",
				fmt.Sprintf("'bridge' must be null when '%s' is not 'vlan' or 'qinq'", v.TypeAttributeName),
			)
			return
		}
	}
}
