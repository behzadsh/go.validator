package rules

import (
	"github.com/thoas/go-funk"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Array checks whether the field under validation is an array or slice.
// This rule accepts no parameters.
//
// Usage: "array".
type Array struct {
	translation.BaseTranslatableRule
}

// Validate checks if the value of the field under validation is an array or slice.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *Array) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	if !funk.IsCollection(value) {
		return NewFailedResult(r.Translate(r.Locale, "validation.array", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}
