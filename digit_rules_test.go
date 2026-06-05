package validation

import (
	"errors"
	"testing"
)

func TestDigits(t *testing.T) {
	tests := []struct {
		value   any
		n       int
		wantErr bool
	}{
		{"1234", 4, false},
		{"0000", 4, false},
		{"12345", 4, true},   // too many
		{"123", 4, true},     // too few
		{"-123", 4, true},    // sign character
		{"12.3", 4, true},    // decimal point
		{"abcd", 4, true},    // letters
		{"", 4, true},
		{nil, 4, true},
		{1234, 4, true},      // non-string
	}
	for _, tt := range tests {
		err := Digits(tt.n).Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("Digits(%d).Validate(%v) error = %v, wantErr %v", tt.n, tt.value, err, tt.wantErr)
		}
		if err != nil && !errors.Is(err, ErrValidationDigits) {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestMinDigits(t *testing.T) {
	tests := []struct {
		value   any
		n       int
		wantErr bool
	}{
		{"1234", 4, false},
		{"12345", 4, false},  // more is fine
		{"123", 4, true},     // too few
		{"-1234", 4, true},   // sign character
		{"", 1, true},
		{nil, 4, true},
		{1234, 4, true},
	}
	for _, tt := range tests {
		err := MinDigits(tt.n).Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("MinDigits(%d).Validate(%v) error = %v, wantErr %v", tt.n, tt.value, err, tt.wantErr)
		}
		if err != nil && !errors.Is(err, ErrValidationMinDigits) {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestMaxDigits(t *testing.T) {
	tests := []struct {
		value   any
		n       int
		wantErr bool
	}{
		{"123", 6, false},
		{"123456", 6, false}, // exactly max
		{"1234567", 6, true}, // too many
		{"", 6, true},        // empty string fails (no digits)
		{nil, 6, true},
		{123, 6, true},
	}
	for _, tt := range tests {
		err := MaxDigits(tt.n).Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("MaxDigits(%d).Validate(%v) error = %v, wantErr %v", tt.n, tt.value, err, tt.wantErr)
		}
		if err != nil && !errors.Is(err, ErrValidationMaxDigits) {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestDigitsBetween(t *testing.T) {
	tests := []struct {
		value   any
		min     int
		max     int
		wantErr bool
	}{
		{"1234", 4, 6, false},
		{"123456", 4, 6, false},
		{"12345", 4, 6, false},
		{"123", 4, 6, true},     // too few
		{"1234567", 4, 6, true}, // too many
		{"-1234", 4, 6, true},   // sign character
		{"", 4, 6, true},
		{nil, 4, 6, true},
		{1234, 4, 6, true},
	}
	for _, tt := range tests {
		err := DigitsBetween(tt.min, tt.max).Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("DigitsBetween(%d,%d).Validate(%v) error = %v, wantErr %v", tt.min, tt.max, tt.value, err, tt.wantErr)
		}
		if err != nil && !errors.Is(err, ErrValidationDigitsBetween) {
			t.Errorf("wrong error type: %v", err)
		}
	}
}
