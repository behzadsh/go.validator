package rules

import (
	"regexp"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Regex checks whether the field under validation match the given regex pattern.
//
// Usage: "regex:pattern".
// Example: "regex:[a-zA-Z0-9]+".
type Regex struct {
	translation.BaseTranslatableRule
	pattern string
}

// Validate checks if the value of the field under validation match the given regex pattern.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
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

// AddParams assigns the provided parameter values to the Regex rule instance.
// The first parameter specifies the `pattern` to compare against (required).
func (r *Regex) AddParams(params []string) {
	r.pattern = params[0]
}

// MinRequiredParams returns the minimum number of required parameters for the Regex rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 1, indicating that the `pattern` parameter is mandatory.
func (*Regex) MinRequiredParams() int {
	return 1
}

// RequiresField returns false as the Regex rule does not require the field to exist.
func (*Regex) RequiresField() bool {
	return false
}
