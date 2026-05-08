# Go Validator
![Coverage](https://img.shields.io/badge/Coverage-98.0%25-brightgreen)
[![Go Report Card](https://goreportcard.com/badge/github.com/behzadsh/go.validator)](https://goreportcard.com/report/github.com/behzadsh/go.validator)

The `go.validator` package provides a simple and convenient way to validate your data.

## Installation

To install the `go.validator` package, run the following command:

```bash
go get -u github.com/behzadsh/go.validator
```

## How to use
You can use this package to validate HTTP request data in any form.

### Validating map data

To validate map data, use the `ValidateMap` function. `ValidateMap` accepts two parameters: the input data
and the validation rules map.

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"

	validation "github.com/behzadsh/go.validator"
)

func main() {
    http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
        var body map[string]any
        decoder := json.NewDecoder(r.Body)
        err := decoder.Decode(&body)
        if err != nil {
            log.Fatal(err)
        }

        res := validation.ValidateMap(body, validation.RulesMap{
            "email": {"required", "email:mx"},
            "password": {"required", "minLength:6"},
            "birthDate": {"dateTime"},
        })
        
        if res.Failed() {
            _ = json.NewEncoder(w).Encode(map[string]any{
                "message": "validation failed",
                "errors": res.Errors.All(), // it will return a map[string][]string, key is the field name and the slice is the list of the field errors.
            })
        }
    })

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}


```

You can also validate other types of data with the following functions:
* `ValidateMapSlice`: Validate a slice of maps (rules are applied to each map).
* `ValidateStruct`: Validate a struct value, just like map validation.
* `ValidateStructSlice`: Same as `ValidateMapSlice`, but for a slice of structs.

### Validating a single variable

```go
package main

import (
	"fmt"
	"log"
	"os"

	validation "github.com/behzadsh/go.validator"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("an argument is required.")
	}

	res := validation.Validate(args[0], []string{"notEmpty", "integer", "between:3,5"})

	if res.Failed() {
		fmt.Println("Validation failed!")
        // the validation errors are stored in res.Errors under the `variable` key when using `validation.Validate()`.
		for _, value := range res.Errors["variable"] {
            fmt.Println(value)
		}
		os.Exit(1)
	}
}
```

## Translation and Localization

The `go.validator` package supports internationalization (i18n) for validation error messages. You can customize translations and locales in several ways.

### Setting Default Locale

You can set a default locale for all validation operations:

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"

    validation "github.com/behzadsh/go.validator"
)

func main() {
    // Set default locale to Spanish
    validation.SetDefaultLocale("es")
    
    http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
        var body map[string]any
        decoder := json.NewDecoder(r.Body)
        err := decoder.Decode(&body)
        if err != nil {
            log.Fatal(err)
        }

        // This will use Spanish locale by default
        res := validation.ValidateMap(body, validation.RulesMap{
            "email": {"required", "email"},
            "password": {"required", "minLength:6"},
        })
        
        if res.Failed() {
            _ = json.NewEncoder(w).Encode(map[string]any{
                "message": "validation failed",
                "errors": res.Errors.All(),
            })
        }
    })

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

### Per-Validation Locale

You can also specify a locale for individual validation calls:

```go
// Validate with French locale
res := validation.ValidateMap(data, rules, "fr")

// Validate struct with German locale  
res := validation.ValidateStruct(user, rules, "de")

// Validate single variable with Italian locale
res := validation.Validate(value, []string{"required", "email"}, "it")
```

### Custom Translation Function

You can provide your own translation function to handle custom translation logic:

```go
package main

import (
    "fmt"
    "strings"

    validation "github.com/behzadsh/go.validator"
    "github.com/behzadsh/go.validator/translation"
)

func main() {
    // Set custom translation function
    translation.SetDefaultTranslatorFunc(func(locale, key string, params ...map[string]string) string {
        var p map[string]string
        if len(params) > 0 {
            p = params[0]
        }

        // Custom translation logic
        switch key {
        case "validation.required":
            switch locale {
            case "es":
                return "El campo :field: es obligatorio."
            case "fr":
                return "Le champ :field: est requis."
            default:
                return "The :field: field is required."
            }
        case "validation.email":
            switch locale {
            case "es":
                return "El campo :field: debe ser una dirección de correo válida."
            case "fr":
                return "Le champ :field: doit être une adresse e-mail valide."
            default:
                return "The :field: field must be a valid email address."
            }
        default:
            return key
        }
    })

    // Set default locale
    validation.SetDefaultLocale("es")

    // Now all validations will use Spanish translations
    res := validation.ValidateMap(map[string]any{
        "email": "",
    }, validation.RulesMap{
        "email": {"required", "email"},
    })

    if res.Failed() {
        fmt.Println(res.Errors.All())
        // Output: map[email:[El campo email es obligatorio.]]
    }
}
```

> For an even easier translation experience, we recommend using the [`go.localization`](https://github.com/behzadsh/go.localization) package.

### Translation Key Format

Validation error messages use translation keys in the format `validation.{ruleName}`. The translation function receives:

- `locale`: The current locale (e.g., "en", "es", "fr")
- `key`: The translation key (e.g., "validation.required", "validation.email")
- `params`: Optional parameters map containing field names and other placeholders

Common placeholders in translation messages:
- `:field:`: The field name being validated
- `:value:`: The actual value being validated
- `:min:`: Minimum value (for rules like `min`, `minLength`)
- `:max:`: Maximum value (for rules like `max`, `maxLength`)

## Custom Rules

You can define your own validation rules and register them with the validator.

### Rule building blocks

- **Rule interface**: implement `Validate(selector string, value any, inputBag bag.InputBag) rules.ValidationResult`.
- **RuleWithParams (optional)**: implement `AddParams(params []string)` and `MinRequiredParams() int` to receive rule parameters (e.g., `between:3,5`).
- **TranslatableRule (optional but recommended)**: embed `translation.BaseTranslatableRule` to get the current locale and a translation function injected automatically.

### Minimal rule (no params, no translations)

```go
package rules

