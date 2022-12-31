package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var greaterThanRuleTestData = map[string]any{
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
	"successfulString": map[string]any{
		"input": map[string]any{
			"selector": "functionName",
			"inputBag": bag.InputBag{
				"functionName": "doSomethingUseful",
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
					"age":      35,
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
	"failedInteger": map[string]any{
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
			"validationFailed": true,
			"validationError":  "The field age must have a value or length greater than 18.",
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
			"validationError":  "The field score must have a value or length greater than 90.",
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
			"validationError":  "The field functionName must have a value or length greater than 3.",
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
			"validationError":  "The field skills must have a value or length greater than 2.",
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
			"validationError":  "The field user must have a value or length greater than 2.",
		},
	},
}

func TestGreaterThanRule(t *testing.T) {
	rule := initGreaterThanRule()

	for name, d := range greaterThanRuleTestData {
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

func initGreaterThanRule() *GreaterThan {
	greaterThanRule := &GreaterThan{}
	greaterThanRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.gt":
			tr := "The field :field: must have a value or length greater than :value:."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		default:
			return key
		}
	})
	return greaterThanRule
}
