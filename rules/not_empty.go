package rules

import (
	"github.com/thoas/go-funk"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// NotEmpty checks the field under validation be a non-empty or non-zero value.
//
// Usage: "notEmpty".
type NotEmpty struct {
	translation.BaseTranslatableRule
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *NotEmpty) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	if funk.IsEmpty(value) {
		return NewFailedResult(r.Translate(r.Locale, "validation.not_empty", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}
