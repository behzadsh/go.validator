package rules

import (
	"regexp"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// MinDigits checks the field under validation has length more than given min digits.
//
// Usage: "minDigits:numberOfDigit".
// Example: "minDigits:5".
type MinDigits struct {
	translation.BaseTranslatableRule
	digitCount string
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *MinDigits) Validate(selector string, value any, _ bag.InputBag) Result {
	strVal := cast.ToString(value)

	ok, err := regexp.MatchString(`^\pN{`+r.digitCount+`,}$`, strVal)
	if !ok || err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.min_digits", map[string]string{
			"field":      selector,
			"digitCount": r.digitCount,
		}))
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *MinDigits) AddParams(params []string) {
	r.digitCount = params[0]
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule accept only 1 parameter, which is number of minDigits and is mandatory.
func (*MinDigits) MinRequiredParams() int {
	return 1
}
