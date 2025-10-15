package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var digitsRuleTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "code",
			"inputBag": bag.InputBag{
				"code": "468432",
			},
			"params": []string{
				"6",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"failedUnequalDigits": map[string]any{
		"input": map[string]any{
			"selector": "code",
			"inputBag": bag.InputBag{
				"code": "46843",
			},
			"params": []string{
				"6",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field code must have exactly 6 digits.",
		},
	},
	"failedNotDigits": map[string]any{
		"input": map[string]any{
			"selector": "code",
			"inputBag": bag.InputBag{
				"code": "SIONDV",
			},
			"params": []string{
				"6",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field code must have exactly 6 digits.",
		},
	},
}

func TestDigitsRule(t *testing.T) {
	rule := initDigitsRule()

	for name, d := range digitsRuleTestData {
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

func initDigitsRule() *Digits {
	digitsRule := &Digits{}
	digitsRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.digits":
			tr := "The field :field: must have exactly :digitCount: digits."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		default:
			return key
		}
	})
	return digitsRule
}

func TestDigits_MinRequiredParams(t *testing.T) {
	rule := initDigitsRule()

	assert.Equal(t, 1, rule.MinRequiredParams())
}
