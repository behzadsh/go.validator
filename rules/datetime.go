package rules

import (
	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// DateTime checks whether the field under validation is a valid datetime string and can be cast to time.Time.
// This rule accepts no parameters.
//
// Usage: "datetime".
type DateTime struct {
	translation.BaseTranslatableRule
}

// Validate checks if the value of the field under validation is a valid datetime string and can be cast to time.Time.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *DateTime) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	_, err := cast.StringToDate(cast.ToString(value))
	if err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.datetime", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}
