package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var notRegexRuleTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "variableName",
			"inputBag": bag.InputBag{
				"variableName": "user_name",
			},
			"params": []string{
				"^[a-zA-Z0-9]+$",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"failed": map[string]any{
		"input": map[string]any{
			"selector": "variableName",
			"inputBag": bag.InputBag{
				"variableName": "userName",
			},
			"params": []string{
				"^[a-zA-Z0-9]+$",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field variableName must not match the regex pattern ^[a-zA-Z0-9]+$.",
		},
	},
	"failedNotString": map[string]any{
		"input": map[string]any{
			"selector": "variableName",
			"inputBag": bag.InputBag{
				"variableName": map[string]any{
					"name": "userName",
				},
			},
			"params": []string{
				"^[a-zA-Z0-9]+$",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
}

func TestNotRegexRule(t *testing.T) {
	rule := initNotRegexRule()

	for name, d := range notRegexRuleTestData {
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

func initNotRegexRule() *NotRegex {
	notRegexRule := &NotRegex{}
	notRegexRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.not_regex":
			tr := "The field :field: must not match the regex pattern :pattern:."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
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
	return notRegexRule
}

func TestNotRegex_MinRequiredParams(t *testing.T) {
	rule := initNotRegexRule()

	assert.Equal(t, 1, rule.MinRequiredParams())
}

func TestNotRegex_RequiresField(t *testing.T) {
	rule := &NotRegex{}
	assert.False(t, rule.RequiresField())
}
