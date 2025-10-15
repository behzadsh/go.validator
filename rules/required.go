package rules

import (
	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Required checks whether the field under validation must exist.
// This rule accepts no parameters.
//
// Usage: "required".
type Required struct {
	translation.BaseTranslatableRule
}

// Validate checks if the value of the field under validation must exist.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *Required) Validate(selector string, _ any, inputBag bag.InputBag) ValidationResult {
	if !inputBag.Has(selector) {
		return NewFailedResult(r.Translate(r.Locale, "validation.required", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}
