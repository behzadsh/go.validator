package rules

import (
	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// RequiredWithout checks whether the field under validation must exist if any of given fields doesn't exist.
//
// Usage: "requiredWithout:otherField[,anotherField,...]".
// Example: "requiredWithout:username".
type RequiredWithout struct {
	translation.BaseTranslatableRule
	otherFields []string
}

// Validate checks if the value of the field under validation must exist if any of given fields doesn't exist.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *RequiredWithout) Validate(selector string, _ any, inputBag bag.InputBag) ValidationResult {
	exists := inputBag.Has(selector)

	if exists {
		return NewSuccessResult()
	}

	for _, field := range r.otherFields {
		if !inputBag.Has(field) {
			return NewFailedResult(r.Translate(r.Locale, "validation.required_without", map[string]string{
				"field":      selector,
				"otherField": field,
			}))
		}
	}

	return NewSuccessResult()
}

// AddParams sets the list of field names that this rule will check for absence,
// by assigning the given parameter values to the RequiredWithout rule instance.
func (r *RequiredWithout) AddParams(params []string) {
	r.otherFields = params
}

// MinRequiredParams returns the minimum number of required parameters for the RequiredWithout rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 1, indicating that at least one other field is required.
func (*RequiredWithout) MinRequiredParams() int {
	return 1
}
