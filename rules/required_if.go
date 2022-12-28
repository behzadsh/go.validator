package rules

import (
	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// RequiredIf check the field under validation exists if the given condition
// is true. The condition is consists of another field name and a value. If
// the value of the other field is equal to the given value, the field under
// validation is required.
// Note that the only supported type for value parameter is string.
//
// Usage: "requiredIf:otherField,value"
// example: "requiredIf:type,user"
type RequiredIf struct {
	translation.BaseTranslatableRule
	otherField, expectedValue string
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *RequiredIf) Validate(selector string, _ any, inputBag bag.InputBag) Result {
	exists := inputBag.Has(selector)
	otherValue, ok := inputBag.Get(r.otherField)
	if !ok {
		return NewFailedResult(r.Translate(r.Locale, "validation.required", map[string]string{
			"field": r.otherField,
		}))
	}

	if !exists && cast.ToString(otherValue) == r.expectedValue {
		return NewFailedResult(r.Translate(r.Locale, "validation.required_if", map[string]string{
			"field":      selector,
			"otherField": r.otherField,
			"value":      r.expectedValue,
		}))
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *RequiredIf) AddParams(params []string) {
	r.otherField = params[0]
	r.expectedValue = params[1]
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule accept 2 parameter, `otherField`, and `value`. both parameters
// are mandatory.
func (r *RequiredIf) MinRequiredParams() int {
	return 2
}
