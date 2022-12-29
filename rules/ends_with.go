package rules

import (
	"strings"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// EndsWith check the field under validation ends with given sub string.
//
// Usage: "endsWith:prefix"
// Example: "endsWith:Model"
type EndsWith struct {
	translation.BaseTranslatableRule
	suffix string
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *EndsWith) Validate(selector string, value any, inputBag bag.InputBag) Result {
	strValue, err := cast.ToStringE(value)
	if err != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.string", map[string]string{
			"field": selector,
		}))
	}

	if !strings.HasSuffix(strValue, r.suffix) {
		return NewFailedResult(r.Translate(r.Locale, "validation.ends_with", map[string]string{
			"field": selector,
			"value": r.suffix,
		}))
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *EndsWith) AddParams(params []string) {
	r.suffix = params[0]
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule needs only one parameter and that is the prefix.
func (r *EndsWith) MinRequiredParams() int {
	return 1
}
