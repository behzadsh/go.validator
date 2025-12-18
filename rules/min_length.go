package rules

import (
	"reflect"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// MinLength checks whether the field under validation length be greater than given length.
//
// Usage: "minLength:value".
// Example: "minLength:10".
type MinLength struct {
	translation.BaseTranslatableRule
	minLength int
}

// Validate checks if the value of the field under validation length be greater than given length.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *MinLength) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
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

	if length < r.minLength {
		return NewFailedResult(r.Translate(r.Locale, "validation.min_length", map[string]string{
			"field": selector,
			"value": cast.ToString(r.minLength),
		}))
	}

	return NewSuccessResult()
}

// AddParams assigns the provided parameter values to the MinLength rule instance.
// The first parameter specifies the `value` to compare against (required).
func (r *MinLength) AddParams(params []string) {
	r.minLength = cast.ToInt(params[0])
}

// MinRequiredParams returns the minimum number of required parameters for the MinLength rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 1, indicating that the `value` parameter is mandatory.
func (*MinLength) MinRequiredParams() int {
	return 1
}

// RequiresField returns false as the MinLength rule does not require the field to exist.
func (*MinLength) RequiresField() bool {
	return false
}
