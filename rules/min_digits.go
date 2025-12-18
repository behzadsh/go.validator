package rules

import (
	"regexp"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// MinDigits checks whether the field under validation has length more than given min digits.
//
// Usage: "minDigits:numberOfDigit".
// Example: "minDigits:5".
type MinDigits struct {
	translation.BaseTranslatableRule
	digitCount string
}

// Validate checks if the value of the field under validation has length more than given min digits.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *MinDigits) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	strVal := cast.ToString(value)

	ok, err := regexp.MatchString(`^\d{`+r.digitCount+`,}$`, strVal)
	if !ok || err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.min_digits", map[string]string{
			"field":      selector,
			"digitCount": r.digitCount,
		}))
	}

	return NewSuccessResult()
}

// AddParams assigns the provided parameter values to the MinDigits rule instance.
// The first parameter specifies the `numberOfDigit` to compare against (required).
func (r *MinDigits) AddParams(params []string) {
	r.digitCount = params[0]
}

// MinRequiredParams returns the minimum number of required parameters for the MinDigits rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 1, indicating that the `numberOfDigit` parameter is mandatory.
func (*MinDigits) MinRequiredParams() int {
	return 1
}

// RequiresField returns false as the MinDigits rule does not require the field to exist.
func (*MinDigits) RequiresField() bool {
	return false
}
