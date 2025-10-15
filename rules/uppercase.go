package rules

import (
	"strings"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Uppercase checks whether the field under validation be uppercase string.
// This rule accepts no parameters.
//
// Usage: "uppercase".
type Uppercase struct {
	translation.BaseTranslatableRule
}

// Validate checks if the value of the field under validation be uppercase string.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *Uppercase) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	strValue := cast.ToString(value)

	if strValue != strings.ToUpper(strValue) {
		return NewFailedResult(r.Translate(r.Locale, "validation.uppercase", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}
