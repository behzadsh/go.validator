package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var afterOrEqualRuleTestData = map[string]any{
	"successfulAfter": map[string]any{
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
	"successfulEqual": map[string]any{
		"input": map[string]any{
			"selector": "end",
			"inputBag": bag.InputBag{
				"start": "2022-01-01",
				"end":   "2022-01-01",
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
	"notAfterOrEqual": map[string]any{
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
			"validationError":  "The field end must be after or equal to field start.",
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

func TestAfterOrEqualRule(t *testing.T) {
	rule := initAfterOrEqualRule()

	for name, d := range afterOrEqualRuleTestData {
		t.Run(name, func(t *testing.T) {
			data := d.(map[string]any)
			input := data["input"].(map[string]any)
			output := data["output"].(map[string]any)
			inputBag := input["inputBag"].(bag.InputBag)

			value, exists := inputBag.Get(input["selector"].(string))

			rule.AddParams(input["params"].([]string))
			res := rule.Validate(input["selector"].(string), value, inputBag, exists)

			assert.Equal(t, output["validationFailed"].(bool), res.Failed())
			assert.Equal(t, output["validationError"].(string), res.Message())
		})
	}
}

func initAfterOrEqualRule() *AfterOrEqual {
	afterOrEqualRule := &AfterOrEqual{}
	afterOrEqualRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.after_or_equal":
			tr := "The field :field: must be after or equal to field :otherField:."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		case "validation.date_time":
			tr := "The field :field: must be a valid date time string."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		case "Validation.required":
			tr := "The field :otherField: is required."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		default:
			return key
		}
	})

	return afterOrEqualRule
}