package rules

import (
	"github.com/thoas/go-funk"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Array checks the field under validation is an array or slice.
// This rule accept no parameters.
//
// Usage: "array".
type Array struct {
	translation.BaseTranslatableRule
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *Array) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	if !funk.IsCollection(value) {
		return NewFailedResult(r.Translate(r.Locale, "validation.array", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}
