package rules

import (
	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// NotIn checks the field under validation not be in one of given values.
//
// Usage: "notIn:value1,value2[,value3,...]
// Example: "notIn:admin,superuser"
type NotIn struct {
	translation.BaseTranslatableRule
	values []string
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *NotIn) Validate(selector string, value any, _ bag.InputBag) Result {
	strValue := cast.ToString(value)
	for _, val := range r.values {
		if strValue == val {
			return NewFailedResult(r.Translate(r.Locale, "validation.not_in", map[string]string{
				"field": selector,
			}))
		}
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *NotIn) AddParams(params []string) {
	r.values = params
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule needs at least 2 parameter or more.
func (r *NotIn) MinRequiredParams() int {
	return 2
}
