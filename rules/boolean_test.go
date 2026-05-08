package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var booleanRuleTestData = map[string]any{
	"successfulBoolean": map[string]any{
		"input": map[string]any{
			"selector": "agree",
			"inputBag": bag.InputBag{
				"agree": true,
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"successfulInt": map[string]any{
		"input": map[string]any{
			"selector": "agree",
			"inputBag": bag.InputBag{
				"agree": 1,
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"successfulString": map[string]any{
		"input": map[string]any{
			"selector": "agree",
			"inputBag": bag.InputBag{
				"agree": "true",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"failedInvalidString": map[string]any{
		"input": map[string]any{
			"selector": "agree",
			"inputBag": bag.InputBag{
				"agree": "invalidString",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field agree must be boolean.",
		},
	},
	"failedNil": map[string]any{
		"input": map[string]any{
			"selector": "agree",
			"inputBag": bag.InputBag{
				"agree": nil,
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field agree must be boolean.",
		},
	},
	"failedNotExists": map[string]any{
		"input": map[string]any{
			"selector": "agree",
			"inputBag": bag.InputBag{},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field agree must be boolean.",
		},
	},
}

func TestBooleanRule(t *testing.T) {
	rule := initBooleanRule()

	for name, d := range booleanRuleTestData {
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

func initBooleanRule() *Boolean {
	booleanRule := &Boolean{}
	booleanRule.AddTranslationFunction(func(local, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.boolean":
			tr := "The field :field: must be boolean."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		default:
			return key
		}
	})

	return booleanRule
}
