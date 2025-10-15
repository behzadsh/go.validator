package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var minLengthRuleTestData = map[string]any{
	"successfulString": map[string]any{
		"input": map[string]any{
			"selector": "functionName",
			"inputBag": bag.InputBag{
				"functionName": "doSomething",
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
	"successfulSlice": map[string]any{
		"input": map[string]any{
			"selector": "skills",
			"inputBag": bag.InputBag{
				"skills": []string{"Go", "Software Engineering", "TDD", "Software Architecture"},
			},
			"params": []string{
				"1",
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
			"selector": "functionName",
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
				"3",
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
				"functionName": "do",
			},
			"params": []string{
				"3",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field functionName must not have a length less than 3.",
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
			"validationError":  "The field skills must not have a length less than 5.",
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
				"5",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field user must not have a length less than 5.",
		},
	},
}

func TestMinLengthRule(t *testing.T) {
	rule := initMinLengthRule()

	for name, d := range minLengthRuleTestData {
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

func initMinLengthRule() *MinLength {
	minLengthRule := &MinLength{}
	minLengthRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.min_length":
			tr := "The field :field: must not have a length less than :value:."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		default:
			return key
		}
	})
	return minLengthRule
}

func TestMinLength_MinRequiredParams(t *testing.T) {
	rule := initMinLengthRule()

	assert.Equal(t, 1, rule.MinRequiredParams())
}
