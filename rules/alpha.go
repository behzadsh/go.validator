package rules

import (
	"regexp"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Alpha checks the field under validation be entirely alphabetic characters.
// This rule accept no parameters.
//
// Usage: "alpha"
type Alpha struct {
	translation.BaseTranslatableRule
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *Alpha) Validate(selector string, value any, _ bag.InputBag) Result {
	strValue, err := cast.ToStringE(value)
	if err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.alpha", map[string]string{
			"field": selector,
		}))
	}

	ok, err := regexp.MatchString(`^[\pL\pM]+$`, strValue)
	if !ok || err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.alpha", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}
