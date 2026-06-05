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
