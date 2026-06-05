package validation

import "reflect"

// Distinct is a Rule that validates the value is a slice or array with no duplicate elements.
//
// Non-slice/array values pass (the rule is irrelevant for scalars). Elements whose type is
// not comparable (e.g. slices, maps) are skipped rather than causing a panic.
//
// Fails if:
//   - the value is a slice/array and contains at least two equal comparable elements
//
// Examples:
//
//	validation.Distinct.Validate([]string{"a", "b", "c"}) // pass
//	validation.Distinct.Validate([]int{1, 2, 1})          // fail — duplicate 1
//	validation.Distinct.Validate("not-a-slice")           // pass — rule irrelevant
//	validation.Distinct.Validate(nil)                     // pass — rule irrelevant
var Distinct Rule = RuleFunc(
	func(value any) error {
		if value == nil {
			return nil
		}

		rv := reflect.ValueOf(value)
		if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
			return nil
		}

		seen := make(map[any]struct{})
		for i := 0; i < rv.Len(); i++ {
			elem := rv.Index(i)
			if !elem.Type().Comparable() {
				continue
			}

			key := elem.Interface()
			if _, exists := seen[key]; exists {
				return ErrValidationDistinct
			}

			seen[key] = struct{}{}
		}

		return nil
	},
)
