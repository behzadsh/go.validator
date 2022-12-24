package rules

import (
	"github.com/thoas/go-funk"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

type NotEmpty struct {
	translation.BaseTranslatableRule
}

func (r *NotEmpty) Validate(_ string, value any, _ bag.InputBag, exists bool) bool {
	return exists && funk.NotEmpty(value)
}

func (r *NotEmpty) Message() string {
	return r.Translate(r.Locale, "validation.not_empty")
}
