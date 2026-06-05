# Available Rules

## Index

<details>
<summary>General</summary>

- [Required](#required)
- [RequiredIf](#requiredif)
- [RequiredUnless](#requiredunless)
- [RequiredWith](#requiredwith)
- [RequiredWithAll](#requiredwithall)
- [RequiredWithout](#requiredwithout)
- [RequiredWithoutAll](#requiredwithoutall)
- [NotEmpty](#notempty)

</details>

<details>
<summary>String</summary>

- [Alpha](#alpha)
- [AlphaDash](#alphadash)
- [AlphaNum](#alphanum)
- [AlphaSpace](#alphaspace)
- [Email](#email)
- [EmailMX](#emailmx)
- [EndsWith](#endswith)
- [Length](#length)
- [Lowercase](#lowercase)
- [MaxLength](#maxlength)
- [MinLength](#minlength)
- [NotRegex](#notregex)
- [Regex](#regex)
- [StartsWith](#startswith)
- [Uppercase](#uppercase)
- [URL](#url)
- [UUID](#uuid)

</details>

<details>
<summary>Number</summary>

- [Between](#between)
- [GT](#gt)
- [GTE](#gte)
- [Integer](#integer)
- [LT](#lt)
- [LTE](#lte)
- [Max](#max)
- [Min](#min)
- [Numeric](#numeric)

</details>

<details>
<summary>Digit</summary>

- [Digits](#digits)
- [DigitsBetween](#digitsbetween)
- [MaxDigits](#maxdigits)
- [MinDigits](#mindigits)

</details>

<details>
<summary>DateTime</summary>

- [After](#after)
- [AfterField](#afterfield)
- [AfterOrEqual](#afterorequal)
- [Before](#before)
- [BeforeField](#beforefield)
- [BeforeOrEqual](#beforeorequal)
- [DateTime](#datetime)
- [DateTimeBetween](#datetimebetween)
- [DateTimeFormat](#datetimeformat)
- [Timezone](#timezone)

</details>

<details>
<summary>Network</summary>

- [IP](#ip)
- [IPv4](#ipv4)
- [IPv6](#ipv6)
- [MACAddress](#macaddress)

</details>

<details>
<summary>Collection</summary>

- [Distinct](#distinct)

</details>

<details>
<summary>Generic</summary>

- [In](#in)
- [NEQ](#neq)
- [NotIn](#notin)

</details>

<details>
<summary>Comparison</summary>

- [Different](#different)
- [SameAs](#sameas)

</details>

---

## General

<a id="required"></a>
### Required

```go
var Required Rule
```

Fails if the value is `nil` or an empty string `""`. Passes for all other types including zero values (`0`, `false`).

```go
validation.New().
    Field("name", validation.Required)
```

---

<a id="requiredif"></a>
### RequiredIf

```go
func RequiredIf(condition string) InputRule
```

Fails if the condition evaluates to `true` and the value is `nil` or `""`.

The condition language supports comparisons (`==`, `!=`, `<`, `>`, `<=`, `>=`), logical operators (`&&`, `||`, `!`), grouping `(...)`, and functions `exists(path)` and `len(path)`.

```go
validation.New().
    Field("vat_number", validation.RequiredIf(`plan == "paid"`)).
    Field("admin_code", validation.RequiredIf(`role == "admin" && exists(org_id)`))
```

---

<a id="requiredunless"></a>
### RequiredUnless

```go
func RequiredUnless(condition string) InputRule
```

The logical complement of `RequiredIf`. Fails if the condition evaluates to `false` and the value is `nil` or `""`.

```go
validation.New().
    Field("reason", validation.RequiredUnless(`status == "approved"`))
```

---

<a id="requiredwith"></a>
### RequiredWith

```go
func RequiredWith(fields ...string) InputRule
```

Fails if **any** of the listed fields are present in the input and the value is `nil` or `""`.

```go
validation.New().
    Field("address", validation.RequiredWith("phone", "mobile"))
```

---

<a id="requiredwithall"></a>
### RequiredWithAll

```go
func RequiredWithAll(fields ...string) InputRule
```

Fails if **all** of the listed fields are present in the input and the value is `nil` or `""`.

```go
validation.New().
    Field("full_name", validation.RequiredWithAll("first_name", "last_name"))
```

---

<a id="requiredwithout"></a>
### RequiredWithout

```go
func RequiredWithout(fields ...string) InputRule
```

Fails if **any** of the listed fields are absent from the input and the value is `nil` or `""`.

```go
validation.New().
    Field("username", validation.RequiredWithout("email"))
```

---

<a id="requiredwithoutall"></a>
### RequiredWithoutAll

```go
func RequiredWithoutAll(fields ...string) InputRule
```

Fails if **all** of the listed fields are absent from the input and the value is `nil` or `""`.

```go
validation.New().
    Field("contact", validation.RequiredWithoutAll("email", "phone"))
```

---

<a id="notempty"></a>
### NotEmpty

```go
var NotEmpty Rule
```

Fails if the value is `nil`, `""`, `0` (any numeric type), `false`, or a zero-value struct.

```go
validation.New().
    Field("count", validation.NotEmpty).   // rejects 0
    Field("active", validation.NotEmpty)   // rejects false
```

---

## String

<a id="alpha"></a>
### Alpha

```go
var Alpha Rule
```

Fails if the value is not a string or contains non-letter characters. Unicode letters are accepted.

```go
validation.New().Field("name", validation.Alpha)
// "hello" → pass, "Ünïcödé" → pass, "hello1" → fail
```

---

<a id="alphadash"></a>
### AlphaDash

```go
var AlphaDash Rule
```

Accepts Unicode letters, digits, underscores `_`, and hyphens `-`.

```go
validation.New().Field("slug", validation.AlphaDash)
// "hello-world_123" → pass, "hello world" → fail
```

---

<a id="alphanum"></a>
### AlphaNum

```go
var AlphaNum Rule
```

Accepts Unicode letters and digits only.

```go
validation.New().Field("code", validation.AlphaNum)
// "abc123" → pass, "abc-123" → fail
```

---

<a id="alphaspace"></a>
### AlphaSpace

```go
var AlphaSpace Rule
```

Accepts Unicode letters and whitespace only.

```go
validation.New().Field("display_name", validation.AlphaSpace)
// "Hello World" → pass, "hello123" → fail
```

---

<a id="email"></a>
### Email

```go
var Email Rule
```

Validates RFC-compliant email format (no network lookup).

```go
validation.New().Field("email", validation.Required, validation.Email)
// "user@example.com" → pass, "notanemail" → fail
```

---

<a id="emailmx"></a>
### EmailMX

```go
var EmailMX Rule
```

Validates email format and performs a live DNS MX record lookup. Requires network access.

```go
validation.New().Field("email", validation.Required, validation.EmailMX)
```

---

<a id="endswith"></a>
### EndsWith

```go
func EndsWith(suffix string) Rule
```

Fails if the value is not a string or does not end with `suffix`.

```go
validation.New().Field("filename", validation.EndsWith(".go"))
// "main.go" → pass, "main.js" → fail
```

---

<a id="length"></a>
### Length

```go
func Length(n int) Rule
```

Fails if the value is not a string or its rune count is not exactly `n`.

```go
validation.New().Field("pin", validation.Length(6))
// "123456" → pass, "12345" → fail
```

---

<a id="lowercase"></a>
### Lowercase

```go
var Lowercase Rule
```

Fails if the value is not a string or contains any uppercase characters.

```go
validation.New().Field("username", validation.Lowercase)
// "hello" → pass, "Hello" → fail
```

---

<a id="maxlength"></a>
### MaxLength

```go
func MaxLength(n int) Rule
```

Fails if the value is not a string or its rune count exceeds `n`.

```go
validation.New().Field("bio", validation.MaxLength(500))
```

---

<a id="minlength"></a>
### MinLength

```go
func MinLength(n int) Rule
```

Fails if the value is not a string or its rune count is less than `n`.

```go
validation.New().Field("password", validation.MinLength(8))
```

---

<a id="notregex"></a>
### NotRegex

```go
func NotRegex(pattern string) Rule
```

Fails if the value is not a string or the string matches the pattern. The pattern is compiled at call time; an invalid pattern causes a `RuleSyntaxError`.

```go
validation.New().Field("username", validation.NotRegex(`\s`))
// "nospaces" → pass, "has spaces" → fail
```

---

<a id="regex"></a>
### Regex

```go
func Regex(pattern string) Rule
```

Fails if the value is not a string or does not match the pattern. The pattern is compiled at call time; an invalid pattern causes a `RuleSyntaxError`.

```go
validation.New().Field("postal_code", validation.Regex(`^\d{5}$`))
// "12345" → pass, "1234" → fail
```

---

<a id="startswith"></a>
### StartsWith

```go
func StartsWith(prefix string) Rule
```

Fails if the value is not a string or does not begin with `prefix`.

```go
validation.New().Field("sku", validation.StartsWith("SKU-"))
// "SKU-001" → pass, "001-SKU" → fail
```

---

<a id="uppercase"></a>
### Uppercase

```go
var Uppercase Rule
```

Fails if the value is not a string or contains any lowercase characters.

```go
validation.New().Field("country_code", validation.Uppercase)
// "US" → pass, "us" → fail
```

---

<a id="url"></a>
### URL

```go
var URL Rule
```

Validates an absolute URL (`https://...`, `http://...`) or a scheme-less URL where `http` is inferred. Rejects bare paths without a host.

```go
validation.New().Field("website", validation.URL)
// "https://example.com" → pass, "not a url" → fail
```

---

<a id="uuid"></a>
### UUID

```go
var UUID Rule
```

Validates UUID format `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx` (case-insensitive, any variant).

```go
validation.New().Field("id", validation.UUID)
// "550e8400-e29b-41d4-a716-446655440000" → pass
```

---

## Number

<a id="between"></a>
### Between

```go
func Between[T number](min, max T) Rule
```

Fails if the value cannot be type-asserted to `T` or falls outside `[min, max]` (inclusive).

```go
validation.New().Field("age", validation.Between[int](18, 65))
// 18 → pass, 65 → pass, 17 → fail
```

---

<a id="gt"></a>
### GT

```go
func GT[T number](v T) Rule
```

Fails if the value is not of type `T` or is not strictly greater than `v` (equal fails).

```go
validation.New().Field("score", validation.GT[int](0))
// 1 → pass, 0 → fail
```

---

<a id="gte"></a>
### GTE

```go
func GTE[T number](v T) Rule
```

Fails if the value is not of type `T` or is less than `v`. Equal values pass. Semantically identical to `Min`.

```go
validation.New().Field("age", validation.GTE[int](18))
// 18 → pass, 17 → fail
```

---

<a id="integer"></a>
### Integer

```go
var Integer Rule
```

Passes if the value's Go type is one of the integer kinds (`int`, `int8`…`int64`, `uint`…`uint64`). `nil` passes. `float64` (the default JSON number type) fails.

```go
validation.New().Field("count", validation.Integer)
// int(5) → pass, float64(5.0) → fail
```

---

<a id="lt"></a>
### LT

```go
func LT[T number](v T) Rule
```

Fails if the value is not of type `T` or is not strictly less than `v` (equal fails).

```go
validation.New().Field("quantity", validation.LT[int](100))
// 99 → pass, 100 → fail
```

---

<a id="lte"></a>
### LTE

```go
func LTE[T number](v T) Rule
```

Fails if the value is not of type `T` or is greater than `v`. Equal values pass. Semantically identical to `Max`.

```go
validation.New().Field("rating", validation.LTE[int](5))
// 5 → pass, 6 → fail
```

---

<a id="max"></a>
### Max

```go
func Max[T number](max T) Rule
```

Fails if the value is not of type `T` or exceeds `max` (inclusive upper bound, i.e. `<=`).

```go
validation.New().Field("rating", validation.Max[int](5))
// 5 → pass, 6 → fail
```

---

<a id="min"></a>
### Min

```go
func Min[T number](min T) Rule
```

Fails if the value is not of type `T` or is below `min` (inclusive lower bound, i.e. `>=`).

```go
validation.New().Field("age", validation.Min[int](18))
// 18 → pass, 17 → fail
```

---

<a id="numeric"></a>
### Numeric

```go
var Numeric Rule
```

Accepts any Go numeric type or a string parseable as a `float64`. Rejects booleans, `nil`, and non-numeric strings.

```go
validation.New().Field("price", validation.Numeric)
// 42, 3.14, "99.5" → pass, "abc" → fail
```

---

## Digit

> Digit rules operate on the **string representation** of a value. Non-string inputs fail. Only the characters `0–9` are counted; signs, decimal points, and other characters cause failure.

<a id="digits"></a>
### Digits

```go
func Digits(n int) Rule
```

Fails if the value is not a string consisting of exactly `n` digit characters.

```go
validation.New().Field("pin", validation.Digits(4))
// "1234" → pass, "123" → fail, "-123" → fail
```

---

<a id="digitsbetween"></a>
### DigitsBetween

```go
func DigitsBetween(min, max int) Rule
```

Fails if the digit count is outside `[min, max]` (inclusive).

```go
validation.New().Field("otp", validation.DigitsBetween(4, 8))
// "1234" → pass, "123" → fail
```

---

<a id="maxdigits"></a>
### MaxDigits

```go
func MaxDigits(n int) Rule
```

Fails if the value is not a string of at most `n` digit characters.

```go
validation.New().Field("code", validation.MaxDigits(6))
// "1234" → pass, "1234567" → fail
```

---

<a id="mindigits"></a>
### MinDigits

```go
func MinDigits(n int) Rule
```

Fails if the value is not a string of at least `n` digit characters.

```go
validation.New().Field("phone", validation.MinDigits(7))
// "1234567" → pass, "123456" → fail
```

---

## DateTime

<a id="after"></a>
### After

```go
func After(ct time.Time) Rule
```

Fails if the value is not a parseable date/time string or is not strictly after `ct` (equal fails).

```go
deadline := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
validation.New().Field("expiry", validation.After(deadline))
// "2024-06-01" → pass, "2024-01-01" → fail (equal)
```

---

<a id="afterfield"></a>
### AfterField

```go
func AfterField(path string) InputRule
```

Cross-field rule. Fails if the value is not strictly after the date/time at `path`. Both fields must be parseable date/time strings.

```go
validation.New().
    Field("start", validation.Required, validation.DateTime).
    Field("end",   validation.Required, validation.AfterField("start"))
```

---

<a id="afterorequal"></a>
### AfterOrEqual

```go
func AfterOrEqual(ct time.Time) Rule
```

Fails if the value is not a parseable date/time string or is strictly before `ct`. Equal values pass.

```go
now := time.Now()
validation.New().Field("date", validation.AfterOrEqual(now))
```

---

<a id="before"></a>
### Before

```go
func Before(ct time.Time) Rule
```

Fails if the value is not a parseable date/time string or is not strictly before `ct` (equal fails).

```go
expiry := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
validation.New().Field("birthday", validation.Before(expiry))
```

---

<a id="beforefield"></a>
### BeforeField

```go
func BeforeField(path string) InputRule
```

Cross-field rule. Fails if the value is not strictly before the date/time at `path`.

```go
validation.New().
    Field("start", validation.Required, validation.BeforeField("end")).
    Field("end",   validation.Required, validation.DateTime)
```

---

<a id="beforeorequal"></a>
### BeforeOrEqual

```go
func BeforeOrEqual(ct time.Time) Rule
```

Fails if the value is not a parseable date/time string or is strictly after `ct`. Equal values pass.

```go
expiry := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
validation.New().Field("valid_until", validation.BeforeOrEqual(expiry))
```

---

<a id="datetime"></a>
### DateTime

```go
var DateTime Rule
```

Fails if the value is not a string or cannot be parsed as a date/time in any common format (RFC3339, ISO 8601, RFC1123, and many more).

```go
validation.New().Field("created_at", validation.DateTime)
// "2024-03-15" → pass, "not-a-date" → fail
```

---

<a id="datetimebetween"></a>
### DateTimeBetween

```go
func DateTimeBetween(min, max time.Time) Rule
```

Fails if the value is not a parseable date/time string or falls outside `[min, max]` (inclusive).

```go
start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
end   := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
validation.New().Field("event_date", validation.DateTimeBetween(start, end))
```

---

<a id="datetimeformat"></a>
### DateTimeFormat

```go
func DateTimeFormat(layout string) Rule
```

Fails if the value is not a string or does not match the given Go time layout exactly.

```go
validation.New().Field("date", validation.DateTimeFormat("2006-01-02"))
// "2024-03-15" → pass, "15/03/2024" → fail
```

---

<a id="timezone"></a>
### Timezone

```go
var Timezone Rule
```

Fails if the value is not a valid IANA timezone name (e.g. `"UTC"`, `"America/New_York"`).

```go
validation.New().Field("tz", validation.Timezone)
// "America/New_York" → pass, "InvalidZone" → fail
```

---

## Network

<a id="ip"></a>
### IP

```go
var IP Rule
```

Fails if the value is not a valid IPv4 or IPv6 address string.

```go
validation.New().Field("remote_addr", validation.IP)
// "192.168.1.1" → pass, "::1" → pass, "999.0.0.1" → fail
```

---

<a id="ipv4"></a>
### IPv4

```go
var IPv4 Rule
```

Fails if the value is not a valid IPv4 address string. IPv6 addresses fail.

```go
validation.New().Field("ip", validation.IPv4)
// "192.168.1.1" → pass, "::1" → fail
```

---

<a id="ipv6"></a>
### IPv6

```go
var IPv6 Rule
```

Fails if the value is not a valid IPv6 address string. IPv4 addresses fail.

```go
validation.New().Field("ip", validation.IPv6)
// "::1" → pass, "192.168.1.1" → fail
```

---

<a id="macaddress"></a>
### MACAddress

```go
var MACAddress Rule
```

Validates a 6-byte MAC address. Accepted formats: `01:23:45:67:89:ab` and `01-23-45-67-89-AB` (case-insensitive). 8-byte EUI-64 addresses are rejected.

```go
validation.New().Field("mac", validation.MACAddress)
// "01:23:45:67:89:ab" → pass, "not-a-mac" → fail
```

---

## Collection

<a id="distinct"></a>
### Distinct

```go
var Distinct Rule
```

Fails if the value is a slice or array containing duplicate elements. Non-slice values and non-comparable element types are skipped (pass).

```go
validation.New().Field("tags", validation.Distinct)
// []string{"a","b","c"} → pass, []int{1,2,1} → fail
```

---

## Generic

<a id="in"></a>
### In

```go
func In[T comparable](slice []T) Rule
```

Fails if the value cannot be type-asserted to `T` or is not found in `slice`. Comparison uses `==`.

```go
validation.New().Field("status", validation.In([]string{"active", "inactive", "pending"}))
// "active" → pass, "deleted" → fail
```

---

<a id="neq"></a>
### NEQ

```go
func NEQ[T comparable](v T) Rule
```

Fails if the value cannot be type-asserted to `T` or equals `v`. Comparison is type-sensitive.

```go
validation.New().Field("role", validation.NEQ[string]("banned"))
// "admin" → pass, "banned" → fail
```

---

<a id="notin"></a>
### NotIn

```go
func NotIn[T comparable](slice []T) Rule
```

Fails if the value cannot be type-asserted to `T` or is found in `slice`.

```go
validation.New().Field("username", validation.NotIn([]string{"admin", "root", "system"}))
// "alice" → pass, "admin" → fail
```

---

## Comparison

<a id="different"></a>
### Different

```go
func Different(path string) InputRule
```

Fails if the value equals the value at `path` in the input. If the referenced field is absent, the rule passes.

```go
validation.New().
    Field("new_password", validation.Required, validation.Different("old_password"))
```

---

<a id="sameas"></a>
### SameAs

```go
func SameAs(path string) InputRule
```

Fails if the value does not equal the value at `path` in the input. If the referenced field is absent, the rule fails. Comparison is type-sensitive (`==`).

```go
validation.New().
    Field("password_confirm", validation.Required, validation.SameAs("password"))
```
