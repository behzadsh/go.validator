package validation

import "time"

// Error is implemented by every validation-failure error returned by the rules in this package.
// Code returns a stable snake_case key suitable for i18n catalog lookups.
// Params returns rule-specific parameters (e.g. {"length": 5} for MinLength(5), or nil for parameter-less rules).
// Custom rules may implement this interface to expose Code and Params on FieldError.
type Error interface {
	error
	Code() string
	Params() map[string]any
}

// ================================================================================================================== //
//                                                     basicError                                                     //
// ================================================================================================================== //

type basicError struct {
	code    string
	message string
}

func (e basicError) Error() string        { return e.message }
func (e basicError) Code() string         { return e.code }
func (basicError) Params() map[string]any { return nil }

// ================================================================================================================== //
//                                                   containsError                                                    //
// ================================================================================================================== //

type containsError struct{ Substring string }

func (containsError) Error() string            { return "contains validation failed" }
func (containsError) Code() string             { return "contains" }
func (e containsError) Params() map[string]any { return map[string]any{"substring": e.Substring} }

// ================================================================================================================== //
//                                                   endsWithError                                                    //
// ================================================================================================================== //

type endsWithError struct{ Suffix string }

func (endsWithError) Error() string            { return "ends with validation failed" }
func (endsWithError) Code() string             { return "ends_with" }
func (e endsWithError) Params() map[string]any { return map[string]any{"suffix": e.Suffix} }

// ================================================================================================================== //
//                                                    lengthError                                                     //
// ================================================================================================================== //

type lengthError struct{ Length int }

func (lengthError) Error() string            { return "length validation failed" }
func (lengthError) Code() string             { return "length" }
func (e lengthError) Params() map[string]any { return map[string]any{"length": e.Length} }

// ================================================================================================================== //
//                                                   maxLengthError                                                   //
// ================================================================================================================== //

type maxLengthError struct{ Length int }

func (maxLengthError) Error() string            { return "max length validation failed" }
func (maxLengthError) Code() string             { return "max_length" }
func (e maxLengthError) Params() map[string]any { return map[string]any{"length": e.Length} }

// ================================================================================================================== //
//                                                   minLengthError                                                   //
// ================================================================================================================== //

type minLengthError struct{ Length int }

func (minLengthError) Error() string            { return "min length validation failed" }
func (minLengthError) Code() string             { return "min_length" }
func (e minLengthError) Params() map[string]any { return map[string]any{"length": e.Length} }

// ================================================================================================================== //
//                                                    notRegexError                                                   //
// ================================================================================================================== //

type notRegexError struct{ Pattern string }

func (notRegexError) Error() string            { return "not regex validation failed" }
func (notRegexError) Code() string             { return "not_regex" }
func (e notRegexError) Params() map[string]any { return map[string]any{"pattern": e.Pattern} }

// ================================================================================================================== //
//                                                     regexError                                                     //
// ================================================================================================================== //

type regexError struct{ Pattern string }

func (regexError) Error() string            { return "regex validation failed" }
func (regexError) Code() string             { return "regex" }
func (e regexError) Params() map[string]any { return map[string]any{"pattern": e.Pattern} }

// ================================================================================================================== //
//                                                  startsWithError                                                   //
// ================================================================================================================== //

type startsWithError struct{ Prefix string }

func (startsWithError) Error() string            { return "starts with validation failed" }
func (startsWithError) Code() string             { return "starts_with" }
func (e startsWithError) Params() map[string]any { return map[string]any{"prefix": e.Prefix} }

// ================================================================================================================== //
//                                                    betweenError                                                    //
// ================================================================================================================== //

type betweenError struct{ Min, Max any }

func (betweenError) Error() string            { return "between validation failed" }
func (betweenError) Code() string             { return "between" }
func (e betweenError) Params() map[string]any { return map[string]any{"min": e.Min, "max": e.Max} }

// ================================================================================================================== //
//                                                      gtError                                                       //
// ================================================================================================================== //

type gtError struct{ Value any }

func (gtError) Error() string            { return "gt validation failed" }
func (gtError) Code() string             { return "gt" }
func (e gtError) Params() map[string]any { return map[string]any{"value": e.Value} }

// ================================================================================================================== //
//                                                      gteError                                                      //
// ================================================================================================================== //

type gteError struct{ Value any }

func (gteError) Error() string            { return "gte validation failed" }
func (gteError) Code() string             { return "gte" }
func (e gteError) Params() map[string]any { return map[string]any{"value": e.Value} }

// ================================================================================================================== //
//                                                      ltError                                                       //
// ================================================================================================================== //

