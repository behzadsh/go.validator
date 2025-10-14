package rules

import (
	"time"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Timezone checks the field under validation is a valid IANA time zone name.
//
// Usage: "timezone".
type Timezone struct {
	translation.BaseTranslatableRule
}

// Validate does the validation process of the rule.
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
