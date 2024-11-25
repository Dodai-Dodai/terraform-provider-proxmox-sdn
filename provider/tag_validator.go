// tag_validator.go
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.Int64 = TagValidator{}

type TagValidator struct {
	TypeAttributeName string
}

func (v TagValidator) Description(_ context.Context) string {
	return fmt.Sprintf("Requires 'tag' to be set when '%s' is 'qinq', otherwise 'tag' must be null", v.TypeAttributeName)
}

func (v TagValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v TagValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	// Get the value of the 'type' attribute
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

	if typeValue == "qinq" {
		// 'tag' は設定されている必要がある
		if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Missing Required Attribute",
				fmt.Sprintf("'tag' must be set when '%s' is 'qinq'", v.TypeAttributeName),
			)
			return
		}
	} else {
		// 'tag' は設定されていてはいけない
		if !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid Attribute",
				fmt.Sprintf("'tag' must be null when '%s' is not 'qinq'", v.TypeAttributeName),
			)
			return
		}
	}
}
