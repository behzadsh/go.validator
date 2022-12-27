package validation

import "github.com/behzadsh/go.validator/bag"

type Rule interface {
	Validate(selector string, value any, inputBag bag.InputBag, exists bool) bool
	Message() string
}

type RuleWithParams interface {
	Rule
	AddParams(params []string)
	MinRequiredParams() int
}

func Register(ruleName string, rule Rule) {
	registry[ruleName] = rule
}
