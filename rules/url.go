package rules

import (
	"net/url"

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

	u, err := url.ParseRequestURI(raw)
	if r.requireScheme {
		if err != nil || u.Scheme == "" || u.Host == "" {
			return NewFailedResult(r.Translate(r.Locale, "validation.url", map[string]string{
				"field": selector,
			}))
		}
	} else {
		// Accept URLs without scheme by attempting to parse with an implied scheme
		if err != nil || u.Host == "" {
			u2, err2 := url.Parse("http://" + raw)
			if err2 != nil || u2.Host == "" {
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
