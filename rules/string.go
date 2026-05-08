package rules

import (
	"reflect"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// String checks whether the field under validation has a string value
// This rule accepts no parameters.
//
// Usage: "string".
type String struct {
	translation.BaseTranslatableRule
}

// Validate checks if the value of the field under validation has a string value.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *String) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	if value == nil {
		return NewSuccessResult()
	}
	v := indirectValue(value)

	switch v.Kind() {
	case reflect.String:
		return NewSuccessResult()
	default:
		return NewFailedResult(r.Translate(r.Locale, "validation.string", map[string]string{
			"field": selector,
		}))
	}
}
