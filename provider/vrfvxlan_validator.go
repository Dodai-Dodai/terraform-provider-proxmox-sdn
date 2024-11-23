// vrfvxlan_validator.go
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure vrfvxlanValidator implements the Int64Validator interface
var _ validator.Int64 = VrfvxlanValidator{}

// vrfvxlanValidator is a custom validator for the 'vrfvxlan' attribute
type VrfvxlanValidator struct {
	TypeAttributeName string
}

// Description returns a plain text description of the validator's behavior
func (v VrfvxlanValidator) Description(_ context.Context) string {
	return fmt.Sprintf("Requires 'vrfvxlan' to be set when '%s' is 'evpn', otherwise 'vrfvxlan' must be null", v.TypeAttributeName)
}

// MarkdownDescription returns a markdown description of the validator's behavior
func (v VrfvxlanValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// ValidateInt64 performs the validation logic
func (v VrfvxlanValidator) ValidateInt64(ctx context.Context, req validator.Int64Request, resp *validator.Int64Response) {
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

	// Check 'vrfvxlan' based on the value of 'type'
	if typeValue == "evpn" {
		// 'vrfvxlan' must be set
		if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Missing Required Attribute",
				fmt.Sprintf("'vrfvxlan' must be set when '%s' is 'evpn'", v.TypeAttributeName),
			)
			return
		}
	} else {
		// 'vrfvxlan' must not be set
		if !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid Attribute",
				fmt.Sprintf("'vrfvxlan' must be null when '%s' is not 'evpn'", v.TypeAttributeName),
			)
			return
		}
	}
}
