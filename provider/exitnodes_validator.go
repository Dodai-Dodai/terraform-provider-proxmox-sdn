// ExitNodes_validator.go
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure ExitNodesValidator implements the Set interface
var _ validator.Set = ExitNodesValidator{}

// ExitNodesValidator is a custom validator for the 'ExitNodes' attribute
type ExitNodesValidator struct {
	TypeAttributeName string
}

// Description returns a plain text description of the validator's behavior
func (v ExitNodesValidator) Description(_ context.Context) string {
	return fmt.Sprintf("'ExitNodes' must be set when '%s' is 'evpn', otherwise 'ExitNodes' must be null", v.TypeAttributeName)
}

// MarkdownDescription returns a markdown description of the validator's behavior
func (v ExitNodesValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// ValidateSet performs the validation logic
func (v ExitNodesValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
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
		// 'ExitNodes' must be set and not empty
		if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() || len(req.ConfigValue.Elements()) == 0 {
			return
		}
	} else {
		// 'ExitNodes' must not be set
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
