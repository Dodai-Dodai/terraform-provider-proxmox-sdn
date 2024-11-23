// tag_validator.go
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure TagValidator implements the Int64Validator interface
var _ validator.Int64 = TagValidator{}

// TagValidator is a custom validator for the 'tag' attribute
type TagValidator struct {
	TypeAttributeName string
}

// Description returns a plain text description of the validator's behavior
func (v TagValidator) Description(_ context.Context) string {
	return fmt.Sprintf("Requires 'tag' to be set when '%s' is 'qinq', otherwise 'tag' must be null", v.TypeAttributeName)
}

// MarkdownDescription returns a markdown description of the validator's behavior
func (v TagValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// ValidateInt64 performs the validation logic
func (v TagValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
	// Get the value of the 'type' attribute
	typeAttrPath := path.Root(v.TypeAttributeName)
	var typeAttr types.String

	diags := req.Config.GetAttribute(ctx, typeAttrPath, &typeAttr)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	// If 'type' is unknown or null, we cannot proceed with validation
	if typeAttr.IsUnknown() || typeAttr.IsNull() {
		return
	}

	typeValue := typeAttr.ValueString()

	// Check 'tag' based on the value of 'type'
	if typeValue == "qinq" {
		// 'tag' must be set
		if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Missing Required Attribute",
				fmt.Sprintf("'tag' must be set when '%s' is 'qinq'", v.TypeAttributeName),
			)
			return
		}
	} else {
		// 'tag' must not be set
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
