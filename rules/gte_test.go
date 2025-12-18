package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var greaterThanEqualRuleTestData = map[string]any{
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
				"score": 80.01,
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
				"functionName": "can",
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
				"skills": []string{"Go", "Software Engineering"},
			},
			"params": []string{
				"2",
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
				},
			},
			"params": []string{
				"2",
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
	"failedInteger": map[string]any{
		"input": map[string]any{
			"selector": "age",
			"inputBag": bag.InputBag{
				"age": 16,
			},
			"params": []string{
				"18",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field age must have a value or length greater than or equal to 18.",
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
			"validationError":  "The field score must have a value or length greater than or equal to 90.",
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
			"validationError":  "The field functionName must have a value or length greater than or equal to 3.",
		},
	},
	"failedSlice": map[string]any{
		"input": map[string]any{
			"selector": "skills",
			"inputBag": bag.InputBag{
				"skills": []string{"Go"},
			},
			"params": []string{
				"2",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field skills must have a value or length greater than or equal to 2.",
		},
	},
	"failedMap": map[string]any{
		"input": map[string]any{
			"selector": "user",
			"inputBag": bag.InputBag{
				"user": map[string]any{
					"userName": "johnDoe",
				},
			},
			"params": []string{
				"2",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field user must have a value or length greater than or equal to 2.",
		},
	},
}

func TestGreaterThanEqualRule(t *testing.T) {
	rule := initGreaterThanEqualRule()

	for name, d := range greaterThanEqualRuleTestData {
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

func initGreaterThanEqualRule() *GreaterThanEqual {
	greaterThanEqualRule := &GreaterThanEqual{}
	greaterThanEqualRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.gte":
			tr := "The field :field: must have a value or length greater than or equal to :value:."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		default:
			return key
		}
	})
	return greaterThanEqualRule
}

func TestGreaterThanEqual_MinRequiredParams(t *testing.T) {
	rule := initGreaterThanEqualRule()

	assert.Equal(t, 1, rule.MinRequiredParams())
}

func TestGreaterThanEqual_RequiresField(t *testing.T) {
	rule := &GreaterThanEqual{}
	assert.False(t, rule.RequiresField())
}
