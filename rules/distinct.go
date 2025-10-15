package rules

import (
	"reflect"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

// Distinct checks whether the field under validation is an array/slice with all unique elements.
// This rule accepts no parameters.
//
// Usage: "distinct".
type Distinct struct {
	translation.BaseTranslatableRule
}

// Validate checks if the value of the field under validation is an array/slice with all unique elements.
// It returns a ValidationResult that indicates success if valid, or the appropriate error message if the check fails.
func (r *Distinct) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
	v := indirectValue(value)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return NewSuccessResult()
	}

	seen := make(map[any]struct{})
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i).Interface()
		// Only comparable keys can be used in a map; fallback to string for others
		key := elem
		if !v.Index(i).CanInterface() {
			continue
		}
		if v.Index(i).CanAddr() && !v.Index(i).Type().Comparable() {
			key = v.Index(i).Interface()
		}
		// Use reflect.Value for non-comparable types by converting to string
		if _, ok := key.(interface{ String() string }); !ok && !v.Index(i).Type().Comparable() {
			key = v.Index(i).String()
		}

		if _, ok := seen[key]; ok {
			return NewFailedResult(r.Translate(r.Locale, "validation.distinct", map[string]string{
				"field": selector,
			}))
		}
		seen[key] = struct{}{}
	}

	return NewSuccessResult()
}
