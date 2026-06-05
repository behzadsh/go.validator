package validation

import "slices"

// In returns a Rule that validates the value is present in the given slice.
//
// Comparison uses ==, so T must be a comparable type. The value must be exactly f type T; a float64 does not match an
// int slice even if numerically equal.
//
// Fails if:
//   - value cannot be type-asserted to T
//   - value is not found in slice
//
// Examples:
//
//	validation.In([]string{"a", "b", "c"}).Validate("a")  // pass
//	validation.In([]string{"a", "b", "c"}).Validate("d")  // fail — not in list
//	validation.In([]int{1, 2, 3}).Validate(2)             // pass
//	validation.In([]int{1, 2, 3}).Validate(float64(2))    // fail — wrong type
func In[T comparable](slice []T) Rule {
	fn := func(value any) error {
		v, ok := value.(T)
		if !ok || !slices.Contains(slice, v) {
			return inError{Values: slice}
		}

		return nil
	}

	return RuleFunc(fn)
}

// NEQ returns a Rule that validates the value is not equal to v.
//
// Comparison is type-sensitive: a float64(1) does not equal an int(1).
//
// Fails if:
//   - value cannot be type-asserted to T
//   - value == v
//
// Examples:
//
//	validation.NEQ[string]("admin").Validate("user")  // pass
//	validation.NEQ[string]("admin").Validate("admin") // fail
//	validation.NEQ[int](0).Validate(1)               // pass
//	validation.NEQ[int](0).Validate(0)               // fail
func NEQ[T comparable](v T) Rule {
	return RuleFunc(
		func(value any) error {
			actual, ok := value.(T)
			if !ok || actual == v {
				return neqError{Value: v}
			}

			return nil
		},
	)
}

// NotIn returns a Rule that validates the value is not present in the given slice.
//
// Like In, comparison uses ==. The value must be exactly of type T.
//
// Fails if:
//   - value cannot be type-asserted to T
//   - value is found in slice
//
// Examples:
//
//	validation.NotIn([]string{"banned", "blocked"}).Validate("allowed") // pass
//	validation.NotIn([]string{"banned", "blocked"}).Validate("banned")  // fail — in list
//	validation.NotIn([]int{0, -1}).Validate(5)                          // pass
//	validation.NotIn([]int{0, -1}).Validate(0)                          // fail — in list
func NotIn[T comparable](slice []T) Rule {
	fn := func(value any) error {
		v, ok := value.(T)
		if !ok || slices.Contains(slice, v) {
			return notInError{Values: slice}
		}

		return nil
	}

	return RuleFunc(fn)
}
