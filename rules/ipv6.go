package rules

import (
	"net"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// IPv6 checks the field under validation is a valid IPv6 address.
//
// Usage: "ipv6".
type IPv6 struct {
	translation.BaseTranslatableRule
}

// Validate does the validation process of the rule.
func (r *IPv6) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	s := cast.ToString(value)
	ip := net.ParseIP(s)
	if ip == nil || ip.To4() != nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.ipv6", map[string]string{
			"field": selector,
		}))
	}
	return NewSuccessResult()
}
