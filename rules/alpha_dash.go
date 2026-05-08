package rules

import (
	"regexp"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// AlphaDash checks whether the field under validation contains only alphanumeric characters, dashes, and underscores.
// This rule accepts no parameters.
//
// Usage: "alphaDash".
type AlphaDash struct {
	translation.BaseTranslatableRule
}

// Validate checks if the value of the field under validation contains only alphanumeric characters, dashes, and underscores.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *AlphaDash) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	ok, err := regexp.MatchString(`^[\pL\pM\pN_-]+$`, cast.ToString(value))
	if !ok || err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.alpha_dash", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}
