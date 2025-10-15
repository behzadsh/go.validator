package rules

import (
	"strings"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// RequiredWithAll checks whether the field under validation must exist if all given fields exist.
//
// Usage: "requiredWithAll:otherField,anotherField[,...]".
// Example: "requiredWithAll:email,username".
type RequiredWithAll struct {
	translation.BaseTranslatableRule
	otherFields []string
}

// Validate checks if the value of the field under validation must exist if all given fields exist.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *RequiredWithAll) Validate(selector string, _ any, inputBag bag.InputBag) ValidationResult {
	exists := inputBag.Has(selector)

	shouldExists := true
	for _, field := range r.otherFields {
		if !inputBag.Has(field) {
			shouldExists = false
			break
		}
	}

	if shouldExists && !exists {
		return NewFailedResult(r.Translate(r.Locale, "validation.required_with_all", map[string]string{
			"field":       selector,
			"otherFields": r.getOtherFieldsConcatenated(),
		}))
	}

	return NewSuccessResult()
}

// AddParams sets the list of field names that this rule will check for presence,
// by assigning the given parameter values to the RequiredWithAll rule instance.
func (r *RequiredWithAll) AddParams(params []string) {
	r.otherFields = params
}

// MinRequiredParams returns the minimum number of required parameters for the RequiredWithAll rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 2, indicating that at least two other fields are required.
func (*RequiredWithAll) MinRequiredParams() int {
	return 2
}

func (r *RequiredWithAll) getOtherFieldsConcatenated() string {
	if len(r.otherFields) == 2 {
		return strings.Join(r.otherFields, " and ")
	}

	output := strings.Join(r.otherFields, ", ")
	i := strings.LastIndex(output, ",")
	return output[:i+1] + " and" + output[i+1:]
}
