package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var maxDigitsRuleTestData = map[string]any{
	"successful": map[string]any{
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
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"failedReachedMaxDigits": map[string]any{
		"input": map[string]any{
			"selector": "code",
			"inputBag": bag.InputBag{
				"code": "468439",
			},
			"params": []string{
				"5",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field code must not have more than 5 digits.",
		},
	},
}

func TestMaxDigitsRule(t *testing.T) {
	rule := initMaxDigitsRule()

	for name, d := range maxDigitsRuleTestData {
		t.Run(name, func(t *testing.T) {
			data := d.(map[string]any)
			input := data["input"].(map[string]any)
			output := data["output"].(map[string]any)
			inputBag := input["inputBag"].(bag.InputBag)

			value, _ := inputBag.Get(input["selector"].(string))

			rule.AddParams(input["params"].([]string))
			res := rule.Validate(input["selector"].(string), value, inputBag)

			assert.Equal(t, output["validationFailed"].(bool), res.Failed())
			assert.Equal(t, output["validationError"].(string), res.Message())
		})
	}
}

func initMaxDigitsRule() *MaxDigits {
	maxDigitsRule := &MaxDigits{}
	maxDigitsRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.max_digits":
			tr := "The field :field: must not have more than :digitCount: digits."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		default:
			return key
		}
	})
	return maxDigitsRule
}
