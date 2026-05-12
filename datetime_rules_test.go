package validation

import (
	"errors"
	"testing"
	"time"
)

var (
	tRef = time.Date(2023, 6, 15, 12, 0, 0, 0, time.UTC)
)

func TestDateTimeFormat(t *testing.T) {
	tests := []struct {
		value   any
		layout  string
		wantErr bool
	}{
		{"2023-06-15", "2006-01-02", false},
		{"2023-06-15T12:00:00Z", time.RFC3339, false},
		{"15/06/2023", "2006-01-02", true},
		{"not a date", "2006-01-02", true},
		{"", "2006-01-02", true},
		{nil, "2006-01-02", true},
		{42, "2006-01-02", true},
	}
	for _, tt := range tests {
		err := DateTimeFormat(tt.layout).Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("DateTimeFormat(%q).Validate(%v) error = %v, wantErr %v", tt.layout, tt.value, err, tt.wantErr)
		}
		if err != nil && !errors.Is(err, ErrValidationDateTimeFormat) {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestAfter(t *testing.T) {
	rule := After(tRef)
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"2024-01-01", false},
		{"2023-06-16", false},
		{"2023-06-15", true},    // equal, not after
		{"2022-01-01", true},
		{"not a date", true},
		{"", true},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := rule.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("After(tRef).Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && !errors.Is(err, ErrValidationAfter) {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestBefore(t *testing.T) {
	rule := Before(tRef)
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"2022-01-01", false},
		{"2023-06-14", false},
		{"2023-06-15T12:00:00Z", true},  // exact equal, not before
		{"2024-01-01", true},
		{"not a date", true},
		{"", true},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := rule.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("Before(tRef).Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && !errors.Is(err, ErrValidationBefore) {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestParseTime(t *testing.T) {
	tests := []struct {
		input  string
		wantOk bool
	}{
		{"2023-06-15", true},
		{"2023-06-15T12:00:00Z", true},
		{"2023-06-15T12:00:00+03:30", true},
		{"Thu, 15 Jun 2023 12:00:00 +0000", true},
		{"not a date", false},
		{"", false},
	}
	for _, tt := range tests {
		_, ok := parseTime(tt.input)
		if ok != tt.wantOk {
			t.Errorf("parseTime(%q) ok = %v, want %v", tt.input, ok, tt.wantOk)
		}
	}
}
