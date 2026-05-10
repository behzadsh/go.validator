// Package validation provides a small, schema-based validator for map and struct input.
//
// Build a schema by chaining Field calls, then call Validate. Each rule is an independent value; the only rule that
// fails on an absent field is Required.
//
//	schema := validation.New().
//		Field("name", validation.Required, validation.MinLength(2)).
//		Field("email", validation.Required, validation.Email)
//
//	if res := schema.Validate(input); res.HasErrors() {
//		for _, e := range res.Errors() {
//			fmt.Println(e.Path, e.Message)
//		}
//	}
//
// The input may be a map[string]any, a struct, or a pointer to a struct.
// Field names in the path follow this resolution order: the first comma-segment of the `json` tag (when present and
// not "-"), then the exported field name.
package validation

// Rule validates a single value and returns nil on success or an error describing the failure.
type Rule interface {
	Validate(value any) error
}

// RuleFunc adapts a plain function to the Rule interface.
type RuleFunc func(value any) error

// Validate satisfies the Rule interface.
func (f RuleFunc) Validate(value any) error { return f(value) }

// InputRule is a Rule that also receives the full validated input so it can
// reference other fields. Use this for cross-field rules such as SameAs.
type InputRule interface {
	Rule
	ValidateWithInput(value any, input *InputBag) error
}

// InputRuleFunc adapts a plain function to the InputRule interface.
type InputRuleFunc func(value any, input *InputBag) error

// Validate satisfies the Rule interface by calling the function with a nil input.
// Schema.Validate always calls ValidateWithInput instead.
func (f InputRuleFunc) Validate(value any) error { return f(value, nil) }

// ValidateWithInput satisfies the InputRule interface.
func (f InputRuleFunc) ValidateWithInput(value any, input *InputBag) error { return f(value, input) }

// Schema describes a set of fields and the rules that apply to each.
//
// A Schema is intended to be fully built up before Validate is called. Once built, Validate is safe to call from
// multiple goroutines concurrently.
type Schema struct {
	fields []fieldRules
}

type fieldRules struct {
	path  string
	rules []Rule
}

// New returns an empty Schema ready to be populated via Field.
func New() *Schema {
	return &Schema{}
}

// Field appends a list of rules for the given dot-notation path.
//
// Field returns the receiver to support chaining.
func (s *Schema) Field(path string, rules ...Rule) *Schema {
	s.fields = append(s.fields, fieldRules{path: path, rules: rules})

	return s
}

// Validate runs every rule against its corresponding field in the input and returns the collected errors.
//
// The input may be a map[string]any, a struct, a pointer to a struct, or any nested combination thereof. The returned
// slice is empty (length zero) when validation succeeds. All rules for a field are executed; validation does not stop
// at the first failure.
func (s *Schema) Validate(input any) (*Result, error) {
	var errs []FieldError
	inputBag := NewInputBag(input)

	for _, f := range s.fields {
		value, _ := inputBag.Lookup(f.path)
		for _, r := range f.rules {
			var err error
			if ir, ok := r.(InputRule); ok {
				err = ir.ValidateWithInput(value, inputBag)
			} else {
				err = r.Validate(value)
			}
			if err != nil {
				errs = append(errs, FieldError{Path: f.path, Err: err, Message: err.Error()})
			}
		}
	}

	return &Result{errors: errs}, nil
}
