package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var requiredWithAllRuleTestData = map[string]any{
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
	"notExistsButOK": map[string]any{
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
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"failed": map[string]any{
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
			"validationError":  "The field email is required when type and username are present.",
		},
	},
	"failed2": map[string]any{
		"input": map[string]any{
			"selector": "email",
			"inputBag": bag.InputBag{
				"type":     "user",
				"username": "goodUser",
				"password": "mySecurePassword",
			},
			"params": []string{
				"type", "username", "password",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field email is required when type, username, and password are present.",
		},
	},
}

func TestRequiredWithAllRule(t *testing.T) {
	rule := initRequiredWithAllRule()

	for name, d := range requiredWithAllRuleTestData {
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

func initRequiredWithAllRule() *RequiredWithAll {
	requiredWithAllRule := &RequiredWithAll{}
	requiredWithAllRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.required_with_all":
			tr := "The field :field: is required when :otherFields: are present."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		default:
			return key
		}
	})
	return requiredWithAllRule
}

func TestRequiredWithAll_MinRequiredParams(t *testing.T) {
	rule := initRequiredWithAllRule()

	assert.Equal(t, 2, rule.MinRequiredParams())
}
