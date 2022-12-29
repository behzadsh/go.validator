package rules

import (
	"reflect"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Between checks the field under validation have a value or a length between
// the given min and max parameters. If the field under validation has a
// numeric value, the value must be between the given min and max. If the field
// under validation is string, slice or map, the length of it will be evaluated.
//
// Usage: "between:min,max"
// Example: "between:2,5"
type Between struct {
	translation.BaseTranslatableRule
	min, max float64
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *Between) Validate(selector string, value any, _ bag.InputBag) Result {
	typeOf := reflect.TypeOf(value)
	if typeOf == nil {
		return NewSuccessResult()
	}

	kind := typeOf.Kind()

	var v float64
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		v = cast.ToFloat64(value)
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
		v = float64(reflect.ValueOf(value).Len())
	default:
		// ignore the rule if not match any of the specified types.
		return NewSuccessResult()
	}

	if v < r.min || v > r.max {
		return NewFailedResult(r.Translate(r.Locale, "validation.between", map[string]string{
			"field": selector,
			"min":   cast.ToString(r.min),
			"max":   cast.ToString(r.max),
		}))
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *Between) AddParams(params []string) {
	r.min = cast.ToFloat64(params[0])
	r.max = cast.ToFloat64(params[1])
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule accept 2 parameter, the first one is the minimum number and
// the second one is the maximum number of the between range. both of the
// parameters are mandatory.
func (r *Between) MinRequiredParams() int {
	return 2
}
