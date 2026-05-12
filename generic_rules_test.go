package validation

import (
	"errors"
	"testing"
)

func TestIn(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		rule := In([]string{"a", "b", "c"})
		tests := []struct {
			value   any
			wantErr bool
		}{
			{"a", false},
			{"b", false},
			{"c", false},
			{"d", true},
			{"", true},
			{nil, true},
			{1, true},
		}
		for _, tt := range tests {
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("In(strings).Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
			}
			if err != nil && !errors.Is(err, ErrValidationIn) {
				t.Errorf("wrong error type: %v", err)
			}
		}
	})

	t.Run("int", func(t *testing.T) {
		rule := In([]int{1, 2, 3})
		tests := []struct {
			value   any
			wantErr bool
		}{
			{1, false},
			{3, false},
			{4, true},
			{float64(1), true},  // wrong type
		}
		for _, tt := range tests {
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("In(ints).Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
			}
		}
	})
}

func TestNotIn(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		rule := NotIn([]string{"banned", "forbidden"})
		tests := []struct {
			value   any
			wantErr bool
		}{
			{"allowed", false},
			{"ok", false},
			{"banned", true},
			{"forbidden", true},
			{nil, true},    // type mismatch returns error per current behavior
			{42, true},     // type mismatch returns error per current behavior
		}
		for _, tt := range tests {
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotIn(strings).Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
			}
		}
	})
}
