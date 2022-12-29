package rules

import (
	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Max checks the field under validation value be less than given value.
//
// Usage: "max:value"
// Example: "max:10"
type Max struct {
	translation.BaseTranslatableRule
	max float64
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *Max) Validate(selector string, value any, inputBag bag.InputBag) Result {
	floatValue := cast.ToFloat64(value)

	if floatValue > r.max {
		return NewFailedResult(r.Translate(r.Locale, "validation.max", map[string]string{
			"field": selector,
			"value": cast.ToString(r.max),
		}))
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *Max) AddParams(params []string) {
	r.max = cast.ToFloat64(params[0])
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule need only one parameter that is the `maxValue`.
func (r *Max) MinRequiredParams() int {
	return 1
}
