package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var regexRuleTestData = map[string]any{
	"successful": map[string]any{
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
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"failed": map[string]any{
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
			"validationFailed": true,
			"validationError":  "The field variableName must match the regex pattern ^[a-zA-Z0-9]+$.",
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
			"validationFailed": true,
			"validationError":  "The field variableName must match the regex pattern ^[a-zA-Z0-9]+$.",
		},
	},
}

func TestRegexRule(t *testing.T) {
	rule := initRegexRule()

	for name, d := range regexRuleTestData {
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

func initRegexRule() *Regex {
	regexRule := &Regex{}
	regexRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.regex":
			tr := "The field :field: must match the regex pattern :pattern:."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		default:
			return key
		}
	})
	return regexRule
}

func TestRegex_MinRequiredParams(t *testing.T) {
	rule := initRegexRule()

	assert.Equal(t, 1, rule.MinRequiredParams())
}
