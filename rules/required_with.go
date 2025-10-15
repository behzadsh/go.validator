package rules

import (
	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// RequiredWith checks whether the field under validation must exist if any of given fields exist.
//
// Usage: "requiredWith:otherField[,anotherField,...]".
// Example: "requiredWith:type".
type RequiredWith struct {
	translation.BaseTranslatableRule
	otherFields []string
}

// Validate checks if the value of the field under validation must exist if any of given fields exist.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *RequiredWith) Validate(selector string, _ any, inputBag bag.InputBag) ValidationResult {
	exists := inputBag.Has(selector)

	if exists {
		return NewSuccessResult()
	}

	for _, field := range r.otherFields {
		if inputBag.Has(field) {
			return NewFailedResult(r.Translate(r.Locale, "validation.required_with", map[string]string{
				"field":      selector,
				"otherField": field,
			}))
		}
	}

	return NewSuccessResult()
}

// AddParams sets the list of field names that this rule will check for presence,
// by assigning the given parameter values to the RequiredWith rule instance.
func (r *RequiredWith) AddParams(params []string) {
	r.otherFields = params
}

// MinRequiredParams returns the minimum number of required parameters for the RequiredWith rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 1, indicating that at least one other field is required.
func (*RequiredWith) MinRequiredParams() int {
	return 1
}
