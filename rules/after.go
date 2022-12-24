package rules

import (
	"github.com/spf13/cast"

	"github.com/behzadsh/go.validator/bag"
	"github.com/behzadsh/go.validator/translation"
)

type After struct {
	translation.BaseTranslatableRule
	otherField string
	message    string
}

func (r *After) Validate(selector string, value any, inputBag bag.InputBag, _ bool) bool {
	r.message = r.Translate(r.Locale, "validation.after", map[string]string{
		"field":      selector,
		"otherField": r.otherField,
	})

	timeValue, err := cast.ToTimeE(value)
	if err != nil {
		r.message = r.Translate(r.Locale, "validation.after_incorrect_date_format", map[string]string{
			"field": selector,
		})
	}

	otherValue, ok := inputBag.Get(r.otherField)
	if !ok {
		r.message = r.Translate(r.Locale, "Validation.after_other_field_not_provided", map[string]string{
			"otherField": r.otherField,
		})
	}

	otherTimeValue, err := cast.ToTimeE(otherValue)
	if err != nil {
		r.message = r.Translate(r.Locale, "validation.after_other_field_incorrect_date_format", map[string]string{
			"otherField": r.otherField,
		})
	}

	return timeValue.After(otherTimeValue)
}

func (r *After) Message() string {
	return r.message
}

func (r *After) AddParams(params []string) {
	r.otherField = params[0]
}

func (r *After) MinRequiredParams() int {
	return 1
}
