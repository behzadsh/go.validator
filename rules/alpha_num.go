package rules

import (
	"regexp"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// AlphaNum checks whether the field under validation contains only alphanumeric characters.
// This rule accepts no parameters.
//
// Usage: "alphaNum".
type AlphaNum struct {
	translation.BaseTranslatableRule
}

// Validate checks if the value of the field under validation contains only alphanumeric characters.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *AlphaNum) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	ok, err := regexp.MatchString(`^[\pL\pM\pN]+$`, cast.ToString(value))
	if !ok || err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.alpha_num", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}

// RequiresField returns false as the AlphaNum rule does not require the field to exist.
func (*AlphaNum) RequiresField() bool {
	return false
}
