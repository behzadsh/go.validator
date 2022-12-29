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

func SetDefaultLocale(locale string) {
	defaultLocale = locale
}

func StopOnFirstFailure() {
	stopOnFirstFailure = true
}

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
		"different":          &rules.Different{},
		"digits":             &rules.Digits{},
		"digitsBetween":      &rules.DigitsBetween{},
		"email":              &rules.Email{},
		"endsWith":           &rules.EndsWith{},
		"gt":                 &rules.GreaterThan{},
		"gte":                &rules.GreaterThanEqual{},
		"in":                 &rules.In{},
		"integer":            &rules.Integer{},
		"length":             &rules.Length{},
		"lowercase":          &rules.Lowercase{},
		"lt":                 &rules.LessThan{},
		"lte":                &rules.LessThanEqual{},
		"max":                &rules.Max{},
		"maxDigits":          &rules.MaxDigits{},
		"maxLength":          &rules.MaxLength{},
		"min":                &rules.Min{},
		"minDigits":          &rules.MinDigits{},
		"minLength":          &rules.MinLength{},
		"neq":                &rules.NotEqual{},
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
		"uppercase":          &rules.Uppercase{},
		"uuid":               &rules.UUID{},
	}
}
