package translation

import "github.com/behzadsh/go.localization"

type Translator interface {
	Translate(locale, key string, params ...map[string]string)
}

type TranslateFunc func(local, key string, params ...map[string]string) string

var defaultTranslationFunc TranslateFunc

func init() {
	translator, _ := lang.NewTranslator(lang.DefaultConfigs())
	defaultTranslationFunc = translator.TranslateBy
}

func SetDefaultTranslatorFunc(fn TranslateFunc) {
	defaultTranslationFunc = fn
}

func GetDefaultTranslatorFunc() TranslateFunc {
	return defaultTranslationFunc
}

type TranslatableRule interface {
	AddLocale(locale string)
	AddTranslationFunction(fn TranslateFunc)
}

type BaseTranslatableRule struct {
	Translate TranslateFunc
	Locale    string
}

func (b *BaseTranslatableRule) AddLocale(locale string) {
	b.Locale = locale
}

func (b *BaseTranslatableRule) AddTranslationFunction(fn TranslateFunc) {
	b.Translate = fn
}
