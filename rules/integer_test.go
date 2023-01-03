package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var integerRuleTestData = map[string]any{
	"successfulInt": map[string]any{
		"input": map[string]any{
			"selector": "age",
			"inputBag": bag.InputBag{
				"age": 25,
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"successfulNotExists": map[string]any{
		"input": map[string]any{
			"selector": "age",
			"inputBag": bag.InputBag{},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"failedFloat": map[string]any{
		"input": map[string]any{
			"selector": "age",
			"inputBag": bag.InputBag{
				"age": 25.9,
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field age must have an integer value.",
		},
	},
	"failedNumericString": map[string]any{
		"input": map[string]any{
			"selector": "year",
			"inputBag": bag.InputBag{
				"year": "1989",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field year must have an integer value.",
		},
	},
	"failedString": map[string]any{
		"input": map[string]any{
			"selector": "age",
			"inputBag": bag.InputBag{
				"age": "John Doe",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field age must have an integer value.",
		},
	},
	"failedSlice": map[string]any{
		"input": map[string]any{
			"selector": "age",
			"inputBag": bag.InputBag{
				"age": []string{"John", "Doe"},
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field age must have an integer value.",
		},
	},
	"failedMap": map[string]any{
		"input": map[string]any{
			"selector": "age",
			"inputBag": bag.InputBag{
				"age": map[string]any{
					"username": "johnDoe",
					"name":     "John Doe",
				},
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field age must have an integer value.",
		},
	},
}

func TestIntegerRule(t *testing.T) {
	rule := initIntegerRule()

	for name, d := range integerRuleTestData {
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

func initIntegerRule() *Integer {
	integerRule := &Integer{}
	integerRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.integer":
			tr := "The field :field: must have an integer value."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		default:
			return key
		}
	})
	return integerRule
}
