package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var differentRuleTestData = map[string]any{
	"success": map[string]any{
		"input": map[string]any{
			"selector": "newPassword",
			"inputBag": bag.InputBag{
				"oldPassword": "mySecurePassword",
				"newPassword": "anotherSecurePassword",
			},
			"params": []string{
				"oldPassword",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"failed": map[string]any{
		"input": map[string]any{
			"selector": "newPassword",
			"inputBag": bag.InputBag{
				"oldPassword": "mySecurePassword",
				"newPassword": "mySecurePassword",
			},
			"params": []string{
				"oldPassword",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field newPassword must be different from the field oldPassword.",
		},
	},
}

func TestDifferentRule(t *testing.T) {
	rule := initDifferentRule()

	for name, d := range differentRuleTestData {
		t.Run(name, func(t *testing.T) {
			data, _ := d.(map[string]any)
			input, _ := data["input"].(map[string]any)
			output, _ := data["output"].(map[string]any)
			inputBag, _ := input["inputBag"].(bag.InputBag)

			value, _ := inputBag.Get(input["selector"].(string))

			rule.AddParams(input["params"].([]string))
			res := rule.Validate(input["selector"].(string), value, inputBag)

			assert.Equal(t, output["validationFailed"].(bool), res.Failed())
			assert.Equal(t, output["validationError"].(string), res.Message())
		})
	}
}

func initDifferentRule() *Different {
	differentRule := &Different{}
	differentRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.different":
			tr := "The field :field: must be different from the field :otherField:."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		default:
			return key
		}
	})
	return differentRule
}

func TestDifferent_MinRequiredParams(t *testing.T) {
	rule := initDifferentRule()

	assert.Equal(t, 1, rule.MinRequiredParams())
}
