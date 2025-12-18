package rules

import (
	"net"
	"net/url"
	"strings"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// URL checks whether the field under validation is a valid URL with scheme and host.
//
// Usage: "url[:scheme]".
// Example: "url".
// Example: "url:scheme".
type URL struct {
	translation.BaseTranslatableRule
	requireScheme bool
}

// Validate checks if the value of the field under validation is a valid URL with scheme and host.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *URL) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	raw := cast.ToString(value)
	if raw == "" {
		return NewSuccessResult()
	}
	// First, try to parse the URL as-is.
	u, err := url.ParseRequestURI(raw)
	if r.requireScheme {
		if err != nil || u.Scheme == "" || !isValidURLHost(u.Host) {
			return NewFailedResult(r.Translate(r.Locale, "validation.url", map[string]string{
				"field": selector,
			}))
		}
	} else {
		// Accept URLs without scheme by attempting to parse with an implied scheme
		if err != nil || !isValidURLHost(u.Host) {
			u2, err2 := url.Parse("http://" + raw)
			if err2 != nil || !isValidURLHost(u2.Host) {
				return NewFailedResult(r.Translate(r.Locale, "validation.url", map[string]string{
					"field": selector,
				}))
			}
		}
	}

	return NewSuccessResult()
}

// AddParams assigns the provided parameter values to the URL rule instance.
// The first parameter specifies the `scheme` to compare against (optional).
// The second parameter, if provided, specifies the `scheme` to compare against (optional).
func (r *URL) AddParams(params []string) {
	for _, p := range params {
		if p == "scheme" {
			r.requireScheme = true
		}
	}
}

// MinRequiredParams returns the minimum number of required parameters for the URL rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 0, indicating that no parameters are required.
func (*URL) MinRequiredParams() int { return 0 }

// isValidURLHost performs additional checks to make sure the host part of the URL
// looks like a real network host and not just arbitrary text.
//
// It accepts:
// - domain-like hosts with at least one dot (e.g. "example.com")
// - "localhost"
// - valid IP addresses (IPv4 / IPv6, with or without port)
func isValidURLHost(host string) bool {
	if host == "" {
		return false
	}

	// Strip port if present (e.g. "example.com:8080" or "[::1]:8080")
	if strings.HasPrefix(host, "[") {
		// For IPv6 in brackets, keep the address inside brackets, drop the port.
		if idx := strings.LastIndex(host, "]"); idx != -1 {
			host = host[1:idx]
		}
	} else if h, _, ok := strings.Cut(host, ":"); ok {
		host = h
	}

	// Accept valid IPs.
	if net.ParseIP(host) != nil {
		return true
	}

	// Accept localhost explicitly.
	if host == "localhost" {
		return true
	}

	// For domain-like hosts, require at least one dot and no empty labels.
	if !strings.Contains(host, ".") {
		return false
	}

	parts := strings.Split(host, ".")
	for _, p := range parts {
		if p == "" {
			return false
		}
	}

	return true
}

// RequiresField returns false as the URL rule does not require the field to exist.
func (*URL) RequiresField() bool {
	return false
}
