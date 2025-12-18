package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var urlRuleTestData = map[string]any{
	"successfulHttp": map[string]any{
		"input": map[string]any{
			"selector": "site",
			"inputBag": bag.InputBag{
				"site": "https://example.com/path?q=1",
			},
			"params": []string{},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"successfulNoScheme": map[string]any{
		"input": map[string]any{
			"selector": "site",
			"inputBag": bag.InputBag{
				"site": "example.com/path?q=1",
			},
			"params": []string{},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"failedNoScheme": map[string]any{
		"input": map[string]any{
			"selector": "site",
			"inputBag": bag.InputBag{
				"site": "example.com",
			},
			"params": []string{"scheme"},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field site must be a valid url.",
		},
	},
	"failedImpliedScheme": map[string]any{
		"input": map[string]any{
			"selector": "site",
			"inputBag": bag.InputBag{
				"site": "/relative/path",
			},
			"params": []string{},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field site must be a valid url.",
		},
	},
	"failedGibberishNoScheme": map[string]any{
		"input": map[string]any{
			"selector": "site",
			"inputBag": bag.InputBag{
				"site": "jhfdskhk",
			},
			"params": []string{},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field site must be a valid url.",
		},
	},
	"failedGibberishWithScheme": map[string]any{
		"input": map[string]any{
			"selector": "site",
			"inputBag": bag.InputBag{
				"site": "http://jhfdskhk",
			},
			"params": []string{},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field site must be a valid url.",
		},
	},
	"emptyInput": map[string]any{
		"input": map[string]any{
			"selector": "site",
			"inputBag": bag.InputBag{
				"site": "",
			},
			"params": []string{},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
}

func TestURLRule(t *testing.T) {
	for name, d := range urlRuleTestData {
		t.Run(
			name, func(t *testing.T) {
				rule := initURLRule()
				data, _ := d.(map[string]any)
				input, _ := data["input"].(map[string]any)
				output, _ := data["output"].(map[string]any)
				inputBag, _ := input["inputBag"].(bag.InputBag)

				value, _ := inputBag.Get(input["selector"].(string))

				if params, ok := input["params"].([]string); ok {
					rule.AddParams(params)
				}
				res := rule.Validate(input["selector"].(string), value, inputBag)

				assert.Equal(t, output["validationFailed"].(bool), res.Failed())
				assert.Equal(t, output["validationError"].(string), res.Message())
			},
		)
	}
}

func initURLRule() *URL {
	urlRule := &URL{}
	urlRule.AddTranslationFunction(
		func(_, key string, params ...map[string]string) string {
			var p map[string]string
			if len(params) > 0 {
				p = params[0]
			}

			switch key {
			case "validation.url":
				tr := "The field :field: must be a valid url."
				for k, v := range p {
					tr = strings.ReplaceAll(tr, ":"+k+":", v)
				}

				return tr
			default:
				return key
			}
		},
	)
	return urlRule
}
