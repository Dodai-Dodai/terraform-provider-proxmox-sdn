// validator.go

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TypeBasedRequiredValidator は、type フィールドの値に基づいて他のフィールドが必須かどうかを検証します。
type TypeBasedRequiredValidator struct {
	TypeField      path.Expression
	ExpectedValue  string
	DependentField path.Expression
}

// TypeBasedRequired は新しい TypeBasedRequiredValidator を返します。
func TypeBasedRequired(typeField path.Expression, expectedValue string, dependentField path.Expression) validator.Object {
	return &TypeBasedRequiredValidator{
		TypeField:      typeField,
		ExpectedValue:  expectedValue,
		DependentField: dependentField,
	}
}

// Description はバリデータの説明を返します。
func (v *TypeBasedRequiredValidator) Description(ctx context.Context) string {
	return fmt.Sprintf("When '%s' is '%s', '%s' must be set", v.TypeField, v.ExpectedValue, v.DependentField)
}

// Validate はバリデーションを実行します。
func (v *TypeBasedRequiredValidator) Validate(ctx context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	var typeValue types.String
	diags := req.Config.GetAttribute(ctx, v.TypeField, &typeValue)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if typeValue.ValueString() != v.ExpectedValue {
		return
	}

	var dependentValue types.Object
	diags = req.Config.GetAttribute(ctx, v.DependentField, &dependentValue)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if dependentValue.IsNull() || dependentValue.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			v.DependentField,
			"Missing Required Configuration",
			fmt.Sprintf("When '%s' is set to '%s', the attribute '%s' must be configured.",
				v.TypeField.Selectors()[0], v.ExpectedValue, v.DependentField.Selectors()[0]),
		)
	}
}
