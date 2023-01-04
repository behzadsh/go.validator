package bag

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type sampleStruct struct {
	UserName    string   `json:"username"`
	Email       string   `json:"email"`
	Age         int      `json:"age"`
	Skills      []string `json:"skills"`
	Experiences []struct {
		JobTitle string    `json:"jobTitle"`
		Salary   int       `json:"salary"`
		From     time.Time `json:"from"`
		To       time.Time `json:"to"`
	} `json:"experiences"`
	Socials struct {
		Twitter  string `json:"twitterHandle"`
		Facebook string `json:"facebookUsername"`
	} `json:"socials"`
}

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

	v, ok = inputBag.Get("glossary.GlossDiv.GlossList.2.Acronym")
	assert.False(t, ok)
	assert.Nil(t, v)

	v, ok = inputBag.Get("version.something")
	assert.False(t, ok)
	assert.Nil(t, v)
}

func TestInputBag_Has(t *testing.T) {
	ok := inputBag.Has("glossary.GlossDiv.GlossList.1.Acronym")
	assert.True(t, ok)

	ok = inputBag.Has("glossary.notExists")
	assert.False(t, ok)

	ok = inputBag.Has("version")
	assert.True(t, ok)

	ok = inputBag.Has("notExists")
	assert.False(t, ok)

	ok = inputBag.Has("glossary.GlossDiv.GlossList.2.Acronym")
	assert.False(t, ok)

	ok = inputBag.Has("version.something")
	assert.False(t, ok)
}

func TestNewInputBagFromStruct(t *testing.T) {
	data := sampleStruct{
		UserName: "username",
		Email:    "email@example.com",
		Age:      30,
		Skills:   []string{"golang", "software engineering"},
		Experiences: []struct {
			JobTitle string    `json:"jobTitle"`
			Salary   int       `json:"salary"`
			From     time.Time `json:"from"`
			To       time.Time `json:"to"`
		}{
			{
				JobTitle: "Full Stack Developer",
				Salary:   50000,
				From:     time.Date(2018, 11, 1, 0, 0, 0, 0, time.UTC),
				To:       time.Date(2019, 8, 12, 0, 0, 0, 0, time.UTC),
			},
			{
				JobTitle: "Backend Developer",
				Salary:   58000,
				From:     time.Date(2019, 8, 12, 0, 0, 0, 0, time.UTC),
				To:       time.Date(2020, 12, 28, 0, 0, 0, 0, time.UTC),
			},
			{
				JobTitle: "Backend Developer",
				Salary:   68000,
				From:     time.Date(2021, 1, 10, 0, 0, 0, 0, time.UTC),
				To:       time.Date(2022, 8, 23, 0, 0, 0, 0, time.UTC),
			},
		},
		Socials: struct {
			Twitter  string `json:"twitterHandle"`
			Facebook string `json:"facebookUsername"`
		}{
			Twitter:  "johnDoe123",
			Facebook: "johnDoe123",
		},
	}

	result := NewInputBagFromStruct(data)
	v, _ := result.Get("username")
	assert.Equal(t, "username", v)

	v, _ = result.Get("skills.0")
	assert.Equal(t, "golang", v)

	v, _ = result.Get("experiences.1.salary")
	assert.EqualValues(t, 58000, v)

	v, _ = result.Get("socials.twitterHandle")
	assert.EqualValues(t, "johnDoe123", v)
}
