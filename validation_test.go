package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateMap(t *testing.T) {
	registerDefaultRules()
	t.Run("success", func(t *testing.T) {
		data := map[string]any{
			"email":    "user@example.com",
			"password": "mySecurePassword",
		}

		res := ValidateMap(data, RulesMap{
			"email":    {"required", "email"},
			"password": {"required", "string"},
		})

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
