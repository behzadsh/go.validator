package rules

import (
	"regexp"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// DigitsBetween checks whether the field under validation has a length between the given min and max parameters.
//
// Usage: "digitsBetween:minDigits,maxDigits".
// Example: "digitsBetween:4,6".
type DigitsBetween struct {
	translation.BaseTranslatableRule
	min, max string
}

// Validate checks if the value of the field under validation has a length between the given min and max parameters.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *DigitsBetween) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	ok, err := regexp.MatchString(`^\d{`+r.min+`,`+r.max+`}$`, cast.ToString(value))
	if !ok || err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.digits_between", map[string]string{
			"field": selector,
			"min":   r.min,
			"max":   r.max,
		}))
	}

	return NewSuccessResult()
}

// AddParams sets the minimum and maximum digit length constraints for the DigitsBetween rule.
// It takes two parameters: the first is assigned as the minimum required digit length, and the second as the maximum
// allowed digit length. No value is returned by this function.
func (r *DigitsBetween) AddParams(params []string) {
	r.min = params[0]
	r.max = params[1]
}

// MinRequiredParams returns the minimum number of required parameters for the DigitsBetween rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 2, indicating that the `min` and `max` parameters are mandatory.
func (*DigitsBetween) MinRequiredParams() int {
	return 2
}
