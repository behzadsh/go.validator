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
		if err != nil && !errors.Is(err, ErrValidationAlpha) {
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

func TestURL(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"https://example.com", false},
		{"http://example.com", false},
		{"http://example.com/path?q=1", false},
		{"example.com", false},
		{"example.com/path", false},
		{"localhost", false},
		{"http://localhost:8080", false},
		{"http://127.0.0.1", false},
		{"example.com:8080", false},
		{"http://[::1]:8080", false},
		{"not a url", true},
		{"", true},
		{"http://", true},
		{42, true},
		{nil, true},
	}
	for _, tt := range tests {
		err := URL.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("URL.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
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
		{"550e8400-e29b-41d4-a716", true},       // too short
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
		if err != nil && !errors.Is(err, ErrValidationUUID) {
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
		if err != nil && !errors.Is(err, ErrValidationRegex) {
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
		if err != nil && !errors.Is(err, ErrValidationNotRegex) {
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
		if err != nil && !errors.Is(err, ErrValidationStartsWith) {
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
		if err != nil && !errors.Is(err, ErrValidationEndsWith) {
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
		if err != nil && !errors.Is(err, ErrValidationLowercase) {
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
		if err != nil && !errors.Is(err, ErrValidationUppercase) {
			t.Errorf("Uppercase.Validate(%v) wrong error: %v", tt.value, err)
		}
	}
}
