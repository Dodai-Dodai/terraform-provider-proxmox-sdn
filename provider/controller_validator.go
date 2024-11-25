// Controller_validator.go
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.String = ControllerValidator{}

type ControllerValidator struct {
	TypeAttributeName string
}

func (v ControllerValidator) Description(_ context.Context) string {
	return fmt.Sprintf("Requires 'Controller' to be set when '%s' is 'vlan' or 'qinq', otherwise 'Controller' must be null", v.TypeAttributeName)
}

func (v ControllerValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v ControllerValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
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
		// 値が設定されなければならない
		if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Missing Required Attribute",
				fmt.Sprintf("'Controller' must be set when '%s' is 'evpn'", v.TypeAttributeName),
			)
			return
		}
	} else {
		// 値が設定されていてはいけない
		if !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid Attribute",
				fmt.Sprintf("'Controller' must be null when '%s' is not 'evpn'", v.TypeAttributeName),
			)
			return
		}
	}
}
