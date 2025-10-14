package rules

import (
	"time"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// DateTimeBetween checks the field under validation is a datetime between two given datetime strings (inclusive).
//
// Usage: "dateTimeBetween:min,max[,timeZone]".
type DateTimeBetween struct {
	translation.BaseTranslatableRule
	min      time.Time
	max      time.Time
	timeZone *time.Location
}

// Validate does the validation process of the rule.
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

// AddParams adds rules parameter values to the rule instance.
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

// MinRequiredParams returns minimum parameter requirement for this rule.
func (*DateTimeBetween) MinRequiredParams() int { return 2 }
