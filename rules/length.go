package rules

import (
	"reflect"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Length checks the field under validation have exact given length.
//
// Usage: "length:value".
// Example: "length:10".
type Length struct {
	translation.BaseTranslatableRule
	expectedLength int
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *Length) Validate(selector string, value any, _ bag.InputBag) Result {
	typeOf := reflect.TypeOf(value)
	if typeOf == nil {
		return NewSuccessResult()
	}

	kind := typeOf.Kind()

	var length int
	switch kind {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
		length = reflect.ValueOf(value).Len()
	default:
		// ignore the rule if not match any of the specified types.
		return NewSuccessResult()
	}

	if length != r.expectedLength {
		return NewFailedResult(r.Translate(r.Locale, "validation.length", map[string]string{
			"field": selector,
			"value": cast.ToString(r.expectedLength),
		}))
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *Length) AddParams(params []string) {
	r.expectedLength = cast.ToInt(params[0])
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule need only one parameter that is the `length`.
func (*Length) MinRequiredParams() int {
	return 1
}