import (
    "strings"
    "github.com/behzadsh/go.validator/bag"
)

// Palindrome checks if the value is a palindrome string.
type Palindrome struct{}

func (r *Palindrome) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
    s, ok := value.(string)
    if !ok {
        return NewFailedResult(":field: must be a string")
    }
    t := strings.ToLower(strings.ReplaceAll(s, " ", ""))
    i, j := 0, len(t)-1
    for i < j {
        if t[i] != t[j] {
            return NewFailedResult(":field: must be a palindrome")
        }
        i++; j--
    }
    return NewSuccessResult()
}
```

### Rule with parameters

```go
package rules

import (
    "github.com/behzadsh/go.validator/bag"
    "github.com/spf13/cast"
)

// MinWords:value — ensures a string has at least :value words.
type MinWords struct {
    min int
}

func (r *MinWords) AddParams(params []string) {
    r.min = cast.ToInt(params[0])
}

func (*MinWords) MinRequiredParams() int { return 1 }

func (r *MinWords) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
    s, ok := value.(string)
    if !ok {
        return NewFailedResult(":field: must be a string")
    }
    words := 0
    inWord := false
    for _, ch := range s {
        if ch == ' ' || ch == '\n' || ch == '\t' { inWord = false; continue }
        if !inWord { words++; inWord = true }
    }
    if words < r.min {
        return NewFailedResult(":field: must have at least :min: words")
    }
    return NewSuccessResult()
}
```

### Adding translations

Embed `translation.BaseTranslatableRule` and use its `Translate` function. The validator injects the current locale and translator before calling `Validate`.

```go
package rules

import (
    "github.com/behzadsh/go.validator/bag"
    "github.com/behzadsh/go.validator/translation"
)

// MaxWords:value — ensures a string has at most :value words.
type MaxWords struct {
    translation.BaseTranslatableRule
    max int
}

func (r *MaxWords) AddParams(params []string) { r.max = cast.ToInt(params[0]) }
func (*MaxWords) MinRequiredParams() int { return 1 }

func (r *MaxWords) Validate(selector string, value any, _ bag.InputBag) ValidationResult {
    s, ok := value.(string)
    if !ok {
        return NewFailedResult(r.Translate(r.Locale, "validation.string", map[string]string{"field": selector}))
    }
    // ... count words into n ...
    n := len(strings.Fields(s))
    if n > r.max {
        return NewFailedResult(r.Translate(r.Locale, "validation.maxWords", map[string]string{
            "field": selector,
            "max":   cast.ToString(r.max),
        }))
    }
    return NewSuccessResult()
}
```

You can provide translations by setting a custom translator (see "Translation and Localization" above) and handling keys like `validation.maxWords`.

### Registering and using your rule

```go
package main

import (
    validation "github.com/behzadsh/go.validator"
    "github.com/behzadsh/go.validator/rules"
)

func init() {
    // Register custom rules once (e.g., app startup). Re-registering a name overrides the previous rule.
    validation.Register("palindrome", &rules.Palindrome{})
    validation.Register("minWords", &rules.MinWords{})
}

func main() {
    data := map[string]any{"title": "able was I ere I saw Elba"}
    res := validation.ValidateMap(data, validation.RulesMap{
        "title": {"required", "palindrome", "minWords:2"},
    })
    if res.Failed() { /* handle errors */ }
}
```

### Notes

- If you declare `MinRequiredParams() > 0` and the user provides fewer parameters, validation will panic with a clear error.
- If a rule name is not registered, validation will panic. Register custom rules before use.
- Embedding `translation.BaseTranslatableRule` is optional but recommended for localized messages.

## Validation Options

### Stop on first failure

By default, the validator evaluates all rules for each field and accumulates all errors. You can change this behavior to stop evaluating further rules for a field after the first failure.

```go
package main

import (
    validation "github.com/behzadsh/go.validator"
)

func init() {
    // Enable once at app startup
    validation.StopOnFirstFailure()
}

func main() {
    data := map[string]any{"age": "abc"}
    // Only the first failing rule for each field is reported
    res := validation.ValidateMap(data, validation.RulesMap{
        "age": {"required", "integer", "between:18,65"},
    })
    _ = res
}
```

Notes:
- This option stops rule evaluation per-field only; other fields are still validated.
- The default is disabled. Calling `StopOnFirstFailure()` enables it globally for subsequent validations.

## Available Rules

The complete list of available rules can be found [here](https://github.com/behzadsh/go.validator/tree/main/rules.md).
