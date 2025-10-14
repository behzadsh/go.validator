package rules

import (
	"net"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// IP checks the field under validation is a valid IP address (v4 or v6).
//
// Usage: "ip".
type IP struct {
	translation.BaseTranslatableRule
}

// Validate does the validation process of the rule.
func (r *IP) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	s := cast.ToString(value)
	if net.ParseIP(s) == nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.ip", map[string]string{
			"field": selector,
		}))
	}
	return NewSuccessResult()
}
