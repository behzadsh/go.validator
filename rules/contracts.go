package rules

import (
	"github.com/behzadsh/go.validator/bag"
)

// ValidationResult is a struct represents a single validation result.
type ValidationResult struct {
	valid   bool
	message string
}

// NewResult generates a ValidationResult instance with given values.
func NewResult(valid bool, message string) ValidationResult {
	return ValidationResult{valid: valid, message: message}
}

// NewFailedResult generates a ValidationResult instance as failed validation
// with given error message.
func NewFailedResult(message string) ValidationResult {
	return ValidationResult{valid: false, message: message}
}

// NewSuccessResult generates a ValidationResult instance as a success result.
func NewSuccessResult() ValidationResult {
	return ValidationResult{valid: true, message: ""}
}

// Message returns the message stored in the ValidationResult instance.
func (r ValidationResult) Message() string {
	return r.message
}

// Failed returns true if the validation has failed.
func (r ValidationResult) Failed() bool {
	return !r.valid
}

// Rule is an interface for rules.
type Rule interface {
	Validate(selector string, value any, inputBag bag.InputBag) ValidationResult
}

// RuleWithParams is an interface for rules that have parameters.
type RuleWithParams interface {
	Rule
	AddParams(params []string)
	MinRequiredParams() int
}
