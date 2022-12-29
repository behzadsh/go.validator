package rules

import (
	"regexp"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// AlphaNum checks the field under validation have alphanumeric characters.
// This rule accept no parameters.
//
// Usage: "alphaNum"
type AlphaNum struct {
	translation.BaseTranslatableRule
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *AlphaNum) Validate(selector string, value any, _ bag.InputBag) Result {
	strValue, err := cast.ToStringE(value)
	if err != nil {
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
