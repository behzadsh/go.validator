package rules

import (
	"net"
	"regexp"
	"strings"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

const emailRegexPattern = "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?" + "(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

// Email checks if the field under validation is a valid email address based on RFC 5322. There is also an optional MX
// record check that can be enabled by passing `mx` as a parameter.
//
// Usage: "email[:mx]".
// Example: "email".
// Example: "email:mx".
type Email struct {
	translation.BaseTranslatableRule
	enableMXCheck bool
	localPart     string
	domainPart    string
}

// Validate checks if the value of the field under validation is a valid email address based on RFC 5322.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *Email) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	if !r.isEmail(cast.ToString(value)) {
		return NewFailedResult(r.Translate(r.Locale, "validation.email", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}

// AddParams assigns the provided parameter values to the Email rule instance.
// The first parameter specifies the `mx` to compare against (optional).
func (r *Email) AddParams(params []string) {
	for _, param := range params {
		if param == "mx" {
			r.enableMXCheck = true
			return
		}
	}
}

// MinRequiredParams returns the minimum number of required parameters for the Email rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 0, indicating that the `mx` parameter is optional.
func (*Email) MinRequiredParams() int {
	return 0
}

func (r *Email) isEmail(emailAddress string) bool {
	if !r.checkFormat(emailAddress) {
		return false
	}

	if r.enableMXCheck {
		return r.checkMXRecord()
	}

	return true
}

func (r *Email) checkFormat(emailAddress string) bool {
	parts := strings.Split(emailAddress, "@")
	if len(parts) != 2 || len(parts[0]) > 64 || len(parts[1]) > 255 {
		return false
	}

	r.localPart = parts[0]
	r.domainPart = parts[1]

	ok, err := regexp.MatchString(emailRegexPattern, emailAddress)
	if !ok || err != nil {
		return false
	}

	return true
}

func (r *Email) checkMXRecord() bool {
	if _, err := net.LookupMX(r.domainPart); err != nil {
		return false
	}

	return true
}
