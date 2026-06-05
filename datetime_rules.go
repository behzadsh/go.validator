package validation

import "time"

var timeFormats = []string{
	"2006-01-02",
	time.RFC3339,
	"2006-01-02T15:04:05",
	time.RFC1123Z,
	time.RFC1123,
	time.RFC822Z,
	time.RFC822,
	time.RFC850,
	"2006-01-02 15:04:05.999999999 -0700 MST",
	"2006-01-02T15:04:05-0700",
	"2006-01-02 15:04:05Z0700",
	"2006-01-02 15:04:05",
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	"2006-01-02 15:04:05Z07:00",
	"02 Jan 2006",
	"2006-01-02 15:04:05 -07:00",
	"2006-01-02 15:04:05 -0700",
	time.Kitchen,
	time.Stamp,
	time.StampMilli,
	time.StampMicro,
	time.StampNano,
}

// After returns a Rule that validates the value is a date/time string occurring strictly after ct.
//
// The string is tried against a broad set of common formats (RFC3339, "2006-01-02", RFC1123, and many more) without
// requiring a specific layout. Equal timestamps are rejected; the value must be strictly after ct.
//
// Fails if:
//   - value is not a string
//   - the string does not match any known date/time format
//   - the parsed time is equal to or before ct
//
// Examples:
//
//	deadline := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
//	validation.After(deadline).Validate("2024-06-01")           // pass
//	validation.After(deadline).Validate("2024-01-01T00:00:00Z") // fail — equal
//	validation.After(deadline).Validate("2023-12-31")           // fail — before
//	validation.After(deadline).Validate("not-a-date")           // fail
func After(ct time.Time) Rule {
	fn := func(value any) error {
		str, ok := value.(string)
		if !ok {
			return ErrValidationAfter
		}

		t, ok := parseTime(str)
		if !ok || !t.After(ct) {
			return ErrValidationAfter
		}

		return nil
	}

	return RuleFunc(fn)
}

// AfterField returns an InputRule that validates the value is a date/time string occurring strictly after the
// date/time at the given field path in the input.
//
// Both the field under validation and the referenced field must be parseable date/time strings.
// If the referenced field is absent or not a string, validation fails.
//
// Fails if:
//   - value is not a string or cannot be parsed
//   - the referenced field is absent, not a string, or cannot be parsed
//   - the parsed time is equal to or before the referenced field's time
//
// Examples:
//
//	schema := validation.New().
//		Field("end", validation.AfterField("start"))
func AfterField(path string) InputRule {
	return InputRuleFunc(
		func(value any, input *InputBag) error {
			str, ok := value.(string)
			if !ok {
				return ErrValidationAfterField
			}

			t, ok := parseTime(str)
			if !ok {
				return ErrValidationAfterField
			}

			otherRaw, found := input.Lookup(path)
			if !found {
				return ErrValidationAfterField
			}

			otherStr, ok := otherRaw.(string)
			if !ok {
				return ErrValidationAfterField
			}

			other, ok := parseTime(otherStr)
			if !ok || !t.After(other) {
				return ErrValidationAfterField
			}

			return nil
		},
	)
}

// AfterOrEqual returns a Rule that validates the value is a date/time string occurring on or after ct.
//
// Fails if:
//   - value is not a string
//   - the string does not match any known date/time format
//   - the parsed time is strictly before ct
//
// Examples:
//
//	deadline := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
//	validation.AfterOrEqual(deadline).Validate("2024-01-01") // pass — equal
//	validation.AfterOrEqual(deadline).Validate("2024-06-01") // pass — after
//	validation.AfterOrEqual(deadline).Validate("2023-12-31") // fail — before
func AfterOrEqual(ct time.Time) Rule {
	return RuleFunc(
		func(value any) error {
			str, ok := value.(string)
			if !ok {
				return ErrValidationAfterOrEqual
			}

			t, ok := parseTime(str)
			if !ok || t.Before(ct) {
				return ErrValidationAfterOrEqual
			}

			return nil
		},
	)
}

// Before returns a Rule that validates the value is a date/time string occurring strictly before ct.
//
// The string is tried against a broad set of common formats (RFC3339, "2006-01-02", RFC1123, and many more) without
// requiring a specific layout. Equal timestamps are rejected; the value must be strictly before ct.
//
// Fails if:
//   - value is not a string
//   - the string does not match any known date/time format
//   - the parsed time is equal to or after ct
//
// Examples:
//
//	expiry := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
//	validation.Before(expiry).Validate("2024-06-01")           // pass
//	validation.Before(expiry).Validate("2025-01-01T00:00:00Z") // fail — equal
//	validation.Before(expiry).Validate("2025-06-01")           // fail — after
//	validation.Before(expiry).Validate("not-a-date")           // fail
func Before(ct time.Time) Rule {
	fn := func(value any) error {
		str, ok := value.(string)
		if !ok {
			return ErrValidationBefore
		}

		t, ok := parseTime(str)
		if !ok || !t.Before(ct) {
			return ErrValidationBefore
		}

		return nil
	}

	return RuleFunc(fn)
}

