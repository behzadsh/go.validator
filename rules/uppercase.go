package rules

import (
	"strings"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Uppercase checks the field under validation be uppercase string.
//
// Usage: "uppercase"
type Uppercase struct {
	translation.BaseTranslatableRule
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *Uppercase) Validate(selector string, value any, inputBag bag.InputBag) Result {
	strValue := cast.ToString(value)

	if strValue != strings.ToUpper(strValue) {
		return NewFailedResult(r.Translate(r.Locale, "validation.uppercase", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}
