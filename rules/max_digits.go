package rules

import (
	"regexp"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// MaxDigits checks the field under validation has length less than given max digits.
//
// Usage: "maxDigits:numberOfDigit".
// Example: "maxDigits:5".
type MaxDigits struct {
	translation.BaseTranslatableRule
	digitCount string
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *MaxDigits) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	strVal := cast.ToString(value)

	ok, err := regexp.MatchString(`^\pN{0,`+r.digitCount+`}$`, strVal)
	if !ok || err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.max_digits", map[string]string{
			"field":      selector,
			"digitCount": r.digitCount,
		}))
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *MaxDigits) AddParams(params []string) {
	r.digitCount = params[0]
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule accept only 1 parameter, which is number of maxDigits and is mandatory.
func (*MaxDigits) MinRequiredParams() int {
	return 1
}
