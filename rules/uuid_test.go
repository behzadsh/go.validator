package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var uuidRuleTestData = map[string]any{
	"successful": map[string]any{
		"input": map[string]any{
			"selector": "id",
			"inputBag": bag.InputBag{
				"id": "c27b23b2-932a-469c-b98f-9150a092d002",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"failed": map[string]any{
		"input": map[string]any{
			"selector": "id",
			"inputBag": bag.InputBag{
				"id": "g27b2yb2-93va-469c-b98f-9150a092h002",
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field id is not a valid uuid.",
		},
	},
}

func TestUUIDRule(t *testing.T) {
	rule := initUUIDRule()

	for name, d := range uuidRuleTestData {
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

func initUUIDRule() *UUID {
	uuidRule := &UUID{}
	uuidRule.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.uuid":
			tr := "The field :field: is not a valid uuid."
			for k, v := range p {
				tr = strings.Replace(tr, ":"+k+":", v, -1)
			}

			return tr
		default:
			return key
		}
	})
	return uuidRule
}
