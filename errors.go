package validation

import "errors"

// General validation errors.
var (
	ErrValidationRequired           = errors.New("required validation failed")
	ErrValidationRequiredIf         = errors.New("required if validation failed")
	ErrValidationNotEmpty           = errors.New("not empty validation failed")
	ErrValidationRequiredUnless     = errors.New("required unless validation failed")
	ErrValidationRequiredWith       = errors.New("required with validation failed")
	ErrValidationRequiredWithAll    = errors.New("required with all validation failed")
	ErrValidationRequiredWithout    = errors.New("required without validation failed")
	ErrValidationRequiredWithoutAll = errors.New("required without all validation failed")
)

// Strings validation errors.
var (
	ErrValidationAlpha      = errors.New("alpha validation failed")
	ErrValidationAlphaDash  = errors.New("alpha dash validation failed")
	ErrValidationAlphaNum   = errors.New("alpha num validation failed")
	ErrValidationAlphaSpace = errors.New("alpha space validation failed")
	ErrValidationASCII      = errors.New("ascii validation failed")
	ErrValidationBase64     = errors.New("base64 validation failed")
	ErrValidationContains   = errors.New("contains validation failed")
	ErrValidationCreditCard = errors.New("credit card validation failed")
	ErrValidationEmail      = errors.New("email validation failed")
	ErrValidationEmailMX    = errors.New("email mx validation failed")
	ErrValidationEndsWith   = errors.New("ends with validation failed")
	ErrValidationHexColor   = errors.New("hex color validation failed")
	ErrValidationJSON       = errors.New("json validation failed")
	ErrValidationJWT        = errors.New("jwt validation failed")
	ErrValidationLength     = errors.New("length validation failed")
	ErrValidationLowercase  = errors.New("lowercase validation failed")
	ErrValidationMaxLength  = errors.New("max length validation failed")
	ErrValidationMinLength  = errors.New("min length validation failed")
	ErrValidationNotRegex   = errors.New("not regex validation failed")
	ErrValidationPhoneE164  = errors.New("phone e164 validation failed")
	ErrValidationRegex      = errors.New("regex validation failed")
	ErrValidationSemver     = errors.New("semver validation failed")
	ErrValidationSlug       = errors.New("slug validation failed")
	ErrValidationStartsWith = errors.New("starts with validation failed")
	ErrValidationUppercase  = errors.New("uppercase validation failed")
	ErrValidationUUID       = errors.New("uuid validation failed")
)

// Numbers validation errors.
var (
	ErrValidationBetween     = errors.New("between validation failed")
	ErrValidationGT          = errors.New("gt validation failed")
	ErrValidationGTE         = errors.New("gte validation failed")
	ErrValidationInteger     = errors.New("integer validation failed")
	ErrValidationLatitude    = errors.New("latitude validation failed")
	ErrValidationLongitude   = errors.New("longitude validation failed")
	ErrValidationLT          = errors.New("lt validation failed")
	ErrValidationLTE         = errors.New("lte validation failed")
	ErrValidationMax         = errors.New("max validation failed")
	ErrValidationMin         = errors.New("min validation failed")
	ErrValidationMultipleOf  = errors.New("multiple of validation failed")
	ErrValidationNegative    = errors.New("negative validation failed")
	ErrValidationNonNegative = errors.New("non negative validation failed")
	ErrValidationNumeric     = errors.New("numeric validation failed")
	ErrValidationPort        = errors.New("port validation failed")
	ErrValidationPositive    = errors.New("positive validation failed")
)

// Digit validation errors.
var (
	ErrValidationDigits        = errors.New("digits validation failed")
	ErrValidationDigitsBetween = errors.New("digits between validation failed")
	ErrValidationMaxDigits     = errors.New("max digits validation failed")
	ErrValidationMinDigits     = errors.New("min digits validation failed")
)

