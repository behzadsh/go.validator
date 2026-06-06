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
		if err != nil && errorCode(err) != "numeric" {
			t.Errorf("Numeric.Validate(%v) wrong error: %v", tt.value, err)
		}
	}
}

func TestBetween(t *testing.T) {
	t.Run(
		"int", func(t *testing.T) {
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
				{float64(5), true}, // wrong type
				{nil, true},
			}
			for _, tt := range tests {
				err := rule.Validate(tt.value)
				if (err != nil) != tt.wantErr {
					t.Errorf("Between[int](1,10).Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
				}
				if err != nil && errorCode(err) != "between" {
					t.Errorf("wrong error type: %v", err)
				}
			}
		},
	)

	t.Run(
		"float64", func(t *testing.T) {
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
				{int(1), true}, // wrong type
			}
			for _, tt := range tests {
				err := rule.Validate(tt.value)
				if (err != nil) != tt.wantErr {
					t.Errorf("Between[float64](0,1).Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
				}
			}
		},
	)
}

func TestMin(t *testing.T) {
	t.Run(
		"int", func(t *testing.T) {
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
				if err != nil && errorCode(err) != "min" {
					t.Errorf("wrong error type: %v", err)
				}
			}
		},
	)
}

func TestMax(t *testing.T) {
	t.Run(
		"int", func(t *testing.T) {
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
				if err != nil && errorCode(err) != "max" {
					t.Errorf("wrong error type: %v", err)
				}
			}
		},
	)
}

func TestGT(t *testing.T) {
	rule := GT[int](18)
	tests := []struct {
		value   any
		wantErr bool
	}{
		{19, false},
		{100, false},
		{18, true}, // equal fails
		{17, true},
		{float64(19), true},
		{nil, true},
	}
	for _, tt := range tests {
		err := rule.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("GT[int](18).Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "gt" {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestGTE(t *testing.T) {
	rule := GTE[int](18)
	tests := []struct {
		value   any
		wantErr bool
	}{
		{18, false}, // equal passes
		{19, false},
		{17, true},
		{float64(18), true},
		{nil, true},
	}
	for _, tt := range tests {
		err := rule.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("GTE[int](18).Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "gte" {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestLT(t *testing.T) {
	rule := LT[int](100)
	tests := []struct {
		value   any
		wantErr bool
	}{
		{99, false},
		{0, false},
		{100, true}, // equal fails
		{101, true},
		{float64(99), true},
		{nil, true},
	}
	for _, tt := range tests {
		err := rule.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("LT[int](100).Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "lt" {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestLTE(t *testing.T) {
	rule := LTE[int](100)
	tests := []struct {
		value   any
		wantErr bool
	}{
		{100, false}, // equal passes
		{99, false},
		{101, true},
		{float64(100), true},
		{nil, true},
	}
	for _, tt := range tests {
		err := rule.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("LTE[int](100).Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "lte" {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestInteger(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{42, false},
		{int8(1), false},
		{int16(2), false},
		{int32(3), false},
		{int64(4), false},
		{uint(5), false},
		{uint8(6), false},
		{uint16(7), false},
		{uint32(8), false},
		{uint64(9), false},
		{nil, true},
		{3.14, true}, // float64
		{"42", true}, // string
		{true, true}, // bool
		{float32(1.0), true},
	}
	for _, tt := range tests {
		err := Integer.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("Integer.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "integer" {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestPositive(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{1, false},
		{0.1, false},
		{float64(5), false},
		{uint(3), false},
		{0, true},
		{-1, true},
		{-0.1, true},
		{nil, true},
		{"5", true},
		{true, true},
	}
	for _, tt := range tests {
		err := Positive.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("Positive.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "positive" {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestNegative(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{-1, false},
		{-0.1, false},
		{float64(-5), false},
		{0, true},
		{1, true},
		{uint(0), true},
		{nil, true},
		{"5", true},
	}
	for _, tt := range tests {
		err := Negative.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("Negative.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "negative" {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestNonNegative(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{0, false},
		{1, false},
		{0.0, false},
		{float64(5), false},
		{uint(0), false},
		{-1, true},
		{-0.1, true},
		{nil, true},
		{"0", true},
	}
	for _, tt := range tests {
		err := NonNegative.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("NonNegative.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "non_negative" {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestMultipleOf(t *testing.T) {
	t.Run(
		"int divisor", func(t *testing.T) {
			rule := MultipleOf[int](3)
			tests := []struct {
				value   any
				wantErr bool
			}{
				{9, false},
				{0, false},
				{-6, false},
				{float64(9), false}, // JSON number
				{8, true},
				{nil, true},
				{"9", true},
			}
			for _, tt := range tests {
				err := rule.Validate(tt.value)
				if (err != nil) != tt.wantErr {
					t.Errorf("MultipleOf[int](3).Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
				}
			}
		},
	)

	t.Run(
		"zero divisor returns syntax error", func(t *testing.T) {
			rule := MultipleOf[int](0)
			err := rule.Validate(9)
			if err == nil {
				t.Error("expected error for zero divisor")
			}
			var rse RuleSyntaxError
			if !errors.As(err, &rse) {
				t.Errorf("expected RuleSyntaxError, got %T", err)
			}
		},
	)
}

func TestPort(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{80, false},
		{443, false},
		{1, false},
		{65535, false},
		{float64(8080), false}, // JSON number
		{0, true},
		{65536, true},
		{-1, true},
		{80.5, true}, // fractional
		{nil, true},
		{"80", true},
	}
	for _, tt := range tests {
		err := Port.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("Port.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "port" {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestLatitude(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{0.0, false},
		{45.0, false},
		{-90.0, false},
		{90.0, false},
		{float64(45), false},
		{90.1, true},
		{-90.1, true},
		{nil, true},
		{"45", true},
	}
	for _, tt := range tests {
		err := Latitude.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("Latitude.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "latitude" {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestLongitude(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{0.0, false},
		{120.5, false},
		{-180.0, false},
		{180.0, false},
		{float64(120), false},
		{180.1, true},
		{-180.1, true},
		{nil, true},
		{"120", true},
	}
	for _, tt := range tests {
		err := Longitude.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("Longitude.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && errorCode(err) != "longitude" {
			t.Errorf("wrong error type: %v", err)
		}
	}
}
