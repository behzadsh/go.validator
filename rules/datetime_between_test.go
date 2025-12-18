package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

func TestDateTimeBetweenRule(t *testing.T) {
	rule := initDateTimeBetweenRule()

	tests := map[string]any{
		"ok": map[string]any{
			"input":  map[string]any{"selector": "d", "inputBag": bag.InputBag{"d": "2022-01-15"}, "params": []string{"2022-01-01", "2022-02-01", "UTC"}},
			"output": map[string]any{"validationFailed": false, "validationError": ""},
		},
		"invalidField": map[string]any{
			"input":  map[string]any{"selector": "d", "inputBag": bag.InputBag{"d": "invalid"}, "params": []string{"2022-01-01", "2022-02-01"}},
			"output": map[string]any{"validationFailed": true, "validationError": "The field d must be a valid date time string."},
		},
		"outOfRange": map[string]any{
			"input":  map[string]any{"selector": "d", "inputBag": bag.InputBag{"d": "2021-12-31"}, "params": []string{"2022-01-01", "2022-02-01"}},
			"output": map[string]any{"validationFailed": true, "validationError": "The field d must be between 2022-01-01T00:00:00Z and 2022-02-01T00:00:00Z."},
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

func initDateTimeBetweenRule() *DateTimeBetween {
	r := &DateTimeBetween{}
	r.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}
		switch key {
		case "validation.datetime":
			tr := "The field :field: must be a valid date time string."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}
			return tr
		case "validation.datetime_between":
			tr := "The field :field: must be between :min: and :max:."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}
			return tr
		default:
			return key
		}
	})
	return r
}

func TestDateTimeBetween_RequiresField(t *testing.T) {
	rule := &DateTimeBetween{}
	assert.False(t, rule.RequiresField())
}
