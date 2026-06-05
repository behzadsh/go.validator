package validation

import (
	"fmt"
	"regexp"
)

// Digits returns a Rule that validates the value is a string consisting of exactly n digit characters (0–9).
//
// Non-digit characters (signs, decimal points, letters) cause failure. The value must be a string.
//
// Fails if:
//   - value is not a string
//   - the string does not consist of exactly n digit characters
//
// Examples:
//
//	validation.Digits(4).Validate("1234")  // pass
//	validation.Digits(4).Validate("12345") // fail — five digits
//	validation.Digits(4).Validate("-123")  // fail — sign character
//	validation.Digits(4).Validate("12.3")  // fail — decimal point
func Digits(n int) Rule {
	re := regexp.MustCompile(fmt.Sprintf(`^\d{%d}$`, n))

	return RuleFunc(func(value any) error {
		str, ok := value.(string)
		if !ok || !re.MatchString(str) {
			return ErrValidationDigits
		}

		return nil
	})
}

// DigitsBetween returns a Rule that validates the value is a string with between min and max digit characters
// (inclusive, 0–9 only).
//
// Fails if:
//   - value is not a string
//   - the digit count is outside [min, max]
//
// Examples:
//
//	validation.DigitsBetween(4, 6).Validate("1234")   // pass
//	validation.DigitsBetween(4, 6).Validate("123456") // pass
//	validation.DigitsBetween(4, 6).Validate("123")    // fail — too few
//	validation.DigitsBetween(4, 6).Validate("1234567") // fail — too many
func DigitsBetween(min, max int) Rule {
	re := regexp.MustCompile(fmt.Sprintf(`^\d{%d,%d}$`, min, max))

	return RuleFunc(func(value any) error {
		str, ok := value.(string)
		if !ok || !re.MatchString(str) {
			return ErrValidationDigitsBetween
		}

		return nil
	})
}

// MaxDigits returns a Rule that validates the value is a string with at most n digit characters (0–9).
//
// Fails if:
//   - value is not a string
//   - the string has more than n digit characters (or any non-digit characters)
//
// Examples:
//
//	validation.MaxDigits(6).Validate("123456") // pass
//	validation.MaxDigits(6).Validate("123")    // pass — fewer is fine
//	validation.MaxDigits(6).Validate("1234567") // fail — seven digits
func MaxDigits(n int) Rule {
	re := regexp.MustCompile(fmt.Sprintf(`^\d{1,%d}$`, n))

	return RuleFunc(func(value any) error {
		str, ok := value.(string)
		if !ok || !re.MatchString(str) {
			return ErrValidationMaxDigits
		}

		return nil
	})
}

// MinDigits returns a Rule that validates the value is a string with at least n digit characters (0–9).
//
// Fails if:
//   - value is not a string
//   - the string has fewer than n digit characters
//
// Examples:
//
//	validation.MinDigits(4).Validate("1234")  // pass
//	validation.MinDigits(4).Validate("12345") // pass — more is fine
//	validation.MinDigits(4).Validate("123")   // fail — only three digits
func MinDigits(n int) Rule {
	re := regexp.MustCompile(fmt.Sprintf(`^\d{%d,}$`, n))

	return RuleFunc(func(value any) error {
		str, ok := value.(string)
		if !ok || !re.MatchString(str) {
			return ErrValidationMinDigits
		}

		return nil
	})
}
