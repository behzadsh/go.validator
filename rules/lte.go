package rules

import (
	"reflect"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// LessThanEqual checks the field under validation has a value or a length
// less than the given value.
//
// Usage: "lte:value"
// Example: "lte:18"
type LessThanEqual struct {
	translation.BaseTranslatableRule
	value float64
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *LessThanEqual) Validate(selector string, value any, inputBag bag.InputBag) Result {
	kind := reflect.TypeOf(value).Kind()

	var floatValue float64
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		floatValue = cast.ToFloat64(value)
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
		floatValue = float64(reflect.ValueOf(value).Len())
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

// AddParams adds rules parameter values to the rule instance.
func (r *LessThanEqual) AddParams(params []string) {
	r.value = cast.ToFloat64(params[0])
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule needs only 1 parameter that is the `value`.
func (r *LessThanEqual) MinRequiredParams() int {
	return 1
}