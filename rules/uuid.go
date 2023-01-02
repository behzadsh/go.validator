package rules

import (
	"regexp"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

const uuidRegex = `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`

// UUID checks field under validation is a valid uuid.
//
// Usage: "uuid".
type UUID struct {
	translation.BaseTranslatableRule
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *UUID) Validate(selector string, value any, _ bag.InputBag) Result {
	strVal := cast.ToString(value)

	ok, err := regexp.MatchString(uuidRegex, strVal)
	if !ok || err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.uuid", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}
