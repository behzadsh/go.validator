package validation

var defaultLocale string

var stopOnFirstFailure bool

func init() {
	// initiate with default rules
	registry = map[string]Rule{
		//
	}

	defaultLocale = "en"
}

func SetDefaultLocale(locale string) {
	defaultLocale = locale
}

func StopOnFirstFailure() {
	stopOnFirstFailure = true
}
