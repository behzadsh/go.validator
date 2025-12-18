package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
)

// RulesMap is a custom type for a map of rules for field selectors.
type RulesMap map[string][]string

// ValidateMap validates each field in the provided input map according to the specified set of rules in rulesMap.
// It returns a Result containing any validation errors found. The optional locale parameter specifies the language
// used for error messages; if not provided, the default locale is used.
func ValidateMap(input map[string]any, rulesMap RulesMap, locale ...string) Result {
	currentLocale := defaultLocale
	if len(locale) > 0 {
		currentLocale = locale[0]
	}

	inputBag := bag.InputBag(input)

	return doValidation(inputBag, rulesMap, currentLocale)
}

// ValidateMapSlice iterates over a slice of input maps, validating each map against the specified rulesMap.
// It returns a Result containing accumulated validation errors for all maps, indexed by their position in the slice.
// The optional locale parameter determines the language for error messages; if not set, the default locale is used.
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

// ValidateStruct validates each field in the provided struct according to the specified set of rules in rulesMap.
// It returns a Result containing any validation errors found. The optional locale parameter specifies the language
// used for error messages; if not provided, the default locale is used.
func ValidateStruct(input any, rulesMap RulesMap, locale ...string) Result {
	currentLocale := defaultLocale
	if len(locale) > 0 {
		currentLocale = locale[0]
	}

	v := reflect.ValueOf(input)
	if v.Kind() != reflect.Struct && (v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct) {
		panic("validation.ValidateStruct only support struct or a pointer to a struct as first parameter")
	}

	inputBag := bag.NewInputBagFromStruct(input)

	return doValidation(inputBag, rulesMap, currentLocale)
}

// ValidateStructSlice iterates over a slice of structs, validating each struct against the specified rulesMap.
// It returns a Result containing accumulated validation errors for all structs, indexed by their position in the slice.
// The optional locale parameter determines the language for error messages; if not set, the default locale is used.
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

// Validate validates the provided input according to the specified set of rules in ruleSlice.
// It returns a Result containing any validation errors found. The optional locale parameter specifies the language
// used for error messages; if not provided, the default locale is used.
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
		val, exists := inputBag.Get(selector)
		for _, ruleStr := range selectorRules {
			ruleName := ruleIndicator(ruleStr)
			rule := ruleName.load(locale)

			if !exists && !rule.RequiresField() {
				continue
			}

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
