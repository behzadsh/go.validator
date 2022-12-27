package validation

import "github.com/behzadsh/go.validator/rules"

var defaultLocale string

var stopOnFirstFailure bool

var registry map[string]rules.Rule

func init() {
	// initiate with default rules
	registry = map[string]rules.Rule{
		"after":        &rules.After{},
		"afterOrEqual": &rules.AfterOrEqual{},
		// "required":  &rules.Required{},
		// "not_empty": &rules.NotEmpty{},
	}

	defaultLocale = "en"
}

func SetDefaultLocale(locale string) {
	defaultLocale = locale
}

func StopOnFirstFailure() {
	stopOnFirstFailure = true
}

func Register(ruleName string, rule rules.Rule) {
	registry[ruleName] = rule
}
