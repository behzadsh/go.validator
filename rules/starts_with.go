package rules

import (
	"strings"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// StartsWith check the field under validation starts with given sub string.
//
// Usage: "startsWith:prefix".
// Example: "startsWith:Model".
type StartsWith struct {
	translation.BaseTranslatableRule
	prefix string
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *StartsWith) Validate(selector string, value any, _ bag.InputBag) Result {
	strValue, err := cast.ToStringE(value)
	if err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.string", map[string]string{
			"field": selector,
		}))
	}

	if !strings.HasPrefix(strValue, r.prefix) {
		return NewFailedResult(r.Translate(r.Locale, "validation.starts_with", map[string]string{
			"field": selector,
			"value": r.prefix,
		}))
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *StartsWith) AddParams(params []string) {
	r.prefix = params[0]
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule needs only one parameter and that is the prefix.
func (*StartsWith) MinRequiredParams() int {
	return 1
}
