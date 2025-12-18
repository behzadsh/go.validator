package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var ipRuleTestData = map[string]any{
	"okIPv4": map[string]any{
		"input": map[string]any{
			"selector": "addr",
			"inputBag": bag.InputBag{
				"addr": "192.168.1.1",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"okIPv6": map[string]any{
		"input": map[string]any{
			"selector": "addr",
			"inputBag": bag.InputBag{
				"addr": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"invalid": map[string]any{
		"input": map[string]any{
			"selector": "addr",
			"inputBag": bag.InputBag{
				"addr": "999.999.999.999",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field addr must be a valid ip.",
		},
	},
	"empty": map[string]any{
		"input": map[string]any{
			"selector": "addr",
			"inputBag": bag.InputBag{
				"addr": "",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field addr must be a valid ip.",
		},
	},
}

func TestIPRule(t *testing.T) {
	for name, d := range ipRuleTestData {
		t.Run(name, func(t *testing.T) {
			rule := initIPRule()
			data, _ := d.(map[string]any)
			input, _ := data["input"].(map[string]any)
			output, _ := data["output"].(map[string]any)
			inputBag, _ := input["inputBag"].(bag.InputBag)

			value, _ := inputBag.Get(input["selector"].(string))
			res := rule.Validate(input["selector"].(string), value, inputBag)

			assert.Equal(t, output["validationFailed"].(bool), res.Failed())
			assert.Equal(t, output["validationError"].(string), res.Message())
		})
	}
}

func initIPRule() *IP {
	r := &IP{}
	r.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.ip":
			tr := "The field :field: must be a valid ip."
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

func TestIP_RequiresField(t *testing.T) {
	rule := &IP{}
	assert.False(t, rule.RequiresField())
}
