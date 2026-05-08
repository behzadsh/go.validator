package rules

import (
	"strings"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// EndsWith checks whether the field under validation ends with the given substring.
//
// Usage: "endsWith:suffix".
// Example: "endsWith:Model".
type EndsWith struct {
	translation.BaseTranslatableRule
	suffix string
}

// Validate checks if the value of the field under validation ends with the given substring.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *EndsWith) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	if !strings.HasSuffix(cast.ToString(value), r.suffix) {
		return NewFailedResult(r.Translate(r.Locale, "validation.ends_with", map[string]string{
			"field": selector,
			"value": r.suffix,
		}))
	}

	return NewSuccessResult()
}

// AddParams assigns the provided parameter values to the EndsWith rule instance.
// The first parameter specifies the `suffix` to compare against (required).
func (r *EndsWith) AddParams(params []string) {
	r.suffix = params[0]
}

// MinRequiredParams returns the minimum number of required parameters for the EndsWith rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 1, indicating that the `suffix` parameter is mandatory.
func (*EndsWith) MinRequiredParams() int {
	return 1
}
