package rules

import (
	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// DateTime checks the field under validation be a valid datetime string, and
// it is castable to time.Time.
//
// Usage: "datetime"
type DateTime struct {
	translation.BaseTranslatableRule
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *DateTime) Validate(selector string, value any, _ bag.InputBag) Result {
	_, err := cast.ToTimeE(value)
	if err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.datetime", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}
