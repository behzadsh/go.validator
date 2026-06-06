package validation

import (
	"errors"
	"math"
	"reflect"
	"strconv"
)

type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

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
			return betweenError{Min: minV, Max: maxV}
		}

		return nil
	}

	return RuleFunc(fn)
}

// GT returns a Rule that validates the value is strictly greater than v.
//
// Unlike Min (which uses >=), GT uses >, so equal values fail.
//
// Fails if:
//   - value cannot be type-asserted to T
//   - value <= v
//
// Examples:
//
//	validation.GT[int](18).Validate(19)  // pass
//	validation.GT[int](18).Validate(18)  // fail — equal
//	validation.GT[int](18).Validate(17)  // fail
func GT[T number](v T) Rule {
	return RuleFunc(
		func(value any) error {
			actual, ok := value.(T)
			if !ok || actual <= v {
				return gtError{Value: v}
			}

			return nil
		},
	)
}

// GTE returns a Rule that validates the value is greater than or equal to v.
//
// GTE is semantically identical to Min; it exists as an explicit named alias.
//
// Fails if:
//   - value cannot be type-asserted to T
//   - value < v
//
// Examples:
//
//	validation.GTE[int](18).Validate(18)  // pass — equal
//	validation.GTE[int](18).Validate(19)  // pass
//	validation.GTE[int](18).Validate(17)  // fail
func GTE[T number](v T) Rule {
	return RuleFunc(
		func(value any) error {
			actual, ok := value.(T)
			if !ok || actual < v {
				return gteError{Value: v}
			}

			return nil
		},
	)
}

// Integer is a Rule that validates the value is an integer type.
//
// Accepted kinds: int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64.
// Note: float64 (the default JSON number type) is rejected — use this rule when you need
// to assert the Go type is strictly integral.
//
// Fails if:
//   - value is not one of the integer kinds above
//
// Examples:
//
//	validation.Integer.Validate(42)          // pass
//	validation.Integer.Validate(uint8(255))  // pass
//	validation.Integer.Validate(3.14)        // fail — float64
//	validation.Integer.Validate("42")        // fail — string
var Integer Rule = RuleFunc(
	func(value any) error {
		switch reflect.ValueOf(value).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return nil
		default:
			return basicError{"integer", "integer validation failed"}
		}
	},
)

// Latitude is a Rule that validates the value is a valid latitude (−90 to 90 inclusive).
//
// Accepts any numeric type or float64 (the default JSON number type).
//
// Fails if:
//   - value is not a numeric type
//   - value < -90 or value > 90
//
// Examples:
//
//	validation.Latitude.Validate(45.0)    // pass
//	validation.Latitude.Validate(-90.0)   // pass — inclusive
//	validation.Latitude.Validate(90.1)    // fail
//	validation.Latitude.Validate("45.0")  // fail — string not accepted
var Latitude Rule = RuleFunc(
	func(value any) error {
		fv, ok := condToFloat(value)
		if !ok || fv < -90 || fv > 90 {
			return basicError{"latitude", "latitude validation failed"}
		}

		return nil
	},
)

// Longitude is a Rule that validates the value is a valid longitude (−180 to 180 inclusive).
//
// Accepts any numeric type or float64 (the default JSON number type).
//
// Fails if:
//   - value is not a numeric type
//   - value < -180 or value > 180
//
// Examples:
//
//	validation.Longitude.Validate(120.5)   // pass
//	validation.Longitude.Validate(-180.0)  // pass — inclusive
//	validation.Longitude.Validate(180.1)   // fail
//	validation.Longitude.Validate("120.5") // fail — string not accepted
var Longitude Rule = RuleFunc(
	func(value any) error {
		fv, ok := condToFloat(value)
		if !ok || fv < -180 || fv > 180 {
			return basicError{"longitude", "longitude validation failed"}
		}

		return nil
	},
)

// LT returns a Rule that validates the value is strictly less than v.
//
// Unlike Max (which uses <=), LT uses <, so equal values fail.
//
// Fails if:
//   - value cannot be type-asserted to T
//   - value >= v
//
// Examples:
//
//	validation.LT[int](100).Validate(99)   // pass
//	validation.LT[int](100).Validate(100)  // fail — equal
//	validation.LT[int](100).Validate(101)  // fail
func LT[T number](v T) Rule {
	return RuleFunc(
		func(value any) error {
			actual, ok := value.(T)
			if !ok || actual >= v {
				return ltError{Value: v}
			}

			return nil
		},
	)
}

