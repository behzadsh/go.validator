package rules

import (
	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Min checks whether the field under validation value be greater than given value.
//
// Usage: "min:value".
// Example: "min:10".
type Min struct {
	translation.BaseTranslatableRule
	min float64
}

// Validate checks if the value of the field under validation value be greater than given value.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *Min) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	floatValue := cast.ToFloat64(value)

	if floatValue < r.min {
		return NewFailedResult(r.Translate(r.Locale, "validation.min", map[string]string{
			"field": selector,
			"value": cast.ToString(r.min),
		}))
	}

	return NewSuccessResult()
}

// AddParams assigns the provided parameter values to the Min rule instance.
// The first parameter specifies the `value` to compare against (required).
func (r *Min) AddParams(params []string) {
	r.min = cast.ToFloat64(params[0])
}

// MinRequiredParams returns the minimum number of required parameters for the Min rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 1, indicating that the `value` parameter is mandatory.
func (*Min) MinRequiredParams() int {
	return 1
}
