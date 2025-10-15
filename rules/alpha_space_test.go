package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var alphaSpaceRuleTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "fullName",
			"inputBag": bag.InputBag{
				"fullName": "John Doe",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"notAlphaSpace": map[string]any{
		"input": map[string]any{
			"selector": "fullName",
			"inputBag": bag.InputBag{
				"fullName": "John Doe 3",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field fullName must only contain letters and spaces.",
		},
	},
}

func TestAlphaSpaceRule(t *testing.T) {
	rule := initAlphaSpaceRule()

	for name, d := range alphaSpaceRuleTestData {
		t.Run(name, func(t *testing.T) {
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

func initAlphaSpaceRule() *AlphaSpace {
	alphaRule := &AlphaSpace{}
	alphaRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.alpha_space":
			tr := "The field :field: must only contain letters and spaces."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		default:
			return key
		}
	})
	return alphaRule
}
