package rules

import (
	"regexp"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// AlphaNum checks the field under validation has alphanumeric characters.
// This rule accept no parameters.
//
// Usage: "alphaNum".
type AlphaNum struct {
	translation.BaseTranslatableRule
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *AlphaNum) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	strValue, ok := value.(string)
	if !ok {
		return NewFailedResult(r.Translate(r.Locale, "validation.alpha_num", map[string]string{
			"field": selector,
		}))
	}

	ok, err := regexp.MatchString(`^[\pL\pM\pN]+$`, strValue)
	if !ok || err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.alpha_num", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}
