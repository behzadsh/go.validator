package rules

import (
	"regexp"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// AlphaSpace checks the field under validation have alphabetic characters
// and spaces.
// This rule accept no parameters.
//
// Usage: "alphaSpace"
type AlphaSpace struct {
	translation.BaseTranslatableRule
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *AlphaSpace) Validate(selector string, value any, _ bag.InputBag, _ bool) Result {
	strValue, err := cast.ToStringE(value)
	if err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.alpha_space", map[string]string{
			"field": selector,
		}))
	}

	ok, err := regexp.MatchString(`^[\pL\pM\s]+$`, strValue)
	if !ok || err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.alpha_space", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}
