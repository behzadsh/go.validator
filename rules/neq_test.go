package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var notEqualRuleTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "username",
			"inputBag": bag.InputBag{
				"username": "goodUser",
			},
			"params": []string{
				"admin",
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
				"admin",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field username could not be admin.",
		},
	},
}

func TestNotEqualRule(t *testing.T) {
	rule := initNotEqualRule()

	for name, d := range notEqualRuleTestData {
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

func initNotEqualRule() *NotEqual {
	notEqualRule := &NotEqual{}
	notEqualRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.neq":
			tr := "The field :field: could not be :value:."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		default:
			return key
		}
	})
	return notEqualRule
}
