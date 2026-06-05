package validation

import (
	"errors"
	"testing"
)

func TestDistinct(t *testing.T) {
	tests := []struct {
		name    string
		value   any
		wantErr bool
	}{
		{"unique strings", []string{"a", "b", "c"}, false},
		{"unique ints", []int{1, 2, 3}, false},
		{"single element", []int{1}, false},
		{"empty slice", []int{}, false},
		{"duplicate ints", []int{1, 2, 1}, true},
		{"duplicate strings", []string{"a", "b", "a"}, true},
		{"non-slice passes", "not-a-slice", false},
		{"nil passes", nil, false},
		{"scalar passes", 42, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Distinct.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Distinct.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
			}
			if err != nil && !errors.Is(err, ErrValidationDistinct) {
				t.Errorf("wrong error type: %v", err)
			}
		})
	}
}

func TestMinSize(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{[]int{1, 2, 3}, false},
		{[]int{1, 2}, false},
		{[]int{1}, true},
		{[]string{}, true},
		{nil, false},
		{"not-a-slice", false},
		{42, false},
	}
	for _, tt := range tests {
		err := MinSize(2).Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("MinSize(2).Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && !errors.Is(err, ErrValidationMinSize) {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestMaxSize(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{[]int{1, 2}, false},
		{[]int{1, 2, 3}, false},
		{[]int{1, 2, 3, 4}, true},
		{[]string{}, false},
		{nil, false},
		{"not-a-slice", false},
		{42, false},
	}
	for _, tt := range tests {
		err := MaxSize(3).Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("MaxSize(3).Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && !errors.Is(err, ErrValidationMaxSize) {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestSize(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{[]int{1, 2, 3}, false},
		{[]int{1, 2}, true},
		{[]int{1, 2, 3, 4}, true},
		{[]string{}, true},
		{nil, false},
		{"not-a-slice", false},
	}
	for _, tt := range tests {
		err := Size(3).Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("Size(3).Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && !errors.Is(err, ErrValidationSize) {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestEach(t *testing.T) {
	t.Run("all pass", func(t *testing.T) {
		err := Each(MinLength(2)).Validate([]string{"ab", "cd", "ef"})
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("one fails", func(t *testing.T) {
		err := Each(MinLength(2)).Validate([]string{"ab", "x", "cd"})
		if err == nil {
			t.Error("expected error, got nil")
		}
		if !errors.Is(err, ErrValidationEach) {
			t.Errorf("wrong error type: %v", err)
		}
	})

	t.Run("nil passes", func(t *testing.T) {
		err := Each(MinLength(2)).Validate(nil)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("non-slice passes", func(t *testing.T) {
		err := Each(MinLength(2)).Validate("not-a-slice")
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("empty slice passes", func(t *testing.T) {
		err := Each(MinLength(2)).Validate([]string{})
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})
}

func TestEach_WithInputRule(t *testing.T) {
	// Verify Each dispatches InputRules correctly (cross-field dispatch keystone).
	schema := New().
		Field("items", Each(MinLength(2), Lowercase))

	t.Run("all valid", func(t *testing.T) {
		res, _ := schema.Validate(map[string]any{"items": []string{"ab", "cd"}})
		if res.HasErrors() {
			t.Errorf("expected no errors, got %v", res.Errors())
		}
	})

	t.Run("element fails rule", func(t *testing.T) {
		res, _ := schema.Validate(map[string]any{"items": []string{"ab", "X"}})
		if !res.HasErrors() {
			t.Error("expected errors, got none")
		}
	})
}
