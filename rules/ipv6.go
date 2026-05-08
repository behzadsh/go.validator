package rules

import (
	"net"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// IPv6 checks whether the field under validation is a valid IPv6 address.
// This rule accepts no parameters.
//
// Usage: "ipv6".
type IPv6 struct {
	translation.BaseTranslatableRule
}

// Validate checks if the value of the field under validation is a valid IPv6 address.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
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
