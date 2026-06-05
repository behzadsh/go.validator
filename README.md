# go.validator

A small, schema-based validator for Go. No globals, no struct tags, no reflection magic in your business code — just a `Schema` you build with chained `Field` calls and validate against any map or struct.

```go
schema := validation.New().
    Field("name", validation.Required, validation.MinLength(2)).
    Field("email", validation.Required, validation.Email)

res, err := schema.Validate(input)
if err != nil {
    log.Fatal(err) // RuleSyntaxError: misconfigured rule, fix at startup
}
if res.HasErrors() {
    for _, e := range res.Errors() {
        fmt.Println(e.Path, e.Message)
    }
}
```

## Install

```bash
go get github.com/behzadsh/go.validator/v2
```

Requires Go 1.21 or later.

## Concepts

Three types are all you need to know:

- **`Rule`** — an interface with one method: `Validate(value any) error`. Rules are values you pass around. Built-in rules are exposed as variables (e.g. `Required`, `Email`) or constructors that return a `Rule` (e.g. `Min(18)`, `MinLength(3)`).
- **`Schema`** — a sequence of (path, rules) pairs. Build it with `New()` and chain `Field` calls.
- **`Result`** — the value returned by `Schema.Validate`. Inspect it via `HasErrors()`, `Errors()`, and `For(path)`. `Result` does **not** implement the `error` interface; iterate `res.Errors()` to handle individual failures. Each entry is a `FieldError{Path, Err, Message, Code, Params}` and *does* implement `error`.

### Absence model

Only the `Required*` and `NotEmpty` rules fail when a field is absent or empty. Every other built-in rule returns `nil` for a missing value. Combine `Required` with another rule to enforce both presence and shape:

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

res, _ := schema.Validate(input)
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

res, _ := schema.Validate(User{Name: "Alice", Email: "alice@example.com", Age: 29})
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

See **[RULES.md](RULES.md)** for the complete rule reference with signatures, fail conditions, and examples.

Quick overview by category:

| Category | Rules |
|---|---|
| General | `Required`, `RequiredIf`, `RequiredUnless`, `RequiredWith`, `RequiredWithAll`, `RequiredWithout`, `RequiredWithoutAll`, `NotEmpty` |
| String | `Alpha`, `AlphaDash`, `AlphaNum`, `AlphaSpace`, `ASCII`, `Base64`, `Contains`, `CreditCard`, `Email`, `EmailMX`, `EndsWith`, `HexColor`, `JSON`, `JWT`, `Length`, `Lowercase`, `MaxLength`, `MinLength`, `NotRegex`, `PhoneE164`, `Regex`, `Semver`, `Slug`, `StartsWith`, `Uppercase`, `URL`, `UUID` |
| Number | `Numeric`, `Integer`, `Min`, `Max`, `GT`, `GTE`, `LT`, `LTE`, `Between`, `Positive`, `Negative`, `NonNegative`, `MultipleOf`, `Port`, `Latitude`, `Longitude` |
| Digit | `Digits`, `MinDigits`, `MaxDigits`, `DigitsBetween` |
| DateTime | `DateTime`, `DateTimeFormat`, `After`, `AfterOrEqual`, `AfterField`, `Before`, `BeforeOrEqual`, `BeforeField`, `DateTimeBetween`, `Timezone` |
| Network | `IP`, `IPv4`, `IPv6`, `CIDR`, `MACAddress` |
| Collection | `Distinct`, `Each`, `Size`, `MinSize`, `MaxSize` |
| Generic | `In`, `NotIn`, `NEQ` |
| Comparison | `SameAs`, `Different` |
| Logical | `Any`, `Not`, `When`, `Unless` |

Every rule except `Required`, `RequiredIf`, `RequiredUnless`, `RequiredWith*`, and `NotEmpty` returns `nil` for a missing value.

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
mustMatchField := func(otherPath string) validation.InputRule {
    return validation.InputRuleFunc(func(value any, input *validation.InputBag) error {
        other, _ := input.Lookup(otherPath)
        if value != other {
            return errors.New("values do not match")
        }
        return nil
    })
}

schema := validation.New().
    Field("password_confirm", validation.Required, mustMatchField("password"))
```

The built-in `SameAs` and `Different` rules cover the most common cross-field comparison patterns.

Rules are values: build them once at startup and reuse across validations and goroutines.

## Error handling

```go
res, err := schema.Validate(input)
if err != nil {
    log.Fatal(err) // RuleSyntaxError: misconfigured rule, fix at startup
}

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

`FieldError` carries `Code` (a stable snake_case key for i18n) and `Params` (rule parameters). Use `Code` to branch on specific failures:

```go
for _, e := range res.Errors() {
    switch e.Code {
    case "email":
        // handle invalid email
    case "min_length":
        // e.Params["length"] holds the minimum
    }
}
```

`FieldError` also implements `error` and exposes the underlying rule error via `Err`, so `errors.As` works for custom rules that return structured errors.

`Result` deliberately does **not** implement `error`. Decide at the API boundary how to surface the collection — most apps return the slice from `Errors()` as JSON to the client, or fail-fast on `HasErrors()`.

## Concurrency

A `Schema` is meant to be built once and called many times. Once construction is complete, `Validate` is safe to call from multiple goroutines: it only reads the schema, and rules are immutable values.

## What this version does not include

- An `And` combinator (unnecessary — multiple rules on a `Field` call are implicitly AND).
- Slice or wildcard paths (`items.*.name`).
- Internationalization. Error messages are plain English strings; use `Code` and `Params` on `FieldError` to build translated messages.

## License

See `LICENSE`.
