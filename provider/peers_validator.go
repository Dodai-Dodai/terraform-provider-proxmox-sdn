// peers_validator.go
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.Set = PeersValidator{}

type PeersValidator struct {
	TypeAttributeName string
}

func (v PeersValidator) Description(_ context.Context) string {
	return fmt.Sprintf("'peers' must be set when '%s' is 'vxlan', otherwise 'peers' must be null", v.TypeAttributeName)
}

func (v PeersValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v PeersValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
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

	if typeValue == "vxlan" {
		// peersは設定されている必要がある
		if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() || len(req.ConfigValue.Elements()) == 0 {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Missing Required Attribute",
				fmt.Sprintf("'peers' must be set and not empty when '%s' is 'vxlan'", v.TypeAttributeName),
			)
			return
		}
	} else {
		// vxlanでない場合、peersは設定されていてはいけない
		if !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() && len(req.ConfigValue.Elements()) != 0 {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid Attribute",
				fmt.Sprintf("'peers' must be null or empty when '%s' is not 'vxlan'", v.TypeAttributeName),
			)
			return
		}
	}
}
