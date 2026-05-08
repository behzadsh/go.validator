package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/bag"
)

var distinctRuleTestData = map[string]any{
	"uniqueStrings": map[string]any{
		"input": map[string]any{
			"selector": "tags",
			"inputBag": bag.InputBag{
				"tags": []string{"a", "b", "c"},
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"duplicates": map[string]any{
		"input": map[string]any{
			"selector": "tags",
			"inputBag": bag.InputBag{
				"tags": []string{"a", "b", "a"},
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field tags must not contain duplicate values.",
		},
	},
	"nonSlice": map[string]any{
		"input": map[string]any{
			"selector": "name",
			"inputBag": bag.InputBag{
				"name": "john",
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"emptySlice": map[string]any{
		"input": map[string]any{
			"selector": "nums",
			"inputBag": bag.InputBag{
				"nums": []int{},
			},
		},
		"output": map[string]any{
			"validationFailed": false,
			"validationError":  "",
		},
	},
	"nonAdjacentDuplicates": map[string]any{
		"input": map[string]any{
			"selector": "nums",
			"inputBag": bag.InputBag{
				"nums": []int{1, 2, 3, 2},
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field nums must not contain duplicate values.",
		},
	},
	"duplicatePointers": map[string]any{
		"input": map[string]any{
			"selector": "ptrs",
			"inputBag": bag.InputBag{
				"ptrs": func() [](*int) {
					a := 5
					b := 6
					pa := &a
					pb := &b
					return [](*int){pa, pb, pa}
				}(),
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field ptrs must not contain duplicate values.",
		},
	},
	"duplicateStructsComparable": map[string]any{
		"input": map[string]any{
			"selector": "items",
			"inputBag": bag.InputBag{
				"items": []struct{ A int }{{A: 1}, {A: 2}, {A: 1}},
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field items must not contain duplicate values.",
		},
	},
	"duplicateNils": map[string]any{
		"input": map[string]any{
			"selector": "vals",
			"inputBag": bag.InputBag{
				"vals": []any{nil, 1, nil},
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field vals must not contain duplicate values.",
		},
	},
	"nonComparableNoStringerDuplicates": map[string]any{
		"input": map[string]any{
			"selector": "items",
			"inputBag": bag.InputBag{
				"items": func() []struct{ B []int } {
					type s struct{ B []int }
					return []struct{ B []int }{
						s{B: []int{3, 4}},
						s{B: []int{3, 4}},
					}
				}(),
			},
		},
		"output": map[string]any{
			"validationFailed": true,
			"validationError":  "The field items must not contain duplicate values.",
		},
	},
}

func TestDistinctRule(t *testing.T) {
	rule := initDistinctRule()

	for name, d := range distinctRuleTestData {
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

func initDistinctRule() *Distinct {
	r := &Distinct{}
	r.AddTranslationFunction(func(_, key string, params ...map[string]string) string {
		var p map[string]string
		if len(params) > 0 {
			p = params[0]
		}

		switch key {
		case "validation.distinct":
			tr := "The field :field: must not contain duplicate values."
			for k, v := range p {
				tr = strings.ReplaceAll(tr, ":"+k+":", v)
			}

			return tr
		default:
			return key
		}
	})
	return r
}
