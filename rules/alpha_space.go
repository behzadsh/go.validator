package rules

import (
	"regexp"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// AlphaSpace checks whether the field under validation contains only alphabetic characters and spaces.
// This rule accepts no parameters.
//
// Usage: "alphaSpace".
type AlphaSpace struct {
	translation.BaseTranslatableRule
}

// Validate checks if the value of the field under validation contains only alphabetic characters and spaces.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *AlphaSpace) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	ok, err := regexp.MatchString(`^[\pL\pM\s]+$`, cast.ToString(value))
	if !ok || err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.alpha_space", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}
