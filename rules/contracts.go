package rules

import "github.com/behzadsh/go.validator/bag"

type Result struct {
	valid   bool
	message string
}

func NewResult(valid bool, message string) Result {
	return Result{valid: valid, message: message}
}

func NewFailedResult(message string) Result {
	return Result{valid: false, message: message}
}

func NewSuccessResult() Result {
	return Result{valid: true, message: ""}
}

func (r Result) Message() string {
	return r.message
}

func (r Result) Failed() bool {
	return !r.valid
}

type Rule interface {
	Validate(selector string, value any, inputBag bag.InputBag) Result
}

type RuleWithParams interface {
	Rule
	AddParams(params []string)
	MinRequiredParams() int
}
