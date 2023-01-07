package rules

import (
	"regexp"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Regex check the field under validation match the given regex pattern.
//
// Usage: "regex:pattern".
// Example: "regex:[a-zA-Z0-9]+".
type Regex struct {
	translation.BaseTranslatableRule
	pattern string
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *Regex) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	ok, err := regexp.MatchString(r.pattern, cast.ToString(value))
	if !ok || err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.regex", map[string]string{
			"field":   selector,
			"pattern": r.pattern,
		}))
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *Regex) AddParams(params []string) {
	r.pattern = params[0]
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule needs only one parameter and that is the regex pattern.
func (*Regex) MinRequiredParams() int {
	return 1
}
