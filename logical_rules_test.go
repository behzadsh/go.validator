package validation

import (
	"errors"
	"testing"
)

func TestNot(t *testing.T) {
	tests := []struct {
		name    string
		rule    Rule
		value   any
		wantErr bool
	}{
		{"not email: non-email passes", Not(Email), "not-an-email", false},
		{"not email: email fails", Not(Email), "user@example.com", true},
		{"not uuid: non-uuid passes", Not(UUID), "plain-string", false},
		{"not uuid: uuid fails", Not(UUID), "550e8400-e29b-41d4-a716-446655440000", true},
		{"not alpha: digit string passes", Not(Alpha), "123", false},
		{"not alpha: alpha string fails", Not(Alpha), "hello", true},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				err := tt.rule.Validate(tt.value)
				if (err != nil) != tt.wantErr {
					t.Errorf("Not.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
				}
				if err != nil && !errors.Is(err, ErrValidationNot) {
					t.Errorf("Not.Validate(%v) wrong error: %v", tt.value, err)
				}
			},
		)
	}
}

func TestNot_WithInputRule(t *testing.T) {
	// Verify Not correctly dispatches InputRules (the cross-field dispatch keystone).
	schema := New().
		Field("password", Required).
		Field("new_password", Not(SameAs("password")))

	t.Run(
		"different values pass", func(t *testing.T) {
			res, _ := schema.Validate(
				map[string]any{
					"password":     "old",
					"new_password": "new",
				},
			)
			if res.HasErrors() {
				t.Errorf("expected no errors, got %v", res.Errors())
			}
		},
	)

	t.Run(
		"same values fail", func(t *testing.T) {
			res, _ := schema.Validate(
				map[string]any{
					"password":     "same",
					"new_password": "same",
				},
			)
			if !res.HasErrors() {
				t.Error("expected errors, got none")
			}
		},
	)
}

func TestAny(t *testing.T) {
	rule := Any(Email, PhoneE164)
	tests := []struct {
		name    string
		value   any
		wantErr bool
	}{
		{"email passes", "user@example.com", false},
		{"phone passes", "+14155552671", false},
		{"neither fails", "notvalid", true},
		{"nil fails", nil, true},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				err := rule.Validate(tt.value)
				if (err != nil) != tt.wantErr {
					t.Errorf("Any.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
				}
				if err != nil && !errors.Is(err, ErrValidationAny) {
					t.Errorf("Any.Validate(%v) wrong error: %v", tt.value, err)
				}
			},
		)
	}
}

func TestAny_WithInputRule(t *testing.T) {
	// Verify Any dispatches InputRules correctly.
	schema := New().
		Field("a", Required).
		Field("b", Any(SameAs("a"), MinLength(10)))

	t.Run(
		"same as a passes", func(t *testing.T) {
			res, _ := schema.Validate(map[string]any{"a": "hello", "b": "hello"})
			if res.HasErrors() {
				t.Errorf("expected no errors, got %v", res.Errors())
			}
		},
	)

	t.Run(
		"long enough passes", func(t *testing.T) {
			res, _ := schema.Validate(map[string]any{"a": "hello", "b": "verylongvalue"})
			if res.HasErrors() {
				t.Errorf("expected no errors, got %v", res.Errors())
			}
		},
	)

	t.Run(
		"neither passes fails", func(t *testing.T) {
			res, _ := schema.Validate(map[string]any{"a": "hello", "b": "other"})
			if !res.HasErrors() {
				t.Error("expected errors, got none")
			}
		},
	)
}

func TestWhen(t *testing.T) {
	schema := New().
		Field("plan", Required).
		Field("vat", When(`plan == "paid"`, MinLength(5)))

	t.Run(
		"condition true: rule applied", func(t *testing.T) {
			res, _ := schema.Validate(map[string]any{"plan": "paid", "vat": "AB"})
			if !res.HasErrors() {
				t.Error("expected errors, got none")
			}
		},
	)

	t.Run(
		"condition true: rule passes", func(t *testing.T) {
			res, _ := schema.Validate(map[string]any{"plan": "paid", "vat": "AB123"})
			if res.HasErrors() {
				t.Errorf("expected no errors, got %v", res.Errors())
			}
		},
	)

	t.Run(
		"condition false: rule skipped", func(t *testing.T) {
			res, _ := schema.Validate(map[string]any{"plan": "free", "vat": "X"})
			if res.HasErrors() {
				t.Errorf("expected no errors, got %v", res.Errors())
			}
		},
	)
}

func TestWhen_InvalidCondition(t *testing.T) {
	schema := New().
		Field("x", When(`!!!invalid`, MinLength(3)))

	res, _ := schema.Validate(map[string]any{"x": "hi"})
	if !res.HasErrors() {
		t.Error("expected syntax error")
	}
	if len(res.Errors()) > 0 {
		var rse RuleSyntaxError
		if !errors.As(res.Errors()[0].Err, &rse) {
			t.Errorf("expected RuleSyntaxError, got %T", res.Errors()[0].Err)
		}
	}
}

func TestUnless(t *testing.T) {
	schema := New().
		Field("status", Required).
		Field("reason", Unless(`status == "approved"`, MinLength(5)))

	t.Run(
		"condition false: rule applied", func(t *testing.T) {
			res, _ := schema.Validate(map[string]any{"status": "pending", "reason": "no"})
			if !res.HasErrors() {
				t.Error("expected errors, got none")
			}
		},
	)

	t.Run(
		"condition false: rule passes", func(t *testing.T) {
			res, _ := schema.Validate(map[string]any{"status": "pending", "reason": "under review"})
			if res.HasErrors() {
				t.Errorf("expected no errors, got %v", res.Errors())
			}
		},
	)

	t.Run(
		"condition true: rule skipped", func(t *testing.T) {
			res, _ := schema.Validate(map[string]any{"status": "approved", "reason": "no"})
			if res.HasErrors() {
				t.Errorf("expected no errors, got %v", res.Errors())
			}
		},
	)
}
