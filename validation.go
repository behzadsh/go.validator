package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
)

type RulesMap map[string][]string

func ValidateMap(input map[string]any, rules RulesMap, locale ...string) Result {
	currentLocale := defaultLocale
	if len(locale) > 0 {
		currentLocale = locale[0]
	}

	inputBag := bag.InputBag(input)

	return doValidation(inputBag, rules, currentLocale)
}

func ValidateMapSlice(input []map[string]any, rules RulesMap, locale ...string) Result {
	currentLocale := defaultLocale
	if len(locale) > 0 {
		currentLocale = locale[0]
	}

	result := NewResult()
	for i, m := range input {
		inputBag := bag.InputBag(m)
		tmpResult := doValidation(inputBag, rules, currentLocale)
		for key, messages := range tmpResult.Errors {
			result.addError(fmt.Sprintf("%d.%s", i, key), messages...)
		}
	}

	return result
}

func ValidateStruct(input any, rules RulesMap, locale ...string) Result {
	currentLocale := defaultLocale
	if len(locale) > 0 {
		currentLocale = locale[0]
	}

	v := reflect.ValueOf(input)
	if v.Kind() != reflect.Struct || (v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct) {
		panic("validation.ValidateStruct only support struct or a pointer to a struct as first parameter")
	}

	inputBag := bag.NewInputBagFromStruct(input)

	return doValidation(inputBag, rules, currentLocale)
}

func ValidateStructSlice(input []any, rules RulesMap, locale ...string) Result {
	currentLocale := defaultLocale
	if len(locale) > 0 {
		currentLocale = locale[0]
	}

	result := NewResult()
	for i, strct := range input {
		tmpResult := ValidateStruct(strct, rules, currentLocale)
		for key, messages := range tmpResult.Errors {
			result.addError(fmt.Sprintf("%d.%s", i, key), messages...)
		}
	}

	return result
}

func Validate(input any, rules []string, locale ...string) Result {
	currentLocale := defaultLocale
	if len(locale) > 0 {
		currentLocale = locale[0]
	}

	return doValidation(bag.InputBag{"variable": input}, RulesMap{"variable": rules}, currentLocale)
}

func doValidation(inputBag bag.InputBag, rules RulesMap, locale string) Result {
	explicitRules := make(RulesMap)

	for fieldSelector, fieldRules := range rules {
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
	if !strings.Contains(selector, "*") {
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

	if len(parts) == 1 {
		parts = append(parts, "")
	}

	var explicitSelectors []string
	for i := 0; i < total; i++ {
		key := fmt.Sprintf("%s.%d%s", parts[0], i, parts[1])
		explicitSelectors = append(explicitSelectors, normalizeFieldSelector(key, input)...)
	}

	return explicitSelectors
}
