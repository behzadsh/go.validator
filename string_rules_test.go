package validation

import (
	"errors"
	"testing"
)

func TestAlpha(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"abc", false},
		{"Ünïcödé", false},
		{"abc123", true},
		{"abc-def", true},
		{"", true},
		{123, true},
		{nil, true},
	}
	for _, tt := range tests {
		err := Alpha.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("Alpha.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "alpha" {
			t.Errorf("Alpha.Validate(%v) wrong error type: %v", tt.value, err)
		}
	}
}

func TestAlphaDash(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"abc", false},
		{"abc-def", false},
		{"abc_def", false},
		{"abc123", false},
		{"abc def", true},
		{"abc@def", true},
		{"", true},
		{42, true},
	}
	for _, tt := range tests {
		err := AlphaDash.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("AlphaDash.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
	}
}

func TestAlphaNum(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"abc123", false},
		{"ABC", false},
		{"abc-123", true},
		{"abc 123", true},
		{"", true},
		{99, true},
	}
	for _, tt := range tests {
		err := AlphaNum.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("AlphaNum.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
	}
}

func TestAlphaSpace(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"hello world", false},
		{"Hello World", false},
		{"Ünïcödé chars", false},
		{"hello123", true},
		{"hello-world", true},
		{"", true},
		{nil, true},
	}
	for _, tt := range tests {
		err := AlphaSpace.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("AlphaSpace.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
	}
}

func TestLength(t *testing.T) {
	tests := []struct {
		value   any
		l       int
		wantErr bool
	}{
		{"hello", 5, false},
		{"héllo", 5, false}, // 5 runes, >5 bytes
		{"héllo", 6, true},  // 5 runes != 6
		{"hi", 5, true},
		{"", 0, false},
		{"", 1, true},
		{42, 5, true},
		{nil, 0, true},
	}
	for _, tt := range tests {
		err := Length(tt.l).Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("Length(%d).Validate(%v) error = %v, wantErr %v", tt.l, tt.value, err, tt.wantErr)
		}
	}
}

func TestMinLength(t *testing.T) {
	tests := []struct {
		value   any
		l       int
		wantErr bool
	}{
		{"hello", 3, false},
		{"hello", 5, false},
		{"héllo", 5, false}, // 5 runes
		{"héllo", 6, true},  // 5 runes < 6
		{"hi", 5, true},
		{"", 1, true},
		{"", 0, false},
		{42, 3, true},
	}
	for _, tt := range tests {
		err := MinLength(tt.l).Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("MinLength(%d).Validate(%v) error = %v, wantErr %v", tt.l, tt.value, err, tt.wantErr)
		}
	}
}

func TestMaxLength(t *testing.T) {
	tests := []struct {
		value   any
		l       int
		wantErr bool
	}{
		{"hi", 5, false},
		{"hello", 5, false},
		{"héllo", 5, false}, // 5 runes <= 5
		{"hello!", 5, true},
		{"héllo!", 5, true}, // 6 runes > 5
		{"", 0, false},
		{"a", 0, true},
		{42, 5, true},
	}
	for _, tt := range tests {
		err := MaxLength(tt.l).Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("MaxLength(%d).Validate(%v) error = %v, wantErr %v", tt.l, tt.value, err, tt.wantErr)
		}
	}
}

func TestEmail(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"user@example.com", false},
		{"user+tag@sub.example.org", false},
		{"user.name@example.co.uk", false},
		{"notanemail", true},
		{"@example.com", true},
		{"user@", true},
		{"user@@example.com", true},
		{"", true},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := Email.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("Email.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
	}
}

func TestEmailMX(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping network test in short mode")
	}

	tests := []struct {
		value   any
		wantErr bool
	}{
		{"user@gmail.com", false},
		{"user@invalid.com", true}, // example.com has no MX records
		{"notanemail", true},
		{nil, true},
	}
	for _, tt := range tests {
		err := EmailMX.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("EmailMX.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
	}
}

func TestUUID(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"550e8400-e29b-41d4-a716-446655440000", false},
		{"550E8400-E29B-41D4-A716-446655440000", false}, // uppercase
		{"00000000-0000-0000-0000-000000000000", false},
		{"not-a-uuid", true},
		{"550e8400-e29b-41d4-a716", true},          // too short
		{"550e8400e29b41d4a716446655440000", true}, // no dashes
		{"", true},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := UUID.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("UUID.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "uuid" {
			t.Errorf("UUID.Validate(%v) wrong error: %v", tt.value, err)
		}
	}
}

