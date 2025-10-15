package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

func TestTimezoneRule(t *testing.T) {
	rule := initTimezoneRule()

	tests := map[string]any{
		"ok": map[string]any{
			"input":  map[string]any{"selector": "tz", "inputBag": bag.InputBag{"tz": "America/New_York"}},
			"output": map[string]any{"validationFailed": false, "validationError": ""},
		},
		"invalid": map[string]any{
			"input":  map[string]any{"selector": "tz", "inputBag": bag.InputBag{"tz": "Not/AZone"}},
			"output": map[string]any{"validationFailed": true, "validationError": "The field tz must be a valid timezone."},
		},
		"empty": map[string]any{
			"input":  map[string]any{"selector": "tz", "inputBag": bag.InputBag{"tz": ""}},
			"output": map[string]any{"validationFailed": true, "validationError": "The field tz must be a valid timezone."},
		},
	}

	for name, d := range tests {
		t.Run(name, func(t *testing.T) {
			data := d.(map[string]any)
			input := data["input"].(map[string]any)
			output := data["output"].(map[string]any)
			inputBag := input["inputBag"].(bag.InputBag)
			value, _ := inputBag.Get(input["selector"].(string))
			res := rule.Validate(input["selector"].(string), value, inputBag)
			assert.Equal(t, output["validationFailed"].(bool), res.Failed())
			assert.Equal(t, output["validationError"].(string), res.Message())
		})
	}
}

func initTimezoneRule() *Timezone {
	r := &Timezone{}
	r.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}
		switch key {
		case "validation.timezone":
			tr := "The field :field: must be a valid timezone."
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
