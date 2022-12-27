package rules

import (
	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Boolean checks the field under validation is boolean or can be cast as
// a boolean value.
//
// Usage: "boolean"
type Boolean struct {
	translation.BaseTranslatableRule
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *Boolean) Validate(selector string, value any, inputBag bag.InputBag, exists bool) Result {
	_, err := cast.ToBoolE(value)
	if err != nil || value == nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.boolean", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}
