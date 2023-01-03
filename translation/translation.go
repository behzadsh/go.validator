package translation

import "github.com/behzadsh/go.localization"

// TranslateFunc is a function type that translate the given key in given locale.
type TranslateFunc func(local, key string, params ...map[string]string) string

var defaultTranslationFunc TranslateFunc

func init() {
	translator, _ := lang.NewTranslator(lang.DefaultConfigs())
	defaultTranslationFunc = translator.TranslateBy
}

// SetDefaultTranslatorFunc sets the default translation function.
func SetDefaultTranslatorFunc(fn TranslateFunc) {
	defaultTranslationFunc = fn
}

// GetDefaultTranslatorFunc returns the default translation function.
func GetDefaultTranslatorFunc() TranslateFunc {
	return defaultTranslationFunc
}

// TranslatableRule is an interface for rules that translate the validation
// error messages.
type TranslatableRule interface {
	AddLocale(locale string)
	AddTranslationFunction(fn TranslateFunc)
}

// BaseTranslatableRule is a semi-abstract struct that do the basic translation
// functionality.
type BaseTranslatableRule struct {
	Translate TranslateFunc
	Locale    string
}

// AddLocale adds the default locale for translation.
func (b *BaseTranslatableRule) AddLocale(locale string) {
	b.Locale = locale
}

// AddTranslationFunction adds the translation function to the base.
func (b *BaseTranslatableRule) AddTranslationFunction(fn TranslateFunc) {
	b.Translate = fn
}
