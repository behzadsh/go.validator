package rules

import (
	"regexp"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// NotRegex check the field under validation does not match the given regex pattern.
//
// Usage: "notRegex:pattern"
// Example: "notRegex:[a-zA-Z0-9]+"
type NotRegex struct {
	translation.BaseTranslatableRule
	pattern string
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *NotRegex) Validate(selector string, value any, inputBag bag.InputBag) Result {
	strValue, err := cast.ToStringE(value)
	if err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.string", map[string]string{
			"field": selector,
		}))
	}

	ok, err := regexp.MatchString(r.pattern, strValue)
	if ok {
		return NewFailedResult(r.Translate(r.Locale, "validation.not_regex", map[string]string{
			"field":   selector,
			"pattern": r.pattern,
		}))
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *NotRegex) AddParams(params []string) {
	r.pattern = params[0]
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule needs only one parameter and that is the regex pattern.
func (r *NotRegex) MinRequiredParams() int {
	return 1
}
