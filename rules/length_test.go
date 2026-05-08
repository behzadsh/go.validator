package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var lengthRuleTestData = map[string]any{
	"successfulString": map[string]any{
		"input": map[string]any{
			"selector": "functionName",
			"inputBag": bag.InputBag{
				"functionName": "doSomething",
			},
			"params": []string{
				"11",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"successfulSlice": map[string]any{
		"input": map[string]any{
			"selector": "skills",
			"inputBag": bag.InputBag{
				"skills": []string{"Go", "Software Engineering", "TDD", "Software Architecture"},
			},
			"params": []string{
				"4",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"successfulMap": map[string]any{
		"input": map[string]any{
			"selector": "user",
			"inputBag": bag.InputBag{
				"user": map[string]any{
					"userName": "johnDoe",
					"name":     "John Doe",
					"age":      35,
				},
			},
			"params": []string{
				"3",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"successfulNotExists": map[string]any{
		"input": map[string]any{
			"selector": "user",
			"inputBag": bag.InputBag{},
			"params": []string{
				"3",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"successfulOtherType": map[string]any{
		"input": map[string]any{
			"selector": "functionName",
			"inputBag": bag.InputBag{
				"functionName": struct{}{},
			},
			"params": []string{
				"10",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"failedString": map[string]any{
		"input": map[string]any{
			"selector": "functionName",
			"inputBag": bag.InputBag{
				"functionName": "doSomethingUseful",
			},
			"params": []string{
				"10",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field functionName must have exact length of 10.",
		},
	},
	"failedSlice": map[string]any{
		"input": map[string]any{
			"selector": "skills",
			"inputBag": bag.InputBag{
				"skills": []string{"Go", "Software Engineering", "TDD", "Software Architecture"},
			},
			"params": []string{
				"5",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field skills must have exact length of 5.",
		},
	},
	"failedMap": map[string]any{
		"input": map[string]any{
			"selector": "user",
			"inputBag": bag.InputBag{
				"user": map[string]any{
					"userName": "johnDoe",
					"name":     "John Doe",
					"age":      35,
				},
			},
			"params": []string{
				"2",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field user must have exact length of 2.",
		},
	},
}

func TestLengthRule(t *testing.T) {
	rule := initLengthRule()

	for name, d := range lengthRuleTestData {
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

func initLengthRule() *Length {
	lengthRule := &Length{}
	lengthRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.length":
			tr := "The field :field: must have exact length of :value:."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		default:
			return key
		}
	})
	return lengthRule
}

func TestLength_MinRequiredParams(t *testing.T) {
	rule := initLengthRule()

	assert.Equal(t, 1, rule.MinRequiredParams())
}
