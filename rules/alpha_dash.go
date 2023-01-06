package rules

import (
	"regexp"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// AlphaDash checks the field under validation has alphanumeric characters,
// as well as dashes and underscore.
// This rule accept no parameters.
//
// Usage: "alphaDash".
type AlphaDash struct {
	translation.BaseTranslatableRule
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *AlphaDash) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	ok, err := regexp.MatchString(`^[\pL\pM\pN_-]+$`, cast.ToString(value))
	if !ok || err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.alpha_dash", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}
