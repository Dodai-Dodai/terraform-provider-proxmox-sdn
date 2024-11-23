// Controller_validator.go
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure ControllerValidator implements the StringValidator interface
var _ validator.String = ControllerValidator{}

// ControllerValidator is a custom validator for the 'Controller' attribute
type ControllerValidator struct {
	TypeAttributeName string
}

// Description returns a plain text description of the validator's behavior
func (v ControllerValidator) Description(_ context.Context) string {
	return fmt.Sprintf("Requires 'Controller' to be set when '%s' is 'vlan' or 'qinq', otherwise 'Controller' must be null", v.TypeAttributeName)
}

// MarkdownDescription returns a markdown description of the validator's behavior
func (v ControllerValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// ValidateString performs the validation logic
func (v ControllerValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
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

	// Check 'Controller' based on the value of 'type'
	if typeValue == "evpn" {
		// 'Controller' must be set
		if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Missing Required Attribute",
				fmt.Sprintf("'Controller' must be set when '%s' is 'evpn'", v.TypeAttributeName),
			)
			return
		}
	} else {
		// 'Controller' must not be set
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
