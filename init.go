package validation

import "github.com/behzadsh/go.validator/rules"

var defaultLocale string

var stopOnFirstFailure bool

var registry map[string]rules.Rule

func init() {
	// initiate with default rules
	registerDefaultRules()

	defaultLocale = "en"
}

// SetDefaultLocale sets the default locale for validation error translations.
func SetDefaultLocale(locale string) {
	defaultLocale = locale
}

// StopOnFirstFailure sets an option to stop validation process when the first
// validation occurs.
func StopOnFirstFailure() {
	stopOnFirstFailure = true
}

// Register registers a rule object with given rule name in rule registry.
func Register(ruleName string, rule rules.Rule) {
	registry[ruleName] = rule
}

func registerDefaultRules() {
	registry = map[string]rules.Rule{
		"after":              &rules.After{},
		"afterOrEqual":       &rules.AfterOrEqual{},
		"alpha":              &rules.Alpha{},
		"alphaDash":          &rules.AlphaDash{},
		"alphaNum":           &rules.AlphaNum{},
		"alphaSpace":         &rules.AlphaSpace{},
		"array":              &rules.Array{},
		"before":             &rules.Before{},
		"beforeOrEqual":      &rules.BeforeOrEqual{},
		"between":            &rules.Between{},
		"boolean":            &rules.Boolean{},
		"dateTime":           &rules.DateTime{},
		"dateTimeAfter":      &rules.DateTimeAfter{},
		"dateTimeBefore":     &rules.DateTimeBefore{},
		"dateTimeBetween":    &rules.DateTimeBetween{},
		"different":          &rules.Different{},
		"digits":             &rules.Digits{},
		"digitsBetween":      &rules.DigitsBetween{},
		"distinct":           &rules.Distinct{},
		"email":              &rules.Email{},
		"endsWith":           &rules.EndsWith{},
		"gt":                 &rules.GreaterThan{},
		"gte":                &rules.GreaterThanEqual{},
		"inArrayField":       &rules.InArrayField{},
		"in":                 &rules.In{},
		"integer":            &rules.Integer{},
		"ip":                 &rules.IP{},
		"ipv4":               &rules.IPv4{},
		"ipv6":               &rules.IPv6{},
		"length":             &rules.Length{},
		"lowercase":          &rules.Lowercase{},
		"lt":                 &rules.LessThan{},
		"lte":                &rules.LessThanEqual{},
		"macAddress":         &rules.MacAddress{},
		"max":                &rules.Max{},
		"maxDigits":          &rules.MaxDigits{},
		"maxLength":          &rules.MaxLength{},
		"min":                &rules.Min{},
		"minDigits":          &rules.MinDigits{},
		"minLength":          &rules.MinLength{},
		"neq":                &rules.NotEqual{},
		"notEmpty":           &rules.NotEmpty{},
		"notIn":              &rules.NotIn{},
		"notRegex":           &rules.NotRegex{},
		"numeric":            &rules.Numeric{},
		"regex":              &rules.Regex{},
		"required":           &rules.Required{},
		"requiredIf":         &rules.RequiredIf{},
		"requiredUnless":     &rules.RequiredUnless{},
		"requiredWith":       &rules.RequiredWith{},
		"requiredWithAll":    &rules.RequiredWithAll{},
		"requiredWithout":    &rules.RequiredWithout{},
		"requiredWithoutAll": &rules.RequiredWithoutAll{},
		"sameAs":             &rules.SameAs{},
		"startsWith":         &rules.StartsWith{},
		"string":             &rules.String{},
		"timezone":           &rules.Timezone{},
		"uppercase":          &rules.Uppercase{},
		"url":                &rules.URL{},
		"uuid":               &rules.UUID{},
	}
}
