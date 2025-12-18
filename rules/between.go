package rules

import (
	"reflect"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Between checks whether the field under validation has a value or a length between the given min and max parameters.
// It will return a validation error if the field under validation has a numeric value, the value must be between the
// given min and max. If the field under validation is string, slice or map, the length of it will be evaluated.
//
// Usage: "between:min,max".
// Example: "between:2,5".
type Between struct {
	translation.BaseTranslatableRule
	min, max float64
}

// Validate checks if the value of the field under validation is between the given min and max parameters.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *Between) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	if value == nil {
		return NewSuccessResult()
	}
	v := indirectValue(value)

	var floatValue float64
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		floatValue = cast.ToFloat64(value)
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
		floatValue = float64(v.Len())
	default:
		// ignore the rule if not match any of the specified types.
		return NewSuccessResult()
	}

	if floatValue < r.min || floatValue > r.max {
		return NewFailedResult(r.Translate(r.Locale, "validation.between", map[string]string{
			"field": selector,
			"min":   cast.ToString(r.min),
			"max":   cast.ToString(r.max),
		}))
	}

	return NewSuccessResult()
}

// AddParams sets the minimum and maximum values for the Between rule based on the provided parameters.
// It expects exactly two parameters: the first for the minimum boundary (min),
// and the second for the maximum boundary (max). The extracted values are converted to float64 and stored.
// Returns nothing.
func (r *Between) AddParams(params []string) {
	r.min = cast.ToFloat64(params[0])
	r.max = cast.ToFloat64(params[1])
}

// MinRequiredParams returns the minimum number of required parameters for the Between rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 2, indicating that the `min` and `max` parameters are mandatory.
func (*Between) MinRequiredParams() int {
	return 2
}

// RequiresField returns false as the Between rule does not require the field to exist.
func (*Between) RequiresField() bool {
	return false
}
