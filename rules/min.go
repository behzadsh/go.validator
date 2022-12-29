package rules

import (
	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Min checks the field under validation value be greater than given value.
//
// Usage: "min:value"
// Example: "min:10"
type Min struct {
	translation.BaseTranslatableRule
	min float64
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *Min) Validate(selector string, value any, inputBag bag.InputBag) Result {
	floatValue := cast.ToFloat64(value)

	if floatValue < r.min {
		return NewFailedResult(r.Translate(r.Locale, "validation.min", map[string]string{
			"field": selector,
			"value": cast.ToString(r.min),
		}))
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *Min) AddParams(params []string) {
	r.min = cast.ToFloat64(params[0])
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule need only one parameter that is the `minValue`.
func (r *Min) MinRequiredParams() int {
	return 1
}
