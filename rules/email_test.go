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
	"successfulFormatIgnoredParam": map[string]any{
		"input": map[string]any{
			"selector": "email",
			"inputBag": bag.InputBag{
				"email": "user@example.com",
			},
			"params": []string{"ignored"},
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
	"failedFormat2": map[string]any{
		"input": map[string]any{
			"selector": "email",
			"inputBag": bag.InputBag{
				"email": "user@something@example.com",
			},
			"params": []string{},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field email must be a valid email.",
		},
	},
	"failedFormat3": map[string]any{
		"input": map[string]any{
			"selector": "email",
			"inputBag": bag.InputBag{
				"email": "veryveryveryveryveryveryveryveryveryveryveryverylongusernameforanemail@example.com",
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
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		default:
			return key
		}
	})
	return emailRule
}

func TestEmail_MinRequiredParams(t *testing.T) {
	rule := initEmailRule()

	assert.Equal(t, 0, rule.MinRequiredParams())
}

func TestEmail_RequiresField(t *testing.T) {
	rule := &Email{}
	assert.False(t, rule.RequiresField())
}
