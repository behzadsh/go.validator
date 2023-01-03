package rules

import (
	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// NotEqual checks the field under validation is not equal to given value.
//
// Usage: "neq:value".
// Example: "neq:admin".
type NotEqual struct {
	translation.BaseTranslatableRule
	value string
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *NotEqual) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	strValue := cast.ToString(value)

	if strValue == r.value {
		return NewFailedResult(r.Translate(r.Locale, "validation.neq", map[string]string{
			"field": selector,
			"value": r.value,
		}))
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *NotEqual) AddParams(params []string) {
	r.value = params[0]
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule only needs one parameter and that is the `value`.
func (*NotEqual) MinRequiredParams() int {
	return 1
}
