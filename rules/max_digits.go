package rules

import (
	"regexp"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// MaxDigits checks whether the field under validation has length less than given max digits.
//
// Usage: "maxDigits:numberOfDigit".
// Example: "maxDigits:5".
type MaxDigits struct {
	translation.BaseTranslatableRule
	digitCount string
}

// Validate checks if the value of the field under validation has length less than given max digits.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *MaxDigits) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	strVal := cast.ToString(value)

	ok, err := regexp.MatchString(`^\d{1,`+r.digitCount+`}$`, strVal)
	if !ok || err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.max_digits", map[string]string{
			"field":      selector,
			"digitCount": r.digitCount,
		}))
	}

	return NewSuccessResult()
}

// AddParams assigns the provided parameter values to the MaxDigits rule instance.
// The first parameter specifies the `numberOfDigit` to compare against (required).
func (r *MaxDigits) AddParams(params []string) {
	r.digitCount = params[0]
}

// MinRequiredParams returns the minimum number of required parameters for the MaxDigits rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 1, indicating that the `numberOfDigit` parameter is mandatory.
func (*MaxDigits) MinRequiredParams() int {
	return 1
}
