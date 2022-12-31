package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var startsWithRuleTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "className",
			"inputBag": bag.InputBag{
				"className": "UserController",
			},
			"params": []string{
				"User",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"failed": map[string]any{
		"input": map[string]any{
			"selector": "className",
			"inputBag": bag.InputBag{
				"className": "AccountController",
			},
			"params": []string{
				"User",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field className must starts with User.",
		},
	},
	"notString": map[string]any{
		"input": map[string]any{
			"selector": "className",
			"inputBag": bag.InputBag{
				"className": map[string]any{
					"name": "UserController",
					"path": "path/to/UserController.php",
				},
			},
			"params": []string{
				"User",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field className must have an string value.",
		},
	},
}

func TestStartsWithRule(t *testing.T) {
	rule := initStartsWithRule()

	for name, d := range startsWithRuleTestData {
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

func initStartsWithRule() *StartsWith {
	startsWithRule := &StartsWith{}
	startsWithRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.starts_with":
			tr := "The field :field: must starts with :value:."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		case "validation.string":
			tr := "The field :field: must have an string value."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		default:
			return key
		}
	})
	return startsWithRule
}
