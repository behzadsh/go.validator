package validation

import (
	"encoding/base64"
	"encoding/json"
	"net"
	"regexp"
	"strings"
	"unicode/utf8"
)

const (
	emailUsernameRegexPattern = "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+$"
	emailDomainRegexPattern   = "^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*\\.[a-zA-Z]{2,}$"
	semverRegexPattern        = `^v?(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`
)

var (
	regexAlpha         = regexp.MustCompile(`^[\pL\pM]+$`)
	regexAlphaDash     = regexp.MustCompile(`^[\pL\pM\pN_-]+$`)
	regexAlphaNum      = regexp.MustCompile(`^[\pL\pM\pN]+$`)
	regexAlphaSpace    = regexp.MustCompile(`^[\pL\pM\s]+$`)
	regexEmailUsername = regexp.MustCompile(emailUsernameRegexPattern)
	regexEmailDomain   = regexp.MustCompile(emailDomainRegexPattern)
	regexHexColor      = regexp.MustCompile(`^#([0-9A-Fa-f]{3}|[0-9A-Fa-f]{6})$`)
	regexJWT           = regexp.MustCompile(`^[A-Za-z0-9_-]+\.[A-Za-z0-9_-]+\.[A-Za-z0-9_-]+$`)
	regexPhoneE164     = regexp.MustCompile(`^\+[1-9]\d{1,14}$`)
	regexSemver        = regexp.MustCompile(semverRegexPattern)
	regexSlug          = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
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

// ASCII is a Rule that validates the value is a string containing only ASCII characters (bytes 0–127).
//
// Fails if:
//   - value is not a string
//   - the string contains any byte with value > 127
//
// Examples:
//
//	validation.ASCII.Validate("hello")   // pass
//	validation.ASCII.Validate("café")    // fail — é is not ASCII
//	validation.ASCII.Validate("hello\t") // pass — tab is ASCII
var ASCII Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok {
			return ErrValidationASCII
		}

		for i := 0; i < len(str); i++ {
			if str[i] > 127 {
				return ErrValidationASCII
			}
		}

		return nil
	},
)

// Base64 returns a Rule that validates the value is a valid standard base64-encoded string (RFC 4648,
// with padding). URL-safe base64 (using - and _) is not accepted.
//
// An empty string passes (it encodes zero bytes).
//
// Fails if:
//   - value is not a string
//   - the string is not valid standard base64
//
// Examples:
//
//	validation.Base64.Validate("aGVsbG8=")     // pass — "hello"
//	validation.Base64.Validate("aGVsbG8")      // fail — missing padding
//	validation.Base64.Validate("not-base64!")  // fail
var Base64 Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok {
			return ErrValidationBase64
		}

		if _, err := base64.StdEncoding.DecodeString(str); err != nil {
			return ErrValidationBase64
		}

		return nil
	},
)

// Contains returns a Rule that validates the value is a string containing the given substring.
//
// Fails if:
//   - value is not a string
//   - the string does not contain sub
//
// Examples:
//
//	validation.Contains("world").Validate("hello world") // pass
//	validation.Contains("world").Validate("hello")       // fail
func Contains(sub string) Rule {
	return RuleFunc(
		func(value any) error {
			str, ok := value.(string)
			if !ok || !strings.Contains(str, sub) {
				return ErrValidationContains
			}

			return nil
		},
	)
}

// CreditCard is a Rule that validates the value is a valid credit card number using the Luhn algorithm.
//
// Spaces and dashes are stripped before validation. Accepted length after stripping: 13–19 digits.
//
// Fails if:
//   - value is not a string
//   - the string (after stripping spaces/dashes) is not 13–19 digits
//   - the Luhn checksum does not pass
//
// Examples:
//
//	validation.CreditCard.Validate("4111111111111111")    // pass — Visa test number
//	validation.CreditCard.Validate("4111-1111-1111-1111") // pass — dashes stripped
//	validation.CreditCard.Validate("1234567890123456")    // fail — invalid Luhn
var CreditCard Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok {
			return ErrValidationCreditCard
		}

		cleaned := strings.NewReplacer(" ", "", "-", "").Replace(str)
		if len(cleaned) < 13 || len(cleaned) > 19 {
			return ErrValidationCreditCard
		}

		for _, c := range cleaned {
			if c < '0' || c > '9' {
				return ErrValidationCreditCard
			}
		}

		if !luhn(cleaned) {
			return ErrValidationCreditCard
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
	return RuleFunc(
		func(value any) error {
			str, ok := value.(string)
			if !ok || !strings.HasSuffix(str, suffix) {
				return ErrValidationEndsWith
			}

			return nil
		},
	)
}

// HexColor is a Rule that validates the value is a valid CSS hex color string.
//
// Accepted formats: #RGB and #RRGGBB (case-insensitive).
//
// Fails if:
//   - value is not a string
//   - the string is not a valid 3- or 6-digit hex color
//
// Examples:
//
//	validation.HexColor.Validate("#fff")     // pass — short form
//	validation.HexColor.Validate("#FF5733")  // pass — long form
//	validation.HexColor.Validate("FF5733")   // fail — missing #
//	validation.HexColor.Validate("#GGHHII")  // fail — not hex digits
var HexColor Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok || !regexHexColor.MatchString(str) {
			return ErrValidationHexColor
		}

		return nil
	},
)

