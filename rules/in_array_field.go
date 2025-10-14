package rules

import (
	"strings"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// InArrayField checks the field under validation exists in another array/slice field.
//
// Usage: "inArrayField:otherField".
// Example: value selected from enum slice field.
type InArrayField struct {
	translation.BaseTranslatableRule
	otherField string
}

// Validate does the validation process of the rule.
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

// AddParams adds rules parameter values to the rule instance.
func (r *InArrayField) AddParams(params []string) {
	r.otherField = params[0]
}

// MinRequiredParams returns minimum parameter requirement for this rule.
func (*InArrayField) MinRequiredParams() int { return 1 }

func toComparableString(v any) string {
	s, ok := v.(string)
	if ok {
		return s
	}
	return indirectValue(v).String()
}
