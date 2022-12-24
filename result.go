package validation

import "github.com/behzadsh/go.validator/bag"

type Result struct {
	Errors bag.ErrorBag
}

func NewResult() Result {
	return Result{
		Errors: make(bag.ErrorBag),
	}
}

func (r Result) Fails() bool {
	return !r.Errors.IsEmpty()
}

func (r Result) AddError(selector string, errorMsg ...string) {
	r.Errors.Add(selector, errorMsg...)
}
