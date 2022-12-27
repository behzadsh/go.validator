package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var alphaRuleTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "name",
			"value":    "John",
			"inputBag": bag.InputBag{
				"name": "John",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"notAlpha": map[string]any{
		"input": map[string]any{
			"selector": "name",
			"value":    "John Doe",
			"inputBag": bag.InputBag{
				"name": "John Doe",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field name must only contain letters.",
		},
	},
}

func TestAlphaRule(t *testing.T) {
	afterRule := initAlphaRule()

	for name, d := range alphaRuleTestData {
		t.Run(name, func(t *testing.T) {
			data := d.(map[string]any)
			input := data["input"].(map[string]any)
			output := data["output"].(map[string]any)
			inputBag := input["inputBag"].(bag.InputBag)

			res := afterRule.Validate(input["selector"].(string), input["value"].(string), inputBag, true)

			assert.Equal(t, output["validationFailed"].(bool), res.Failed())
			assert.Equal(t, output["validationError"].(string), res.Message())
		})
	}
}

func initAlphaRule() *Alpha {
	alphaRule := &Alpha{}
	alphaRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.alpha":
			tr := "The field :field: must only contain letters."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		default:
			return key
		}
	})
	return alphaRule
}
