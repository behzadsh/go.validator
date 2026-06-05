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
		t.Run(
			tt.name, func(t *testing.T) {
				rule := RequiredIf(tt.condition)
				bag := NewInputBag(tt.input)
				err := rule.ValidateWithInput(tt.value, bag)

				if (err != nil) != tt.wantErr {
					t.Fatalf(
						"RequiredIf(%q).ValidateWithInput(%v) error = %v, wantErr %v",
						tt.condition,
						tt.value,
						err,
						tt.wantErr,
					)
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
			},
		)
	}
}

func TestRequiredUnless(t *testing.T) {
	tests := []struct {
		name      string
		condition string
		value     any
		input     map[string]any
		wantErr   bool
		wantIs    error
	}{
		{
			name:      "condition true → field optional, nil passes",
			condition: `type == "guest"`,
			value:     nil,
			input:     map[string]any{"type": "guest"},
			wantErr:   false,
		},
		{
			name:      "condition false → field required, nil fails",
			condition: `type == "guest"`,
			value:     nil,
			input:     map[string]any{"type": "member"},
			wantErr:   true,
			wantIs:    ErrValidationRequiredUnless,
		},
		{
			name:      "condition false → field required, present passes",
			condition: `type == "guest"`,
			value:     "value",
			input:     map[string]any{"type": "member"},
			wantErr:   false,
		},
		{
			name:      "invalid condition → RuleSyntaxError",
			condition: "unknown(field)",
			value:     nil,
			input:     map[string]any{},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				rule := RequiredUnless(tt.condition)
				bag := NewInputBag(tt.input)
				err := rule.ValidateWithInput(tt.value, bag)

				if (err != nil) != tt.wantErr {
					t.Fatalf("error = %v, wantErr %v", err, tt.wantErr)
				}
				if tt.wantIs != nil && !errors.Is(err, tt.wantIs) {
					t.Errorf("error = %v, want errors.Is %v", err, tt.wantIs)
				}
			},
		)
	}
}

func TestRequiredWith(t *testing.T) {
	tests := []struct {
		name    string
		fields  []string
		value   any
		input   map[string]any
		wantErr bool
	}{
		{"any present, value present → pass", []string{"phone"}, "val", map[string]any{"phone": "123"}, false},
		{"any present, value nil → fail", []string{"phone"}, nil, map[string]any{"phone": "123"}, true},
		{"none present, value nil → pass", []string{"phone"}, nil, map[string]any{}, false},
		{"second present, value nil → fail", []string{"email", "phone"}, nil, map[string]any{"phone": "123"}, true},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				bag := NewInputBag(tt.input)
				err := RequiredWith(tt.fields...).ValidateWithInput(tt.value, bag)
				if (err != nil) != tt.wantErr {
					t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				}
				if err != nil && !errors.Is(err, ErrValidationRequiredWith) {
					t.Errorf("wrong error type: %v", err)
				}
			},
		)
	}
}

func TestRequiredWithAll(t *testing.T) {
	tests := []struct {
		name    string
		fields  []string
		value   any
		input   map[string]any
		wantErr bool
	}{
		{"all present, value present → pass", []string{"a", "b"}, "val", map[string]any{"a": 1, "b": 2}, false},
		{"all present, value nil → fail", []string{"a", "b"}, nil, map[string]any{"a": 1, "b": 2}, true},
		{"one missing, value nil → pass", []string{"a", "b"}, nil, map[string]any{"a": 1}, false},
		{"none present, value nil → pass", []string{"a", "b"}, nil, map[string]any{}, false},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				bag := NewInputBag(tt.input)
				err := RequiredWithAll(tt.fields...).ValidateWithInput(tt.value, bag)
				if (err != nil) != tt.wantErr {
					t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				}
				if err != nil && !errors.Is(err, ErrValidationRequiredWithAll) {
					t.Errorf("wrong error type: %v", err)
				}
			},
		)
	}
}

func TestRequiredWithout(t *testing.T) {
	tests := []struct {
		name    string
		fields  []string
		value   any
		input   map[string]any
		wantErr bool
	}{
		{"one absent, value nil → fail", []string{"email", "phone"}, nil, map[string]any{"email": "a@b.com"}, true},
		{
			"one absent, value present → pass",
			[]string{"email", "phone"},
			"val",
			map[string]any{"email": "a@b.com"},
			false,
		},
		{
			"all present, value nil → pass",
			[]string{"email", "phone"},
			nil,
			map[string]any{"email": "a@b.com", "phone": "123"},
			false,
		},
		{"all absent, value nil → fail", []string{"email", "phone"}, nil, map[string]any{}, true},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				bag := NewInputBag(tt.input)
				err := RequiredWithout(tt.fields...).ValidateWithInput(tt.value, bag)
				if (err != nil) != tt.wantErr {
					t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				}
				if err != nil && !errors.Is(err, ErrValidationRequiredWithout) {
					t.Errorf("wrong error type: %v", err)
				}
			},
		)
	}
}

func TestRequiredWithoutAll(t *testing.T) {
	tests := []struct {
		name    string
		fields  []string
		value   any
		input   map[string]any
		wantErr bool
	}{
		{"all absent, value nil → fail", []string{"email", "phone"}, nil, map[string]any{}, true},
		{"all absent, value present → pass", []string{"email", "phone"}, "val", map[string]any{}, false},
		{"one present, value nil → pass", []string{"email", "phone"}, nil, map[string]any{"email": "a@b.com"}, false},
		{
			"all present, value nil → pass",
			[]string{"email", "phone"},
			nil,
			map[string]any{"email": "a@b.com", "phone": "123"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				bag := NewInputBag(tt.input)
				err := RequiredWithoutAll(tt.fields...).ValidateWithInput(tt.value, bag)
				if (err != nil) != tt.wantErr {
					t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				}
				if err != nil && !errors.Is(err, ErrValidationRequiredWithoutAll) {
					t.Errorf("wrong error type: %v", err)
				}
			},
		)
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
