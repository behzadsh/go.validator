package validation

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
