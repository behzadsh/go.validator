package validation

import (
	"strconv"
)

type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

// Numeric is a Rule that validate the value is a number, or it can be converted to a number.
//
// Accepted types: int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, complex64,
// complex128, and any string whose content is a valid decimal number.
//
// Fails if:
//   - value is nil or a boolean
//   - value is a string that cannot be parsed as float64 (e.g. "abc", "1.2.3")
//   - value is any other non-numeric type (slice, map, struct, etc.)
//
// Examples:
//
//	validation.Numeric.Validate(42)       // pass
//	validation.Numeric.Validate(3.14)     // pass
//	validation.Numeric.Validate("99.5")   // pass — parseable string
//	validation.Numeric.Validate("abc")    // fail — not a number
//	validation.Numeric.Validate(true)     // fail — boolean not accepted
var Numeric Rule = RuleFunc(
	func(value any) error {
		switch v := value.(type) {
		case int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64,
			float32, float64,
			complex64, complex128:
			return nil
		case string:
			_, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return ErrValidationNumeric
			}

			return nil
		default:
			return ErrValidationNumeric
		}
	},
)

// Between returns a Rule that validates the value is between minV and maxV inclusive.
//
// The type parameter T must be instantiated explicitly because Go cannot infer it from the min/max arguments alone when
// the field value is typed as any.
//
// Fails if:
//   - value cannot be type-asserted to T
//   - value < min
//   - value > max
//
// Examples:
//
//	validation.Between[int](1, 100).Validate(50)        // pass
//	validation.Between[int](1, 100).Validate(1)         // pass — inclusive
//	validation.Between[int](1, 100).Validate(100)       // pass — inclusive
//	validation.Between[int](1, 100).Validate(0)         // fail — below min
//	validation.Between[int](1, 100).Validate(101)       // fail — above max
//	validation.Between[float64](0.0, 1.0).Validate(0.5) // pass
func Between[T number](minV, maxV T) Rule {
	fn := func(value any) error {
		v, ok := value.(T)
		if !ok || v < minV || v > maxV {
			return ErrValidationBetween
		}

		return nil
	}

	return RuleFunc(fn)
}

// Min returns a Rule that validates the value is at least minV.
//
// The type parameter T must be instantiated explicitly.
//
// Fails if:
//   - value cannot be type-asserted to T
//   - value < min
//
// Examples:
//
//	validation.Min[int](18).Validate(18)        // pass — equal to min
//	validation.Min[int](18).Validate(21)        // pass
//	validation.Min[int](18).Validate(17)        // fail — below min
//	validation.Min[float64](0.5).Validate(0.4)  // fail
func Min[T number](minV T) Rule {
	fn := func(value any) error {
		v, ok := value.(T)
		if !ok || v < minV {
			return ErrValidationMin
		}

		return nil
	}

	return RuleFunc(fn)
}

// Max returns a Rule that validates the value is at most maxV.
//
// The type parameter T must be instantiated explicitly.
//
// Fails if:
//   - value cannot be type-asserted to T
//   - value > max
//
// Examples:
//
//	validation.Max[int](100).Validate(100)      // pass — equal to max
//	validation.Max[int](100).Validate(50)       // pass
//	validation.Max[int](100).Validate(101)      // fail — above max
//	validation.Max[float64](1.0).Validate(1.1)  // fail
func Max[T number](maxV T) Rule {
	fn := func(value any) error {
		v, ok := value.(T)
		if !ok || v > maxV {
			return ErrValidationMax
		}

		return nil
	}

	return RuleFunc(fn)
}
