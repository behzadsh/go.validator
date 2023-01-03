package rules

import (
	"strings"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// RequiredWithoutAll check the field under validation exists if all given
// fields not exist.
//
// Usage: "requiredWithout:otherField,anotherField[,...]".
// example: "requiredWithout:phone,username".
type RequiredWithoutAll struct {
	translation.BaseTranslatableRule
	otherFields []string
}

// Validate does the validation process of the rule. See struct documentation
// for more details.
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

// AddParams adds rules parameter values to the rule instance.
func (r *RequiredWithoutAll) AddParams(params []string) {
	r.otherFields = params
}

// MinRequiredParams returns minimum parameter requirement for this rule.
// This rule needs at least 2 parameters which represent the name of other
// field or fields.
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
