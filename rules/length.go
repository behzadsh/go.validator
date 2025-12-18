package rules

import (
	"reflect"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Length checks whether the field under validation has an exact given length.
//
// Usage: "length:value".
// Example: "length:10".
type Length struct {
	translation.BaseTranslatableRule
	expectedLength int
}

// Validate checks if the value of the field under validation has an exact given length.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *Length) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	if value == nil {
		return NewSuccessResult()
	}
	v := indirectValue(value)

	var length int
	switch v.Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
		length = v.Len()
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

// AddParams assigns the provided parameter values to the Length rule instance.
// The first parameter specifies the `value` to compare against (required).
func (r *Length) AddParams(params []string) {
	r.expectedLength = cast.ToInt(params[0])
}

// MinRequiredParams returns the minimum number of required parameters for the Length rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 1, indicating that the `value` parameter is mandatory.
func (*Length) MinRequiredParams() int {
	return 1
}

// RequiresField returns false as the Length rule does not require the field to exist.
func (*Length) RequiresField() bool {
	return false
}
