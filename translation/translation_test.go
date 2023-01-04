package translation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetDefaultTranslatorFunc(t *testing.T) {
	defaultTranslationFunc = nil
	SetDefaultTranslatorFunc(func(local, key string, params ...map[string]string) string {
		return "it works!"
	})

	assert.Equal(t, "it works!", defaultTranslationFunc("", ""))
	defaultTranslationFunc = nil
}

func TestGetDefaultTranslatorFunc(t *testing.T) {
	assert.Nil(t, GetDefaultTranslatorFunc())

	SetDefaultTranslatorFunc(func(local, key string, params ...map[string]string) string {
		return "it works!"
	})

	assert.Equal(t, "it works!", GetDefaultTranslatorFunc()("", ""))
}

func TestBaseTranslatableRule_AddLocale(t *testing.T) {
	r := &BaseTranslatableRule{}

	r.AddLocale("en")
	assert.Equal(t, "en", r.Locale)
}

func TestBaseTranslatableRule_AddTranslationFunction(t *testing.T) {
	r := &BaseTranslatableRule{}
	r.AddTranslationFunction(func(local, key string, params ...map[string]string) string {
		return "it works!"
	})

	assert.Equal(t, "it works!", r.Translate("", ""))
}
