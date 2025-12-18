package rules

import (
	"reflect"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Integer checks whether the field under validation has an integer value.
// This rule accepts no parameters.
//
// Usage: "integer".
type Integer struct {
	translation.BaseTranslatableRule
}

// Validate checks if the value of the field under validation has an integer value.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *Integer) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	if value == nil {
		return NewSuccessResult()
	}
	v := indirectValue(value)

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return NewSuccessResult()
	default:
		return NewFailedResult(r.Translate(r.Locale, "validation.integer", map[string]string{
			"field": selector,
		}))
	}
}

// RequiresField returns false as the Integer rule does not require the field to exist.
func (*Integer) RequiresField() bool {
	return false
}
