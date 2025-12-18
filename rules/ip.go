package rules

import (
	"net"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// IP checks whether the field under validation is a valid IP address (v4 or v6).
// This rule accepts no parameters.
//
// Usage: "ip".
type IP struct {
	translation.BaseTranslatableRule
}

// Validate checks if the value of the field under validation is a valid IP address (v4 or v6).
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *IP) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	s := cast.ToString(value)
	if net.ParseIP(s) == nil {
		return NewFailedResult(r.Translate(r.Locale, "validation.ip", map[string]string{
			"field": selector,
		}))
	}
	return NewSuccessResult()
}

// RequiresField returns false as the IP rule does not require the field to exist.
func (*IP) RequiresField() bool {
	return false
}
