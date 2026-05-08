package rules

import (
	"regexp"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

const uuidRegex = `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`

// UUID checks whether the field under validation is a valid uuid.
// This rule accepts no parameters.
//
// Usage: "uuid".
type UUID struct {
	translation.BaseTranslatableRule
}

// Validate checks if the value of the field under validation is a valid uuid.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *UUID) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	strVal := cast.ToString(value)

	ok, err := regexp.MatchString(uuidRegex, strVal)
	if !ok || err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.uuid", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}