// LTE returns a Rule that validates the value is less than or equal to v.
//
// LTE is semantically identical to Max; it exists as an explicit named alias.
//
// Fails if:
//   - value cannot be type-asserted to T
//   - value > v
//
// Examples:
//
//	validation.LTE[int](100).Validate(100)  // pass — equal
//	validation.LTE[int](100).Validate(99)   // pass
//	validation.LTE[int](100).Validate(101)  // fail
func LTE[T number](v T) Rule {
	return RuleFunc(
		func(value any) error {
			actual, ok := value.(T)
			if !ok || actual > v {
				return lteError{Value: v}
			}

			return nil
		},
	)
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
			return maxError{Value: maxV}
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
			return minError{Value: minV}
		}

		return nil
	}

	return RuleFunc(fn)
}

// MultipleOf returns a Rule that validates the value is a multiple of n.
//
// Accepts any numeric type; comparison is done in float64, n must not be zero.
// If n is zero, Schema.Validate returns a RuleSyntaxError.
//
// Fails if:
//   - value is not a numeric type
//   - value is not evenly divisible by n
//
// Examples:
//
//	validation.MultipleOf[int](3).Validate(9)          // pass
//	validation.MultipleOf[int](3).Validate(float64(9)) // pass — JSON number accepted
//	validation.MultipleOf[int](3).Validate(8)          // fail
func MultipleOf[T number](n T) Rule {
	return RuleFunc(
		func(value any) error {
			if float64(n) == 0 {
				return RuleSyntaxError{Rule: "MultipleOf", Err: errors.New("divisor must not be zero")}
			}

			fv, ok := condToFloat(value)
			if !ok {
				return multipleOfError{Value: n}
			}

			if math.Mod(fv, float64(n)) != 0 {
				return multipleOfError{Value: n}
			}

			return nil
		},
	)
}

// Negative is a Rule that validates the value is strictly less than zero.
//
// Accepts any numeric type.
//
// Fails if:
//   - value is not a numeric type
//   - value >= 0
//
// Examples:
//
//	validation.Negative.Validate(-1)   // pass
//	validation.Negative.Validate(-0.5) // pass
//	validation.Negative.Validate(0)    // fail — zero is not negative
//	validation.Negative.Validate(1)    // fail
var Negative Rule = RuleFunc(
	func(value any) error {
		fv, ok := condToFloat(value)
		if !ok || fv >= 0 {
			return basicError{"negative", "negative validation failed"}
		}

		return nil
	},
)

// NonNegative is a Rule that validates the value is greater than or equal to zero.
//
// Accepts any numeric type.
//
// Fails if:
//   - value is not a numeric type
//   - value < 0
//
// Examples:
//
//	validation.NonNegative.Validate(0)    // pass — zero is non-negative
//	validation.NonNegative.Validate(5)    // pass
//	validation.NonNegative.Validate(-1)   // fail
//	validation.NonNegative.Validate(-0.1) // fail
var NonNegative Rule = RuleFunc(
	func(value any) error {
		fv, ok := condToFloat(value)
		if !ok || fv < 0 {
			return basicError{"non_negative", "non negative validation failed"}
		}

		return nil
	},
)

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
				return basicError{"numeric", "numeric validation failed"}
			}

			return nil
		default:
			return basicError{"numeric", "numeric validation failed"}
		}
	},
)

// Port is a Rule that validates the value is a valid TCP/UDP port number (1–65535).
//
// Accepts any numeric type. Fractional values (e.g. float64(80.5)) fail.
//
// Fails if:
//   - value is not a numeric type
//   - value is fractional
//   - value < 1 or value > 65535
//
// Examples:
//
//	validation.Port.Validate(80)             // pass
//	validation.Port.Validate(float64(8080))  // pass — JSON number accepted
//	validation.Port.Validate(0)              // fail — port 0 is reserved
//	validation.Port.Validate(65536)          // fail
//	validation.Port.Validate(80.5)           // fail — fractional
var Port Rule = RuleFunc(
	func(value any) error {
		fv, ok := condToFloat(value)
		if !ok || fv != math.Trunc(fv) || fv < 1 || fv > 65535 {
			return basicError{"port", "port validation failed"}
		}

		return nil
	},
)

// Positive is a Rule that validates the value is strictly greater than zero.
//
// Accepts any numeric type.
//
// Fails if:
//   - value is not a numeric type
//   - value <= 0
//
// Examples:
//
//	validation.Positive.Validate(1)    // pass
//	validation.Positive.Validate(0.1)  // pass
//	validation.Positive.Validate(0)    // fail — zero is not positive
//	validation.Positive.Validate(-1)   // fail
var Positive Rule = RuleFunc(
	func(value any) error {
		fv, ok := condToFloat(value)
		if !ok || fv <= 0 {
			return basicError{"positive", "positive validation failed"}
		}

		return nil
	},
)
