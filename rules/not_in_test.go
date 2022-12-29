package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var notInRuleTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "username",
			"inputBag": bag.InputBag{
				"username": "goodUser",
			},
			"params": []string{
				"admin", "superuser",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"failed": map[string]any{
		"input": map[string]any{
			"selector": "username",
			"inputBag": bag.InputBag{
				"username": "admin",
			},
			"params": []string{
				"admin", "superuser",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The selected username is invalid.",
		},
	},
}

func TestNotInRule(t *testing.T) {
	rule := initNotInRule()

	for name, d := range notInRuleTestData {
		t.Run(name, func(t *testing.T) {
			data := d.(map[string]any)
			input := data["input"].(map[string]any)
			output := data["output"].(map[string]any)
			inputBag := input["inputBag"].(bag.InputBag)

			value, _ := inputBag.Get(input["selector"].(string))

			rule.AddParams(input["params"].([]string))
			res := rule.Validate(input["selector"].(string), value, inputBag)

			assert.Equal(t, output["validationFailed"].(bool), res.Failed())
			assert.Equal(t, output["validationError"].(string), res.Message())
		})
	}
}

func initNotInRule() *NotIn {
	notInRule := &NotIn{}
	notInRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.not_in":
			tr := "The selected :field: is invalid."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		default:
			return key
		}
	})
	return notInRule
}
