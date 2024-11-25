// PrimaryExitNode_validator.go
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.String = PrimaryExitNodeValidator{}

type PrimaryExitNodeValidator struct {
	TypeAttributeName string
}

func (v PrimaryExitNodeValidator) Description(_ context.Context) string {
	return fmt.Sprintf("'PrimaryExitNode''%s' is 'evpn'; otherwise, 'PrimaryExitNode' must be null", v.TypeAttributeName)
}

func (v PrimaryExitNodeValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v PrimaryExitNodeValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
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
		// 値は設定されていなくてよい
		if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
			return
		}
	} else {
		// evpnでない場合、PrimaryExitNodeは設定されていてはいけない
		if !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid Attribute",
				fmt.Sprintf("'PrimaryExitNode' must be null when '%s' is not 'evpn'", v.TypeAttributeName),
			)
			return
		}
	}
}
