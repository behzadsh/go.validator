package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var requiredRuleTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "name",
			"inputBag": bag.InputBag{
				"name": "John",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"failed": map[string]any{
		"input": map[string]any{
			"selector": "age",
			"inputBag": bag.InputBag{
				"name": "John Doe",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field age is required.",
		},
	},
}

func TestRequiredRule(t *testing.T) {
	rule := initRequiredRule()

	for name, d := range requiredRuleTestData {
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

func initRequiredRule() *Required {
	requiredRule := &Required{}
	requiredRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.required":
			tr := "The field :field: is required."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		default:
			return key
		}
	})
	return requiredRule
}
