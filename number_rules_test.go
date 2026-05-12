package validation

import (
	"errors"
	"testing"
)

func TestNumeric(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{int(1), false},
		{int8(2), false},
		{int16(3), false},
		{int32(4), false},
		{int64(5), false},
		{uint(6), false},
		{uint8(7), false},
		{uint16(8), false},
		{uint32(9), false},
		{uint64(10), false},
		{float32(1.5), false},
		{float64(2.5), false},
		{complex64(1 + 2i), false},
		{complex128(3 + 4i), false},
		{"3.14", false},
		{"42", false},
		{"not a number", true},
		{"", true},
		{true, true},
		{nil, true},
		{[]int{1}, true},
	}
	for _, tt := range tests {
		err := Numeric.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("Numeric.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && !errors.Is(err, ErrValidationNumeric) {
			t.Errorf("Numeric.Validate(%v) wrong error: %v", tt.value, err)
		}
	}
}

func TestBetween(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		rule := Between[int](1, 10)
		tests := []struct {
			value   any
			wantErr bool
		}{
			{1, false},
			{5, false},
			{10, false},
			{0, true},
			{11, true},
			{float64(5), true},  // wrong type
			{nil, true},
		}
		for _, tt := range tests {
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Between[int](1,10).Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
			}
			if err != nil && !errors.Is(err, ErrValidationBetween) {
				t.Errorf("wrong error type: %v", err)
			}
		}
	})

	t.Run("float64", func(t *testing.T) {
		rule := Between[float64](0.0, 1.0)
		tests := []struct {
			value   any
			wantErr bool
		}{
			{0.0, false},
			{0.5, false},
			{1.0, false},
			{-0.1, true},
			{1.1, true},
			{int(1), true},  // wrong type
		}
		for _, tt := range tests {
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Between[float64](0,1).Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
			}
		}
	})
}

func TestMin(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		rule := Min[int](18)
		tests := []struct {
			value   any
			wantErr bool
		}{
			{18, false},
			{21, false},
			{17, true},
			{0, true},
			{float64(20), true},
			{nil, true},
		}
		for _, tt := range tests {
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Min[int](18).Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
			}
			if err != nil && !errors.Is(err, ErrValidationMin) {
				t.Errorf("wrong error type: %v", err)
			}
		}
	})
}

func TestMax(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		rule := Max[int](100)
		tests := []struct {
			value   any
			wantErr bool
		}{
			{100, false},
			{50, false},
			{101, true},
			{float64(50), true},
			{nil, true},
		}
		for _, tt := range tests {
			err := rule.Validate(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Max[int](100).Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
			}
			if err != nil && !errors.Is(err, ErrValidationMax) {
				t.Errorf("wrong error type: %v", err)
			}
		}
	})
}