// JSON is a Rule that validates the value is a string containing valid JSON.
//
// Any valid JSON value is accepted: object, array, string, number, boolean, or null.
// An empty string fails.
//
// Fails if:
//   - value is not a string
//   - the string is not valid JSON
//
// Examples:
//
//	validation.JSON.Validate(`{"key":"value"}`) // pass
//	validation.JSON.Validate(`[1,2,3]`)         // pass
//	validation.JSON.Validate(`null`)            // pass
//	validation.JSON.Validate(`"hello"`)         // pass
//	validation.JSON.Validate(`{invalid}`)       // fail
//	validation.JSON.Validate(``)                // fail
var JSON Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok || !json.Valid([]byte(str)) {
			return ErrValidationJSON
		}

		return nil
	},
)

// JWT is a Rule that validates the value is a string with a valid JWT format.
//
// A JWT must consist of exactly three dot-separated base64url-encoded segments (header.payload.signature).
// Padding characters are not accepted (standard JWT compact serialization). The content of each segment
// is not decoded or validated.
//
// Fails if:
//   - value is not a string
//   - the string does not match the three-segment JWT format
//
// Examples:
//
//	validation.JWT.Validate("eyJ.eyJ.sig") // pass
//	validation.JWT.Validate("notajwt")     // fail — only one segment
//	validation.JWT.Validate("a.b")         // fail — only two segments
var JWT Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok || !regexJWT.MatchString(str) {
			return ErrValidationJWT
		}

		return nil
	},
)

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

	return RuleFunc(
		func(value any) error {
			if err != nil {
				return RuleSyntaxError{Rule: "NotRegex", Err: err}
			}

			str, ok := value.(string)
			if !ok || re.MatchString(str) {
				return ErrValidationNotRegex
			}

			return nil
		},
	)
}

// PhoneE164 is a Rule that validates the value is a phone number in E.164 format.
//
// E.164 format: a leading +, followed by a country code digit (1–9), followed by 1–14 more digits.
// Total digits (excluding +): 2–15.
//
// Fails if:
//   - value is not a string
//   - the string does not match E.164 format
//
// Examples:
//
//	validation.PhoneE164.Validate("+14155552671")  // pass — US number
//	validation.PhoneE164.Validate("+441234567890") // pass — UK number
//	validation.PhoneE164.Validate("14155552671")   // fail — missing +
//	validation.PhoneE164.Validate("+0123456789")   // fail — country code starts with 0
var PhoneE164 Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok || !regexPhoneE164.MatchString(str) {
			return ErrValidationPhoneE164
		}

		return nil
	},
)

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

	return RuleFunc(
		func(value any) error {
			if err != nil {
				return RuleSyntaxError{Rule: "Regex", Err: err}
			}

			str, ok := value.(string)
			if !ok || !re.MatchString(str) {
				return ErrValidationRegex
			}

			return nil
		},
	)
}

// Semver is a Rule that validates the value is a valid semantic version string (semver.org).
//
// Both prefixed ("v1.2.3") and un-prefixed ("1.2.3") forms are accepted. Supports pre-release
// identifiers (-alpha.1) and build metadata (+001). To enforce the "v" prefix, compose with
// StartsWith("v").
//
// Fails if:
//   - value is not a string
//   - the string is not a valid semver (e.g. leading zeros in numeric identifiers, missing components)
//
// Examples:
//
//	validation.Semver.Validate("1.0.0")               // pass
//	validation.Semver.Validate("v1.0.0")              // pass — v prefix accepted
//	validation.Semver.Validate("2.1.3-alpha.1")       // pass
//	validation.Semver.Validate("1.0.0+build.123")     // pass
//	validation.Semver.Validate("1.0")                 // fail — missing patch
//	validation.Semver.Validate("01.0.0")              // fail — leading zero
var Semver Rule = RuleFunc(func(value any) error {
	str, ok := value.(string)
	if !ok || !regexSemver.MatchString(str) {
		return ErrValidationSemver
	}

	return nil
})

// Slug is a Rule that validates the value is a URL-friendly slug.
//
// A slug consists of lowercase ASCII letters, digits, and hyphens. It must not start or end with
// a hyphen and must not contain consecutive hyphens.
//
// Fails if:
//   - value is not a string
//   - the string contains uppercase letters, non-ASCII characters, spaces, or invalid hyphens
//   - the string is empty
//
// Examples:
//
//	validation.Slug.Validate("hello-world")   // pass
//	validation.Slug.Validate("my-post-123")   // pass
//	validation.Slug.Validate("Hello-World")   // fail — uppercase
//	validation.Slug.Validate("-leading")      // fail — leading hyphen
//	validation.Slug.Validate("double--dash")  // fail — consecutive hyphens
var Slug Rule = RuleFunc(
	func(value any) error {
		str, ok := value.(string)
		if !ok || !regexSlug.MatchString(str) {
			return ErrValidationSlug
		}

		return nil
	},
)

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
	return RuleFunc(
		func(value any) error {
			str, ok := value.(string)
			if !ok || !strings.HasPrefix(str, prefix) {
				return ErrValidationStartsWith
			}

			return nil
		},
	)
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

func luhn(number string) bool {
	sum := 0
	nDigits := len(number)
	parity := nDigits % 2
	for i := 0; i < nDigits; i++ {
		digit := int(number[i] - '0')
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	return sum%10 == 0
}

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

