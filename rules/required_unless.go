package rules

import (
	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// RequiredUnless check the field under validation exists unless the given
// condition is true. The condition is consists of another field name and a
// value. Unless the value of the other field is equal to the given value, the
// field under validation is required.
// Note that the only supported type for value parameter is string.
//
// Usage: "requiredUnless:otherField,value"
// example: "requiredUnless:type,user"
type RequiredUnless struct {
	translation.BaseTranslatableRule
	otherField, expectedValue string
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *RequiredUnless) Validate(selector string, _ any, inputBag bag.InputBag) Result {
	exists := inputBag.Has(selector)
	otherValue, _ := inputBag.Get(r.otherField)

	if !exists && cast.ToString(otherValue) != r.expectedValue {
		return NewFailedResult(r.Translate(r.Locale, "validation.required_unless", map[string]string{
			"field":      selector,
			"otherField": r.otherField,
			"value":      r.expectedValue,
		}))
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *RequiredUnless) AddParams(params []string) {
	r.otherField = params[0]
	r.expectedValue = params[1]
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule accept 2 parameter, `otherField`, and `value`. both parameters
// are mandatory.
func (r *RequiredUnless) MinRequiredParams() int {
	return 2
}
