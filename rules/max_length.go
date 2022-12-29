package rules

import (
	"reflect"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// MaxLength checks the field under validation length did not reach given length.
//
// Usage: "maxLength:value"
// Example: "maxLength:10"
type MaxLength struct {
	translation.BaseTranslatableRule
	maxLength int
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *MaxLength) Validate(selector string, value any, inputBag bag.InputBag) Result {
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

	if length > r.maxLength {
		return NewFailedResult(r.Translate(r.Locale, "validation.max_length", map[string]string{
			"field": selector,
			"value": cast.ToString(r.maxLength),
		}))
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *MaxLength) AddParams(params []string) {
	r.maxLength = cast.ToInt(params[0])
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule need only one parameter that is the `length`.
func (r *MaxLength) MinRequiredParams() int {
	return 1
}
