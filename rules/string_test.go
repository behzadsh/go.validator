package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var stringRuleTestData = map[string]any{
	"successfulString": map[string]any{
		"input": map[string]any{
			"selector": "name",
			"inputBag": bag.InputBag{
				"name": "John Doe",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"successfulEmptyString": map[string]any{
		"input": map[string]any{
			"selector": "name",
			"inputBag": bag.InputBag{
				"name": "",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"successfulNotExists": map[string]any{
		"input": map[string]any{
			"selector": "name",
			"inputBag": bag.InputBag{},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"failedFloat": map[string]any{
		"input": map[string]any{
			"selector": "name",
			"inputBag": bag.InputBag{
				"name": 25.9,
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field name must have an string value.",
		},
	},
	"failedInteger": map[string]any{
		"input": map[string]any{
			"selector": "name",
			"inputBag": bag.InputBag{
				"name": 1989,
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field name must have an string value.",
		},
	},
	"failedSlice": map[string]any{
		"input": map[string]any{
			"selector": "name",
			"inputBag": bag.InputBag{
				"name": []string{"John", "Doe"},
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field name must have an string value.",
		},
	},
	"failedMap": map[string]any{
		"input": map[string]any{
			"selector": "name",
			"inputBag": bag.InputBag{
				"name": map[string]any{
					"username": "johnDoe",
					"name":     "John Doe",
				},
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field name must have an string value.",
		},
	},
}

func TestStringRule(t *testing.T) {
	rule := initStringRule()

	for name, d := range stringRuleTestData {
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

func initStringRule() *String {
	stringRule := &String{}
	stringRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.string":
			tr := "The field :field: must have an string value."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		default:
			return key
		}
	})
	return stringRule
}
