package rules

import (
	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Numeric checks the field under validation has a numeric value
//
// Usage: "numeric".
type Numeric struct {
	translation.BaseTranslatableRule
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *Numeric) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	_, err := cast.ToFloat64E(value)
	if err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.numeric", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}
