// DisableARPNdSuppression_validator.go
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure DisableARPNdSuppressionValidator implements the StringValidator interface
var _ validator.Bool = DisableARPNdSuppressionValidator{}

// DisableARPNdSuppressionValidator is a custom validator for the 'DisableARPNdSuppression' attribute
type DisableARPNdSuppressionValidator struct {
	TypeAttributeName string
}

// Description returns a plain text description of the validator's behavior
func (v DisableARPNdSuppressionValidator) Description(_ context.Context) string {
	return fmt.Sprintf("'DisableARPNdSuppression''%s' is 'evpn'; otherwise, 'DisableARPNdSuppression' must be null", v.TypeAttributeName)
}

// MarkdownDescription returns a markdown description of the validator's behavior
func (v DisableARPNdSuppressionValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// ValidateBool performs the validation logic
func (v DisableARPNdSuppressionValidator) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
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
		// 'DisableARPNdSuppression' can be set but must be one of the allowed values
		if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
			// It's acceptable for 'DisableARPNdSuppression' to be null when 'type' is 'evpn'
			return
		}
	} else {
		// 'DisableARPNdSuppression' must not be set
		if !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid Attribute",
				fmt.Sprintf("'DisableARPNdSuppression' must be null when '%s' is not 'evpn'", v.TypeAttributeName),
			)
			return
		}
	}
}
