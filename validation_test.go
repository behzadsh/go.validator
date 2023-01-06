package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func resetConfigs() {
	registerDefaultRules()

	defaultLocale = "en"
	stopOnFirstFailure = false
}

func TestValidateMap(t *testing.T) {
	resetConfigs()
	t.Run("success", func(t *testing.T) {
		data := map[string]any{
			"email":    "user@example.com",
			"password": "mySecurePassword",
		}

		res := ValidateMap(data, RulesMap{
			"email":    {"required", "email"},
			"password": {"required", "string"},
		}, "en")

		assert.False(t, res.Failed())
		assert.Empty(t, res.Errors)
	})

	t.Run("failed", func(t *testing.T) {
		data := map[string]any{
			"email": "invalidEmail",
		}

		res := ValidateMap(data, RulesMap{
			"email":    {"required", "email"},
			"password": {"required", "string"},
		})

		assert.True(t, res.Failed())
		assert.NotEmpty(t, res.Errors)
		assert.True(t, res.Errors.Has("email"))
		assert.Len(t, res.Errors["email"], 1)
		assert.Equal(t, "validation.email", res.Errors["email"][0])
		assert.True(t, res.Errors.Has("password"))
		assert.Len(t, res.Errors["password"], 1)
		assert.Equal(t, "validation.required", res.Errors["password"][0])
	})
}

func TestValidateMapSlice(t *testing.T) {
	resetConfigs()
	t.Run("success", func(t *testing.T) {
		data := []map[string]any{
			{
				"email":    "user@example.com",
				"password": "mySecurePassword",
			},
			{
				"email":    "user2@example.com",
				"password": "mySecurePa$$word",
			},
		}

		res := ValidateMapSlice(data, RulesMap{
			"email":    {"required", "email"},
			"password": {"required", "string"},
		}, "en")

		assert.False(t, res.Failed())
		assert.Empty(t, res.Errors)
	})

	t.Run("failed", func(t *testing.T) {
		data := []map[string]any{
			{
				"email": "invalidEmail",
			},
			{
				"password": false,
			},
		}

		res := ValidateMapSlice(data, RulesMap{
			"email":    {"required", "email"},
			"password": {"required", "string"},
		})

		assert.True(t, res.Failed())
		assert.NotEmpty(t, res.Errors)
		assert.True(t, res.Errors.Has("0.email"))
		assert.Len(t, res.Errors["0.email"], 1)
		assert.Equal(t, "validation.email", res.Errors["0.email"][0])
		assert.True(t, res.Errors.Has("0.password"))
		assert.Len(t, res.Errors["0.password"], 1)
		assert.Equal(t, "validation.required", res.Errors["0.password"][0])
		assert.NotEmpty(t, res.Errors)
		assert.True(t, res.Errors.Has("1.email"))
		assert.Len(t, res.Errors["1.email"], 2)
		assert.Equal(t, "validation.required", res.Errors["1.email"][0])
		assert.Equal(t, "validation.email", res.Errors["1.email"][1])
		assert.True(t, res.Errors.Has("1.password"))
		assert.Len(t, res.Errors["1.password"], 1)
		assert.Equal(t, "validation.string", res.Errors["1.password"][0])
	})
}

func TestValidateStruct(t *testing.T) {
	resetConfigs()
	t.Run("success", func(t *testing.T) {
		data := struct {
			Title string `json:"title"`
			Desc  string `json:"description"`
		}{
			Title: "My Title",
			Desc:  "This is a sample description for testing a struct data.",
		}

		res := ValidateStruct(data, RulesMap{
			"title":       {"required", "string", "maxLength:50"},
			"description": {"required", "string", "minLength:50"},
		}, "en")

		assert.False(t, res.Failed())
		assert.Empty(t, res.Errors)
	})

	t.Run("failed", func(t *testing.T) {
		data := struct {
			Title string `json:"title"`
			Desc  string `json:"description"`
		}{
			Title: "This is a sample description for testing a struct data.",
			Desc:  "My Title",
		}

		res := ValidateStruct(data, RulesMap{
			"title":       {"required", "string", "maxLength:50"},
			"description": {"required", "string", "minLength:50"},
		}, "en")

		assert.True(t, res.Failed())
		assert.NotEmpty(t, res.Errors)
		assert.True(t, res.Errors.Has("title"))
		assert.Len(t, res.Errors["title"], 1)
		assert.Equal(t, "validation.max_length", res.Errors["title"][0])
		assert.True(t, res.Errors.Has("description"))
		assert.Len(t, res.Errors["description"], 1)
		assert.Equal(t, "validation.min_length", res.Errors["description"][0])
	})

	t.Run("panic", func(t *testing.T) {
		data := map[string]any{
			"title":       "My Title",
			"description": "This is a sample description for testing a struct data.",
		}

		assert.Panics(t, func() {
			ValidateStruct(data, RulesMap{
				"title":       {"required", "string", "maxLength:50"},
				"description": {"required", "string", "minLength:50"},
			})
		})
	})
}

