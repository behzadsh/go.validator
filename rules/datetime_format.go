package rules

import (
	"time"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// DateTimeFormat checks the field under validation be a valid datetime string
// and match the given format.
//
// Usage: "datetimeFormat:format"
// Example: "datetimeFormat:2006-01-02T15:04:05Z07:00"
type DateTimeFormat struct {
	translation.BaseTranslatableRule
	layout string
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *DateTimeFormat) Validate(selector string, value any, _ bag.InputBag) Result {
	_, err := time.Parse(r.layout, cast.ToString(value))
	if err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.datetime_format", map[string]string{
			"field":  selector,
			"format": r.layout,
		}))
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *DateTimeFormat) AddParams(params []string) {
	r.layout = params[0]
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule accept only one parameter, the time layout, which is a mandatory
// parameter.
func (r *DateTimeFormat) MinRequiredParams() int {
	return 1
}
