// bridge_validator.go
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure BridgeValidator implements the StringValidator interface
var _ validator.String = BridgeValidator{}

// BridgeValidator is a custom validator for the 'bridge' attribute
type BridgeValidator struct {
	TypeAttributeName string
}

// Description returns a plain text description of the validator's behavior
func (v BridgeValidator) Description(_ context.Context) string {
	return fmt.Sprintf("Requires 'bridge' to be set when '%s' is 'vlan' or 'qinq', otherwise 'bridge' must be null", v.TypeAttributeName)
}

// MarkdownDescription returns a markdown description of the validator's behavior
func (v BridgeValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// ValidateString performs the validation logic
func (v BridgeValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
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

	// Check 'bridge' based on the value of 'type'
	if typeValue == "vlan" || typeValue == "qinq" {
		// 'bridge' must be set
		if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Missing Required Attribute",
				fmt.Sprintf("'bridge' must be set when '%s' is 'vlan' or 'qinq'", v.TypeAttributeName),
			)
			return
		}
	} else {
		// 'bridge' must not be set
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
