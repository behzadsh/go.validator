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

func parseTime(str string) (time.Time, bool) {
	for _, format := range timeFormats {
		t, err := time.Parse(format, str)
		if err == nil {
			return t, true
		}
	}

	return time.Time{}, false
}
