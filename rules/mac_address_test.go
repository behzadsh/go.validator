package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

func TestMacAddressRule(t *testing.T) {
	rule := initMacRule()

	tests := map[string]any{
		"okColon": map[string]any{
			"input":  map[string]any{"selector": "mac", "inputBag": bag.InputBag{"mac": "01:23:45:67:89:ab"}},
			"output": map[string]any{"validationFailed": false, "validationError": ""},
		},
		"okDash": map[string]any{
			"input":  map[string]any{"selector": "mac", "inputBag": bag.InputBag{"mac": "01-23-45-67-89-AB"}},
			"output": map[string]any{"validationFailed": false, "validationError": ""},
		},
		"invalid": map[string]any{
			"input":  map[string]any{"selector": "mac", "inputBag": bag.InputBag{"mac": "0123.4567.89ab"}},
			"output": map[string]any{"validationFailed": true, "validationError": "The field mac must be a valid mac address."},
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

func initMacRule() *MacAddress {
	r := &MacAddress{}
	r.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}
		switch key {
		case "validation.mac_address":
			tr := "The field :field: must be a valid mac address."
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
