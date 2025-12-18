package rules

import (
	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// NotIn checks whether the field under validation not be in one of given values.
//
// Usage: "notIn:value1,value2[,value3,...]".
// Example: "notIn:admin,superuser".
type NotIn struct {
	translation.BaseTranslatableRule
	values []string
}

// Validate checks if the value of the field under validation not be in one of given values.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *NotIn) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	strValue := cast.ToString(value)
	for _, val := range r.values {
		if strValue == val {
			return NewFailedResult(r.Translate(r.Locale, "validation.not_in", map[string]string{
				"field": selector,
			}))
		}
	}

	return NewSuccessResult()
}

// AddParams assigns the provided parameter values to the NotIn rule instance.
// The first parameter specifies the `value1` to compare against (required),
// and the second parameter, if provided, specifies the `value2` to compare against (optional),
// and so on.
func (r *NotIn) AddParams(params []string) {
	r.values = params
}

// MinRequiredParams returns the minimum number of required parameters for the NotIn rule.
// It specifies how many parameters must be provided when configuring this rule.
// Returns 2, indicating that the `value1` and `value2` parameters are mandatory.
func (*NotIn) MinRequiredParams() int {
	return 2
}

// RequiresField returns false as the NotIn rule does not require the field to exist.
func (*NotIn) RequiresField() bool {
	return false
}
