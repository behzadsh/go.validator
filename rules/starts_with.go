package rules

import (
	"strings"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// StartsWith checks whether the field under validation starts with given sub string.
//
// Usage: "startsWith:prefix".
// Example: "startsWith:Model".
type StartsWith struct {
	translation.BaseTranslatableRule
	prefix string
}

// Validate checks if the value of the field under validation starts with given sub string.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *StartsWith) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	if !strings.HasPrefix(cast.ToString(value), r.prefix) {
		return NewFailedResult(r.Translate(r.Locale, "validation.starts_with", map[string]string{
			"field": selector,
			"value": r.prefix,
		}))
	}

	return NewSuccessResult()
}

// AddParams assigns the provided parameter values to the StartsWith rule instance.
// The first parameter specifies the `prefix` to compare against (required).
func (r *StartsWith) AddParams(params []string) {
	r.prefix = params[0]
}

// MinRequiredParams returns the minimum number of required parameters for the StartsWith rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 1, indicating that the `prefix` parameter is mandatory.
func (*StartsWith) MinRequiredParams() int {
	return 1
}
