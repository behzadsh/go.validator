package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var sameAsRuleTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "passwordConfirmation",
			"inputBag": bag.InputBag{
				"password":             "mySecurePassword",
				"passwordConfirmation": "mySecurePassword",
			},
			"params": []string{
				"password",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"notTheSame": map[string]any{
		"input": map[string]any{
			"selector": "passwordConfirmation",
			"inputBag": bag.InputBag{
				"password":             "mySecurePassword",
				"passwordConfirmation": "anotherSecurePassword",
			},
			"params": []string{
				"password",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field passwordConfirmation and password must be match.",
		},
	},
	"notTheSameType": map[string]any{
		"input": map[string]any{
			"selector": "passwordConfirmation",
			"inputBag": bag.InputBag{
				"password":             "123456",
				"passwordConfirmation": 123456,
			},
			"params": []string{
				"password",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field passwordConfirmation and password must be match.",
		},
	},
}

func TestSameAsRule(t *testing.T) {
	rule := initSameAsRule()

	for name, d := range sameAsRuleTestData {
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

func initSameAsRule() *SameAs {
	sameAsRule := &SameAs{}
	sameAsRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.same_as":
			tr := "The field :field: and :otherField: must be match."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		default:
			return key
		}
	})
	return sameAsRule
}
