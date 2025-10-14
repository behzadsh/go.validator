package rules

import (
	"net/url"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// URL checks the field under validation is a valid URL with scheme and host.
//
// Usage: "url".
type URL struct {
	translation.BaseTranslatableRule
	requireScheme bool
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
func (r *URL) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	raw := cast.ToString(value)
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

// AddParams adds rules parameter values to the rule instance.
// Optional parameters:
// - "scheme": require scheme presence in the URL.
func (r *URL) AddParams(params []string) {
	for _, p := range params {
		if p == "scheme" {
			r.requireScheme = true
		}
	}
}

// MinRequiredParams returns minimum parameter requirement for this rule.
func (*URL) MinRequiredParams() int { return 0 }
