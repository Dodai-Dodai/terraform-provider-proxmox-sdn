// PrimaryExitNode_validator.go
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure PrimaryExitNodeValidator implements the StringValidator interface
var _ validator.String = PrimaryExitNodeValidator{}

// PrimaryExitNodeValidator is a custom validator for the 'PrimaryExitNode' attribute
type PrimaryExitNodeValidator struct {
	TypeAttributeName string
}

// Description returns a plain text description of the validator's behavior
func (v PrimaryExitNodeValidator) Description(_ context.Context) string {
	return fmt.Sprintf("'PrimaryExitNode''%s' is 'evpn'; otherwise, 'PrimaryExitNode' must be null", v.TypeAttributeName)
}

// MarkdownDescription returns a markdown description of the validator's behavior
func (v PrimaryExitNodeValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// ValidateString performs the validation logic
func (v PrimaryExitNodeValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
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

	if typeValue == "evpn" {
		// 'PrimaryExitNode' can be set but must be one of the allowed values
		if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
			// It's acceptable for 'PrimaryExitNode' to be null when 'type' is 'evpn'
			return
		}
	} else {
		// 'PrimaryExitNode' must not be set
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
