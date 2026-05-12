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
var Required Rule = RuleFunc(
	func(value any) error {
		if value == nil {
			return ErrValidationRequired
		}

		s, ok := value.(string)
		if ok && s == "" {
			return ErrValidationRequired
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
				return ErrValidationRequiredIf
			}
			if s, isStr := value.(string); isStr && s == "" {
				return ErrValidationRequiredIf
			}
		}

		return nil
	}

	return InputRuleFunc(fn)
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
var NotEmpty Rule = RuleFunc(
	func(value any) error {
		if value == nil {
			return ErrValidationNotEmpty
		}

		if reflect.ValueOf(value).IsZero() {
			return ErrValidationNotEmpty
		}

		return nil
	},
)
