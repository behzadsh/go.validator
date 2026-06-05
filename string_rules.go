package validation

import (
	"net"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

const (
	emailUsernameRegexPattern = "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+$"
	emailDomainRegexPattern   = "^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*\\.[a-zA-Z]{2,}$"
)

var (
	regexAlpha         = regexp.MustCompile(`^[\pL\pM]+$`)
	regexAlphaDash     = regexp.MustCompile(`^[\pL\pM\pN_-]+$`)
	regexAlphaNum      = regexp.MustCompile(`^[\pL\pM\pN]+$`)
	regexAlphaSpace    = regexp.MustCompile(`^[\pL\pM\s]+$`)
	regexEmailUsername = regexp.MustCompile(emailUsernameRegexPattern)
	regexEmailDomain   = regexp.MustCompile(emailDomainRegexPattern)
	regexUUID          = regexp.MustCompile(`(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
)

// Alpha is a Rule that validates that the value is a string containing only Unicode letters.
//
// Fails if:
//   - value is not a string
//   - value contains digits, spaces, punctuation, or any non-letter character
//   - value is an empty string
//
// Examples:
//
//	validation.Alpha.Validate("hello")      // pass
//	validation.Alpha.Validate("Ünïcödé")    // pass — Unicode letters accepted
//	validation.Alpha.Validate("hello1")     // fail — contains digit
//	validation.Alpha.Validate("hi there")   // fail — contains space
var Alpha Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok || !regexAlpha.MatchString(str) {
			return ErrValidationAlpha
		}

		return nil
	},
)

// AlphaDash is a Rule that validates that the value is a string containing only Unicode letters, digits, underscores,
// and dashes.
//
// Fails if:
//   - value is not a string
//   - value contains spaces, punctuation other than _ and -, or any other non-alphanumeric character
//   - value is an empty string
//
// Examples:
//
//	validation.AlphaDash.Validate("hello-world") // pass
//	validation.AlphaDash.Validate("hello_world") // pass
//	validation.AlphaDash.Validate("hello123")    // pass
//	validation.AlphaDash.Validate("hello world") // fail — space not allowed
//	validation.AlphaDash.Validate("hello@world") // fail — @ not allowed
var AlphaDash Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok || !regexAlphaDash.MatchString(str) {
			return ErrValidationAlphaDash
		}

		return nil
	},
)

// AlphaNum is a Rule that validates that the value is a string containing only Unicode letters and digits.
//
// Fails if:
//   - value is not a string
//   - value contains spaces, dashes, underscores, or any non-alphanumeric character
//   - value is an empty string
//
// Examples:
//
//	validation.AlphaNum.Validate("hello123") // pass
//	validation.AlphaNum.Validate("ABC")      // pass
//	validation.AlphaNum.Validate("hello-1")  // fail — dash not allowed
//	validation.AlphaNum.Validate("hello 1")  // fail — space not allowed
var AlphaNum Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok || !regexAlphaNum.MatchString(str) {
			return ErrValidationAlphaNum
		}

		return nil
	},
)

// AlphaSpace is a Rule that validates that the value is a string containing only Unicode letters and whitespace.
//
// Fails if:
//   - value is not a string
//   - value contains digits, punctuation, or any non-letter non-whitespace character
//   - value is an empty string
//
// Examples:
//
//	validation.AlphaSpace.Validate("hello world") // pass
//	validation.AlphaSpace.Validate("Ünïcödé")     // pass
//	validation.AlphaSpace.Validate("hello1")      // fail — digit not allowed
//	validation.AlphaSpace.Validate("hello-world") // fail — dash not allowed
var AlphaSpace Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok || !regexAlphaSpace.MatchString(str) {
			return ErrValidationAlphaSpace
		}

		return nil
	},
)

// Email is a Rule that validates the value is a well-formed email address.
//
// Fails if:
//   - value is not a string
//   - value has no "@" or more than one "@"
//   - the username portion contains characters outside the allowed RFC set
//   - the domain portion has no dot, empty labels, or a TLD shorter than two characters
//
// Examples:
//
//	validation.Email.Validate("user@example.com")        // pass
//	validation.Email.Validate("user+tag@sub.example.org") // pass
//	validation.Email.Validate("notanemail")              // fail — no @
//	validation.Email.Validate("user@")                   // fail — empty domain
//	validation.Email.Validate("@example.com")            // fail — empty username
var Email Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok || !isEmail(str) {
			return ErrValidationEmail
		}

		return nil
	},
)

// EmailMX is a Rule that validates the value is a well-formed email address whose domain has at least one MX record.
//
// Fails if:
//   - value is not a string or is not a valid email format (returns ErrValidationEmail)
//   - the domain has no MX records (returns ErrValidationEmailMX)
//
// Note: this rule performs a network call on every invocation. Avoid in hot paths; cache results externally if needed.
//
// Examples:
//
//	validation.EmailMX.Validate("user@gmail.com")   // pass — gmail.com has MX records
//	validation.EmailMX.Validate("user@invalid.com") // fail — example.com has no MX records
//	validation.EmailMX.Validate("notanemail")       // fail — format invalid
var EmailMX Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok || !isEmail(str) {
			return ErrValidationEmail
		}

		domain := strings.SplitN(str, "@", 2)[1]
		if _, err := net.LookupMX(domain); err != nil {
			return ErrValidationEmailMX
		}

		return nil
	},
)

// EndsWith returns a Rule that validates the value is a string ending with the given suffix.
//
// Fails if:
//   - value is not a string
//   - the string does not end with suffix
//
// Examples:
//
//	validation.EndsWith(".go").Validate("main.go") // pass
//	validation.EndsWith(".go").Validate("main.js") // fail
func EndsWith(suffix string) Rule {
	return RuleFunc(func(value any) error {
		str, ok := value.(string)
		if !ok || !strings.HasSuffix(str, suffix) {
			return ErrValidationEndsWith
		}

		return nil
	})
}

// Length returns a Rule that validates the string's rune count is exactly equal to l.
//
// Rune count is used, not byte length, so multibyte characters such as "é" count as one.
//
// Fails if:
//   - value is not a string
//   - the string has fewer or more runes than l
//
// Examples:
//
//	validation.Length(5).Validate("hello")    // pass — 5 runes
//	validation.Length(5).Validate("héllo")    // pass — 5 runes (é is one rune)
//	validation.Length(5).Validate("hi")       // fail — 2 runes
//	validation.Length(5).Validate("too long") // fail — 8 runes
func Length(l int) Rule {
	return RuleFunc(
		func(value any) error {
			str, ok := value.(string)
			if !ok || utf8.RuneCountInString(str) != l {
				return ErrValidationLength
			}

			return nil
		},
	)
}

// Lowercase is a Rule that validates the value is a string containing only lowercase characters.
//
// Fails if:
//   - value is not a string
//   - the string contains any uppercase character
//
// Examples:
//
//	validation.Lowercase.Validate("hello")       // pass
//	validation.Lowercase.Validate("hello world") // pass
//	validation.Lowercase.Validate("Hello")       // fail
//	validation.Lowercase.Validate("HELLO")       // fail
var Lowercase Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok || str != strings.ToLower(str) {
			return ErrValidationLowercase
		}

		return nil
	},
)

// MaxLength returns a Rule that validates the string's rune count is at most l.
//
// Rune count is used, not byte length, so multibyte characters such as "é" count as one.
//
// Fails if:
//   - value is not a string
//   - the string has more than l runes
//
// Examples:
//
//	validation.MaxLength(10).Validate("hello")        // pass — 5 runes <= 10
//	validation.MaxLength(3).Validate("too long")      // fail — 8 runes > 3
//	validation.MaxLength(3).Validate("héé")           // pass — 3 runes
func MaxLength(l int) Rule {
	return RuleFunc(
		func(value any) error {
			str, ok := value.(string)
			if !ok || utf8.RuneCountInString(str) > l {
				return ErrValidationMaxLength
			}

			return nil
		},
	)
}

// MinLength returns a Rule that validates the string's rune count is at least l.
//
// Rune count is used, not byte length, so multibyte characters such as "é" count as one.
//
// Fails if:
//   - value is not a string
//   - the string has fewer than l runes
//
// Examples:
//
//	validation.MinLength(3).Validate("hello") // pass — 5 runes >= 3
//	validation.MinLength(3).Validate("hi")    // fail — 2 runes < 3
//	validation.MinLength(3).Validate("héé")   // pass — 3 runes
func MinLength(l int) Rule {
	return RuleFunc(
		func(value any) error {
			str, ok := value.(string)
			if !ok || utf8.RuneCountInString(str) < l {
				return ErrValidationMinLength
			}

			return nil
		},
	)
}

// NotRegex returns a Rule that validates the value is a string that does NOT match the given regular expression.
//
// The pattern is compiled once at call time. If the pattern is invalid, Schema.Validate returns a
// RuleSyntaxError — treat that as a programming error and fix the schema at startup.
//
// Fails if:
//   - value is not a string
//   - the string matches the pattern
//
// Examples:
//
//	validation.NotRegex(`\s`).Validate("nospaces")    // pass — no whitespace
//	validation.NotRegex(`\s`).Validate("has spaces")  // fail — contains whitespace
func NotRegex(pattern string) Rule {
	re, err := regexp.Compile(pattern)

	return RuleFunc(func(value any) error {
		if err != nil {
			return RuleSyntaxError{Rule: "NotRegex", Err: err}
		}

		str, ok := value.(string)
		if !ok || re.MatchString(str) {
			return ErrValidationNotRegex
		}

		return nil
	})
}

// Regex returns a Rule that validates the value is a string matching the given regular expression.
//
// The pattern is compiled once at call time. If the pattern is invalid, Schema.Validate returns a
// RuleSyntaxError — treat that as a programming error and fix the schema at startup.
//
// Fails if:
//   - value is not a string
//   - the string does not match the pattern
//
// Examples:
//
//	validation.Regex(`^\d{4}$`).Validate("1234")  // pass
//	validation.Regex(`^\d{4}$`).Validate("12345") // fail — too many digits
//	validation.Regex(`^\d{4}$`).Validate("abcd")  // fail — not digits
func Regex(pattern string) Rule {
	re, err := regexp.Compile(pattern)

	return RuleFunc(func(value any) error {
		if err != nil {
			return RuleSyntaxError{Rule: "Regex", Err: err}
		}

		str, ok := value.(string)
		if !ok || !re.MatchString(str) {
			return ErrValidationRegex
		}

		return nil
	})
}

// StartsWith returns a Rule that validates the value is a string beginning with the given prefix.
//
// Fails if:
//   - value is not a string
//   - the string does not start with prefix
//
// Examples:
//
//	validation.StartsWith("SKU-").Validate("SKU-001") // pass
//	validation.StartsWith("SKU-").Validate("001-SKU") // fail
func StartsWith(prefix string) Rule {
	return RuleFunc(func(value any) error {
		str, ok := value.(string)
		if !ok || !strings.HasPrefix(str, prefix) {
			return ErrValidationStartsWith
		}

		return nil
	})
}

// Uppercase is a Rule that validates the value is a string containing only uppercase characters.
//
// Fails if:
//   - value is not a string
//   - the string contains any lowercase character
//
// Examples:
//
//	validation.Uppercase.Validate("HELLO")       // pass
//	validation.Uppercase.Validate("HELLO WORLD") // pass
//	validation.Uppercase.Validate("Hello")       // fail
//	validation.Uppercase.Validate("hello")       // fail
var Uppercase Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok || str != strings.ToUpper(str) {
			return ErrValidationUppercase
		}

		return nil
	},
)

// URL is a Rule that validates that the value is a string that can be parsed as valid absolute URL; schema-less URLs
// are also accepted.
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

// UUID is a Rule that validates the value is a valid UUID string (any variant, case-insensitive).
//
// Fails if:
//   - value is not a string
//   - the string does not match the UUID format xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
//
// Examples:
//
//	validation.UUID.Validate("550e8400-e29b-41d4-a716-446655440000") // pass
//	validation.UUID.Validate("not-a-uuid")                           // fail
var UUID Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok || !regexUUID.MatchString(str) {
			return ErrValidationUUID
		}

		return nil
	},
)

func isEmail(str string) bool {
	parts := strings.Split(str, "@")
	if len(parts) != 2 {
		return false
	}

	if !regexEmailUsername.MatchString(parts[0]) {
		return false
	}

	if !regexEmailDomain.MatchString(parts[1]) {
		return false
	}

	return true
}

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
