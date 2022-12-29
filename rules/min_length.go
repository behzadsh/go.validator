package rules

import (
	"reflect"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// MinLength checks the field under validation length be greater than given length.
//
// Usage: "minLength:value"
// Example: "minLength:10"
type MinLength struct {
	translation.BaseTranslatableRule
	minLength int
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *MinLength) Validate(selector string, value any, inputBag bag.InputBag) Result {
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

	if length < r.minLength {
		return NewFailedResult(r.Translate(r.Locale, "validation.min_length", map[string]string{
			"field": selector,
			"value": cast.ToString(r.minLength),
		}))
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *MinLength) AddParams(params []string) {
	r.minLength = cast.ToInt(params[0])
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule need only one parameter that is the `length`.
func (r *MinLength) MinRequiredParams() int {
	return 1
}