func TestRegex(t *testing.T) {
	tests := []struct {
		value   any
		pattern string
		wantErr bool
	}{
		{"1234", `^\d{4}$`, false},
		{"12345", `^\d{4}$`, true},
		{"abcd", `^\d{4}$`, true},
		{"hello", `^[a-z]+$`, false},
		{"Hello", `^[a-z]+$`, true},
		{"", `^\d{4}$`, true},
		{nil, `^\d{4}$`, true},
		{42, `^\d{4}$`, true},
	}
	for _, tt := range tests {
		err := Regex(tt.pattern).Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("Regex(%q).Validate(%v) error = %v, wantErr %v", tt.pattern, tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "regex" {
			t.Errorf("Regex(%q).Validate(%v) wrong error: %v", tt.pattern, tt.value, err)
		}
	}
}

func TestRegexInvalidPattern(t *testing.T) {
	err := Regex(`[invalid`).Validate("test")
	if err == nil {
		t.Fatal("expected error for invalid pattern, got nil")
	}
	var syntaxErr RuleSyntaxError
	if !errors.As(err, &syntaxErr) {
		t.Errorf("expected RuleSyntaxError, got %T: %v", err, err)
	}
}

func TestNotRegex(t *testing.T) {
	tests := []struct {
		value   any
		pattern string
		wantErr bool
	}{
		{"nospaces", `\s`, false},
		{"has spaces", `\s`, true},
		{"hello", `\d`, false},
		{"hello1", `\d`, true},
		{"", `\s`, false},
		{nil, `\s`, true},
		{42, `\s`, true},
	}
	for _, tt := range tests {
		err := NotRegex(tt.pattern).Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("NotRegex(%q).Validate(%v) error = %v, wantErr %v", tt.pattern, tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "not_regex" {
			t.Errorf("NotRegex(%q).Validate(%v) wrong error: %v", tt.pattern, tt.value, err)
		}
	}
}

func TestStartsWith(t *testing.T) {
	tests := []struct {
		value   any
		prefix  string
		wantErr bool
	}{
		{"SKU-001", "SKU-", false},
		{"001-SKU", "SKU-", true},
		{"SKU-", "SKU-", false},
		{"", "SKU-", true},
		{"", "", false},
		{nil, "SKU-", true},
		{42, "SKU-", true},
	}
	for _, tt := range tests {
		err := StartsWith(tt.prefix).Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("StartsWith(%q).Validate(%v) error = %v, wantErr %v", tt.prefix, tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "starts_with" {
			t.Errorf("StartsWith(%q).Validate(%v) wrong error: %v", tt.prefix, tt.value, err)
		}
	}
}

func TestEndsWith(t *testing.T) {
	tests := []struct {
		value   any
		suffix  string
		wantErr bool
	}{
		{"main.go", ".go", false},
		{"main.js", ".go", true},
		{".go", ".go", false},
		{"", ".go", true},
		{"", "", false},
		{nil, ".go", true},
		{42, ".go", true},
	}
	for _, tt := range tests {
		err := EndsWith(tt.suffix).Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("EndsWith(%q).Validate(%v) error = %v, wantErr %v", tt.suffix, tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "ends_with" {
			t.Errorf("EndsWith(%q).Validate(%v) wrong error: %v", tt.suffix, tt.value, err)
		}
	}
}

func TestLowercase(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"hello", false},
		{"hello world", false},
		{"hello123", false},
		{"Hello", true},
		{"HELLO", true},
		{"hEllo", true},
		{"", false},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := Lowercase.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("Lowercase.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "lowercase" {
			t.Errorf("Lowercase.Validate(%v) wrong error: %v", tt.value, err)
		}
	}
}

func TestUppercase(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"HELLO", false},
		{"HELLO WORLD", false},
		{"HELLO123", false},
		{"Hello", true},
		{"hello", true},
		{"HELLo", true},
		{"", false},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := Uppercase.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("Uppercase.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "uppercase" {
			t.Errorf("Uppercase.Validate(%v) wrong error: %v", tt.value, err)
		}
	}
}

func TestASCII(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"hello", false},
		{"hello world", false},
		{"abc123!@#", false},
		{"\t\n\r", false},
		{"café", true},
		{"Ünïcödé", true},
		{"", false},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := ASCII.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("ASCII.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "ascii" {
			t.Errorf("ASCII.Validate(%v) wrong error: %v", tt.value, err)
		}
	}
}