// BeforeField returns an InputRule that validates the value is a date/time string occurring strictly before the
// date/time at the given field path in the input.
//
// Both the field under validation and the referenced field must be parseable date/time strings.
// If the referenced field is absent or not a string, validation fails.
//
// Fails if:
//   - value is not a string or cannot be parsed
//   - the referenced field is absent, not a string, or cannot be parsed
//   - the parsed time is equal to or after the referenced field's time
//
// Examples:
//
//	schema := validation.New().
//		Field("start", validation.BeforeField("end"))
func BeforeField(path string) InputRule {
	return InputRuleFunc(
		func(value any, input *InputBag) error {
			str, ok := value.(string)
			if !ok {
				return ErrValidationBeforeField
			}

			t, ok := parseTime(str)
			if !ok {
				return ErrValidationBeforeField
			}

			otherRaw, found := input.Lookup(path)
			if !found {
				return ErrValidationBeforeField
			}

			otherStr, ok := otherRaw.(string)
			if !ok {
				return ErrValidationBeforeField
			}

			other, ok := parseTime(otherStr)
			if !ok || !t.Before(other) {
				return ErrValidationBeforeField
			}

			return nil
		},
	)
}

// BeforeOrEqual returns a Rule that validates the value is a date/time string occurring on or before ct.
//
// Fails if:
//   - value is not a string
//   - the string does not match any known date/time format
//   - the parsed time is strictly after ct
//
// Examples:
//
//	expiry := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
//	validation.BeforeOrEqual(expiry).Validate("2025-01-01") // pass — equal
//	validation.BeforeOrEqual(expiry).Validate("2024-06-01") // pass — before
//	validation.BeforeOrEqual(expiry).Validate("2025-06-01") // fail — after
func BeforeOrEqual(ct time.Time) Rule {
	return RuleFunc(
		func(value any) error {
			str, ok := value.(string)
			if !ok {
				return ErrValidationBeforeOrEqual
			}

			t, ok := parseTime(str)
			if !ok || t.After(ct) {
				return ErrValidationBeforeOrEqual
			}

			return nil
		},
	)
}

// DateTime is a Rule that validates the value is a recognizable date/time string.
//
// The string is tried against a broad set of common formats (the same set used by After and Before).
// No specific layout is required; any of the supported formats will pass.
//
// Fails if:
//   - value is not a string
//   - the string does not match any known date/time format
//
// Examples:
//
//	validation.DateTime.Validate("2024-03-15")              // pass
//	validation.DateTime.Validate("2024-03-15T10:00:00Z")    // pass
//	validation.DateTime.Validate("not-a-date")              // fail
var DateTime Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok {
			return ErrValidationDateTime
		}

		if _, ok := parseTime(str); !ok {
			return ErrValidationDateTime
		}

		return nil
	},
)

// DateTimeBetween returns a Rule that validates the value is a date/time string occurring between min and max
// (inclusive on both ends).
//
// Fails if:
//   - value is not a string
//   - the string does not match any known date/time format
//   - the parsed time is before min or after max
//
// Examples:
//
//	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
//	end   := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
//	validation.DateTimeBetween(start, end).Validate("2024-06-01") // pass
//	validation.DateTimeBetween(start, end).Validate("2024-01-01") // pass — equal to min
//	validation.DateTimeBetween(start, end).Validate("2023-12-31") // fail — before min
//	validation.DateTimeBetween(start, end).Validate("2025-01-01") // fail — after max
func DateTimeBetween(minV, maxV time.Time) Rule {
	return RuleFunc(
		func(value any) error {
			str, ok := value.(string)
			if !ok {
				return ErrValidationDateTimeBetween
			}

			t, ok := parseTime(str)
			if !ok || t.Before(minV) || t.After(maxV) {
				return ErrValidationDateTimeBetween
			}

			return nil
		},
	)
}

// DateTimeFormat returns a Rule that validates the value is a string matching the given time layout.
//
// The layout uses Go's reference time: Mon Jan 2 15:04:05 MST 2006 (i.e. "2006-01-02" for ISO dates, time.RFC3339 for
// full timestamps, etc.).
//
// Fails if:
//   - value is not a string
//   - the string does not match the layout (wrong format, wrong separators, out-of-range values)
//
// Examples:
//
//	validation.DateTimeFormat("2006-01-02").Validate("2024-03-15")           // pass
//	validation.DateTimeFormat(time.RFC3339).Validate("2024-03-15T10:00:00Z") // pass
//	validation.DateTimeFormat("2006-01-02").Validate("15/03/2024")           // fail — wrong format
//	validation.DateTimeFormat("2006-01-02").Validate("not-a-date")           // fail
func DateTimeFormat(layout string) Rule {
	fn := func(value any) error {
		str, ok := value.(string)
		if !ok {
			return ErrValidationDateTimeFormat
		}

		if _, err := time.Parse(layout, str); err != nil {
			return ErrValidationDateTimeFormat
		}

		return nil
	}

	return RuleFunc(fn)
}

// Timezone is a Rule that validates the value is a valid IANA timezone name.
//
// Validation uses time.LoadLocation which recognizes IANA names (e.g. "UTC", "America/New_York",
// "Europe/London") as well as fixed-offset zones (e.g. "UTC+5").
//
// Fails if:
//   - value is not a string
//   - the string is not a valid IANA timezone name
//
// Examples:
//
//	validation.Timezone.Validate("UTC")               // pass
//	validation.Timezone.Validate("America/New_York")  // pass
//	validation.Timezone.Validate("Europe/London")     // pass
//	validation.Timezone.Validate("InvalidZone")       // fail
var Timezone Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok || str == "" {
			return ErrValidationTimezone
		}

		if _, err := time.LoadLocation(str); err != nil {
			return ErrValidationTimezone
		}

		return nil
	},
)

func parseTime(str string) (time.Time, bool) {
	for _, format := range timeFormats {
		t, err := time.Parse(format, str)
		if err == nil {
			return t, true
		}
	}

	return time.Time{}, false
}
