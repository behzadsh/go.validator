package rules

import (
	"strings"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// RequiredWithoutAll checks whether the field under validation must exist if all given fields not exist.
//
// Usage: "requiredWithoutAll:otherField,anotherField[,...]".
// Example: "requiredWithoutAll:phone,username".
type RequiredWithoutAll struct {
	translation.BaseTranslatableRule
	otherFields []string
}

// Validate checks if the value of the field under validation must exist if all given fields not exist.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *RequiredWithoutAll) Validate(selector string, _ any, inputBag bag.InputBag) ValidationResult {
	exists := inputBag.Has(selector)

	shouldExists := true
	for _, field := range r.otherFields {
		if inputBag.Has(field) {
			shouldExists = false
			break
		}
	}

	if shouldExists && !exists {
		return NewFailedResult(r.Translate(r.Locale, "validation.required_without_all", map[string]string{
			"field":       selector,
			"otherFields": r.getOtherFieldsConcatenated(),
		}))
	}

	return NewSuccessResult()
}

// AddParams sets the list of field names that this rule will check for absence,
// by assigning the given parameter values to the RequiredWithoutAll rule instance.
func (r *RequiredWithoutAll) AddParams(params []string) {
	r.otherFields = params
}

// MinRequiredParams returns the minimum number of required parameters for the RequiredWithoutAll rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 2, indicating that at least two other fields are required.
func (*RequiredWithoutAll) MinRequiredParams() int {
	return 2
}

func (r *RequiredWithoutAll) getOtherFieldsConcatenated() string {
	if len(r.otherFields) == 2 {
		return strings.Join(r.otherFields, " and ")
	}

	output := strings.Join(r.otherFields, ", ")
	i := strings.LastIndex(output, ",")
	return output[:i+1] + " and" + output[i+1:]
}
