package validation

import (
	"fmt"
	"strings"

	"github.com/spf13/cast"
	"github.com/thoas/go-funk"

	"github.com/behzadsh/go.validator/translation"
)

type ruleIndicator string

func (r ruleIndicator) load(locale string) Rule {
	name, params := r.parseRuleParams()
	rule, ok := registry[name]
	if !ok || rule == nil {
		panic(fmt.Errorf("rule %s is not registered", name))
	}

	if ruleWithParams, ok := rule.(RuleWithParams); ok {
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
