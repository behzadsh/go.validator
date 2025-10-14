package rules

import (
	"time"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// DateTimeAfter checks the field under validation is a datetime after the given datetime string.
//
// Usage: "dateTimeAfter:value[,timeZone]".
type DateTimeAfter struct {
	translation.BaseTranslatableRule
	threshold time.Time
	timeZone  *time.Location
}

// Validate does the validation process of the rule.
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

// AddParams adds rules parameter values to the rule instance.
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

// MinRequiredParams returns minimum parameter requirement for this rule.
func (*DateTimeAfter) MinRequiredParams() int { return 1 }
