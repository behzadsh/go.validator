package rules

import (
	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// In checks the field under validation be in one of given values.
//
// Usage: "in:value1,value2[,value3,...]
// Example: "in:EUR,USD,GBP"
type In struct {
	translation.BaseTranslatableRule
	values []string
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *In) Validate(selector string, value any, _ bag.InputBag) Result {
	strValue := cast.ToString(value)
	for _, val := range r.values {
		if strValue == val {
			return NewSuccessResult()
		}
	}

	return NewFailedResult(r.Translate(r.Locale, "validation.in", map[string]string{
		"field": selector,
	}))
}

// AddParams adds rules parameter values to the rule instance.
func (r *In) AddParams(params []string) {
	r.values = params
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule needs at least 2 parameter or more.
func (r *In) MinRequiredParams() int {
	return 2
}
