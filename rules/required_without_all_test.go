package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var requiredWithoutAllRuleTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "email",
			"inputBag": bag.InputBag{
				"email": "user@example.com",
			},
			"params": []string{
				"phone", "username",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"notExistsButOK": map[string]any{
		"input": map[string]any{
			"selector": "email",
			"inputBag": bag.InputBag{
				"username": "goodUser",
			},
			"params": []string{
				"phone", "username",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"failed": map[string]any{
		"input": map[string]any{
			"selector": "email",
			"inputBag": bag.InputBag{},
			"params": []string{
				"phone", "username",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field email is required when phone and username are not present.",
		},
	},
}

func TestRequiredWithoutAllRule(t *testing.T) {
	rule := initRequiredWithoutAllRule()

	for name, d := range requiredWithoutAllRuleTestData {
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

func initRequiredWithoutAllRule() *RequiredWithoutAll {
	requiredWithoutAllRule := &RequiredWithoutAll{}
	requiredWithoutAllRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.required_without_all":
			tr := "The field :field: is required when :otherFields: are not present."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		default:
			return key
		}
	})
	return requiredWithoutAllRule
}
