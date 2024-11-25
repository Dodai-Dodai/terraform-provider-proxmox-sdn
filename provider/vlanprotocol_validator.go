// vlanprotocol_validator.go
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.String = VLANProtocolValidator{}

type VLANProtocolValidator struct {
	TypeAttributeName string
}

func (v VLANProtocolValidator) Description(_ context.Context) string {
	return fmt.Sprintf("'vlanprotocol' must be '802.1ad' or '802.1q' when '%s' is 'qinq'; otherwise, 'vlanprotocol' must be null", v.TypeAttributeName)
}

func (v VLANProtocolValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v VLANProtocolValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
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

	// この値のみ許可される
	allowedValues := []string{"802.1ad", "802.1q"}

	if typeValue == "qinq" {
		// この値は設定されてなくてもいい
		if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
			return
		}

		vlanProtocolValue := req.ConfigValue.ValueString()

		// 値のチェック
		isValid := false
		for _, val := range allowedValues {
			if vlanProtocolValue == val {
				isValid = true
				break
			}
		}

		if !isValid {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid 'vlanprotocol' Value",
				fmt.Sprintf("'vlanprotocol' must be one of %v when '%s' is 'qinq', got: %s", allowedValues, v.TypeAttributeName, vlanProtocolValue),
			)
			return
		}
	} else {
		// qinqでない場合、vlanprotocolは設定されていてはいけない
		if !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid Attribute",
				fmt.Sprintf("'vlanprotocol' must be null when '%s' is not 'qinq'", v.TypeAttributeName),
			)
			return
		}
	}
}
