package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var minDigitsRuleTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "code",
			"inputBag": bag.InputBag{
				"code": "46843",
			},
			"params": []string{
				"5",
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
				"code": "4684",
			},
			"params": []string{
				"5",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field code must not have less than 5 digits.",
		},
	},
}

func TestMinDigitsRule(t *testing.T) {
	rule := initMinDigitsRule()

	for name, d := range minDigitsRuleTestData {
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

func initMinDigitsRule() *MinDigits {
	minDigitsRule := &MinDigits{}
	minDigitsRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.min_digits":
			tr := "The field :field: must not have less than :digitCount: digits."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		default:
			return key
		}
	})
	return minDigitsRule
}

func TestMinDigits_MinRequiredParams(t *testing.T) {
	rule := initMinDigitsRule()

	assert.Equal(t, 1, rule.MinRequiredParams())
}

func TestMinDigits_RequiresField(t *testing.T) {
	rule := &MinDigits{}
	assert.False(t, rule.RequiresField())
}
