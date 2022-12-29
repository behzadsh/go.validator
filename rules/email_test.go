package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var emailRuleTestData = map[string]any{
	"successfulFormat": map[string]any{
		"input": map[string]any{
			"selector": "email",
			"inputBag": bag.InputBag{
				"email": "user@example.com",
			},
			"params": []string{},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"successfulMXRecord": map[string]any{
		"input": map[string]any{
			"selector": "email",
			"inputBag": bag.InputBag{
				"email": "user@gmail.com",
			},
			"params": []string{"mx"},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"failedFormat": map[string]any{
		"input": map[string]any{
			"selector": "email",
			"inputBag": bag.InputBag{
				"email": "é&ààà@example.com",
			},
			"params": []string{},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field email must be a valid email.",
		},
	},
	"failedMXRecord": map[string]any{
		"input": map[string]any{
			"selector": "email",
			"inputBag": bag.InputBag{
				"email": "admin@notarealdomain12345.com",
			},
			"params": []string{"mx"},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field email must be a valid email.",
		},
	},
}

func TestEmailRule(t *testing.T) {
	rule := initEmailRule()

	for name, d := range emailRuleTestData {
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

func initEmailRule() *Email {
	emailRule := &Email{}
	emailRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.email":
			tr := "The field :field: must be a valid email."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		default:
			return key
		}
	})
	return emailRule
}
