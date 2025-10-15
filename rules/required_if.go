package rules

import (
	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// RequiredIf checks whether the field under validation must exist if the given condition is true.
// The condition consists of another field name and a value. If the value of the other field is equal to the given
// value, the field under validation is required.
// Note that the only supported type for the value parameter is string.
//
// Usage: "requiredIf:otherField,value".
// Example: "requiredIf:type,user".
type RequiredIf struct {
	translation.BaseTranslatableRule
	otherField, expectedValue string
}

// Validate checks if the value of the field under validation must exist if the given condition is true.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *RequiredIf) Validate(selector string, _ any, inputBag bag.InputBag) ValidationResult {
	exists := inputBag.Has(selector)
	otherValue, _ := inputBag.Get(r.otherField)

	if !exists && cast.ToString(otherValue) == r.expectedValue {
		return NewFailedResult(r.Translate(r.Locale, "validation.required_if", map[string]string{
			"field":      selector,
			"otherField": r.otherField,
			"value":      r.expectedValue,
		}))
	}

	return NewSuccessResult()
}

// AddParams assigns the provided parameter values to the RequiredIf rule instance.
// The first parameter specifies the `otherField` to compare against (required),
// and the second parameter, if provided, specifies the `value` to compare against (optional).
func (r *RequiredIf) AddParams(params []string) {
	r.otherField = params[0]
	r.expectedValue = params[1]
}

// MinRequiredParams returns the minimum number of required parameters for the RequiredIf rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 2, indicating that the `otherField` and `value` parameters are mandatory.
func (*RequiredIf) MinRequiredParams() int {
	return 2
}
