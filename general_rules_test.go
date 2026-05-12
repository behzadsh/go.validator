package validation

import (
	"errors"
	"testing"
)

func TestRequired(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"hello", false},
		{"0", false},
		{0, false},
		{false, false},
		{nil, true},
		{"", true},
	}
	for _, tt := range tests {
		err := Required.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("Required.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && !errors.Is(err, ErrValidationRequired) {
			t.Errorf("Required.Validate(%v) wrong error: %v", tt.value, err)
		}
	}
}

func TestRequiredIf(t *testing.T) {
	tests := []struct {
		name      string
		condition string
		value     any
		input     map[string]any
		wantErr   bool
		wantIs    error
	}{
		{
			name:      "condition true, value present",
			condition: `role == "admin"`,
			value:     "something",
			input:     map[string]any{"role": "admin"},
			wantErr:   false,
		},
		{
			name:      "condition true, value nil",
			condition: `role == "admin"`,
			value:     nil,
			input:     map[string]any{"role": "admin"},
			wantErr:   true,
			wantIs:    ErrValidationRequiredIf,
		},
		{
			name:      "condition true, value empty string",
			condition: `role == "admin"`,
			value:     "",
			input:     map[string]any{"role": "admin"},
			wantErr:   true,
			wantIs:    ErrValidationRequiredIf,
		},
		{
			name:      "condition false, value nil",
			condition: `role == "admin"`,
			value:     nil,
			input:     map[string]any{"role": "user"},
			wantErr:   false,
		},
		{
			name:      "compound condition true",
			condition: `role == "admin" && exists(plan)`,
			value:     nil,
			input:     map[string]any{"role": "admin", "plan": "pro"},
			wantErr:   true,
			wantIs:    ErrValidationRequiredIf,
		},
		{
			name:      "syntax error returns RuleSyntaxError",
			condition: "unknown(field)",
			value:     nil,
			input:     map[string]any{},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := RequiredIf(tt.condition)
			bag := NewInputBag(tt.input)
			err := rule.ValidateWithInput(tt.value, bag)

			if (err != nil) != tt.wantErr {
				t.Fatalf("RequiredIf(%q).ValidateWithInput(%v) error = %v, wantErr %v", tt.condition, tt.value, err, tt.wantErr)
			}
			if tt.wantIs != nil && !errors.Is(err, tt.wantIs) {
				t.Errorf("error = %v, want errors.Is %v", err, tt.wantIs)
			}
			if tt.condition == "unknown(field)" && err != nil {
				var syntaxErr RuleSyntaxError
				if !errors.As(err, &syntaxErr) {
					t.Errorf("expected RuleSyntaxError, got %T: %v", err, err)
				}
			}
		})
	}
}

func TestNotEmpty(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"hello", false},
		{1, false},
		{true, false},
		{0.1, false},
		{nil, true},
		{"", true},
		{0, true},
		{false, true},
		{0.0, true},
	}
	for _, tt := range tests {
		err := NotEmpty.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("NotEmpty.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && !errors.Is(err, ErrValidationNotEmpty) {
			t.Errorf("NotEmpty.Validate(%v) wrong error: %v", tt.value, err)
		}
	}
}
