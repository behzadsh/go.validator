package rules

import (
	"strings"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Lowercase checks whether the field under validation is a lowercase string.
// This rule accepts no parameters.
//
// Usage: "lowercase".
type Lowercase struct {
	translation.BaseTranslatableRule
}

// Validate checks if the value of the field under validation is a lowercase string.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *Lowercase) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	strValue := cast.ToString(value)

	if strValue != strings.ToLower(strValue) {
		return NewFailedResult(r.Translate(r.Locale, "validation.lowercase", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}

// RequiresField returns false as the Lowercase rule does not require the field to exist.
func (*Lowercase) RequiresField() bool {
	return false
}
