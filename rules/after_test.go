package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var afterRuleTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "end",
			"inputBag": bag.InputBag{
				"start": "2022-01-01",
				"end":   "2022-05-01",
			},
			"params": []string{
				"start", "America/New_York",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"notAfter": map[string]any{
		"input": map[string]any{
			"selector": "end",
			"inputBag": bag.InputBag{
				"start": "2023-01-01",
				"end":   "2022-05-01",
			},
			"params": []string{
				"start", "America/New_York",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field end must be after field start.",
		},
	},
	"invalidFieldValue": map[string]any{
		"input": map[string]any{
			"selector": "end",
			"inputBag": bag.InputBag{
				"start": "2022-01-01",
				"end":   "invalid",
			},
			"params": []string{
				"start", "America/New_York",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field end must be a valid date time string.",
		},
	},
	"OtherFieldNotProvided": map[string]any{
		"input": map[string]any{
			"selector": "end",
			"inputBag": bag.InputBag{
				"end": "2022-05-01",
			},
			"params": []string{
				"start", "America/New_York",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field start is required.",
		},
	},
	"invalidOtherFieldValue": map[string]any{
		"input": map[string]any{
			"selector": "end",
			"inputBag": bag.InputBag{
				"start": "invalid",
				"end":   "2022-01-01",
			},
			"params": []string{
				"start", "America/New_York",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field start must be a valid date time string.",
		},
	},
}

func TestAfterRule(t *testing.T) {
	rule := initAfterRule()

	for name, d := range afterRuleTestData {
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

func initAfterRule() *After {
	afterRule := &After{}
	afterRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.after":
			tr := "The field :field: must be after field :otherField:."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		case "validation.datetime":
			tr := "The field :field: must be a valid date time string."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		case "Validation.required":
			tr := "The field :field: is required."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		default:
			return key
		}
	})
	return afterRule
}

func TestAfter_MinRequiredParams(t *testing.T) {
	rule := initAfterRule()

	assert.Equal(t, 1, rule.MinRequiredParams())
}
