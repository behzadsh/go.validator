package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var notEmptyRuleTestData = map[string]any{
	"emptyString": map[string]any{
		"input": map[string]any{
			"selector": "name",
			"inputBag": bag.InputBag{
				"name": "",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field name must not be empty.",
		},
	},
	"nilValue": map[string]any{
		"input": map[string]any{
			"selector": "name",
			"inputBag": bag.InputBag{
				"name": nil,
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field name must not be empty.",
		},
	},
	"emptySlice": map[string]any{
		"input": map[string]any{
			"selector": "skills",
			"inputBag": bag.InputBag{
				"skills": []string{},
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field skills must not be empty.",
		},
	},
	"emptyBoolean": map[string]any{
		"input": map[string]any{
			"selector": "agreed",
			"inputBag": bag.InputBag{
				"agreed": false,
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field agreed must not be empty.",
		},
	},
	"notEmpty": map[string]any{
		"input": map[string]any{
			"selector": "name",
			"inputBag": bag.InputBag{
				"name": "John",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
}

func TestNotEmptyRule(t *testing.T) {
	rule := initNotEmptyRule()

	for name, d := range notEmptyRuleTestData {
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

func initNotEmptyRule() *NotEmpty {
	notEmptyRule := &NotEmpty{}
	notEmptyRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.not_empty":
			tr := "The field :field: must not be empty."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		default:
			return key
		}
	})
	return notEmptyRule
}

func TestNotEmpty_RequiresField(t *testing.T) {
	rule := &NotEmpty{}
	assert.True(t, rule.RequiresField())
}
