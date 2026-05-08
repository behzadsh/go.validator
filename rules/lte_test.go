package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var lessThanEqualRuleTestData = map[string]any{
	"successfulInteger": map[string]any{
		"input": map[string]any{
			"selector": "age",
			"inputBag": bag.InputBag{
				"age": 18,
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
				"score": 79.99,
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
	"successfulString": map[string]any{
		"input": map[string]any{
			"selector": "functionName",
			"inputBag": bag.InputBag{
				"functionName": "doSomething",
			},
			"params": []string{
				"12",
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
	"successfulNotExist": map[string]any{
		"input": map[string]any{
			"selector": "age",
			"inputBag": bag.InputBag{},
			"params": []string{
				"18",
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
				"age": 19,
			},
			"params": []string{
				"18",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field age must have a value or length less than or equal to 18.",
		},
	},
	"failedFloat": map[string]any{
		"input": map[string]any{
			"selector": "score",
			"inputBag": bag.InputBag{
				"score": 86.58,
			},
			"params": []string{
				"80.50",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field score must have a value or length less than or equal to 80.5.",
		},
	},
	"failedString": map[string]any{
		"input": map[string]any{
			"selector": "functionName",
			"inputBag": bag.InputBag{
				"functionName": "doSomething",
			},
			"params": []string{
				"10",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field functionName must have a value or length less than or equal to 10.",
		},
	},
	"failedSlice": map[string]any{
		"input": map[string]any{
			"selector": "skills",
			"inputBag": bag.InputBag{
				"skills": []string{"Go", "Software Engineering", "TDD", "Software Architecture"},
			},
			"params": []string{
				"3",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field skills must have a value or length less than or equal to 3.",
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
			"validationError":  "The field user must have a value or length less than or equal to 2.",
		},
	},
	"successfulOtherTypes": map[string]any{
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
}

func TestLessThanEqualRule(t *testing.T) {
	rule := initLessThanEqualRule()

	for name, d := range lessThanEqualRuleTestData {
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

func initLessThanEqualRule() *LessThanEqual {
	lessThanEqualRule := &LessThanEqual{}
	lessThanEqualRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.lte":
			tr := "The field :field: must have a value or length less than or equal to :value:."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		default:
			return key
		}
	})
	return lessThanEqualRule
}

func TestLessThanEqual_MinRequiredParams(t *testing.T) {
	rule := initLessThanEqualRule()

	assert.Equal(t, 1, rule.MinRequiredParams())
}
