package rules

import (
	"time"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// DateTimeAfter checks whether the field under validation is a datetime string that occurs after the given datetime value.
//
// Usage: "dateTimeAfter:value[,timeZone]".
// Example: "dateTimeAfter:2021-01-01".
// Example: "dateTimeAfter:2021-01-01,America/New_York".
type DateTimeAfter struct {
	translation.BaseTranslatableRule
	threshold time.Time
	timeZone  *time.Location
}

// Validate checks if the value of the field under validation is a datetime string that occurs after the given datetime value.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *DateTimeAfter) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	timeValue, err := cast.ToTimeInDefaultLocationE(value, r.timeZone)
	if err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.datetime", map[string]string{
			"field": selector,
		}))
	}

	if !timeValue.After(r.threshold) {
		return NewFailedResult(r.Translate(r.Locale, "validation.datetime_after", map[string]string{
			"field": selector,
			"value": r.threshold.Format(time.RFC3339),
		}))
	}

	return NewSuccessResult()
}

// AddParams assigns the provided parameter values to the DateTimeAfter rule instance.
// The first parameter specifies the `value` to compare against (required),
// and the second parameter, if provided, sets the time zone for parsing date/time values (optional).
func (r *DateTimeAfter) AddParams(params []string) {
	r.timeZone = time.UTC
	if len(params) > 1 {
		if tz, err := time.LoadLocation(params[1]); err == nil {
			r.timeZone = tz
		}
	}

	// parse threshold in given timezone
	t, err := cast.ToTimeInDefaultLocationE(params[0], r.timeZone)
	if err == nil {
		r.threshold = t
	}
}

// MinRequiredParams returns the minimum number of required parameters for the DateTimeAfter rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 1, indicating that the `value` parameter is mandatory, while the `timeZone` parameter is optional.
func (*DateTimeAfter) MinRequiredParams() int { return 1 }

// RequiresField returns false as the DateTimeAfter rule does not require the field to exist.
func (*DateTimeAfter) RequiresField() bool {
	return false
}
