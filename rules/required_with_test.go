package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var requiredWithRuleTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "email",
			"inputBag": bag.InputBag{
				"email":    "user@example.com",
				"type":     "user",
				"username": "goodUser",
			},
			"params": []string{
				"type", "username",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"successfulNotRequired": map[string]any{
		"input": map[string]any{
			"selector": "email",
			"inputBag": bag.InputBag{},
			"params": []string{
				"type", "username",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"oneOfTheOtherFieldsExists": map[string]any{
		"input": map[string]any{
			"selector": "email",
			"inputBag": bag.InputBag{
				"username": "goodUser",
			},
			"params": []string{
				"type", "username",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field email is required when username is present.",
		},
	},
	"allOfTheOtherFieldsExists": map[string]any{
		"input": map[string]any{
			"selector": "email",
			"inputBag": bag.InputBag{
				"type":     "user",
				"username": "goodUser",
			},
			"params": []string{
				"type", "username",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field email is required when type is present.",
		},
	},
}

func TestRequiredWithRule(t *testing.T) {
	rule := initRequiredWithRule()

	for name, d := range requiredWithRuleTestData {
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

func initRequiredWithRule() *RequiredWith {
	requiredWithRule := &RequiredWith{}
	requiredWithRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.required_with":
			tr := "The field :field: is required when :otherField: is present."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		default:
			return key
		}
	})
	return requiredWithRule
}

func TestRequiredWith_MinRequiredParams(t *testing.T) {
	rule := initRequiredWithRule()

	assert.Equal(t, 1, rule.MinRequiredParams())
}
