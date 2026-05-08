package rules

import (
	"time"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// DateTimeBefore checks whether the field under validation is a datetime string that occurs before the given datetime value.
//
// Usage: "dateTimeBefore:value[,timeZone]".
// Example: "dateTimeBefore:2021-01-01".
// Example: "dateTimeBefore:2021-01-01,America/New_York".
type DateTimeBefore struct {
	translation.BaseTranslatableRule
	threshold time.Time
	timeZone  *time.Location
}

// Validate checks if the value of the field under validation is a datetime string that occurs before the given datetime value.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *DateTimeBefore) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	timeValue, err := cast.ToTimeInDefaultLocationE(value, r.timeZone)
	if err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.datetime", map[string]string{
			"field": selector,
		}))
	}

	if !timeValue.Before(r.threshold) {
		return NewFailedResult(r.Translate(r.Locale, "validation.datetime_before", map[string]string{
			"field": selector,
			"value": r.threshold.Format(time.RFC3339),
		}))
	}

	return NewSuccessResult()
}

// AddParams assigns the provided parameter values to the DateTimeBefore rule instance.
// The first parameter specifies the `value` to compare against (required),
// and the second parameter, if provided, sets the time zone for parsing date/time values (optional).
func (r *DateTimeBefore) AddParams(params []string) {
	r.timeZone = time.UTC
	if len(params) > 1 {
		if tz, err := time.LoadLocation(params[1]); err == nil {
			r.timeZone = tz
		}
	}

	t, err := cast.ToTimeInDefaultLocationE(params[0], r.timeZone)
	if err == nil {
		r.threshold = t
	}
}

// MinRequiredParams returns the minimum number of required parameters for the DateTimeBefore rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 1, indicating that the `value` parameter is mandatory, while the `timeZone` parameter is optional.
func (*DateTimeBefore) MinRequiredParams() int { return 1 }
