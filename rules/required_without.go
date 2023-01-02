package rules

import (
	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// RequiredWithout check the field under validation exists if any of given
// fields doesn't exist.
//
// Usage: "requiredWithout:otherField[,anotherField,...]".
// example: "requiredWithout:username".
type RequiredWithout struct {
	translation.BaseTranslatableRule
	otherFields []string
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *RequiredWithout) Validate(selector string, _ any, inputBag bag.InputBag) Result {
	exists := inputBag.Has(selector)

	if !exists {
		for _, field := range r.otherFields {
			if !inputBag.Has(field) {
				return NewFailedResult(r.Translate(r.Locale, "validation.required_without", map[string]string{
					"field":      selector,
					"otherField": field,
				}))
			}
		}
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *RequiredWithout) AddParams(params []string) {
	r.otherFields = params
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule needs at least 1 parameter which represent the name of other
// field or fields.
func (*RequiredWithout) MinRequiredParams() int {
	return 1
}
