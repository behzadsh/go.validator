package rules

import (
	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Boolean checks whether the field under validation is boolean or can be cast as a boolean value.
// This rule accepts no parameters.
//
// Usage: "boolean".
type Boolean struct {
	translation.BaseTranslatableRule
}

// Validate checks if the value of the field under validation is a boolean or can be cast as a boolean value.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *Boolean) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	_, err := cast.ToBoolE(value)
	if err != nil || value == nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.boolean", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}

// RequiresField returns false as the Boolean rule does not require the field to exist.
func (*Boolean) RequiresField() bool {
	return false
}
