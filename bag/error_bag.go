package bag

// ErrorBag is a custom type that holds a map of field selectors and their related validation errors.
type ErrorBag map[string][]string

// IsEmpty checks whether the error bag map is empty.
func (b ErrorBag) IsEmpty() bool {
	return len(b) == 0
}

// All returns all the errors in the error bag map.
func (b ErrorBag) All() map[string][]string {
	return b
}

// Add adds an error message (or many error messages) for the given selector.
func (b ErrorBag) Add(selector string, msg ...string) {
	b[selector] = append(b[selector], msg...)
}

// FirstOf returns the first error message of the given selector (if it exists).
func (b ErrorBag) FirstOf(selector string) string {
	msg, ok := b[selector]
	if !ok {
		return ""
	}

	return msg[0]
}

// Has checks if there is an error for the given selector in the error bag.
func (b ErrorBag) Has(selector string) bool {
	v, ok := b[selector]

	return ok && v != nil && len(v) > 0
}
