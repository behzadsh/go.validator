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
