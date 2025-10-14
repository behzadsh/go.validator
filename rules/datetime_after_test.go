package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

func TestDateTimeAfterRule(t *testing.T) {
	rule := initDateTimeAfterRule()

	tests := map[string]any{
		"ok": map[string]any{
			"input":  map[string]any{"selector": "start", "inputBag": bag.InputBag{"start": "2022-02-01"}, "params": []string{"2022-01-01", "America/New_York"}},
			"output": map[string]any{"validationFailed": false, "validationError": ""},
		},
		"invalidField": map[string]any{
			"input":  map[string]any{"selector": "start", "inputBag": bag.InputBag{"start": "invalid"}, "params": []string{"2022-01-01"}},
			"output": map[string]any{"validationFailed": true, "validationError": "The field start must be a valid date time string."},
		},
		"notAfter": map[string]any{
			"input":  map[string]any{"selector": "start", "inputBag": bag.InputBag{"start": "2021-12-31"}, "params": []string{"2022-01-01"}},
			"output": map[string]any{"validationFailed": true, "validationError": "The field start must be after 2022-01-01T00:00:00Z."},
		},
	}

	for name, d := range tests {
		t.Run(name, func(t *testing.T) {
			data := d.(map[string]any)
			input := data["input"].(map[string]any)
			output := data["output"].(map[string]any)
			inputBag := input["inputBag"].(bag.InputBag)
			value, _ := inputBag.Get(input["selector"].(string))
			rule.AddParams(input["params"].([]string))
			res := rule.Validate(input["selector"].(string), value, inputBag)
			assert.Equal(t, output["validationFailed"].(bool), res.Failed())
			assert.Equal(t, output["validationError"].(string), res.Message())
		})
	}
}

func initDateTimeAfterRule() *DateTimeAfter {
	r := &DateTimeAfter{}
	r.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}
		switch key {
		case "validation.datetime":
			tr := "The field :field: must be a valid date time string."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v, )
			}
			return tr
		case "validation.datetime_after":
			tr := "The field :field: must be after :value:."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v, )
			}
			return tr
		case "validation.after":
			tr := "The field :field: must be after field :otherField:."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v, )
			}
			return tr
		default:
			return key
		}
	})
	return r
}
