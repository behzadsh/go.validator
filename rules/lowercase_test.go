package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var lowercaseRuleTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "name",
			"inputBag": bag.InputBag{
				"name": "john",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"notLowercase": map[string]any{
		"input": map[string]any{
			"selector": "name",
			"inputBag": bag.InputBag{
				"name": "John Doe",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field name must be lowercase.",
		},
	},
}

func TestLowercaseRule(t *testing.T) {
	rule := initLowercaseRule()

	for name, d := range lowercaseRuleTestData {
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

func initLowercaseRule() *Lowercase {
	lowercaseRule := &Lowercase{}
	lowercaseRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.lowercase":
			tr := "The field :field: must be lowercase."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		default:
			return key
		}
	})
	return lowercaseRule
}
