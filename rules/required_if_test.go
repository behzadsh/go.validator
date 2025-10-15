package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var requiredIfRuleTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "email",
			"inputBag": bag.InputBag{
				"email": "user@example.com",
				"type":  "user",
			},
			"params": []string{
				"type", "user",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"otherFieldNotExist": map[string]any{
		"input": map[string]any{
			"selector": "email",
			"inputBag": bag.InputBag{
				"email": "user@example.com",
			},
			"params": []string{
				"type", "user",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"fieldNotExist": map[string]any{
		"input": map[string]any{
			"selector": "email",
			"inputBag": bag.InputBag{
				"type": "user",
			},
			"params": []string{
				"type", "user",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field email is required when type is user.",
		},
	},
}

func TestRequiredIfRule(t *testing.T) {
	rule := initRequiredIfRule()

	for name, d := range requiredIfRuleTestData {
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

func initRequiredIfRule() *RequiredIf {
	requiredIfRule := &RequiredIf{}
	requiredIfRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.required_if":
			tr := "The field :field: is required when :otherField: is :value:."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		default:
			return key
		}
	})
	return requiredIfRule
}

func TestRequiredIf_MinRequiredParams(t *testing.T) {
	rule := initRequiredIfRule()

	assert.Equal(t, 2, rule.MinRequiredParams())
}
