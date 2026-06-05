package validation

import (
	"errors"
	"testing"
	"time"
)

var tRef = time.Date(2023, 6, 15, 12, 0, 0, 0, time.UTC)

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
		{"2023-06-15", true}, // equal, not after
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
		{"2023-06-15T12:00:00Z", true}, // exact equal, not before
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

func TestDateTime(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"2024-03-15", false},
		{"2024-03-15T10:00:00Z", false},
		{"Thu, 15 Jun 2023 12:00:00 +0000", false},
		{"not-a-date", true},
		{"", true},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := DateTime.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("DateTime.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && !errors.Is(err, ErrValidationDateTime) {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestAfterOrEqual(t *testing.T) {
	rule := AfterOrEqual(tRef)
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"2024-01-01", false},
		{"2023-06-15T12:00:00Z", false}, // exactly equal to tRef passes
		{"2022-01-01", true},            // before fails
		{"2023-06-15", true},            // midnight < noon tRef → before
		{"not a date", true},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := rule.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("AfterOrEqual(tRef).Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && !errors.Is(err, ErrValidationAfterOrEqual) {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestBeforeOrEqual(t *testing.T) {
	rule := BeforeOrEqual(tRef)
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"2022-01-01", false},
		{"2023-06-15", false}, // equal passes
		{"2024-01-01", true},  // after fails
		{"not a date", true},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := rule.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("BeforeOrEqual(tRef).Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && !errors.Is(err, ErrValidationBeforeOrEqual) {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestDateTimeBetween(t *testing.T) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
	rule := DateTimeBetween(start, end)
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"2024-06-15", false},
		{"2024-01-01", false}, // equal to min
		{"2024-12-31", false}, // equal to max
		{"2023-12-31", true},  // before min
		{"2025-01-01", true},  // after max
		{"not a date", true},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := rule.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("DateTimeBetween.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && !errors.Is(err, ErrValidationDateTimeBetween) {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestAfterField(t *testing.T) {
	rule := AfterField("start")

	t.Run(
		"pass: end after start", func(t *testing.T) {
			bag := NewInputBag(map[string]any{"start": "2024-01-01", "end": "2024-06-01"})
			if err := rule.ValidateWithInput("2024-06-01", bag); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		},
	)

	t.Run(
		"fail: end equal to start", func(t *testing.T) {
			bag := NewInputBag(map[string]any{"start": "2024-01-01"})
			err := rule.ValidateWithInput("2024-01-01", bag)
			if err == nil {
				t.Error("expected error for equal times")
			}
			if !errors.Is(err, ErrValidationAfterField) {
				t.Errorf("wrong error type: %v", err)
			}
		},
	)

	t.Run(
		"fail: end before start", func(t *testing.T) {
			bag := NewInputBag(map[string]any{"start": "2024-06-01"})
			if err := rule.ValidateWithInput("2024-01-01", bag); err == nil {
				t.Error("expected error")
			}
		},
	)

	t.Run(
		"fail: start field absent", func(t *testing.T) {
			bag := NewInputBag(map[string]any{})
			if err := rule.ValidateWithInput("2024-01-01", bag); err == nil {
				t.Error("expected error for absent field")
			}
		},
	)

	t.Run(
		"fail: non-string value", func(t *testing.T) {
			bag := NewInputBag(map[string]any{"start": "2024-01-01"})
			if err := rule.ValidateWithInput(42, bag); err == nil {
				t.Error("expected error")
			}
		},
	)
}

func TestBeforeField(t *testing.T) {
	rule := BeforeField("end")

	t.Run(
		"pass: start before end", func(t *testing.T) {
			bag := NewInputBag(map[string]any{"end": "2024-06-01"})
			if err := rule.ValidateWithInput("2024-01-01", bag); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		},
	)

	t.Run(
		"fail: start equal to end", func(t *testing.T) {
			bag := NewInputBag(map[string]any{"end": "2024-01-01"})
			err := rule.ValidateWithInput("2024-01-01", bag)
			if err == nil {
				t.Error("expected error for equal times")
			}
			if !errors.Is(err, ErrValidationBeforeField) {
				t.Errorf("wrong error type: %v", err)
			}
		},
	)

	t.Run(
		"fail: start after end", func(t *testing.T) {
			bag := NewInputBag(map[string]any{"end": "2024-01-01"})
			if err := rule.ValidateWithInput("2024-06-01", bag); err == nil {
				t.Error("expected error")
			}
		},
	)
}

func TestTimezone(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"UTC", false},
		{"America/New_York", false},
		{"Europe/London", false},
		{"Asia/Tokyo", false},
		{"InvalidZone", true},
		{"", true},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := Timezone.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("Timezone.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && !errors.Is(err, ErrValidationTimezone) {
			t.Errorf("wrong error type: %v", err)
		}
	}
}
