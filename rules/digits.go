package rules

import (
	"regexp"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Digits checks whether the field under validation has an exact number of digits.
//
// Usage: "digits:numberOfDigit".
// Example: "digits:5".
type Digits struct {
	translation.BaseTranslatableRule
	digitCount string
}

// Validate checks if the value of the field under validation has an exact number of digits.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *Digits) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	ok, err := regexp.MatchString(`^\d{`+r.digitCount+`}$`, cast.ToString(value))
	if !ok || err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.digits", map[string]string{
			"field":      selector,
			"digitCount": r.digitCount,
		}))
	}

	return NewSuccessResult()
}

// AddParams assigns the provided parameter values to the Digits rule instance.
// The first parameter specifies the `numberOfDigit` to compare against (required).
func (r *Digits) AddParams(params []string) {
	r.digitCount = params[0]
}

// MinRequiredParams returns the minimum number of required parameters for the Digits rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 1, indicating that the `numberOfDigit` parameter is mandatory.
func (*Digits) MinRequiredParams() int {
	return 1
}
