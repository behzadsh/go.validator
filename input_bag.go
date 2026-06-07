package validation

import (
	"reflect"
	"strings"
)

// InputBag wraps the raw input passed to Schema.Validate and provides path-based field access. Paths use dot notation,
// e.g. "user.profile.email", and can traverse nested maps, structs, and pointers in any combination.
//
// InputBag is constructed once per Validate call and then passed read-only to every InputRule. Callers outside the
// validation package typically obtain one from the input parameter of InputRuleFunc — there is no need to construct
// an InputBag directly.
type InputBag struct {
	input any
}

// NewInputBag wraps input in an InputBag. The input may be a map[string]any, a struct, a pointer to a struct,
// or any nested combination thereof.
func NewInputBag(input any) *InputBag {
	return &InputBag{input: input}
}

// Lookup resolves a dot-notation path against the wrapped input and returns the value at that path together with
// a boolean indicating whether the path was found.
//
// Returning (nil, false) means the path does not exist in the input.
// Returning (nil, true) means the path exists but its value is nil.
//
// Supported segment transitions at each dot:
//   - map[string]any  → key lookup
//   - any other map with string keys → reflect-based key lookup
//   - struct or *struct → field resolved by json tag, then by Go field name
//   - pointer / interface → automatically dereferenced
//
// Returning false when any segment is missing or a non-traversable value, e.g. a scalar, is encountered before the path
// is fully consumed.
func (b *InputBag) Lookup(path string) (any, bool) {
	if path == "" {
		return nil, false
	}

	current := b.input
	for _, segment := range strings.Split(path, ".") {
		var ok bool
		current, ok = step(current, segment)
		if !ok {
			return nil, false
		}
	}

	return current, true
}

// step advances one segment of a dot-notation path against the current value. It handles map[string]any directly for
// performance, then falls back to reflection for other map types and structs.
func step(current any, segment string) (any, bool) {
	if current == nil {
		return nil, false
	}

	if m, ok := current.(map[string]any); ok {
		v, exists := m[segment]
		return v, exists
	}

	rv := reflect.ValueOf(current)
	for rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return nil, false
		}
		rv = rv.Elem()
	}

	switch rv.Kind() {
	case reflect.Map:
		if rv.Type().Key().Kind() != reflect.String {
			return nil, false
		}
		v := rv.MapIndex(reflect.ValueOf(segment))
		if !v.IsValid() {
			return nil, false
		}
		return v.Interface(), true
	case reflect.Struct:
		return structField(rv, segment)
	default:
		return nil, false
	}
}

// structField resolves a single path segment against a struct value.
//
// Resolution order:
//  1. The first comma-segment of the `json` tag (skipped when the tag is "-").
//  2. The exported Go field name.
//
// Anonymous (embedded) struct fields are searched recursively so that promoted fields are visible to the same rules
// as top-level ones.
func structField(rv reflect.Value, name string) (any, bool) {
	if rv.Kind() != reflect.Struct {
		return nil, false
	}

	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		if field.Anonymous {
			fv := rv.Field(i)
			if fv.Kind() == reflect.Pointer {
				if fv.IsNil() {
					continue
				}
				fv = fv.Elem()
			}
			if fv.Kind() == reflect.Struct {
				if v, ok := structField(fv, name); ok {
					return v, true
				}
			}

			continue
		}
		if !field.IsExported() {
			continue
		}
		if tag, ok := field.Tag.Lookup("json"); ok {
			tagName := strings.Split(tag, ",")[0]
			if tagName == "-" {
				continue
			}
			if tagName != "" && tagName == name {
				return derefField(rv.Field(i))
			}
		}
		if field.Name == name {
			return derefField(rv.Field(i))
		}
	}

	return nil, false
}

func derefField(fv reflect.Value) (any, bool) {
	for fv.Kind() == reflect.Pointer {
		if fv.IsNil() {
			return nil, false
		}
		fv = fv.Elem()
	}
	return fv.Interface(), true
}
