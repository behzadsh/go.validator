package rules

import (
	"time"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// DateTimeBetween checks whether the field under validation is a datetime string that occurs between two given datetime values (inclusive).
//
// Usage: "dateTimeBetween:min,max[,timeZone]".
// Example: "dateTimeBetween:2021-01-01,2021-01-02".
// Example: "dateTimeBetween:2021-01-01,2021-01-02,America/New_York".
type DateTimeBetween struct {
	translation.BaseTranslatableRule
	min      time.Time
	max      time.Time
	timeZone *time.Location
}

// Validate checks if the value of the field under validation is a datetime string that occurs between two given datetime values (inclusive).
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *DateTimeBetween) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	timeValue, err := cast.ToTimeInDefaultLocationE(value, r.timeZone)
	if err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.datetime", map[string]string{
			"field": selector,
		}))
	}

	if timeValue.Before(r.min) || timeValue.After(r.max) {
		return NewFailedResult(r.Translate(r.Locale, "validation.datetime_between", map[string]string{
			"field": selector,
			"min":   r.min.Format(time.RFC3339),
			"max":   r.max.Format(time.RFC3339),
		}))
	}

	return NewSuccessResult()
}

// AddParams assigns the provided parameter values to the DateTimeBetween rule instance.
// The first parameter specifies the `min` value to compare against (required),
// the second parameter specifies the `max` value to compare against (required),
// and the third parameter, if provided, sets the time zone for parsing date/time values (optional).
func (r *DateTimeBetween) AddParams(params []string) {
	r.timeZone = time.UTC
	if len(params) > 2 {
		if tz, err := time.LoadLocation(params[2]); err == nil {
			r.timeZone = tz
		}
	}

	if t, err := cast.ToTimeInDefaultLocationE(params[0], r.timeZone); err == nil {
		r.min = t
	}
	if t, err := cast.ToTimeInDefaultLocationE(params[1], r.timeZone); err == nil {
		r.max = t
	}
}

// MinRequiredParams returns the minimum number of required parameters for the DateTimeBetween rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 2, indicating that the `min` and `max` parameters are mandatory, while the `timeZone` parameter is optional.
func (*DateTimeBetween) MinRequiredParams() int { return 2 }

// RequiresField returns false as the DateTimeBetween rule does not require the field to exist.
func (*DateTimeBetween) RequiresField() bool {
	return false
}