type ltError struct{ Value any }

func (ltError) Error() string            { return "lt validation failed" }
func (ltError) Code() string             { return "lt" }
func (e ltError) Params() map[string]any { return map[string]any{"value": e.Value} }

// ================================================================================================================== //
//                                                      lteError                                                      //
// ================================================================================================================== //

type lteError struct{ Value any }

func (lteError) Error() string            { return "lte validation failed" }
func (lteError) Code() string             { return "lte" }
func (e lteError) Params() map[string]any { return map[string]any{"value": e.Value} }

// ================================================================================================================== //
//                                                      maxError                                                      //
// ================================================================================================================== //

type maxError struct{ Value any }

func (maxError) Error() string            { return "max validation failed" }
func (maxError) Code() string             { return "max" }
func (e maxError) Params() map[string]any { return map[string]any{"value": e.Value} }

// ================================================================================================================== //
//                                                      minError                                                      //
// ================================================================================================================== //

type minError struct{ Value any }

func (minError) Error() string            { return "min validation failed" }
func (minError) Code() string             { return "min" }
func (e minError) Params() map[string]any { return map[string]any{"value": e.Value} }

// ================================================================================================================== //
//                                                  multipleOfError                                                   //
// ================================================================================================================== //

type multipleOfError struct{ Value any }

func (multipleOfError) Error() string            { return "multiple of validation failed" }
func (multipleOfError) Code() string             { return "multiple_of" }
func (e multipleOfError) Params() map[string]any { return map[string]any{"value": e.Value} }

// ================================================================================================================== //
//                                                    digitsError                                                     //
// ================================================================================================================== //

type digitsError struct{ Digits int }

func (digitsError) Error() string            { return "digits validation failed" }
func (digitsError) Code() string             { return "digits" }
func (e digitsError) Params() map[string]any { return map[string]any{"digits": e.Digits} }

// ================================================================================================================== //
//                                                 digitsBetweenError                                                 //
// ================================================================================================================== //

type digitsBetweenError struct{ Min, Max int }

func (digitsBetweenError) Error() string { return "digits between validation failed" }
func (digitsBetweenError) Code() string  { return "digits_between" }
func (e digitsBetweenError) Params() map[string]any {
	return map[string]any{"min": e.Min, "max": e.Max}
}

// ================================================================================================================== //
//                                                   maxDigitsError                                                   //
// ================================================================================================================== //

type maxDigitsError struct{ Digits int }

func (maxDigitsError) Error() string            { return "max digits validation failed" }
func (maxDigitsError) Code() string             { return "max_digits" }
func (e maxDigitsError) Params() map[string]any { return map[string]any{"digits": e.Digits} }

// ================================================================================================================== //
//                                                   minDigitsError                                                   //
// ================================================================================================================== //

type minDigitsError struct{ Digits int }

func (minDigitsError) Error() string            { return "min digits validation failed" }
func (minDigitsError) Code() string             { return "min_digits" }
func (e minDigitsError) Params() map[string]any { return map[string]any{"digits": e.Digits} }

// ================================================================================================================== //
//                                                     afterError                                                     //
// ================================================================================================================== //

type afterError struct{ Time time.Time }

func (afterError) Error() string            { return "after validation failed" }
func (afterError) Code() string             { return "after" }
func (e afterError) Params() map[string]any { return map[string]any{"time": e.Time} }

// ================================================================================================================== //
//                                                  afterFieldError                                                   //
// ================================================================================================================== //

type afterFieldError struct{ Field string }

func (afterFieldError) Error() string            { return "after field validation failed" }
func (afterFieldError) Code() string             { return "after_field" }
func (e afterFieldError) Params() map[string]any { return map[string]any{"field": e.Field} }

// ================================================================================================================== //
//                                                 afterOrEqualError                                                  //
// ================================================================================================================== //

type afterOrEqualError struct{ Time time.Time }

func (afterOrEqualError) Error() string            { return "after or equal validation failed" }
func (afterOrEqualError) Code() string             { return "after_or_equal" }
func (e afterOrEqualError) Params() map[string]any { return map[string]any{"time": e.Time} }

// ================================================================================================================== //
//                                                    beforeError                                                     //
// ================================================================================================================== //

type beforeError struct{ Time time.Time }

func (beforeError) Error() string            { return "before validation failed" }
func (beforeError) Code() string             { return "before" }
func (e beforeError) Params() map[string]any { return map[string]any{"time": e.Time} }

// ================================================================================================================== //
//                                                 beforeFieldError                                                   //
// ================================================================================================================== //

