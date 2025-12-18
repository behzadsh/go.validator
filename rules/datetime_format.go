package rules

import (
	"time"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// DateTimeFormat checks whether the field under validation is a valid datetime string that matches the given format.
//
// Usage: "datetimeFormat:format".
// Example: "datetimeFormat:2006-01-02T15:04:05Z07:00".
type DateTimeFormat struct {
	translation.BaseTranslatableRule
	layout string
}

// Validate checks if the value of the field under validation is a valid datetime string that matches the given format.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *DateTimeFormat) Validate(selector string, value any, inputBag bag.InputBag) ValidationResult {
	if !inputBag.Has(selector) {
		return NewSuccessResult()
	}

	_, err := time.Parse(r.layout, cast.ToString(value))
	if err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.datetime_format", map[string]string{
			"field":  selector,
			"format": r.layout,
		}))
	}

	return NewSuccessResult()
}

// AddParams assigns the provided parameter values to the DateTimeFormat rule instance.
// The first parameter specifies the `format` to compare against (required).
func (r *DateTimeFormat) AddParams(params []string) {
	r.layout = params[0]
}

// MinRequiredParams returns the minimum number of required parameters for the DateTimeFormat rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 1, indicating that the `format` parameter is mandatory.
func (*DateTimeFormat) MinRequiredParams() int {
	return 1
}

// RequiresField returns false as the DateTimeFormat rule does not require the field to exist.
func (*DateTimeFormat) RequiresField() bool {
	return false
}
