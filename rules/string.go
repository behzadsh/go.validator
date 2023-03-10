package rules

import (
	"reflect"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// String checks the field under validation has a string value
//
// Usage: "string".
type String struct {
	translation.BaseTranslatableRule
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
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
