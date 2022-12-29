package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var alphaNumRuleTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "title",
			"inputBag": bag.InputBag{
				"title": "user1",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"notAlphaNum": map[string]any{
		"input": map[string]any{
			"selector": "title",
			"inputBag": bag.InputBag{
				"title": "user-content 1",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field title must only contain letters and numbers.",
		},
	},
}

func TestAlphaNumRule(t *testing.T) {
	rule := initAlphaNumRule()

	for name, d := range alphaNumRuleTestData {
		t.Run(name, func(t *testing.T) {
			data := d.(map[string]any)
			input := data["input"].(map[string]any)
			output := data["output"].(map[string]any)
			inputBag := input["inputBag"].(bag.InputBag)

			value, _ := inputBag.Get(input["selector"].(string))

			res := rule.Validate(input["selector"].(string), value, inputBag)

			assert.Equal(t, output["validationFailed"].(bool), res.Failed())
			assert.Equal(t, output["validationError"].(string), res.Message())
		})
	}
}

func initAlphaNumRule() *AlphaNum {
	alphaRule := &AlphaNum{}
	alphaRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.alpha_num":
			tr := "The field :field: must only contain letters and numbers."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		default:
			return key
		}
	})
	return alphaRule
}
