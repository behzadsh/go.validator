package rules

import (
	"strings"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// In checks whether the field under validation is in one of the given values.
//
// Usage: "in:value1,value2[,value3,...].
// Example: "in:EUR,USD,GBP".
type In struct {
	translation.BaseTranslatableRule
	values []string
}

// Validate checks if the value of the field under validation is in one of the given values.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *In) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	strValue := strings.ToLower(cast.ToString(value))
	for _, val := range r.values {
		val = strings.ToLower(val)
		if strValue == val {
			return NewSuccessResult()
		}
	}

	return NewFailedResult(r.Translate(r.Locale, "validation.in", map[string]string{
		"field": selector,
	}))
}

// AddParams assigns the provided parameter values to the In rule instance.
// The first parameter specifies the `value1` to compare against (required),
// and the second parameter, if provided, specifies the `value2` to compare against (optional),
// and so on.
func (r *In) AddParams(params []string) {
	r.values = params
}

// MinRequiredParams returns the minimum number of required parameters for the In rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 2, indicating that the `value1` and `value2` parameters are mandatory.
func (*In) MinRequiredParams() int {
	return 2
}
