package rules

import (
	"regexp"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Digits checks the field under validation has an exact length digits.
//
// Usage: "digits:numberOfDigit".
// Example: "digits:5".
type Digits struct {
	translation.BaseTranslatableRule
	digitCount string
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *Digits) Validate(selector string, value any, _ bag.InputBag) Result {
	strVal := cast.ToString(value)

	ok, err := regexp.MatchString(`^\pN{`+r.digitCount+`}$`, strVal)
	if !ok || err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.digits", map[string]string{
			"field":      selector,
			"digitCount": r.digitCount,
		}))
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *Digits) AddParams(params []string) {
	r.digitCount = params[0]
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule accept only 1 parameter, which is number of digits and is mandatory.
func (*Digits) MinRequiredParams() int {
	return 1
}
