package bag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var inputBag = InputBag{
	"glossary": map[string]any{
		"title": "example glossary",
		"GlossDiv": map[string]any{
			"title": "S",
			"GlossList": []map[string]any{
				{
					"ID":        "SGML",
					"SortAs":    "SGML",
					"GlossTerm": "Standard Generalized Markup Language",
					"Acronym":   "SGML",
					"Abbrev":    "ISO 8879:1986",
					"GlossDef": map[string]any{
						"para":         "A meta-markup language, used to create markup languages such as DocBook.",
						"GlossSeeAlso": []string{"GML", "XML"},
					},
					"GlossSee": "markup",
				},
				{
					"ID":        "SSML",
					"SortAs":    "SSML",
					"GlossTerm": "Standard Simplified Markup Language",
					"Acronym":   "SSML",
					"Abbrev":    "ISO 18879:1986",
					"GlossDef": map[string]any{
						"para":         "A meta-markup language, used to create markup languages such as DocBook.",
						"GlossSeeAlso": []string{"SML", "XML"},
					},
					"GlossSee": "markup",
				},
			},
		},
	},
	"version": 1.0,
}

func TestInputBag_Get(t *testing.T) {
	v, ok := inputBag.Get("glossary.GlossDiv.GlossList.1.Acronym")

	assert.True(t, ok)
	assert.Equal(t, "SSML", v)

	v, ok = inputBag.Get("glossary.notExists")

	assert.False(t, ok)
	assert.Nil(t, v)

	v, ok = inputBag.Get("version")

	assert.True(t, ok)
	assert.Equal(t, 1.0, v)

	v, ok = inputBag.Get("notExists")

	assert.False(t, ok)
	assert.Nil(t, v)
}
