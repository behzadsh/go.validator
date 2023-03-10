package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var beforeOrEqualRuleTestData = map[string]any{
	"successfulBefore": map[string]any{
		"input": map[string]any{
			"selector": "start",
			"inputBag": bag.InputBag{
				"start": "2022-01-01",
				"end":   "2022-05-01",
			},
			"params": []string{
				"end", "America/New_York",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"successfulEqual": map[string]any{
		"input": map[string]any{
			"selector": "start",
			"inputBag": bag.InputBag{
				"start": "2022-01-01",
				"end":   "2022-01-01",
			},
			"params": []string{
				"end", "America/New_York",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"notBeforeOrEqual": map[string]any{
		"input": map[string]any{
			"selector": "start",
			"inputBag": bag.InputBag{
				"start": "2023-01-01",
				"end":   "2022-05-01",
			},
			"params": []string{
				"end", "America/New_York",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field start must be before or equal to field end.",
		},
	},
	"invalidFieldValue": map[string]any{
		"input": map[string]any{
			"selector": "start",
			"inputBag": bag.InputBag{
				"start": "invalid",
				"end":   "2022-01-01",
			},
			"params": []string{
				"end", "America/New_York",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field start must be a valid date time string.",
		},
	},
	"OtherFieldNotProvided": map[string]any{
		"input": map[string]any{
			"selector": "start",
			"inputBag": bag.InputBag{
				"start": "2022-05-01",
			},
			"params": []string{
				"end", "America/New_York",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field end is required.",
		},
	},
	"invalidOtherFieldValue": map[string]any{
		"input": map[string]any{
			"selector": "start",
			"inputBag": bag.InputBag{
				"start": "2022-01-01",
				"end":   "invalid",
			},
			"params": []string{
				"end", "America/New_York",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field end must be a valid date time string.",
		},
	},
}

func TestBeforeOrEqualRule(t *testing.T) {
	rule := initBeforeOrEqualRule()

	for name, d := range beforeOrEqualRuleTestData {
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

func initBeforeOrEqualRule() *BeforeOrEqual {
	beforeOrEqualRule := &BeforeOrEqual{}
	beforeOrEqualRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.before_or_equal":
			tr := "The field :field: must be before or equal to field :otherField:."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		case "validation.datetime":
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

	return beforeOrEqualRule
}

func TestBeforeOrEqual_MinRequiredParams(t *testing.T) {
	rule := initBeforeOrEqualRule()

	assert.Equal(t, 1, rule.MinRequiredParams())
}
