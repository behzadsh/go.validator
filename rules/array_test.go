package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var arrayRuleTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "skills",
			"inputBag": bag.InputBag{
				"skills": []string{"go", "mongodb"},
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"notArray": map[string]any{
		"input": map[string]any{
			"selector": "skills",
			"inputBag": bag.InputBag{
				"skills": map[string]any{
					"backend":  "go",
					"database": "mongo",
				},
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field skills must be an array or slice.",
		},
	},
}

func TestArrayRule(t *testing.T) {
	rule := initArrayRule()

	for name, d := range arrayRuleTestData {
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

func initArrayRule() *Array {
	alphaRule := &Array{}
	alphaRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.array":
			tr := "The field :field: must be an array or slice."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		default:
			return key
		}
	})
	return alphaRule
}

func TestArray_RequiresField(t *testing.T) {
	rule := &Array{}
	assert.False(t, rule.RequiresField())
}
