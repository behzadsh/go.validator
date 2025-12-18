package rules

import (
	"reflect"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// LessThanEqual checks whether the field under validation has a value or a length less than the given value.
//
// Usage: "lte:value".
// Example: "lte:18".
type LessThanEqual struct {
	translation.BaseTranslatableRule
	value float64
}

// Validate checks if the value of the field under validation has a value or a length less than the given value.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *LessThanEqual) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
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

	if floatValue > r.value {
		return NewFailedResult(r.Translate(r.Locale, "validation.lte", map[string]string{
			"field": selector,
			"value": cast.ToString(r.value),
		}))
	}

	return NewSuccessResult()
}

// AddParams assigns the provided parameter values to the LessThanEqual rule instance.
// The first parameter specifies the `value` to compare against (required).
func (r *LessThanEqual) AddParams(params []string) {
	r.value = cast.ToFloat64(params[0])
}

// MinRequiredParams returns the minimum number of required parameters for the LessThanEqual rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 1, indicating that the `value` parameter is mandatory.
func (*LessThanEqual) MinRequiredParams() int {
	return 1
}

// RequiresField returns false as the LessThanEqual rule does not require the field to exist.
func (*LessThanEqual) RequiresField() bool {
	return false
}
