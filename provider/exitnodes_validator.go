// ExitNodes_validator.go
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.Set = ExitNodesValidator{}

type ExitNodesValidator struct {
	TypeAttributeName string
}

func (v ExitNodesValidator) Description(_ context.Context) string {
	return fmt.Sprintf("'ExitNodes' must be set when '%s' is 'evpn', otherwise 'ExitNodes' must be null", v.TypeAttributeName)
}

func (v ExitNodesValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v ExitNodesValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
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
		// 値が設定されていなくてもいい
		if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() || len(req.ConfigValue.Elements()) == 0 {
			return
		}
	} else {
		// 値が設定されていないこと
		if !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() && len(req.ConfigValue.Elements()) != 0 {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid Attribute",
				fmt.Sprintf("'ExitNodes' must be null or empty when '%s' is not 'evpn'", v.TypeAttributeName),
			)
			return
		}
	}
}
