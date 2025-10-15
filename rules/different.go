package rules

import (
	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Different checks whether the field under validation has a different value than the specified field.
//
// Usage: "different:otherField".
// Example" "different:oldPassword".
type Different struct {
	translation.BaseTranslatableRule
	otherField string
}

// Validate checks if the value of the field under validation is different from the value of the specified field.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *Different) Validate(selector string, value any, inputBag bag.InputBag) ValidationResult {
	otherValue, _ := inputBag.Get(r.otherField)

	if value == otherValue {
		return NewFailedResult(r.Translate(r.Locale, "validation.different", map[string]string{
			"field":      selector,
			"otherField": r.otherField,
		}))
	}

	return NewSuccessResult()
}

// AddParams assigns the provided parameter values to the Different rule instance.
// The first parameter specifies the `otherField` to compare against (required).
func (r *Different) AddParams(params []string) {
	r.otherField = params[0]
}

// MinRequiredParams returns the minimum number of required parameters for the Different rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 1, indicating that the `otherField` parameter is mandatory.
func (*Different) MinRequiredParams() int {
	return 1
}
