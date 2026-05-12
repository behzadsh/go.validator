# go.validator

A small, schema-based validator for Go. No globals, no struct tags, no reflection magic in your business code — just a `Schema` you build with chained `Field` calls and validate against any map or struct.

```go
schema := validation.New().
    Field("name", validation.Required, validation.MinLength(2)).
    Field("email", validation.Required, validation.Email)

if res := schema.Validate(input); res.HasErrors() {
    for _, e := range res.Errors() {
        fmt.Println(e.Path, e.Message)
    }
}
```

## Install

```bash
go get github.com/behzadsh/go.validator/v2
```

Requires Go 1.21 or later. Zero third-party dependencies.

## Concepts

Three types are all you need to know:

- **`Rule`** — an interface with one method: `Validate(value any) error`. Rules are values you pass around. Built-in rules are exposed as variables (e.g. `Required`, `Email`) or constructors that return a `Rule` (e.g. `Min(18)`, `MinLength(3)`).
- **`Schema`** — a sequence of (path, rules) pairs. Build it with `New()` and chain `Field` calls.
- **`Result`** — the value returned by `Schema.Validate`. Inspect it via `HasErrors()`, `Errors()`, and `For(path)`. `Result` does **not** implement the `error` interface; iterate `res.Errors()` to handle individual failures. Each entry is a `FieldError{Path, Err, Message}` and *does* implement `error`.

### Absence model

Only `Required` and `RequiredIf` fail when a field is absent or empty. Every other built-in rule returns `nil` for a missing value. Combine `Required` with another rule to enforce both presence and shape:

```go
.Field("email", validation.Required, validation.Email)
```

Without `Required`, a missing `email` is accepted; if present, it must be a valid email.

## Validating a map

```go
input := map[string]any{
    "name":  "Alice",
    "email": "alice@example.com",
    "age":   29,
}

schema := validation.New().
    Field("name", validation.Required, validation.MinLength(2)).
    Field("email", validation.Required, validation.Email).
    Field("age", validation.Min[int](18), validation.Max[int](120))

res := schema.Validate(input)
```

Nested keys use dot-notation:

```go
schema := validation.New().
    Field("profile.handle", validation.Required, validation.AlphaNum)

input := map[string]any{
    "profile": map[string]any{"handle": "alice99"},
}
```

## Validating a struct

The same schema works against a struct or `*struct`. Field names in the path resolve in this order:

1. The first comma-segment of the `json` struct tag, if present and not `"-"`.
2. The exported Go field name.

A `json:"-"` tag hides the field from the validator. Embedded (anonymous) struct fields are searched recursively.

```go
type User struct {
    Name    string  `json:"name"`
    Email   string  `json:"email"`
    Age     int     `json:"age"`
    Profile Profile `json:"profile"`
}

type Profile struct {
    Handle string `json:"handle"`
}

schema := validation.New().
    Field("name", validation.Required, validation.MinLength(2)).
    Field("email", validation.Required, validation.Email).
    Field("age", validation.Min[int](18)).
    Field("profile.handle", validation.Required)

res := schema.Validate(User{Name: "Alice", Email: "alice@example.com", Age: 29})
```

A field without a `json` tag is reachable by its Go name:

```go
type Item struct {
    SKU string
}

validation.New().Field("SKU", validation.Required).Validate(Item{})
```

A `nil` `*struct` is treated as if every field were absent. Required fields will fail; other rules will pass.

## Built-in rules

### General

| Rule | Description |
|---|---|
| `Required` | Fails when value is `nil` or `""`. |
| `RequiredIf(condition)` | Fails when value is `nil` or `""` and the condition expression evaluates to true. |
| `NotEmpty` | Fails when value is `nil` or the zero value for its type (`0`, `false`, `""`, etc.). |

### String

| Rule | Description |
|---|---|
| `Alpha` | Value must be a string containing only Unicode letters. |
| `AlphaDash` | Value must be a string containing only Unicode letters, digits, underscores, and dashes. |
| `AlphaNum` | Value must be a string containing only Unicode letters and digits. |
| `AlphaSpace` | Value must be a string containing only Unicode letters and whitespace. |
| `Email` | Value must be a well-formed email address. No DNS lookup. |
| `EmailMX` | Value must be a well-formed email address whose domain has at least one MX record. |
| `URL` | Value must parse as a valid absolute URL; scheme-less URLs are accepted as http. |
| `Length(n)` | String rune count must be exactly `n`. |
| `MinLength(n)` | String rune count must be at least `n`. |
| `MaxLength(n)` | String rune count must be at most `n`. |

