package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var minRuleTestData = map[string]any{
	"successfulInteger": map[string]any{
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
				"80",
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
				"30",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field age must not have a value less than 30.",
		},
	},
	"failedFloat": map[string]any{
		"input": map[string]any{
			"selector": "score",
			"inputBag": bag.InputBag{
				"score": 86.58,
			},
			"params": []string{
				"90",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field score must not have a value less than 90.",
		},
	},
}

func TestMinRule(t *testing.T) {
	rule := initMinRule()

	for name, d := range minRuleTestData {
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

func initMinRule() *Min {
	minRule := &Min{}
	minRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.min":
			tr := "The field :field: must not have a value less than :value:."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		default:
			return key
		}
	})
	return minRule
}

func TestMin_MinRequiredParams(t *testing.T) {
	rule := initMinRule()

	assert.Equal(t, 1, rule.MinRequiredParams())
}

func TestMin_RequiresField(t *testing.T) {
	rule := &Min{}
	assert.False(t, rule.RequiresField())
}
