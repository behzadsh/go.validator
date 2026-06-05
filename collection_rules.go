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
				return basicError{"distinct", "distinct validation failed"}
			}

			seen[key] = struct{}{}
		}

		return nil
	},
)

// Each returns a Rule that applies the given rules to every element of a slice or array.
//
// Non-slice/array values and nil pass (the rule is irrelevant for scalars). Validation stops at the first failing
// element and returns basicError{"each", "each validation failed"}; the index and inner error are not propagated.
//
// Fails if:
//   - any element fails any of the given rules
//
// Examples:
//
//	validation.Each(validation.MinLength(2)).Validate([]string{"ab", "cd"}) // pass
//	validation.Each(validation.MinLength(2)).Validate([]string{"ab", "x"})  // fail — "x" fails
//	validation.Each(validation.Positive).Validate([]int{1, 2, 3})           // pass
//	validation.Each(validation.Positive).Validate([]int{1, -1, 3})          // fail
func Each(rules ...Rule) Rule {
	return InputRuleFunc(
		func(value any, input *InputBag) error {
			if value == nil {
				return nil
			}

			rv := reflect.ValueOf(value)
			if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
				return nil
			}

			for i := 0; i < rv.Len(); i++ {
				elem := rv.Index(i).Interface()
				for _, r := range rules {
					if err := applyRule(r, elem, input); err != nil {
						return basicError{"each", "each validation failed"}
					}
				}
			}

			return nil
		},
	)
}

// MaxSize returns a Rule that validates a slice or array has at most n elements.
//
// Non-slice/array values and nil pass.
//
// Fails if:
//   - value is a slice/array with more than n elements
//
// Examples:
//
//	validation.MaxSize(3).Validate([]int{1, 2, 3})    // pass — exactly 3
//	validation.MaxSize(3).Validate([]int{1, 2, 3, 4}) // fail — 4 elements
//	validation.MaxSize(3).Validate(nil)                // pass
func MaxSize(n int) Rule {
	return RuleFunc(
		func(value any) error {
			if value == nil {
				return nil
			}

			rv := reflect.ValueOf(value)
			if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
				return nil
			}

			if rv.Len() > n {
				return maxSizeError{Size: n}
			}

			return nil
		},
	)
}

// MinSize returns a Rule that validates a slice or array has at least n elements.
//
// Non-slice/array values and nil pass.
//
// Fails if:
//   - value is a slice/array with fewer than n elements
//
// Examples:
//
//	validation.MinSize(2).Validate([]int{1, 2})    // pass — exactly 2
//	validation.MinSize(2).Validate([]int{1})        // fail — only 1 element
//	validation.MinSize(2).Validate(nil)             // pass
func MinSize(n int) Rule {
	return RuleFunc(
		func(value any) error {
			if value == nil {
				return nil
			}

			rv := reflect.ValueOf(value)
			if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
				return nil
			}

			if rv.Len() < n {
				return minSizeError{Size: n}
			}

			return nil
		},
	)
}

// Size returns a Rule that validates a slice or array has exactly n elements.
//
// Non-slice/array values and nil pass.
//
// Fails if:
//   - value is a slice/array whose length is not exactly n
//
// Examples:
//
//	validation.Size(3).Validate([]int{1, 2, 3}) // pass
//	validation.Size(3).Validate([]int{1, 2})    // fail — 2 elements
//	validation.Size(3).Validate(nil)             // pass
func Size(n int) Rule {
	return RuleFunc(
		func(value any) error {
			if value == nil {
				return nil
			}

			rv := reflect.ValueOf(value)
			if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
				return nil
			}

			if rv.Len() != n {
				return sizeError{Size: n}
			}

			return nil
		},
	)
}
