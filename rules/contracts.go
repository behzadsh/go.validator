package rules

import (
	"github.com/behzadsh/go.validator/bag"
)

// ValidationResult is a struct that represents a single validation result.
type ValidationResult struct {
	valid   bool
	message string
}

// NewResult generates a ValidationResult instance with the given values.
func NewResult(valid bool, message string) ValidationResult {
	return ValidationResult{valid: valid, message: message}
}

// NewFailedResult generates a ValidationResult instance as failed validation with the given error message.
func NewFailedResult(message string) ValidationResult {
	return ValidationResult{valid: false, message: message}
}

// NewSuccessResult generates a ValidationResult instance as a success result with an empty message.
func NewSuccessResult() ValidationResult {
	return ValidationResult{valid: true, message: ""}
}

// Message returns the message stored in the ValidationResult instance.
// It returns the message associated with the validation result.
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
	RequiresField() bool
}

// RuleWithParams is an interface for rules that have parameters.
type RuleWithParams interface {
	Rule
	AddParams(params []string)
	MinRequiredParams() int
}
