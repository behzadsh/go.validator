package rules

import (
	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Required checks the field under validation exists.
//
// Usage: "required".
type Required struct {
	translation.BaseTranslatableRule
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *Required) Validate(selector string, _ any, inputBag bag.InputBag) Result {
	if !inputBag.Has(selector) {
		return NewFailedResult(r.Translate(r.Locale, "validation.required", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}
