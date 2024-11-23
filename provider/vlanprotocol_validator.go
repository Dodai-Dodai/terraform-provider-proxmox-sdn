// vlanprotocol_validator.go
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure VLANProtocolValidator implements the StringValidator interface
var _ validator.String = VLANProtocolValidator{}

// VLANProtocolValidator is a custom validator for the 'vlanprotocol' attribute
type VLANProtocolValidator struct {
	TypeAttributeName string
}

// Description returns a plain text description of the validator's behavior
func (v VLANProtocolValidator) Description(_ context.Context) string {
	return fmt.Sprintf("'vlanprotocol' must be '802.1ad' or '802.1q' when '%s' is 'qinq'; otherwise, 'vlanprotocol' must be null", v.TypeAttributeName)
}

// MarkdownDescription returns a markdown description of the validator's behavior
func (v VLANProtocolValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// ValidateString performs the validation logic
func (v VLANProtocolValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
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

	// Allowed VLANProtocol values
	allowedValues := []string{"802.1ad", "802.1q"}

	if typeValue == "qinq" {
		// 'vlanprotocol' can be set but must be one of the allowed values
		if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
			// It's acceptable for 'vlanprotocol' to be null when 'type' is 'qinq'
			return
		}

		vlanProtocolValue := req.ConfigValue.ValueString()

		// Validate the value
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
		// 'vlanprotocol' must not be set
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
