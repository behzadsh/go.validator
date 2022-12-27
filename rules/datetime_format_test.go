package rules

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var dateTimeFormatTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "loggedAt",
			"inputBag": bag.InputBag{
				"loggedAt": "2022-12-28T23:54:34+03:30",
			},
			"params": []string{
				time.RFC3339,
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"incorrectFormat": map[string]any{
		"input": map[string]any{
			"selector": "loggedAt",
			"inputBag": bag.InputBag{
				"loggedAt": "2022-12-28 23:54:34",
			},
			"params": []string{
				time.RFC3339,
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "the field loggedAt must be in 2006-01-02T15:04:05Z07:00 format.",
		},
	},
}

func TestDateTimeFormatRule(t *testing.T) {
	rule := initDateTimeFormatRule()

	for name, d := range dateTimeFormatTestData {
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

func initDateTimeFormatRule() *DateTimeFormat {
	dateTimeFormatRule := &DateTimeFormat{}
	dateTimeFormatRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.datetime_format":
			tr := "the field :field: must be in :format: format."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		default:
			return key
		}
	})
	return dateTimeFormatRule
}