type beforeFieldError struct{ Field string }

func (beforeFieldError) Error() string            { return "before field validation failed" }
func (beforeFieldError) Code() string             { return "before_field" }
func (e beforeFieldError) Params() map[string]any { return map[string]any{"field": e.Field} }

// ================================================================================================================== //
//                                                beforeOrEqualError                                                  //
// ================================================================================================================== //

type beforeOrEqualError struct{ Time time.Time }

func (beforeOrEqualError) Error() string            { return "before or equal validation failed" }
func (beforeOrEqualError) Code() string             { return "before_or_equal" }
func (e beforeOrEqualError) Params() map[string]any { return map[string]any{"time": e.Time} }

// ================================================================================================================== //
//                                               dateTimeBetweenError                                                 //
// ================================================================================================================== //

type dateTimeBetweenError struct{ Min, Max time.Time }

func (dateTimeBetweenError) Error() string { return "datetime between validation failed" }
func (dateTimeBetweenError) Code() string  { return "date_time_between" }
func (e dateTimeBetweenError) Params() map[string]any {
	return map[string]any{"min": e.Min, "max": e.Max}
}

// ================================================================================================================== //
//                                               dateTimeFormatError                                                  //
// ================================================================================================================== //

type dateTimeFormatError struct{ Format string }

func (dateTimeFormatError) Error() string            { return "datetime format validation failed" }
func (dateTimeFormatError) Code() string             { return "date_time_format" }
func (e dateTimeFormatError) Params() map[string]any { return map[string]any{"format": e.Format} }

// ================================================================================================================== //
//                                                    maxSizeError                                                    //
// ================================================================================================================== //

type maxSizeError struct{ Size int }

func (maxSizeError) Error() string            { return "max size validation failed" }
func (maxSizeError) Code() string             { return "max_size" }
func (e maxSizeError) Params() map[string]any { return map[string]any{"size": e.Size} }

// ================================================================================================================== //
//                                                    minSizeError                                                    //
// ================================================================================================================== //

type minSizeError struct{ Size int }

func (minSizeError) Error() string            { return "min size validation failed" }
func (minSizeError) Code() string             { return "min_size" }
func (e minSizeError) Params() map[string]any { return map[string]any{"size": e.Size} }

// ================================================================================================================== //
//                                                     sizeError                                                      //
// ================================================================================================================== //

type sizeError struct{ Size int }

func (sizeError) Error() string            { return "size validation failed" }
func (sizeError) Code() string             { return "size" }
func (e sizeError) Params() map[string]any { return map[string]any{"size": e.Size} }

// ================================================================================================================== //
//                                                      inError                                                       //
// ================================================================================================================== //

type inError struct{ Values any }

func (inError) Error() string            { return "in validation failed" }
func (inError) Code() string             { return "in" }
func (e inError) Params() map[string]any { return map[string]any{"values": e.Values} }

// ================================================================================================================== //
//                                                     notInError                                                     //
// ================================================================================================================== //

type notInError struct{ Values any }

func (notInError) Error() string            { return "not in validation failed" }
func (notInError) Code() string             { return "not_in" }
func (e notInError) Params() map[string]any { return map[string]any{"values": e.Values} }

// ================================================================================================================== //
//                                                      neqError                                                      //
// ================================================================================================================== //

type neqError struct{ Value any }

func (neqError) Error() string            { return "neq validation failed" }
func (neqError) Code() string             { return "neq" }
func (e neqError) Params() map[string]any { return map[string]any{"value": e.Value} }

// ================================================================================================================== //
//                                                    sameAsError                                                     //
// ================================================================================================================== //

type sameAsError struct{ Field string }

func (sameAsError) Error() string            { return "same as validation failed" }
func (sameAsError) Code() string             { return "same_as" }
func (e sameAsError) Params() map[string]any { return map[string]any{"field": e.Field} }

// ================================================================================================================== //
//                                                   differentError                                                   //
// ================================================================================================================== //

type differentError struct{ Field string }

func (differentError) Error() string            { return "different validation failed" }
func (differentError) Code() string             { return "different" }
func (e differentError) Params() map[string]any { return map[string]any{"field": e.Field} }

// FieldError describes a single validation failure for a single field.
//
// Err holds the underlying error returned by the rule; Message is its pre-rendered string form, kept on the struct so
// that callers do not need to invoke Err.Error() in hot paths or templates.
// Code and Params are populated when the rule error implements ValidationError; they are empty/nil for custom rules
// that do not implement that interface.
type FieldError struct {
	Path    string
	Err     error
	Message string
	Code    string
	Params  map[string]any
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
