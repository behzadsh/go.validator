package validation

import (
	"errors"
	"testing"
)

func TestSameAs(t *testing.T) {
	rule := SameAs("password")

	tests := []struct {
		name    string
		value   any
		input   map[string]any
		wantErr bool
	}{
		{
			name:    "matching strings",
			value:   "secret",
			input:   map[string]any{"password": "secret"},
			wantErr: false,
		},
		{
			name:    "mismatched strings",
			value:   "wrong",
			input:   map[string]any{"password": "secret"},
			wantErr: true,
		},
		{
			name:    "referenced field absent",
			value:   "secret",
			input:   map[string]any{},
			wantErr: true,
		},
		{
			name:    "type mismatch",
			value:   "1",
			input:   map[string]any{"password": 1},
			wantErr: true,
		},
		{
			name:    "both nil",
			value:   nil,
			input:   map[string]any{"password": nil},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bag := NewInputBag(tt.input)
			err := rule.ValidateWithInput(tt.value, bag)
			if (err != nil) != tt.wantErr {
				t.Errorf("SameAs(%q).ValidateWithInput(%v) error = %v, wantErr %v", "password", tt.value, err, tt.wantErr)
			}
			if err != nil && !errors.Is(err, ErrValidationSameAs) {
				t.Errorf("wrong error type: %v", err)
			}
		})
	}
}

func TestDifferent(t *testing.T) {
	rule := Different("old_password")

	tests := []struct {
		name    string
		value   any
		input   map[string]any
		wantErr bool
	}{
		{
			name:    "different values",
			value:   "newpass",
			input:   map[string]any{"old_password": "oldpass"},
			wantErr: false,
		},
		{
			name:    "same values",
			value:   "samepass",
			input:   map[string]any{"old_password": "samepass"},
			wantErr: true,
		},
		{
			name:    "referenced field absent — passes",
			value:   "anything",
			input:   map[string]any{},
			wantErr: false,
		},
		{
			name:    "type mismatch — passes (1 != \"1\")",
			value:   "1",
			input:   map[string]any{"old_password": 1},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bag := NewInputBag(tt.input)
			err := rule.ValidateWithInput(tt.value, bag)
			if (err != nil) != tt.wantErr {
				t.Errorf("Different(%q).ValidateWithInput(%v) error = %v, wantErr %v", "old_password", tt.value, err, tt.wantErr)
			}
			if err != nil && !errors.Is(err, ErrValidationDifferent) {
				t.Errorf("wrong error type: %v", err)
			}
		})
	}
}
