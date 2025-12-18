package rules

import (
	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Numeric checks whether the field under validation has a numeric value.
// This rule accepts no parameters.
//
// Usage: "numeric".
type Numeric struct {
	translation.BaseTranslatableRule
}

// Validate checks if the value of the field under validation has a numeric value.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *Numeric) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	_, err := cast.ToFloat64E(value)
	if err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.numeric", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}

// RequiresField returns false as the Numeric rule does not require the field to exist.
func (*Numeric) RequiresField() bool {
	return false
}
