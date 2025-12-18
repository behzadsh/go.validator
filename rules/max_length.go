package rules

import (
	"reflect"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// MaxLength checks whether the field under validation length did not reach given length.
//
// Usage: "maxLength:value".
// Example: "maxLength:10".
type MaxLength struct {
	translation.BaseTranslatableRule
	maxLength int
}

// Validate checks if the value of the field under validation length did not reach given length.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *MaxLength) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
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

	if length > r.maxLength {
		return NewFailedResult(r.Translate(r.Locale, "validation.max_length", map[string]string{
			"field": selector,
			"value": cast.ToString(r.maxLength),
		}))
	}

	return NewSuccessResult()
}

// AddParams assigns the provided parameter values to the MaxLength rule instance.
// The first parameter specifies the `value` to compare against (required).
func (r *MaxLength) AddParams(params []string) {
	r.maxLength = cast.ToInt(params[0])
}

// MinRequiredParams returns the minimum number of required parameters for the MaxLength rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 1, indicating that the `value` parameter is mandatory.
func (*MaxLength) MinRequiredParams() int {
	return 1
}

// RequiresField returns false as the MaxLength rule does not require the field to exist.
func (*MaxLength) RequiresField() bool {
	return false
}
