package rules

import (
	"strings"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// InArrayField checks whether the field under validation exists in another array/slice field.
//
// Usage: "inArrayField:otherField".
// Example: value selected from enum slice field.
type InArrayField struct {
	translation.BaseTranslatableRule
	otherField string
}

// Validate checks if the value of the field under validation exists in another array/slice field.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *InArrayField) Validate(selector string, value any, inputBag bag.InputBag) ValidationResult {
	other, ok := inputBag.Get(r.otherField)
	if !ok {
		return NewFailedResult(r.Translate(r.Locale, "Validation.required", map[string]string{
			"otherField": r.otherField,
		}))
	}

	v := indirectValue(other)
	if v.Len() == 0 {
		return NewFailedResult(r.Translate(r.Locale, "validation.in", map[string]string{
			"field": selector,
		}))
	}

	strVal := toComparableString(value)
	for i := 0; i < v.Len(); i++ {
		if strings.EqualFold(strVal, toComparableString(v.Index(i).Interface())) {
			return NewSuccessResult()
		}
	}

	return NewFailedResult(r.Translate(r.Locale, "validation.in", map[string]string{
		"field": selector,
	}))
}

// AddParams assigns the provided parameter values to the InArrayField rule instance.
// The first parameter specifies the `otherField` to compare against (required).
func (r *InArrayField) AddParams(params []string) {
	r.otherField = params[0]
}

// MinRequiredParams returns the minimum number of required parameters for the InArrayField rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 1, indicating that the `otherField` parameter is mandatory.
func (*InArrayField) MinRequiredParams() int { return 1 }

// RequiresField returns false as the InArrayField rule does not require the field to exist.
func (*InArrayField) RequiresField() bool {
	return false
}

func toComparableString(v any) string {
	s, ok := v.(string)
	if ok {
		return s
	}
	return indirectValue(v).String()
}
