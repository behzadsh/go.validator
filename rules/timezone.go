package rules

import (
	"time"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Timezone checks whether the field under validation is a valid IANA time zone name.
// This rule accepts no parameters.
//
// Usage: "timezone".
type Timezone struct {
	translation.BaseTranslatableRule
}

// Validate checks if the value of the field under validation is a valid IANA time zone name.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *Timezone) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	name := cast.ToString(value)
	if name == "" {
		return NewFailedResult(r.Translate(r.Locale, "validation.timezone", map[string]string{
			"field": selector,
		}))
	}
	if _, err := time.LoadLocation(name); err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.timezone", map[string]string{
			"field": selector,
		}))
	}
	return NewSuccessResult()
}
