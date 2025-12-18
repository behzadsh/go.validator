package rules

import (
	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// SameAs checks whether the field under validation has the value same as the other given field.
//
// Usage: "sameAs:otherField".
// Example: "sameAs:password".
type SameAs struct {
	translation.BaseTranslatableRule
	otherField string
}

// Validate checks if the value of the field under validation has the value same as the other given field.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *SameAs) Validate(selector string, value any, inputBag bag.InputBag) ValidationResult {
	otherValue, _ := inputBag.Get(r.otherField)

	if otherValue != value {
		return NewFailedResult(r.Translate(r.Locale, "validation.same_as", map[string]string{
			"field":      selector,
			"otherField": r.otherField,
		}))
	}

	return NewSuccessResult()
}

// AddParams assigns the provided parameter values to the SameAs rule instance.
// The first parameter specifies the `otherField` to compare against (required).
func (r *SameAs) AddParams(params []string) {
	r.otherField = params[0]
}

// MinRequiredParams returns the minimum number of required parameters for the SameAs rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 1, indicating that the `otherField` parameter is mandatory.
func (*SameAs) MinRequiredParams() int {
	return 1
}

// RequiresField returns false as the SameAs rule does not require the field to exist.
func (*SameAs) RequiresField() bool {
	return false
}
