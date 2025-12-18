package rules

import (
	"github.com/thoas/go-funk"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// NotEmpty checks whether the field under validation be a non-empty or non-zero value.
// This rule accepts no parameters.
//
// Usage: "notEmpty".
type NotEmpty struct {
	translation.BaseTranslatableRule
}

// Validate checks if the value of the field under validation be a non-empty or non-zero value.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *NotEmpty) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	if funk.IsEmpty(value) {
		return NewFailedResult(r.Translate(r.Locale, "validation.not_empty", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}

// RequiresField returns true as the NotEmpty rule requires the field to exist.
func (*NotEmpty) RequiresField() bool {
	return true
}
