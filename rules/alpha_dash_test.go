package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var alphaDashRuleTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "title",
			"inputBag": bag.InputBag{
				"title": "user-content_1",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"notAlphaDash": map[string]any{
		"input": map[string]any{
			"selector": "title",
			"inputBag": bag.InputBag{
				"title": "user-content 1",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field title must only contain letters, numbers, dashes and underscores.",
		},
	},
}

func TestAlphaDashRule(t *testing.T) {
	rule := initAlphaDashRule()

	for name, d := range alphaDashRuleTestData {
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

func initAlphaDashRule() *AlphaDash {
	alphaRule := &AlphaDash{}
	alphaRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.alpha_dash":
			tr := "The field :field: must only contain letters, numbers, dashes and underscores."
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
