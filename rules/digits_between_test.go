package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var digitsBetweenRuleTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "code",
			"inputBag": bag.InputBag{
				"code": "46843",
			},
			"params": []string{
				"4", "6",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"failedLessThanMinDigits": map[string]any{
		"input": map[string]any{
			"selector": "code",
			"inputBag": bag.InputBag{
				"code": "464",
			},
			"params": []string{
				"4", "6",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field code must be between 4 and 6 digits.",
		},
	},
}

func TestDigitsBetweenRule(t *testing.T) {
	rule := initDigitsBetweenRule()

	for name, d := range digitsBetweenRuleTestData {
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

func initDigitsBetweenRule() *DigitsBetween {
	digitsBetweenRule := &DigitsBetween{}
	digitsBetweenRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.digits_between":
			tr := "The field :field: must be between :min: and :max: digits."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		default:
			return key
		}
	})
	return digitsBetweenRule
}
