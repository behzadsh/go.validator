package validation

import "reflect"

// Required is a Rule that validates the value exists, by checking value is not nil or empty string.
//
// Fails if:
//   - value is nil
//   - value is a string equal to ""
//
// Passes for any other type, including zero values such as 0 or false — use NotEmpty if those should also be rejected.
//
// Example:
//
//	schema := validation.New().
//		Field("name", validation.Required).
//		Field("email", validation.Required, validation.Email)
var Required Rule = presenceRuleFunc(
	func(value any) error {
		if value == nil {
			return basicError{"required", "required validation failed"}
		}

		s, ok := value.(string)
		if ok && s == "" {
			return basicError{"required", "required validation failed"}
		}

		return nil
	},
)

// RequiredIf returns a Rule that validates the value exists if the condition evaluated to true.
//
// Fails if:
//   - the condition is true AND value is nil
//   - the condition is true AND value is a string equal to ""
//
// Returns RuleSyntaxError if the condition string is malformed — treat that as a programming error and fix the schema
// at startup.
//
// The condition language supports:
//   - comparisons: ==  !=  <  >  <=  >=
//   - logical:     &&  ||  !
//   - grouping:    ( expr )
//   - functions:   exists(path), len(path)
//
// Field paths follow dot notation and can traverse nested maps and structs,
// e.g. "category.id", "order.shipping.country".
// String literals must be quoted ("admin" or 'admin').
// Unquoted identifiers are resolved as field paths in the input.
//
// Examples:
//
//	validation.RequiredIf(`plan == "paid"`)
//	validation.RequiredIf(`role == "admin" && plan != "free"`)
//	validation.RequiredIf(`exists(role) && len(tags) > 0`)
//	validation.RequiredIf(`(status == "active" || status == "pending") && verified == true`)
//	validation.RequiredIf(`category.id == 10`)
//	validation.RequiredIf(`exists(order.shipping.address) && order.shipping.country == "US"`)
func RequiredIf(condition string) InputRule {
	fn := func(value any, input *InputBag) error {
		ok, err := evalCondition(condition, input)
		if err != nil {
			return RuleSyntaxError{Rule: "RequiredIf", Err: err}
		}

		if ok {
			if value == nil {
				return basicError{"required_if", "required if validation failed"}
			}
			if s, isStr := value.(string); isStr && s == "" {
				return basicError{"required_if", "required if validation failed"}
			}
		}

		return nil
	}

	return presenceInputRuleFunc(fn)
}

// RequiredUnless returns an InputRule that validates the value exists unless the given condition evaluates to true.
//
// It is the logical complement of RequiredIf: the field is required when the condition is FALSE.
// The condition language is identical to RequiredIf (comparisons, &&, ||, !, exists(), len()).
//
// Returns RuleSyntaxError if the condition string is malformed.
//
// Examples:
//
//	validation.RequiredUnless(`type == "guest"`)
//	validation.RequiredUnless(`role == "admin"`)
func RequiredUnless(condition string) InputRule {
	fn := func(value any, input *InputBag) error {
		ok, err := evalCondition(condition, input)
		if err != nil {
			return RuleSyntaxError{Rule: "RequiredUnless", Err: err}
		}

		if !ok {
			if value == nil {
				return basicError{"required_unless", "required unless validation failed"}
			}
			if s, isStr := value.(string); isStr && s == "" {
				return basicError{"required_unless", "required unless validation failed"}
			}
		}

		return nil
	}

	return presenceInputRuleFunc(fn)
}

// RequiredWith returns an InputRule that validates the value exists if any of the given fields are present in the input.
//
// The field under validation is optional unless at least one of the listed fields is present.
//
// Examples:
//
//	validation.RequiredWith("phone", "mobile")
func RequiredWith(fields ...string) InputRule {
	return presenceInputRuleFunc(
		func(value any, input *InputBag) error {
			for _, f := range fields {
				if _, found := input.Lookup(f); found {
					if value == nil {
						return basicError{"required_with", "required with validation failed"}
					}
					if s, ok := value.(string); ok && s == "" {
						return basicError{"required_with", "required with validation failed"}
					}

					return nil
				}
			}

			return nil
		},
	)
}

// RequiredWithAll returns an InputRule that validates the value exists if all of the given fields are present in
// the input.
//
// The field under validation is optional unless every listed field is present.
//
// Examples:
//
//	validation.RequiredWithAll("first_name", "last_name")
func RequiredWithAll(fields ...string) InputRule {
	return presenceInputRuleFunc(
		func(value any, input *InputBag) error {
			for _, f := range fields {
				if _, found := input.Lookup(f); !found {
					return nil
				}
			}

			if value == nil {
				return basicError{"required_with_all", "required with all validation failed"}
			}
			if s, ok := value.(string); ok && s == "" {
				return basicError{"required_with_all", "required with all validation failed"}
			}

			return nil
		},
	)
}

// RequiredWithout returns an InputRule that validates the value exists if any of the given fields are absent from
// the input.
//
// The field under validation is optional only when all listed fields are present.
//
// Examples:
//
//	validation.RequiredWithout("email", "phone")
func RequiredWithout(fields ...string) InputRule {
	return presenceInputRuleFunc(
		func(value any, input *InputBag) error {
			for _, f := range fields {
				if _, found := input.Lookup(f); !found {
					if value == nil {
						return basicError{"required_without", "required without validation failed"}
					}
					if s, ok := value.(string); ok && s == "" {
						return basicError{"required_without", "required without validation failed"}
					}

					return nil
				}
			}

			return nil
		},
	)
}

// RequiredWithoutAll returns an InputRule that validates the value exists if all of the given fields are absent from
// the input.
//
// The field under validation is optional as long as at least one of the listed fields is present.
//
// Examples:
//
//	validation.RequiredWithoutAll("email", "phone")
func RequiredWithoutAll(fields ...string) InputRule {
	return presenceInputRuleFunc(
		func(value any, input *InputBag) error {
			for _, f := range fields {
				if _, found := input.Lookup(f); found {
					return nil
				}
			}

			if value == nil {
				return basicError{"required_without_all", "required without all validation failed"}
			}
			if s, ok := value.(string); ok && s == "" {
				return basicError{"required_without_all", "required without all validation failed"}
			}

			return nil
		},
	)
}

// NotEmpty is a Rule that validates the value is not an empty or zero value.
//
// Fails if:
//   - value is nil
//   - value is "" (empty string)
//   - value is 0 (any numeric type)
//   - value is false
//   - value is a zero-value struct
//
// Unlike Required, NotEmpty rejects all zero values, not just nil and "".
//
// Example:
//
//	schema := validation.New().
//		Field("count", validation.NotEmpty). // rejects 0
//		Field("active", validation.NotEmpty) // rejects false
var NotEmpty Rule = presenceRuleFunc(
	func(value any) error {
		if value == nil {
			return basicError{"not_empty", "not empty validation failed"}
		}

		if reflect.ValueOf(value).IsZero() {
			return basicError{"not_empty", "not empty validation failed"}
		}

		return nil
	},
)
