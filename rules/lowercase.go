package rules

import (
	"strings"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Lowercase checks the field under validation be lowercase string.
//
// Usage: "lowercase".
type Lowercase struct {
	translation.BaseTranslatableRule
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *Lowercase) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	strValue := cast.ToString(value)

	if strValue != strings.ToLower(strValue) {
		return NewFailedResult(r.Translate(r.Locale, "validation.lowercase", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}
