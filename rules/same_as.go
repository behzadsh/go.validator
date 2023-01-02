package rules

import (
	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// SameAs check the field under validation has the value same as the other given field.
//
// Usage: "sameAs:otherField".
// Example: "sameAs:password".
type SameAs struct {
	translation.BaseTranslatableRule
	otherField string
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *SameAs) Validate(selector string, value any, inputBag bag.InputBag) Result {
	otherValue, _ := inputBag.Get(r.otherField)

	if otherValue != value {
		return NewFailedResult(r.Translate(r.Locale, "validation.same_as", map[string]string{
			"field":      selector,
			"otherField": r.otherField,
		}))
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *SameAs) AddParams(params []string) {
	r.otherField = params[0]
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule needs only one parameter and that is the regex pattern.
func (*SameAs) MinRequiredParams() int {
	return 1
}
