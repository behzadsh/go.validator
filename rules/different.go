package rules

import (
	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Different checks the field under validation has different value than
// the given field.
//
// Usage: "different:otherField"
// Example" "different:oldPassword"
type Different struct {
	translation.BaseTranslatableRule
	otherField string
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *Different) Validate(selector string, value any, inputBag bag.InputBag) Result {
	otherValue, _ := inputBag.Get(r.otherField)

	if value == otherValue {
		return NewFailedResult(r.Translate(r.Locale, "validation.different", map[string]string{
			"field":      selector,
			"otherField": r.otherField,
		}))
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *Different) AddParams(params []string) {
	r.otherField = params[0]
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule accept only one parameter, and that is `otherField`. this
// parameter is mandatory.
func (r *Different) MinRequiredParams() int {
	return 1
}
