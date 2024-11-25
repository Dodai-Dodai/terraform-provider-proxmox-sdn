// AdvertiseSubnets_validator.go
package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.Bool = AdvertiseSubnetsValidator{}

type AdvertiseSubnetsValidator struct {
	TypeAttributeName string
}

func (v AdvertiseSubnetsValidator) Description(_ context.Context) string {
	return fmt.Sprintf("'AdvertiseSubnets''%s' is 'evpn'; otherwise, 'AdvertiseSubnets' must be null", v.TypeAttributeName)
}

func (v AdvertiseSubnetsValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v AdvertiseSubnetsValidator) ValidateBool(ctx context.Context, req validator.BoolRequest, resp *validator.BoolResponse) {
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

	if typeValue == "evpn" {
		// 値が設定されていなくてもよい
		if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
			return
		}
	} else {
		// 値が設定されていてはいけない
		if !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid Attribute",
				fmt.Sprintf("'AdvertiseSubnets' must be null when '%s' is not 'evpn'", v.TypeAttributeName),
			)
			return
		}
	}
}
