package validation

import (
	"net"
	"net/url"
	"strings"
)

// CIDR is a Rule that validates the value is a valid CIDR notation string (e.g. "192.168.0.0/24" or "2001:db8::/32").
//
// Fails if:
//   - value is not a string
//   - the string cannot be parsed as a CIDR block
//
// Examples:
//
//	validation.CIDR.Validate("192.168.0.0/24")  // pass
//	validation.CIDR.Validate("2001:db8::/32")   // pass — IPv6
//	validation.CIDR.Validate("192.168.0.1")     // fail — no prefix length
//	validation.CIDR.Validate("not-cidr")        // fail
var CIDR Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok {
			return ErrValidationCIDR
		}

		if _, _, err := net.ParseCIDR(str); err != nil {
			return ErrValidationCIDR
		}

		return nil
	},
)

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

// URL is a Rule that validates that the value is a string that can be parsed as a valid absolute URL;
// scheme-less URLs are also accepted.
//
// The host must be a valid domain, IP address, or "localhost".
//
// Fails if:
//   - value is not a string
//   - value is an empty string
//   - value has no resolvable host (e.g. "http://")
//   - value contains characters that make it unparseable (e.g. unencoded spaces)
//
// Examples:
//
//	validation.URL.Validate("https://example.com")   // pass
//	validation.URL.Validate("example.com/path")      // pass — scheme inferred
//	validation.URL.Validate("http://localhost:8080") // pass
//	validation.URL.Validate("http://[::1]:8080/api") // pass — IPv6
//	validation.URL.Validate("not a url")             // fail — unparseable
//	validation.URL.Validate("http://")               // fail — no host
var URL Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok {
			return ErrValidationURL
		}

		if u, err := url.ParseRequestURI(str); err == nil && isValidURLHost(u.Host) {
			return nil
		}
		if u, err := url.ParseRequestURI("http://" + str); err == nil && isValidURLHost(u.Host) {
			return nil
		}

		return ErrValidationURL
	},
)

func isValidURLHost(host string) bool {
	if host == "" {
		return false
	}

	// Strip port if present (e.g. "example.com:8080" or "[::1]:8080")
	if strings.HasPrefix(host, "[") {
		// IPv6 in brackets: keep address, drop port.
		if idx := strings.LastIndex(host, "]"); idx != -1 {
			host = host[1:idx]
		}
	} else if h, _, ok := strings.Cut(host, ":"); ok {
		host = h
	}

	if net.ParseIP(host) != nil {
		return true
	}

	if host == "localhost" {
		return true
	}

	if !strings.Contains(host, ".") {
		return false
	}

	parts := strings.Split(host, ".")
	for _, p := range parts {
		if p == "" {
			return false
		}
	}

	return true
}
