package rules

import (
	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

type Required struct {
	translation.BaseTranslatableRule
}

func (r *Required) Validate(_ string, _ any, _ bag.InputBag, exists bool) bool {
	return exists
}

func (r *Required) Message() string {
	return r.Translate(r.Locale, "validation.required")
}
