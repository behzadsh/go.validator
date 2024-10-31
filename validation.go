package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/cast"
	"github.com/thoas/go-funk"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/rules"
	"github.com/behzadsh/go.validator/translation"
)

type ruleIndicator string

// RulesMap is a custom type for a map of rules for field selectors.
type RulesMap map[string][]string

func (r ruleIndicator) load(locale string) rules.Rule {
	name, params := r.parseRuleParams()
	rule, ok := registry[name]
	if !ok || rule == nil {
		panic(fmt.Errorf("rule %s is not registered", name))
	}

	if ruleWithParams, ok := rule.(rules.RuleWithParams); ok {
		paramNum := len(params)
		minRequiredParams := ruleWithParams.MinRequiredParams()

		if paramNum < minRequiredParams {
			panic(fmt.Errorf("rule %s need at least %d parameter, got %d", name, minRequiredParams, paramNum))
		}

		ruleWithParams.AddParams(params)
	}

	if translatableRule, ok := rule.(translation.TranslatableRule); ok {
		translatableRule.AddLocale(locale)
		translatableRule.AddTranslationFunction(translation.GetDefaultTranslatorFunc())
	}

	return rule
}

func (r ruleIndicator) parseRuleParams() (string, []string) {
	parts := strings.SplitN(string(r), ":", 2)
	if len(parts) == 1 {
		return parts[0], nil
	}

	return parts[0], cast.ToStringSlice(funk.Map(strings.Split(parts[1], ","), func(s string) string {
		return strings.TrimSpace(s)
	}))
}

// ValidateMap validate the given input map with given rules map.
func ValidateMap(input map[string]any, rulesMap RulesMap, locale ...string) Result {
	currentLocale := defaultLocale
	if len(locale) > 0 {
		currentLocale = locale[0]
	}

	inputBag := bag.InputBag(input)

	return doValidation(inputBag, rulesMap, currentLocale)
}

// ValidateMapSlice validates the given slice of maps with given rules map.
func ValidateMapSlice(input []map[string]any, rulesMap RulesMap, locale ...string) Result {
	currentLocale := defaultLocale
	if len(locale) > 0 {
		currentLocale = locale[0]
	}

	result := NewResult()
	for i, m := range input {
		inputBag := bag.InputBag(m)
		tmpResult := doValidation(inputBag, rulesMap, currentLocale)
		for key, messages := range tmpResult.Errors {
			result.addError(fmt.Sprintf("%d.%s", i, key), messages...)
		}
	}

	return result
}

// ValidateStruct validates the given struct with given rules map.
func ValidateStruct(input any, rulesMap RulesMap, locale ...string) Result {
	currentLocale := defaultLocale
	if len(locale) > 0 {
		currentLocale = locale[0]
	}

	v := reflect.ValueOf(input)
	if v.Kind() != reflect.Struct && !(v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct) {
		panic("validation.ValidateStruct only support struct or a pointer to a struct as first parameter")
	}

	inputBag := bag.NewInputBagFromStruct(input)

	return doValidation(inputBag, rulesMap, currentLocale)
}

// ValidateStructSlice validates the given slice of struct with given rules map.
func ValidateStructSlice(input []any, rulesMap RulesMap, locale ...string) Result {
	currentLocale := defaultLocale
	if len(locale) > 0 {
		currentLocale = locale[0]
	}

	result := NewResult()
	for i, strct := range input {
		tmpResult := ValidateStruct(strct, rulesMap, currentLocale)
		for key, messages := range tmpResult.Errors {
			result.addError(fmt.Sprintf("%d.%s", i, key), messages...)
		}
	}

	return result
}

// Validate validate the given input with given validation rules.
func Validate(input any, ruleSlice []string, locale ...string) Result {
	currentLocale := defaultLocale
	if len(locale) > 0 {
		currentLocale = locale[0]
	}

	return doValidation(bag.InputBag{"variable": input}, RulesMap{"variable": ruleSlice}, currentLocale)
}

func doValidation(inputBag bag.InputBag, rulesMap RulesMap, locale string) Result {
	explicitRules := make(RulesMap)

	for fieldSelector, fieldRules := range rulesMap {
		for _, explicitFieldSelector := range normalizeFieldSelector(fieldSelector, inputBag) {
			explicitRules[explicitFieldSelector] = fieldRules
		}
	}

	result := NewResult()
	for selector, selectorRules := range explicitRules {
		val, _ := inputBag.Get(selector)
		for _, ruleStr := range selectorRules {
			ruleName := ruleIndicator(ruleStr)
			rule := ruleName.load(locale)
			ruleResult := rule.Validate(selector, val, inputBag)
			if ruleResult.Failed() {
				result.addError(selector, ruleResult.Message())
				if stopOnFirstFailure {
					break
				}
			}
		}
	}

	return result
}

func normalizeFieldSelector(selector string, input bag.InputBag) []string {
	if !strings.Contains(selector, ".*") {
		return []string{selector}
	}

	parts := strings.SplitN(selector, ".*", 2)
	total := 0
	val, ok := input.Get(parts[0])
	if ok {
		temp, err := cast.ToSliceE(val)
		if err == nil {
			total = len(temp)
		}
	}

	var explicitSelectors []string
	for i := 0; i < total; i++ {
		key := fmt.Sprintf("%s.%d%s", parts[0], i, parts[1])
		explicitSelectors = append(explicitSelectors, normalizeFieldSelector(key, input)...)
	}

	return explicitSelectors
}
