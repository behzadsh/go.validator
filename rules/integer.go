package rules

import (
	"reflect"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Integer checks the field under validation has a integer value
//
// Usage: "integer".
type Integer struct {
	translation.BaseTranslatableRule
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *Integer) Validate(selector string, value any, _ bag.InputBag) Result {
	typeOf := reflect.TypeOf(value)
	if typeOf == nil {
		return NewSuccessResult()
	}

	kind := typeOf.Kind()

	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return NewSuccessResult()
	default:
		return NewFailedResult(r.Translate(r.Locale, "validation.integer", map[string]string{
			"field": selector,
		}))
	}
}
