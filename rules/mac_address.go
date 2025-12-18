package rules

import (
	"regexp"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

var macRegex = regexp.MustCompile(`^(?i)([0-9A-F]{2}([-:])){5}([0-9A-F]{2})$`)

// MacAddress checks whether the field under validation is a valid MAC address.
// Supports formats like 01:23:45:67:89:ab or 01-23-45-67-89-ab.
// This rule accepts no parameters.
//
// Usage: "macAddress".
type MacAddress struct {
	translation.BaseTranslatableRule
}

// Validate checks if the value of the field under validation is a valid MAC address.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *MacAddress) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	s := cast.ToString(value)
	if !macRegex.MatchString(s) {
		return NewFailedResult(r.Translate(r.Locale, "validation.mac_address", map[string]string{
			"field": selector,
		}))
	}
	return NewSuccessResult()
}

// RequiresField returns false as the MacAddress rule does not require the field to exist.
func (*MacAddress) RequiresField() bool {
	return false
}
