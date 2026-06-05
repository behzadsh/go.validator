package validation

import (
	"errors"
	"testing"
)

func TestMACAddress(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"01:23:45:67:89:ab", false},
		{"01-23-45-67-89-AB", false},
		{"FF:FF:FF:FF:FF:FF", false},
		{"not-a-mac", true},
		{"01:23:45:67:89", true},          // too short
		{"01:23:45:67:89:ab:cd:ef", true}, // 8-byte EUI-64
		{"", true},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := MACAddress.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("MACAddress.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && !errors.Is(err, ErrValidationMACAddress) {
			t.Errorf("MACAddress.Validate(%v) wrong error: %v", tt.value, err)
		}
	}
}

func TestIP(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"192.168.1.1", false},
		{"10.0.0.1", false},
		{"0.0.0.0", false},
		{"255.255.255.255", false},
		{"::1", false},
		{"2001:db8::1", false},
		{"fe80::1", false},
		{"999.0.0.1", true},
		{"192.168.1", true},
		{"not-an-ip", true},
		{"", true},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := IP.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("IP.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && !errors.Is(err, ErrValidationIP) {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestIPv4(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"192.168.1.1", false},
		{"0.0.0.0", false},
		{"255.255.255.255", false},
		{"::1", true},           // IPv6 fails
		{"2001:db8::1", true},   // IPv6 fails
		{"999.0.0.1", true},
		{"not-an-ip", true},
		{"", true},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := IPv4.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("IPv4.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && !errors.Is(err, ErrValidationIPv4) {
			t.Errorf("wrong error type: %v", err)
		}
	}
}

func TestIPv6(t *testing.T) {
	tests := []struct {
		value   any
		wantErr bool
	}{
		{"::1", false},
		{"2001:db8::1", false},
		{"fe80::1", false},
		{"192.168.1.1", true},  // IPv4 fails
		{"0.0.0.0", true},      // IPv4 fails
		{"not-an-ip", true},
		{"", true},
		{nil, true},
		{42, true},
	}
	for _, tt := range tests {
		err := IPv6.Validate(tt.value)
		if (err != nil) != tt.wantErr {
			t.Errorf("IPv6.Validate(%v) error = %v, wantErr %v", tt.value, err, tt.wantErr)
		}
		if err != nil && !errors.Is(err, ErrValidationIPv6) {
			t.Errorf("wrong error type: %v", err)
		}
	}
}
