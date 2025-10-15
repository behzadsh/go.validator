package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var dateTimeRuleTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "birthday",
			"inputBag": bag.InputBag{
				"birthday": "1989-05-01",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"failed": map[string]any{
		"input": map[string]any{
			"selector": "birthday",
			"inputBag": bag.InputBag{
				"birthday": "invalid datetime string",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field birthday must be a valid date time string.",
		},
	},
}

func TestDateTimeRule(t *testing.T) {
	rule := initDateTimeRule()

	for name, d := range dateTimeRuleTestData {
		t.Run(name, func(t *testing.T) {
			data, _ := d.(map[string]any)
			input, _ := data["input"].(map[string]any)
			output, _ := data["output"].(map[string]any)
			inputBag, _ := input["inputBag"].(bag.InputBag)

			value, _ := inputBag.Get(input["selector"].(string))
			res := rule.Validate(input["selector"].(string), value, inputBag)

			assert.Equal(t, output["validationFailed"].(bool), res.Failed())
			assert.Equal(t, output["validationError"].(string), res.Message())
		})
	}
}

func initDateTimeRule() *DateTime {
	datetimeRule := &DateTime{}
	datetimeRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.datetime":
			tr := "The field :field: must be a valid date time string."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		default:
			return key
		}
	})

	return datetimeRule
}
