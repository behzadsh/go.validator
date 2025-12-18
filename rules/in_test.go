package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var inRuleTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "currency",
			"inputBag": bag.InputBag{
				"currency": "USD",
			},
			"params": []string{
				"EUR", "USD", "GBP", "JPY", "CHF",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"failed": map[string]any{
		"input": map[string]any{
			"selector": "currency",
			"inputBag": bag.InputBag{
				"currency": "HKD",
			},
			"params": []string{
				"EUR", "USD", "GBP", "JPY", "CHF",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The selected currency is invalid.",
		},
	},
}

func TestInRule(t *testing.T) {
	rule := initInRule()

	for name, d := range inRuleTestData {
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

func initInRule() *In {
	inRule := &In{}
	inRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.in":
			tr := "The selected :field: is invalid."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		default:
			return key
		}
	})
	return inRule
}

func TestIn_MinRequiredParams(t *testing.T) {
	rule := initInRule()

	assert.Equal(t, 2, rule.MinRequiredParams())
}

func TestIn_RequiresField(t *testing.T) {
	rule := &In{}
	assert.False(t, rule.RequiresField())
}