func TestValidateStructSlice(t *testing.T) {
	resetConfigs()
	t.Run("success", func(t *testing.T) {
		data := []any{
			struct {
				Title string `json:"title"`
				Desc  string `json:"description"`
			}{
				Title: "My Title",
				Desc:  "This is a sample description for testing a struct data.",
			},
			struct {
				Title string `json:"title"`
				Desc  string `json:"description"`
			}{
				Title: "Another Title",
				Desc:  "This is another sample description for testing a struct data.",
			},
		}

		res := ValidateStructSlice(data, RulesMap{
			"title":       {"required", "string", "maxLength:50"},
			"description": {"required", "string", "minLength:50"},
		}, "en")

		assert.False(t, res.Failed())
		assert.Empty(t, res.Errors)
	})

	t.Run("failed", func(t *testing.T) {
		data := []any{
			struct {
				Title string `json:"title"`
				Desc  string `json:"description"`
			}{
				Title: "This is a sample description for testing a struct data.",
				Desc:  "My Title",
			},
			struct {
				Title string `json:"title"`
				Desc  string `json:"description"`
			}{
				Title: "My Title",
			},
		}

		res := ValidateStructSlice(data, RulesMap{
			"title":       {"notEmpty", "string", "maxLength:50"},
			"description": {"notEmpty", "string", "minLength:50"},
		}, "en")

		assert.True(t, res.Failed())
		assert.NotEmpty(t, res.Errors)
		assert.True(t, res.Errors.Has("0.title"))
		assert.Len(t, res.Errors["0.title"], 1)
		assert.Equal(t, "validation.max_length", res.Errors["0.title"][0])
		assert.True(t, res.Errors.Has("0.description"))
		assert.Len(t, res.Errors["0.description"], 1)
		assert.Equal(t, "validation.min_length", res.Errors["0.description"][0])
		assert.False(t, res.Errors.Has("1.title"))
		assert.True(t, res.Errors.Has("1.description"))
		assert.Len(t, res.Errors["1.description"], 2)
		assert.Equal(t, "validation.not_empty", res.Errors["1.description"][0])
		assert.Equal(t, "validation.min_length", res.Errors["1.description"][1])
	})

	t.Run("panic", func(t *testing.T) {
		data := map[string]any{
			"title":       "My Title",
			"description": "This is a sample description for testing a struct data.",
		}

		assert.Panics(t, func() {
			ValidateStruct(data, RulesMap{
				"title":       {"required", "string", "maxLength:50"},
				"description": {"required", "string", "minLength:50"},
			})
		})
	})
}

func TestValidate(t *testing.T) {
	resetConfigs()
	t.Run("success", func(t *testing.T) {
		data := "fa7a0f59-7c7c-4c51-b2d5-32964a090eac"

		res := Validate(data, []string{"uuid"})

		assert.False(t, res.Failed())
		assert.Empty(t, res.Errors)
	})

	t.Run("failed", func(t *testing.T) {
		data := 10

		res := Validate(data, []string{"string", "uuid"}, "en")

		assert.True(t, res.Failed())
		assert.NotEmpty(t, res.Errors)
		assert.True(t, res.Errors.Has("variable"))
		assert.Equal(t, "validation.string", res.Errors["variable"][0])
		assert.Equal(t, "validation.uuid", res.Errors["variable"][1])
	})
}

func TestSpecialValidation(t *testing.T) {
	resetConfigs()
	t.Run("nestedValidation", func(t *testing.T) {
		data := map[string]any{
			"map": map[string]any{
				"field1": "value1",
				"field2": "value2",
			},
			"array": []any{"val1", "val2"},
			"mapArray": []map[string]any{
				{
					"field1": "value1",
					"field2": "value2",
				},
				{
					"field1": "value3",
					"field2": "value4",
				},
			},
		}

		res := ValidateMap(data, RulesMap{
			"map.field1":        {"required"},
			"array.*":           {"string"},
			"mapArray.*.field2": {"required"},
		})

		assert.False(t, res.Failed())
		assert.Empty(t, res.Errors)
	})

	t.Run("stopOnFailure", func(t *testing.T) {
		StopOnFirstFailure()
		data := map[string]any{
			"password": "mySecurePassword",
		}

		res := ValidateMap(data, RulesMap{
			"email":    {"required", "email"},
			"password": {"required", "string"},
		})

		assert.True(t, res.Failed())
		assert.NotEmpty(t, res.Errors)
		assert.True(t, res.Errors.Has("email"))
		assert.Len(t, res.Errors["email"], 1)
		assert.Equal(t, "validation.required", res.Errors["email"][0])
	})

	t.Run("insufficientParam", func(t *testing.T) {
		resetConfigs()
		data := map[string]any{
			"code": 13839,
		}

		assert.Panics(t, func() {
			ValidateMap(data, RulesMap{
				"code": {"digitsBetween:5"},
			})
		})
	})

	t.Run("unregisteredRule", func(t *testing.T) {
		resetConfigs()
		data := map[string]any{
			"code": 13839,
		}

		assert.Panics(t, func() {
			ValidateMap(data, RulesMap{
				"code": {"unregistered:10"},
			})
		})
	})
}
