package validation

import "github.com/behzadsh/go.validator/bag"

// Result is a struct consist of validation errors.
type Result struct {
	Errors bag.ErrorBag
}

// NewResult generate a new result object.
func NewResult() Result {
	return Result{
		Errors: make(bag.ErrorBag),
	}
}

// Failed returns true if there is an error.
func (r Result) Failed() bool {
	return !r.Errors.IsEmpty()
}

func (r Result) addError(selector string, errorMsg ...string) {
	r.Errors.Add(selector, errorMsg...)
}
