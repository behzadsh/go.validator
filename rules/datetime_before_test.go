package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

func TestDateTimeBeforeRule(t *testing.T) {
	rule := initDateTimeBeforeRule()

	tests := map[string]any{
		"ok": map[string]any{
			"input":  map[string]any{"selector": "end", "inputBag": bag.InputBag{"end": "2021-12-31"}, "params": []string{"2022-01-01", "America/New_York"}},
			"output": map[string]any{"validationFailed": false, "validationError": ""},
		},
		"invalidField": map[string]any{
			"input":  map[string]any{"selector": "end", "inputBag": bag.InputBag{"end": "invalid"}, "params": []string{"2022-01-01"}},
			"output": map[string]any{"validationFailed": true, "validationError": "The field end must be a valid date time string."},
		},
		"notBefore": map[string]any{
			"input":  map[string]any{"selector": "end", "inputBag": bag.InputBag{"end": "2022-02-01"}, "params": []string{"2022-01-01"}},
			"output": map[string]any{"validationFailed": true, "validationError": "The field end must be before 2022-01-01T00:00:00Z."},
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

func initDateTimeBeforeRule() *DateTimeBefore {
	r := &DateTimeBefore{}
	r.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}
		switch key {
		case "validation.datetime":
			tr := "The field :field: must be a valid date time string."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}
			return tr
		case "validation.datetime_before":
			tr := "The field :field: must be before :value:."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}
			return tr
		default:
			return key
		}
	})
	return r
}
