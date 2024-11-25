// RouteTargetImport_validator.go
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.String = RouteTargetImportValidator{}

type RouteTargetImportValidator struct {
	TypeAttributeName string
}

func (v RouteTargetImportValidator) Description(_ context.Context) string {
	return fmt.Sprintf("'RouteTargetImport''%s' is 'evpn'; otherwise, 'RouteTargetImport' must be null", v.TypeAttributeName)
}

func (v RouteTargetImportValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v RouteTargetImportValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
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
		// この値は、書かなくてもいい
		if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
			return
		}
	} else {
		// evpnでない場合、RouteTargetImportは設定されていてはいけない
		if !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid Attribute",
				fmt.Sprintf("'RouteTargetImport' must be null when '%s' is not 'evpn'", v.TypeAttributeName),
			)
			return
		}
	}
}
