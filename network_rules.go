package validation

import "net"

// IP is a Rule that validates the value is a valid IP address (v4 or v6).
//
// Fails if:
//   - value is not a string
//   - the string cannot be parsed as a valid IPv4 or IPv6 address
//
// Examples:
//
//	validation.IP.Validate("192.168.1.1")  // pass — IPv4
//	validation.IP.Validate("::1")          // pass — IPv6
//	validation.IP.Validate("999.0.0.1")    // fail
//	validation.IP.Validate("not-an-ip")    // fail
var IP Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok || net.ParseIP(str) == nil {
			return ErrValidationIP
		}

		return nil
	},
)

// IPv4 is a Rule that validates the value is a valid IPv4 address.
//
// Fails if:
//   - value is not a string
//   - the string is not a valid IPv4 address (IPv6 addresses fail)
//
// Examples:
//
//	validation.IPv4.Validate("192.168.1.1") // pass
//	validation.IPv4.Validate("::1")         // fail — IPv6
//	validation.IPv4.Validate("not-an-ip")   // fail
var IPv4 Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok {
			return ErrValidationIPv4
		}

		ip := net.ParseIP(str)
		if ip == nil || ip.To4() == nil {
			return ErrValidationIPv4
		}

		return nil
	},
)

// IPv6 is a Rule that validates the value is a valid IPv6 address.
//
// Fails if:
//   - value is not a string
//   - the string is not a valid IPv6 address (IPv4 addresses fail)
//
// Examples:
//
//	validation.IPv6.Validate("::1")           // pass
//	validation.IPv6.Validate("2001:db8::1")   // pass
//	validation.IPv6.Validate("192.168.1.1")   // fail — IPv4
//	validation.IPv6.Validate("not-an-ip")     // fail
var IPv6 Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok {
			return ErrValidationIPv6
		}

		ip := net.ParseIP(str)
		if ip == nil || ip.To4() != nil {
			return ErrValidationIPv6
		}

		return nil
	},
)

// MACAddress is a Rule that validates the value is a valid 6-byte MAC address.
//
// Accepted formats (via net.ParseMAC): 01:23:45:67:89:ab and 01-23-45-67-89-ab (case-insensitive).
// 8-byte EUI-64 addresses are rejected.
//
// Fails if:
//   - value is not a string
//   - the string is not a valid 6-byte MAC address
//
// Examples:
//
//	validation.MACAddress.Validate("01:23:45:67:89:ab") // pass
//	validation.MACAddress.Validate("01-23-45-67-89-AB") // pass
//	validation.MACAddress.Validate("not-a-mac")         // fail
var MACAddress Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok {
			return ErrValidationMACAddress
		}

		hw, err := net.ParseMAC(str)
		if err != nil || len(hw) != 6 {
			return ErrValidationMACAddress
		}

		return nil
	},
)