func TestBase64(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"aGVsbG8=", false}, // "hello"
		{"dGVzdA==", false}, // "test"
		{"", false},         // empty is valid base64
		{"aGVsbG8", true},   // missing padding
		{"not-base64!", true},
		{"aGVsbG8===", true}, // bad padding
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := Base64.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("Base64.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "base64" {
			t.Errorf("Base64.Validate(%v) wrong error: %v", tt.value, err)
		}
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"hello world", false},
		{"hello", false},
		{"world", true},
		{"", true},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := Contains("hello").Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("Contains(\"hello\").Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "contains" {
			t.Errorf("Contains.Validate(%v) wrong error: %v", tt.value, err)
		}
	}
}

func TestCreditCard(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"4111111111111111", false},    // Visa test
		{"4111-1111-1111-1111", false}, // dashes stripped
		{"4111 1111 1111 1111", false}, // spaces stripped
		{"5500005555555559", false},    // Mastercard test
		{"378282246310005", false},     // Amex test (15 digits)
		{"1234567890123456", true},     // invalid Luhn
		{"411111111111111", true},      // too short (15 chars but bad Luhn)
		{"", true},
		{"not-a-card", true},
		{nil, true},
		{4111111111111111, true}, // not a string
	}
	for _, tt := range tests {
		err := CreditCard.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("CreditCard.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "credit_card" {
			t.Errorf("CreditCard.Validate(%v) wrong error: %v", tt.value, err)
		}
	}
}

func TestHexColor(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"#fff", false},
		{"#FFF", false},
		{"#FF5733", false},
		{"#aabbcc", false},
		{"FF5733", true},   // missing #
		{"#GGG", true},     // invalid hex digits
		{"#12345", true},   // 5 digits
		{"#1234567", true}, // 7 digits
		{"", true},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := HexColor.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("HexColor.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "hex_color" {
			t.Errorf("HexColor.Validate(%v) wrong error: %v", tt.value, err)
		}
	}
}

func TestJSON(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{`{"key":"value"}`, false},
		{`[1,2,3]`, false},
		{`null`, false},
		{`"hello"`, false},
		{`123`, false},
		{`true`, false},
		{`false`, false},
		{`{invalid}`, true},
		{``, true},
		{`{`, true},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := JSON.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("JSON.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "json" {
			t.Errorf("JSON.Validate(%v) wrong error: %v", tt.value, err)
		}
	}
}

func TestJWT(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIn0.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c", false},
		{"eyJ.eyJ.sig", false},
		{"a.b.c", false},
		{"notajwt", true},
		{"a.b", true},     // only two segments
		{"a.b.c.d", true}, // four segments
		{"", true},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := JWT.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("JWT.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "jwt" {
			t.Errorf("JWT.Validate(%v) wrong error: %v", tt.value, err)
		}
	}
}

func TestPhoneE164(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"+14155552671", false},
		{"+441234567890", false},
		{"+1", true},                // too short (only 1 digit after +)
		{"14155552671", true},       // missing +
		{"+0123456789", true},       // country code starts with 0
		{"+1234567890123456", true}, // too long (16 digits)
		{"", true},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := PhoneE164.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("PhoneE164.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "phone_e164" {
			t.Errorf("PhoneE164.Validate(%v) wrong error: %v", tt.value, err)
		}
	}
}

func TestSemver(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"1.0.0", false},
		{"0.0.1", false},
		{"v1.0.0", false}, // v prefix accepted
		{"v2.1.3-alpha.1", false},
		{"2.1.3-alpha.1", false},
		{"1.0.0+build.123", false}, // build metadata accepted
		{"1.0.0-beta.1+exp.sha.5114f85", false},
		{"1.0", true},     // missing patch
		{"1.0.0.0", true}, // extra component
		{"01.0.0", true},  // leading zero
		{"", true},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := Semver.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("Semver.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "semver" {
			t.Errorf("Semver.Validate(%v) wrong error: %v", tt.value, err)
		}
	}
}

func TestSlug(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"hello", false},
		{"hello-world", false},
		{"my-post-123", false},
		{"abc", false},
		{"Hello-World", true},  // uppercase
		{"-leading", true},     // leading hyphen
		{"trailing-", true},    // trailing hyphen
		{"double--dash", true}, // consecutive hyphens
		{"hello world", true},  // space
		{"", true},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := Slug.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("Slug.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "slug" {
			t.Errorf("Slug.Validate(%v) wrong error: %v", tt.value, err)
		}
	}
}
