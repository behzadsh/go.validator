package rules

import (
	"net"
	"regexp"
	"strings"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

const emailRegexPattern = "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?" +
	"(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

// Email checks the field under validation is a valid email address based on
// RFC 53222. There is also an optional mx record check, you can enable by
// passing `mx` as parameter.
//
// Usage: "email[:mx].
// Example: "email".
// Example: "email:mx".
type Email struct {
	translation.BaseTranslatableRule
	enableMXCheck bool
	localPart     string
	domainPart    string
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *Email) Validate(selector string, value any, _ bag.InputBag) Result {
	strValue := cast.ToString(value)

	if !r.isEmail(strValue) {
		return NewFailedResult(r.Translate(r.Locale, "validation.email", map[string]string{
			"field": selector,
		}))
	}

	return NewSuccessResult()
}

// AddParams adds rules parameter values to the rule instance.
func (r *Email) AddParams(params []string) {
	for _, param := range params {
		if param == "mx" {
			r.enableMXCheck = true
			return
		}
	}
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule accept `mx` as the only and optional parameter. By passing this
// parameter, an extra step will be added to validation process. The `mx`
// validation checks the email domain for mx record.
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
