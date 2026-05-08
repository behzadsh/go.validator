package rules

import (
	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// NotEqual checks whether the field under validation is not equal to given value.
//
// Usage: "neq:value".
// Example: "neq:admin".
type NotEqual struct {
	translation.BaseTranslatableRule
	value string
}

// Validate checks if the value of the field under validation is not equal to given value.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *NotEqual) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	if cast.ToString(value) == r.value {
		return NewFailedResult(r.Translate(r.Locale, "validation.neq", map[string]string{
			"field": selector,
			"value": r.value,
		}))
	}

	return NewSuccessResult()
}

// AddParams assigns the provided parameter values to the NotEqual rule instance.
// The first parameter specifies the `value` to compare against (required).
func (r *NotEqual) AddParams(params []string) {
	r.value = params[0]
}

// MinRequiredParams returns the minimum number of required parameters for the NotEqual rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 1, indicating that the `value` parameter is mandatory.
func (*NotEqual) MinRequiredParams() int {
	return 1
}