### Numeric

| Rule | Description |
|---|---|
| `Numeric` | Value must be any numeric type or a string that parses as a float64. |
| `Min[T](n)` | Value must be at least `n`. |
| `Max[T](n)` | Value must be at most `n`. |
| `Between[T](min, max)` | Value must be between `min` and `max` inclusive. |

### Date / time

| Rule | Description |
|---|---|
| `DateTimeFormat(layout)` | Value must be a string matching the given Go time layout. |
| `After(t)` | Value must be a date/time string strictly after `t`. |
| `Before(t)` | Value must be a date/time string strictly before `t`. |

`After` and `Before` accept a broad set of common date/time formats (RFC3339, `2006-01-02`, RFC1123, and many more) without requiring a layout to be specified.

### Generic

| Rule | Description |
|---|---|
| `In[T](slice)` | Value must equal one of the given slice elements. |
| `NotIn[T](slice)` | Value must not equal any of the given slice elements. |

Every rule except `Required`, `RequiredIf`, and `NotEmpty` returns `nil` for a missing value.

## RequiredIf

`RequiredIf` accepts a small expression language for cross-field conditions:

- **Comparisons:** `field == value`, `field != value`, `field < value`, `field > value`, `field <= value`, `field >= value`
- **Logical:** `expr && expr`, `expr || expr`, `!expr`
- **Grouping:** `(expr)`
- **Functions:** `exists(path)`, `len(path) == n`

String literals must be quoted (`"admin"` or `'admin'`). Unquoted identifiers are looked up as field paths.

```go
schema := validation.New().
    Field("billing_address", validation.RequiredIf(`plan == "paid"`)).
    Field("company_name",    validation.RequiredIf(`role == "business" && exists(vat_number)`)).
    Field("note",            validation.RequiredIf(`(status == "active" || status == "pending") && verified == true`))
```

## Custom rules

Any value that implements `Rule` is acceptable. The fastest path is `RuleFunc`:

```go
isCorporate := validation.RuleFunc(func(v any) error {
    if v == nil {
        return nil
    }
    s, ok := v.(string)
    if !ok || !strings.HasSuffix(s, "@acme.corp") {
        return errors.New("must be a corporate email")
    }
    return nil
})

schema := validation.New().
    Field("email", validation.Required, validation.Email, isCorporate)
```

For rules that need to read other fields, implement `InputRule` or use `InputRuleFunc`:

```go
sameAs := func(otherPath string) validation.InputRule {
    return validation.InputRuleFunc(func(value any, input *validation.InputBag) error {
        other, _ := input.Lookup(otherPath)
        if value != other {
            return errors.New("values do not match")
        }
        return nil
    })
}

schema := validation.New().
    Field("password_confirm", validation.Required, sameAs("password"))
```

Rules are values: build them once at startup and reuse across validations and goroutines.

## Error handling

```go
res := schema.Validate(input)

if res.HasErrors() {
    for _, e := range res.Errors() {
        log.Printf("%s: %s", e.Path, e.Message)
    }
}

// Get every error for a single path
for _, e := range res.For("email") {
    log.Println(e.Message)
}
```

`FieldError` implements the `error` interface and exposes the underlying rule error via `Err`, so `errors.Is` and `errors.As` work:

```go
for _, e := range res.Errors() {
    if errors.Is(e, validation.ErrValidationEmail) {
        // handle email error
    }
}
```

`Result` deliberately does **not** implement `error`. Decide at the API boundary how to surface the collection — most apps return the slice from `Errors()` as JSON to the client, or fail-fast on `HasErrors()`.

## Concurrency

A `Schema` is meant to be built once and called many times. Once construction is complete, `Validate` is safe to call from multiple goroutines: it only reads the schema, and rules are immutable values.

## What this version does not include

- Composition helpers (`And`, `Or`, `Not`, `Each`).
- Slice or wildcard paths (`items.*.name`).
- Internationalization. Error messages are plain English strings.

## License

See `LICENSE`.
