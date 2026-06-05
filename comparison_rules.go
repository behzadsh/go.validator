package validation

// SameAs returns an InputRule that validates the value is equal to the value at the given field path.
//
// Comparison uses ==, so both value and type must match (e.g. string "1" != int 1).
// If the referenced field is absent, validation fails.
//
// Fails if:
//   - the referenced field is absent
//   - value != the referenced field's value
//
// Examples:
//
//	schema := validation.New().
//		Field("password_confirm", validation.Required, validation.SameAs("password"))
func SameAs(path string) InputRule {
	return InputRuleFunc(
		func(value any, input *InputBag) error {
			other, found := input.Lookup(path)
			if !found || value != other {
				return sameAsError{Field: path}
			}

			return nil
		},
	)
}

// Different returns an InputRule that validates the value is not equal to the value at the given field path.
//
// Comparison uses ==, so both value and type are considered.
// If the referenced field is absent, the rule passes (no value to compare against).
//
// Fails if:
//   - value == the referenced field's value
//
// Examples:
//
//	schema := validation.New().
//		Field("new_password", validation.Required, validation.Different("old_password"))
func Different(path string) InputRule {
	return InputRuleFunc(
		func(value any, input *InputBag) error {
			other, found := input.Lookup(path)
			if found && value == other {
				return differentError{Field: path}
			}

			return nil
		},
	)
}
