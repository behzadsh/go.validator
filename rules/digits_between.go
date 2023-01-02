package rules

import (
	"regexp"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// DigitsBetween checks the field under validation has length between given min and max.
//
// Usage: "digitsBetween:minDigits,maxDigits".
// Example: "digitsBetween:4,6".
type DigitsBetween struct {
	translation.BaseTranslatableRule
	min, max string
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *DigitsBetween) Validate(selector string, value any, _ bag.InputBag) Result {
	strVal := cast.ToString(value)

	ok, err := regexp.MatchString(`^\pN{`+r.min+`,`+r.max+`}$`, strVal)
	if !ok || err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.digits_between", map[string]string{
			"field": selector,
			"min":   r.min,
			"max":   r.max,
		}))
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *DigitsBetween) AddParams(params []string) {
	r.min = params[0]
	r.max = params[1]
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule accept 2 parameters, min and max.
func (*DigitsBetween) MinRequiredParams() int {
	return 2
}
