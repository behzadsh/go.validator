package rules

import (
	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Max checks whether the field under validation value be less than given value.
//
// Usage: "max:value".
// Example: "max:10".
type Max struct {
	translation.BaseTranslatableRule
	max float64
}

// Validate checks if the value of the field under validation value be less than given value.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *Max) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	if cast.ToFloat64(value) > r.max {
		return NewFailedResult(r.Translate(r.Locale, "validation.max", map[string]string{
			"field": selector,
			"value": cast.ToString(r.max),
		}))
	}

	return NewSuccessResult()
}

// AddParams assigns the provided parameter values to the Max rule instance.
// The first parameter specifies the `value` to compare against (required).
func (r *Max) AddParams(params []string) {
	r.max = cast.ToFloat64(params[0])
}

// MinRequiredParams returns the minimum number of required parameters for the Max rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 1, indicating that the `value` parameter is mandatory.
func (*Max) MinRequiredParams() int {
	return 1
}

// RequiresField returns false as the Max rule does not require the field to exist.
func (*Max) RequiresField() bool {
	return false
}