// Date/Time validation errors.
var (
	ErrValidationAfter           = errors.New("after validation failed")
	ErrValidationAfterField      = errors.New("after field validation failed")
	ErrValidationAfterOrEqual    = errors.New("after or equal validation failed")
	ErrValidationBefore          = errors.New("before validation failed")
	ErrValidationBeforeField     = errors.New("before field validation failed")
	ErrValidationBeforeOrEqual   = errors.New("before or equal validation failed")
	ErrValidationDateTime        = errors.New("datetime validation failed")
	ErrValidationDateTimeBetween = errors.New("datetime between validation failed")
	ErrValidationDateTimeFormat  = errors.New("datetime format validation failed")
	ErrValidationTimezone        = errors.New("timezone validation failed")
)

// Network validation errors.
var (
	ErrValidationCIDR       = errors.New("cidr validation failed")
	ErrValidationIP         = errors.New("ip validation failed")
	ErrValidationIPv4       = errors.New("ipv4 validation failed")
	ErrValidationIPv6       = errors.New("ipv6 validation failed")
	ErrValidationMACAddress = errors.New("mac address validation failed")
	ErrValidationURL        = errors.New("url validation failed")
)

// Collection validation errors.
var (
	ErrValidationDistinct = errors.New("distinct validation failed")
	ErrValidationEach     = errors.New("each validation failed")
	ErrValidationMaxSize  = errors.New("max size validation failed")
	ErrValidationMinSize  = errors.New("min size validation failed")
	ErrValidationSize     = errors.New("size validation failed")
)

// Logical validation errors.
var (
	ErrValidationAny = errors.New("any validation failed")
	ErrValidationNot = errors.New("not validation failed")
)

// Generic validation errors.
var (
	ErrValidationIn    = errors.New("in validation failed")
	ErrValidationNotIn = errors.New("not in validation failed")
	ErrValidationNEQ   = errors.New("neq validation failed")
)

// Comparison validation errors.
var (
	ErrValidationSameAs    = errors.New("same as validation failed")
	ErrValidationDifferent = errors.New("different validation failed")
)

// FieldError describes a single validation failure for a single field.
//
// Err holds the underlying error returned by the rule; Message is its pre-rendered string form, kept on the struct so
// that callers do not need to invoke Err.Error() in hot paths or templates.
type FieldError struct {
	Path    string
	Err     error
	Message string
}

// Error implements the error interface.
func (e FieldError) Error() string {
	return e.Path + ": " + e.Message
}

// Unwrap returns the underlying rule error so that errors.Is / errors.As work.
func (e FieldError) Unwrap() error {
	return e.Err
}

// RuleSyntaxError is returned by Schema.Validate when a rule is misconfigured; for example, an invalid regex pattern
// supplied to a rule constructor. It signals a programming error rather than a validation failure; callers should treat
// it as fatal and fix the schema at startup.
type RuleSyntaxError struct {
	Rule string
	Err  error
}

// Error implements the error interface.
func (e RuleSyntaxError) Error() string {
	if e.Rule != "" {
		return "rule " + e.Rule + ": " + e.Err.Error()
	}
	return "rule syntax error: " + e.Err.Error()
}

// Unwrap returns the underlying cause.
func (e RuleSyntaxError) Unwrap() error { return e.Err }

// Result is the value returned by Schema.Validate. It carries the collection of validation failures.
// Result deliberately does not implement the error interface; call Errors to obtain the slice or HasErrors to test it.
type Result struct {
	errors []FieldError
}

// Errors returns the collected field errors. The returned slice is nil when validation succeeded.
func (r *Result) Errors() []FieldError {
	return r.errors
}

// HasErrors reports whether validation produced any errors.
func (r *Result) HasErrors() bool {
	return len(r.errors) > 0
}

// For returns every FieldError whose path equals the given path.
func (r *Result) For(path string) []FieldError {
	var out []FieldError
	for _, fe := range r.errors {
		if fe.Path == path {
			out = append(out, fe)
		}
	}
	return out
}
