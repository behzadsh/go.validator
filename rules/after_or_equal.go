package rules

import (
	"time"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// AfterOrEqual checks the field under validation be a value after or equal to
// the value of given field. It will return validation error if one or both of
// the field are not a valid datetime string. It also will return validation
// error if the other field could not be found in input bag.
//
// Usage: "afterOrEqual:otherField[,timeZoneString]
// Example: "afterOrEqual:start,America/New_York"
type AfterOrEqual struct {
	translation.BaseTranslatableRule
	otherField string
	timeZone   *time.Location
	message    string
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *AfterOrEqual) Validate(selector string, value any, inputBag bag.InputBag, _ bool) Result {
	timeValue, err := cast.ToTimeInDefaultLocationE(value, r.timeZone)
	if err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.datetime", map[string]string{
			"field": selector,
		}))
	}

	otherValue, ok := inputBag.Get(r.otherField)
	if !ok {
		return NewFailedResult(r.Translate(r.Locale, "Validation.required", map[string]string{
			"otherField": r.otherField,
		}))
	}

	otherTimeValue, err := cast.ToTimeInDefaultLocationE(otherValue, r.timeZone)
	if err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.datetime", map[string]string{
			"field": r.otherField,
		}))
	}

	result := timeValue.After(otherTimeValue) || timeValue.Equal(otherTimeValue)

	if !result {
		return NewFailedResult(r.Translate(r.Locale, "validation.after_or_equal", map[string]string{
			"field":      selector,
			"otherField": r.otherField,
		}))
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *AfterOrEqual) AddParams(params []string) {
	r.otherField = params[0]
	r.timeZone = time.UTC

	if len(params) > 1 {
		if tz, err := time.LoadLocation(params[1]); err != nil {
			r.timeZone = tz
		}
	}
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule accept 2 parameter, the first one, `otherField`, is mandatory
// and the second one, `timeZoneString` is optional.
func (r *AfterOrEqual) MinRequiredParams() int {
	return 1
}
