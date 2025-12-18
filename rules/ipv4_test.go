package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

func TestIPv4Rule(t *testing.T) {
	rule := initIPv4Rule()

	tests := map[string]any{
		"ok": map[string]any{
			"input":  map[string]any{"selector": "addr", "inputBag": bag.InputBag{"addr": "8.8.8.8"}},
			"output": map[string]any{"validationFailed": false, "validationError": ""},
		},
		"invalidV6": map[string]any{
			"input":  map[string]any{"selector": "addr", "inputBag": bag.InputBag{"addr": "2001:db8::1"}},
			"output": map[string]any{"validationFailed": true, "validationError": "The field addr must be a valid ipv4."},
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

func initIPv4Rule() *IPv4 {
	r := &IPv4{}
	r.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}
		switch key {
		case "validation.ipv4":
			tr := "The field :field: must be a valid ipv4."
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

func TestIPv4_RequiresField(t *testing.T) {
	rule := &IPv4{}
	assert.False(t, rule.RequiresField())
}
