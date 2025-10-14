package rules

import (
	"net"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// IPv4 checks the field under validation is a valid IPv4 address.
//
// Usage: "ipv4".
type IPv4 struct {
	translation.BaseTranslatableRule
}

// Validate does the validation process of the rule.
func (r *IPv4) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	s := cast.ToString(value)
	ip := net.ParseIP(s)
	if ip == nil || ip.To4() == nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.ipv4", map[string]string{
			"field": selector,
		}))
	}
	return NewSuccessResult()
}
