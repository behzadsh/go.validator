package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var maxRuleTestData = map[string]any{
	"successfulInteger": map[string]any{
		"input": map[string]any{
			"selector": "age",
			"inputBag": bag.InputBag{
				"age": 25,
			},
			"params": []string{
				"30",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"successfulFloat": map[string]any{
		"input": map[string]any{
			"selector": "score",
			"inputBag": bag.InputBag{
				"score": 86.58,
			},
			"params": []string{
				"100",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"failedInteger": map[string]any{
		"input": map[string]any{
			"selector": "age",
			"inputBag": bag.InputBag{
				"age": 25,
			},
			"params": []string{
				"18",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field age must not have a value more than 18.",
		},
	},
	"failedFloat": map[string]any{
		"input": map[string]any{
			"selector": "score",
			"inputBag": bag.InputBag{
				"score": 86.58,
			},
			"params": []string{
				"50",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field score must not have a value more than 50.",
		},
	},
}

func TestMaxRule(t *testing.T) {
	rule := initMaxRule()

	for name, d := range maxRuleTestData {
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

func initMaxRule() *Max {
	maxRule := &Max{}
	maxRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.max":
			tr := "The field :field: must not have a value more than :value:."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		default:
			return key
		}
	})
	return maxRule
}

func TestMax_MinRequiredParams(t *testing.T) {
	rule := initMaxRule()

	assert.Equal(t, 1, rule.MinRequiredParams())
}
