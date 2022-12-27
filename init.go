package validation

import "github.com/behzadsh/go.validator/rules"

var defaultLocale string

var stopOnFirstFailure bool

var registry map[string]Rule

func init() {
	// initiate with default rules
	registry = map[string]Rule{
		"after":     &rules.After{},
		"required":  &rules.Required{},
		"not_empty": &rules.NotEmpty{},
	}

	defaultLocale = "en"
}

func SetDefaultLocale(locale string) {
	defaultLocale = locale
}

func StopOnFirstFailure() {
	stopOnFirstFailure = true
}
