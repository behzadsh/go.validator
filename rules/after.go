package rules

import (
	"time"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// After checks whether the field under validation has a value that is after the value of the given field. It will
// return a validation error if one or both of the fields are not valid datetime strings. It will also return a
// validation error if the other field cannot be found in the input bag.
//
// Usage: "after:otherField[,timeZoneString]".
// Example: "after:start".
// Example: "after:start,America/New_York".
type After struct {
	translation.BaseTranslatableRule
	otherField string
	timeZone   *time.Location
}

// Validate checks if the value of the field under validation is a datetime string that occurs after the datetime value
// of another specified field. It returns a ValidationResult that indicates success if valid, or the appropriate error
// message if the check fails, the datetime formats are invalid, or the other field is missing.
func (r *After) Validate(selector string, value any, inputBag bag.InputBag) ValidationResult {
	timeValue, err := cast.ToTimeInDefaultLocationE(value, r.timeZone)
	if err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.datetime", map[string]string{
			"field": selector,
		}))
	}

	otherValue, ok := inputBag.Get(r.otherField)
	if !ok {
		return NewFailedResult(r.Translate(r.Locale, "Validation.required", map[string]string{
			"field": r.otherField,
		}))
	}

	otherTimeValue, err := cast.ToTimeInDefaultLocationE(otherValue, r.timeZone)
	if err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.datetime", map[string]string{
			"field": r.otherField,
		}))
	}

	result := timeValue.After(otherTimeValue)

	if !result {
		return NewFailedResult(r.Translate(r.Locale, "validation.after", map[string]string{
			"field":      selector,
			"otherField": r.otherField,
		}))
	}

	return NewSuccessResult()
}

// AddParams assigns the provided parameter values to the After rule instance.
// The first parameter specifies the `otherField` to compare against (required),
// and the second parameter, if provided, sets the time zone for parsing date/time values (optional).
func (r *After) AddParams(params []string) {
	r.otherField = params[0]
	r.timeZone = time.UTC

	if len(params) > 1 {
		if tz, err := time.LoadLocation(params[1]); err == nil {
			r.timeZone = tz
		}
	}
}

// MinRequiredParams returns the minimum number of required parameters for the After rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 1, indicating that the `otherField` parameter is mandatory, while the `timeZoneString` parameter is optional.
func (*After) MinRequiredParams() int {
	return 1
}
