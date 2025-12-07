# Available Rules

## Index

<details>
<summary>Click to expand</summary>

- [After](#after)
- [After Or Equal](#afterOrEqual)
- [Alpha](#alpha)
- [Alpha Dash](#alphaDash)
- [Alpha Num](#alphaNum)
- [Alpha Space](#alphaSpace)
- [Array](#array)
- [Before](#before)
- [Before Or Equal](#beforeOrEqual)
- [Between](#between)
- [Boolean](#boolean)
- [Datetime](#datetime)
- [Datetime Format](#datetimeFormat)
- [DateTime After](#dateTimeAfter)
- [DateTime Before](#dateTimeBefore)
- [DateTime Between](#dateTimeBetween)
- [Different](#different)
- [Digits](#digits)
- [Digits Between](#digitsBetween)
- [Distinct](#distinct)
- [Email](#email)
- [Ends With](#endsWith)
- [Greater Than](#gt)
- [Greater Than or Equal](#gte)
- [In](#in)
- [InArrayField](#inArrayField)
- [Integer](#integer)
- [IP](#ip)
- [IPV4](#ipv4)
- [IPV6](#ipv6)
- [Length](#length)
- [Lowercase](#lowercase)
- [Less Than](#lt)
- [Less Than or Equal](#lte)
- [Max](#max)
- [Max Digits](#maxDigits)
- [Max Length](#maxLength)
- [Min](#min)
- [Min Digits](#minDigits)
- [Min Length](#minLength)
- [Not Equal](#neq)
- [Not Empty](#notEmpty)
- [Not In](#notIn)
- [Not Regex](#notRegex)
- [Numeric](#numeric)
- [Regex](#regex)
- [Required](#required)
- [Required If](#requiredIf)
- [Required Unless](#requiredUnless)
- [Required With](#requiredWith)
- [Required With All](#requiredWithAll)
- [Required Without](#requiredWithout)
- [Required Without All](#requiredWithoutAll)
- [Same As](#sameAs)
- [Starts With](#startsWith)
- [String](#string)
- [Timezone](#timezone)
- [Uppercase](#uppercase)
- [URL](#url)
- [UUID](#uuid)

</details>

<a id="after"></a>
## After: `after:otherField[,timeZone]`
This rule checks the field under validation is a datetime value after the given datetime. The `after` rule accepts two parameters, `otherField` and `timeZone`. The `otherField` is a mandatory parameter that specifies the field to compare the current field value.

> Note that when using this rule, the field under validation must pass the `datetime` rule, and the `otherField` must be present and also pass the `datetime` rule.

```go
rulesMap := validation.RulesMap{
	"end": {"after:start"} 
}
```

In this example, then `end` and `start` must pass the `datetime` rule, and also `start` must be present.

> The `after` rule does not imply the `required` rule on the field under validation.

### Timezone parameter
If you want to compare the time values in a specified time zone, you can pass your desired time zone string as the second parameter.

```go
rulesMap := validation.RulesMap{
	"end": {"after:start,America/New_York"} 
}
```

> The default value for the time zone parameter is **UTC**.

### Translation

| Key              | Params            |
|------------------|-------------------|
| validation.after | field, otherField |

| Param name | Description                                             |
|------------|---------------------------------------------------------|
| field      | The field under validation.                             |
| OtherField | The other field that the current field is compare with. |

<a id="afterOrEqual"></a>
## After Or Equal: `afterOrEqual:otherField[,timeZone]`
This rule checks the field under validation is a datetime value after or equal to the given datetime. The `afterOrEqual` rule accepts two parameters, `otherField` and `timeZone`. The `otherField` is a mandatory parameter that specifies the field to compare the current field value.

> Note that when using this rule, the field under validation must pass the `datetime` rule, and the `otherField` must be present and also pass the `datetime` rule.

```go
rulesMap := validation.RulesMap{
	"end": {"afterOrEqual:start"} 
}
```

In this example, then `end` and `start` must pass the `datetime` rule, and also `start` must be present.

> The `afterOrEqual` rule does not imply the `required` rule on the field under validation.

### Timezone parameter
If you want to compare the time values in a specified time zone, you can pass your desired time zone string as the second parameter.

```go
rulesMap := validation.RulesMap{
	"end": {"afterOrEqual:start,America/New_York"} 
}
```

> The default value for the time zone parameter is **UTC**.

### Translation

| Key                       | Params            |
|---------------------------|-------------------|
| validation.after_or_equal | field, otherField |

| Param name | Description                                             |
|------------|---------------------------------------------------------|
| field      | The field under validation.                             |
| OtherField | The other field that the current field is compare with. |

<a id="alpha"></a>
## Alpha: `alpha`
This rule checks the field under validation is entirely alphabetic characters.

```go
rulesMap := validation.RulesMap{
	"name": {"alpha"} 
}
```

### Translation

| Key              | Params |
|------------------|--------|
| validation.alpha | field  |

| Param name | Description                                             |
|------------|---------------------------------------------------------|
| field      | The field under validation.                             |

<a id="alphaDash"></a>
## Alpha Dash: `alphaDash`
This rule checks the field under validation has alphabetic characters as well as dashes and underscores.

```go
rulesMap := validation.RulesMap{
	"filename": {"alphaDash"} 
}
```

### Translation

| Key                   | Params |
|-----------------------|--------|
| validation.alpha_dash | field  |

| Param name | Description                                             |
|------------|---------------------------------------------------------|
| field      | The field under validation.                             |

<a id="alphaNum"></a>
## Alpha Num: `alphaNum`
This rule checks the field under validation has alphanumeric characters.

```go
rulesMap := validation.RulesMap{
	"title": {"alphaNum"} 
}
```

### Translation

| Key                  | Params |
|----------------------|--------|
| validation.alpha_num | field  |

| Param name | Description                                             |
|------------|---------------------------------------------------------|
| field      | The field under validation.                             |

<a id="alphaSpace"></a>
## Alpha Space: `alphaSpace`
This rule checks the field under validation has alphabetic characters as well as spaces.

```go
rulesMap := validation.RulesMap{
	"title": {"alphaSpace"} 
}
```

### Translation

| Key                    | Params |
|------------------------|--------|
| validation.alpha_space | field  |

| Param name | Description                                             |
|------------|---------------------------------------------------------|
| field      | The field under validation.                             |

<a id="array"></a>
## Array: `array`
This rule checks the field under validation is an array or slice.

```go
rulesMap := validation.RulesMap{
	"title": {"array"} 
}
```

### Translation

| Key              | Params |
|------------------|--------|
| validation.array | field  |

| Param name | Description                                             |
|------------|---------------------------------------------------------|
| field      | The field under validation.                             |

<a id="before"></a>
## Before: `before:otherField[,timeZone]`
This rule checks the field under validation is a datetime value before the given datetime. The `before` rule accepts two parameters, `otherField` and `timeZone`. The `otherField` is a mandatory parameter that specifies the field to compare the current field value.

> Note that when using this rule, the field under validation must pass the `datetime` rule, and the `otherField` must be present and also pass the `datetime` rule.

```go
rulesMap := validation.RulesMap{
	"start": {"before:end"} 
}
```

In this example, then `start` and `end` must pass the `datetime` rule, and also `end` must be present.

> The `before` rule does not imply the `required` rule on the field under validation.

### Timezone parameter
If you want to compare the time values in a specified time zone, you can pass your desired time zone string as the second parameter.

```go
rulesMap := validation.RulesMap{
	"end": {"before:start,America/New_York"} 
}
```

> The default value for the time zone parameter is **UTC**.

### Translation

| Key               | Params            |
|-------------------|-------------------|
| validation.before | field, otherField |

| Param name | Description                                             |
|------------|---------------------------------------------------------|
| field      | The field under validation.                             |
| OtherField | The other field that the current field is compare with. |

<a id="beforeOrEqual"></a>
## Before Or Equal: `beforeOrEqual:otherField[,timeZone]`
This rule checks the field under validation is a datetime value before or equal to the given datetime. The `beforeOrEqual` rule accepts two parameters, `otherField` and `timeZone`. The `otherField` is a mandatory parameter that specifies the field to compare the current field value.

> Note that when using this rule, the field under validation must pass the `datetime` rule, and the `otherField` must be present and also pass the `datetime` rule.

```go
rulesMap := validation.RulesMap{
	"start": {"beforeOrEqual:end"} 
}
```

In this example, then `start` and `end` must pass the `datetime` rule, and also `end` must be present.

> The `beforeOrEqual` rule does not imply the `required` rule on the field under validation.

### Timezone parameter
If you want to compare the time values in a specified time zone, you can pass your desired time zone string as the second parameter.

```go
rulesMap := validation.RulesMap{
    "start": {"beforeOrEqual:end,America/New_York"} 
}
```

> The default value for the time zone parameter is **UTC**.

### Translation

| Key                        | Params            |
|----------------------------|-------------------|
| validation.before_or_equal | field, otherField |

| Param name | Description                                             |
|------------|---------------------------------------------------------|
| field      | The field under validation.                             |
| OtherField | The other field that the current field is compare with. |

<a id="between"></a>
## Between: `between:min,max`
This rule acts a little differently based on the type of field value.

 * If the value is numeric, the rule will check the value is between the given min and max.
 * If the value is a string, slice, array, or map, the rule will check the length of the value is between the given min and max.
 * Otherwise, the rule checks nothing.

```go
rulesMap := validation.RulesMap{
	"age": {"between:18,30"},       // assume age is an integer value
	"score": {"between:8.5,10"},    // assume score is a float value
	"title": {"between:5,50"},      // assume title is a string value
	"skills": {"between:2,5"},      // assume skills is a slice of strings
}
```

> The `between` rule only checks the condition when the field under validation is present.

### Translation

| Key                | Params          |
|--------------------|-----------------|
| validation.between | field, min, max |

| Param name | Description                        |
|------------|------------------------------------|
| field      | The field under validation.        |
| min        | The minimum value set in the rule. |
| max        | The maximum value set in the rule. |

<a id="boolean"></a>
## Boolean: `boolean`
This rule checks the field under validation has a boolean value.

```go
rulesMap := validation.RulesMap{
	"accept": {"boolean"} 
}
```

### Translation

| Key                | Params |
|--------------------|--------|
| validation.boolean | field  |

| Param name | Description                                             |
|------------|---------------------------------------------------------|
| field      | The field under validation.                             |

<a id="datetime"></a>
## Datetime: `datetime`
This rule checks the field under validation is a valid datetime string.

```go
rulesMap := validation.RulesMap{
	"start": {"datetime"} 
}
```

### Translation

| Key                 | Params |
|---------------------|--------|
| validation.datetime | field  |

| Param name | Description                                             |
|------------|---------------------------------------------------------|
| field      | The field under validation.                             |

<a id="datetimeFormat"></a>
## Datetime Format: `datetimeFormat:format`
This rule checks the field under validation matches the given datetime layout format.

```go
rulesMap := validation.RulesMap{
	"start": {"datetimeFormat:2006-01-02T15:04:05Z07:00"}
}
```

> The `datetimeFormat` rule only checks the condition when the field under validation is present.

### Translation

| Key                        | Params        |
|----------------------------|---------------|
| validation.datetime_format | field, format |

| Param name | Description                 |
|------------|-----------------------------|
| field      | The field under validation. |
| format     | The datetime layout format. |

## DateTime After: `dateTimeAfter:value[,timeZone]`
This rule checks the field under validation is a datetime after the given datetime value.

```go
rulesMap := validation.RulesMap{
    "start": {"dateTimeAfter:2022-01-01,America/New_York"},
}
```

> The default value for the time zone parameter is **UTC**.

### Translation

| Key                       | Params       |
|---------------------------|--------------|
| validation.datetime_after | field, value |

| Param name | Description                        |
|------------|------------------------------------|
| field      | The field under validation.        |
| value      | The threshold datetime string.     |

## DateTime Before: `dateTimeBefore:value[,timeZone]`
This rule checks the field under validation is a datetime before the given datetime value.

```go
rulesMap := validation.RulesMap{
    "end": {"dateTimeBefore:2022-01-01,America/New_York"},
}
```

> The default value for the time zone parameter is **UTC**.

### Translation

| Key                        | Params       |
|----------------------------|--------------|
| validation.datetime_before | field, value |

| Param name | Description                        |
|------------|------------------------------------|
| field      | The field under validation.        |
| value      | The threshold datetime string.     |

## DateTime Between: `dateTimeBetween:min,max[,timeZone]`
This rule checks the field under validation is a datetime value between the given min and max.

```go
rulesMap := validation.RulesMap{
    "date": {"dateTimeBetween:2022-01-01,2022-02-01,UTC"},
}
```

> The default value for the time zone parameter is **UTC**.

### Translation

| Key                          | Params          |
|------------------------------|-----------------|
| validation.datetime_between  | field, min, max |

| Param name | Description                        |
|------------|------------------------------------|
| field      | The field under validation.        |
| min        | The minimum datetime string.       |
| max        | The maximum datetime string.       |

<a id="different"></a>
## Different: `different:otherField`
This rule checks the field under validation has a different value than the given field.

```go
rulesMap := validation.RulesMap{
	"newPassword": {"different:oldPassword"} 
}
```

### Translation

| Key                  | Params            |
|----------------------|-------------------|
| validation.different | field, otherField |

| Param name | Description                                        |
|------------|----------------------------------------------------|
| field      | The field under validation.                        |
| otherField | The other field that current field is compared to. |

<a id="digits"></a>
## Digits: `digits:count`
This rule checks the field under validation is all digits and has exact digits as given count.

```go
rulesMap := validation.RulesMap{
	"code": {"digits:6"} 
}
```

### Translation

| Key               | Params             |
|-------------------|--------------------|
| validation.digits | field, digitsCount |

| Param name  | Description                       |
|-------------|-----------------------------------|
| field       | The field under validation.       |
| digitsCount | The number of digits set in rule. |

<a id="digitsBetween"></a>
## Digits Between: `digitsBetween:min,max`
This rule checks the field under validation is all digits and has digits between the given min and max.

```go
rulesMap := validation.RulesMap{
    "code": {"digitsBetween:4,6"} 
}
```

### Translation

| Key                       | Params          |
|---------------------------|-----------------|
| validation.digits_between | field, min, max |

| Param name | Description                               |
|------------|-------------------------------------------|
| field      | The field under validation.               |
| min        | The minimum digits count set in the rule. |
| max        | The maximum digits count set in the rule. |

<a id="email"></a>
## Email: `email[:mx]`
This rule checks the field under validation has a valid email format.

```go
rulesMap := validation.RulesMap{
	"email": {"email"}
}
```

There is also an optional parameter called `mx` which doing extra check on email address and checks the
MX record in given email host dns records.

```go
rulesMap := validation.RulesMap{
	"email": {"email:mx"}
}
```

### Translation

| Key              | Params |
|------------------|--------|
| validation.email | field  |

| Param name | Description                               |
|------------|-------------------------------------------|
| field      | The field under validation.               |

## Distinct: `distinct`
This rule checks the field under validation is an array/slice with unique elements.

```go
rulesMap := validation.RulesMap{
    "tags": {"distinct"},
}
```

### Translation

| Key                 | Params |
|---------------------|--------|
| validation.distinct | field  |

| Param name | Description                 |
|------------|-----------------------------|
| field      | The field under validation. |

<a id="endsWith"></a>
## Ends With: `ends_with:suffix`
This rules checks the field under validation ends with the given suffix string.

```go
rulesMap := validation.RulesMap{
	"className": {"endsWith:Model"}
}
```

> The `suffix` parameter is case-sensitive. 

### Translation

| Key                  | Params       |
|----------------------|--------------|
| validation.ends_with | field, value |

| Param name | Description                   |
|------------|-------------------------------|
| field      | The field under validation.   |
| value      | The suffix value set in rule. |

<a id="gt"></a>
## Greater Than: `gt:value`
This rule acts a little differently based on the type of field value.

* If the value is numeric, the rule will check the value is greater than given value.
* If the value is a string, slice, array, or map, the rule will check the length of the value is greater than given value.
* Otherwise, the rule checks nothing.

```go
rulesMap := validation.RulesMap{
	"age": {"gt:18"},       // assume age is an integer value
	"score": {"gt:8.5"},    // assume score is a float value
	"title": {"gt:5"},      // assume title is a string value
	"skills": {"gt:2"},     // assume skills is a slice of strings
}
```

### Translation

| Key           | Params       |
|---------------|--------------|
| validation.gt | field, value |

| Param name | Description                 |
|------------|-----------------------------|
| field      | The field under validation. |
| value      | The min value set in rule.  |

<a id="gte"></a>
## Greater Than or Equal: `gte:value`
This rule acts a little differently based on the type of field value.

* If the value is numeric, the rule will check the value is greater than or equal to given value.
* If the value is a string, slice, array, or map, the rule will check the length of the value is greater than or equal to given value.
* Otherwise, the rule checks nothing.

```go
rulesMap := validation.RulesMap{
	"age": {"gte:18"},       // assume age is an integer value
	"score": {"gte:8.5"},    // assume score is a float value
	"title": {"gte:5"},      // assume title is a string value
	"skills": {"gte:2"},     // assume skills is a slice of strings
}
```

### Translation

| Key            | Params       |
|----------------|--------------|
| validation.gte | field, value |

| Param name | Description                 |
|------------|-----------------------------|
| field      | The field under validation. |
| value      | The min value set in rule.  |

<a id="in"></a>
## In: `in:value1,value2[,value3,...]`
This rule check the field under validation exists in one of the given options. 

> This rule is case-insensitive.

```go
rulesMap := validation.RulesMap{
	"currency": {"in:usd,eur,gbp"},
}
```

### Translation

| Key           | Params |
|---------------|--------|
| validation.in | field  |

| Param name | Description                 |
|------------|-----------------------------|
| field      | The field under validation. |

## In Array Field: `inArrayField:otherField`
This rule checks the field under validation exists in another array/slice field.

```go
rulesMap := validation.RulesMap{
    "choice": {"inArrayField:options"},
}
```

### Translation

| Key                 | Params            |
|---------------------|-------------------|
| validation.in       | field             |
| Validation.required | otherField        |

| Param name | Description                                 |
|------------|---------------------------------------------|
| field      | The field under validation.                  |
| otherField | The other field that contains valid options. |

<a id="integer"></a>
## Integer: `integer`
This rule checks the field under validation is an integer.

```go
rulesMap := validation.RulesMap{
	"age": {"integer"},
}
```

### Translation

| Key                | Params |
|--------------------|--------|
| validation.integer | field  |

| Param name | Description                 |
|------------|-----------------------------|
| field      | The field under validation. |

## IP: `ip`
This rule checks the field under validation is a valid IP address (v4 or v6).

```go
rulesMap := validation.RulesMap{
    "addr": {"ip"},
}
```

### Translation

| Key            | Params |
|----------------|--------|
| validation.ip  | field  |

| Param name | Description                 |
|------------|-----------------------------|
| field      | The field under validation. |

## IPV4: `ipv4`
This rule checks the field under validation is a valid IPv4 address.

```go
rulesMap := validation.RulesMap{
    "addr": {"ipv4"},
}
```

### Translation

| Key              | Params |
|------------------|--------|
| validation.ipv4  | field  |

| Param name | Description                 |
|------------|-----------------------------|
| field      | The field under validation. |

## IPV6: `ipv6`
This rule checks the field under validation is a valid IPv6 address.

```go
rulesMap := validation.RulesMap{
    "addr": {"ipv6"},
}
```

### Translation

| Key              | Params |
|------------------|--------|
| validation.ipv6  | field  |

| Param name | Description                 |
|------------|-----------------------------|
| field      | The field under validation. |

<a id="length"></a>
## Length: `length:value`
This rule checks the field under validation length is as the given value.

```go
rulesMap := validation.RulesMap{
	"skills": {"length:3"},
}
```

> The `length` rule only apply on strings, arrays, slices and maps.

### Translation

| Key               | Params       |
|-------------------|--------------|
| validation.length | field, value |

| Param name | Description                  |
|------------|------------------------------|
| field      | The field under validation.  |
| value      | The given value in the rule. |


<a id="lowercase"></a>
## Lowercase: `lowercase`
This rule checks the field under validation is all lowercase.

```go
rulesMap := validation.RulesMap{
	"name": {"lowercase"},
}
```

### Translation

| Key                  | Params |
|----------------------|--------|
| validation.lowercase | field  |

| Param name | Description                 |
|------------|-----------------------------|
| field      | The field under validation. |

<a id="lt"></a>
## Less Than: `lt:value`
This rule acts a little differently based on the type of field value.

* If the value is numeric, the rule will check the value is less than given value.
* If the value is a string, slice, array, or map, the rule will check the length of the value is less than given value.
* Otherwise, the rule checks nothing.

```go
rulesMap := validation.RulesMap{
	"age": {"lt:18"},       // assume age is an integer value
	"score": {"lt:8.5"},    // assume score is a float value
	"title": {"lt:5"},      // assume title is a string value
	"skills": {"lt:2"},     // assume skills is a slice of strings
}
```

### Translation

| Key           | Params       |
|---------------|--------------|
| validation.lt | field, value |

| Param name | Description                 |
|------------|-----------------------------|
| field      | The field under validation. |
| value      | The max value set in rule.  |

<a id="lte"></a>
## Less Than or Equal: `lte:value`
This rule acts a little differently based on the type of field value.

* If the value is numeric, the rule will check the value is less than or equal to given value.
* If the value is a string, slice, array, or map, the rule will check the length of the value is less than or equal to given value.
* Otherwise, the rule checks nothing.

```go
rulesMap := validation.RulesMap{
	"age": {"lte:18"},       // assume age is an integer value
	"score": {"lte:8.5"},    // assume score is a float value
	"title": {"lte:5"},      // assume title is a string value
	"skills": {"lte:2"},     // assume skills is a slice of strings
}
```

### Translation

| Key            | Params       |
|----------------|--------------|
| validation.lte | field, value |

| Param name | Description                 |
|------------|-----------------------------|
| field      | The field under validation. |
| value      | The max value set in rule.  |

<a id="max"></a>
## Max: `max:value`
This rule checks the field under validation value be less than given value.

```go
rulesMap := validation.RulesMap{
	"age": {"max:30"}
}
```

### Translation

| Key            | Params       |
|----------------|--------------|
| validation.max | field, value |

| Param name | Description                 |
|------------|-----------------------------|
| field      | The field under validation. |
| value      | The max value set in rule.  |

<a id="maxDigits"></a>
## Max Digits: `maxDigits:value`
This rule checks the field under validation has length less than given max digits.

```go
rulesMap := validation.RulesMap{
	"code": {"maxDigits:5"}
}
```

### Translation

| Key                   | Params             |
|-----------------------|--------------------|
| validation.max_digits | field, digitsCount |

| Param name  | Description                           |
|-------------|---------------------------------------|
| field       | The field under validation.           |
| digitsCount | The max number of digits set in rule. |

<a id="maxLength"></a>
## Max Length: `maxLength:value`
This rule checks the field under validation has length less than given value.

```go
rulesMap := validation.RulesMap{
	"title": {"maxLength:50"}
}
```

### Translation

| Key                   | Params       |
|-----------------------|--------------|
| validation.max_length | field, value |

| Param name | Description                     |
|------------|---------------------------------|
| field      | The field under validation.     |
| value      | The maximum length set in rule. |

<a id="min"></a>
## Min: `min:value`
This rule checks the field under validation value be greater than given value.

```go
rulesMap := validation.RulesMap{
	"age": {"min:30"}
}
```

### Translation

| Key            | Params       |
|----------------|--------------|
| validation.min | field, value |

| Param name | Description                 |
|------------|-----------------------------|
| field      | The field under validation. |
| value      | The min value set in rule.  |

<a id="minDigits"></a>
## Min Digits: `minDigits:value`
This rule checks the field under validation has length greater than given min digits.

```go
rulesMap := validation.RulesMap{
	"code": {"minDigits:5"}
}
```

### Translation

| Key                   | Params             |
|-----------------------|--------------------|
| validation.min_digits | field, digitsCount |

| Param name  | Description                           |
|-------------|---------------------------------------|
| field       | The field under validation.           |
| digitsCount | The min number of digits set in rule. |

<a id="minLength"></a>
## Min Length: `minLength:value`
This rule checks the field under validation has length greater than given value.

```go
rulesMap := validation.RulesMap{
	"title": {"minLength:10"}
}
```

### Translation

| Key                   | Params       |
|-----------------------|--------------|
| validation.min_length | field, value |

| Param name | Description                     |
|------------|---------------------------------|
| field      | The field under validation.     |
| value      | The minimum length set in rule. |

<a id="neq"></a>
## Not Equal: `neq:value`
This rule checks the field under validation is not equal to given value.

```go
rulesMap := validation.RulesMap{
	"username": {"neq:admin"}
}
```

### Translation

| Key            | Params       |
|----------------|--------------|
| validation.neq | field, value |

| Param name | Description                     |
|------------|---------------------------------|
| field      | The field under validation.     |
| value      | The given value in the rule.    |

<a id="notEmpty"></a>
## Not Empty: `notEmpty`

This rule checks the field under validation presents and have a non-empty or non-zero value.

```go
rulesMap := validation.RulesMap{
	"username": {"notEmpty"}
}
```

### Translation

| Key                  | Params |
|----------------------|--------|
| validation.not_empty | field  |

| Param name | Description                     |
|------------|---------------------------------|
| field      | The field under validation.     |

<a id="notIn"></a>
## Not In: `notIn:value1,value2[,value3,...]`

This rule check the field under validation not exists in one of the given options.

```go
rulesMap := validation.RulesMap{
	"username": {"notIn:admin,god,superuser"},
}
```

### Translation

| Key               | Params |
|-------------------|--------|
| validation.not_in | field  |

| Param name | Description                 |
|------------|-----------------------------|
| field      | The field under validation. |

<a id="notRegex"></a>
## Not Regex: `notRegex:pattern`
This rule checks the field under validation does not match the given regex pattern.

```go
rulesMap := validation.RulesMap{
	"username": {"notRegex:[\\s]+"},
}
```

### Translation

| Key                  | Params         |
|----------------------|----------------|
| validation.not_regex | field, pattern |

| Param name | Description                    |
|------------|--------------------------------|
| field      | The field under validation.    |
| pattern    | The given pattern in the rule. |

<a id="numeric"></a>
## Numeric: `numeric`
This rule checks the field under validation has a numeric value.


```go
rulesMap := validation.RulesMap{
	"amount": {"numeric"},
}
```

### Translation

| Key                | Params |
|--------------------|--------|
| validation.numeric | field  |

| Param name | Description                 |
|------------|-----------------------------|
| field      | The field under validation. |

<a id="regex"></a>
## Regex: `regex:pattern`
This rule checks the field under validation matches the given regex pattern.

```go
rulesMap := validation.RulesMap{
	"username": {"regex:[a-zA-z0-9\\._]+"},
}
```

### Translation

| Key                  | Params         |
|----------------------|----------------|
| validation.regex     | field, pattern |

| Param name | Description                    |
|------------|--------------------------------|
| field      | The field under validation.    |
| pattern    | The given pattern in the rule. |

<a id="required"></a>
## Required: `required`
This rule checks the field under validation exists.

```go
rulesMap := validation.RulesMap{
	"username": {"required"},
}
```

### Translation

| Key                 | Params |
|---------------------|--------|
| validation.required | field  |

| Param name | Description                 |
|------------|-----------------------------|
| field      | The field under validation. |

<a id="requiredIf"></a>
## Required If: `requiredIf:otherField,value`
This rule check the field under validation exists if the given condition is true.
The condition is consists of another field name and a value. If the value of the other field is equal to
the given value, the field under validation is required.

> Note that the only supported type for value parameter is string.

```go
rulesMap := validation.RulesMap{
	"username": {"requiredIf:type,user"},
}
```

### Translation

| Key                    | Params                   |
|------------------------|--------------------------|
| validation.required_if | field, otherField, value |

| Param name | Description                          |
|------------|--------------------------------------|
| field      | The field under validation.          |
| otherField | The other field to check its value.  |
| value      | The given value for the other value. |

<a id="requiredUnless"></a>
## Required Unless: `requiredUnless:otherField,value`
This rule check the field under validation exists unless the given condition is true.
The condition is consists of another field name and a value. Unless the value of the other field is equal to
the given value, the field under validation is required.

> Note that the only supported type for value parameter is string.

```go
rulesMap := validation.RulesMap{
	"username": {"requiredUnless:type,admin"},
}
```

### Translation

| Key                        | Params                   |
|----------------------------|--------------------------|
| validation.required_unless | field, otherField, value |

| Param name | Description                          |
|------------|--------------------------------------|
| field      | The field under validation.          |
| otherField | The other field to check its value.  |
| value      | The given value for the other value. |

<a id="requiredWith"></a>
## Required With: `requiredWith:otherField[,anotherField,...]`
This rule check the field under validation exists if any of given fields exist.

```go
rulesMap := validation.RulesMap{
	"username": {"requiredWith:phone,email"},
}
```

### Translation

| Key                      | Params            |
|--------------------------|-------------------|
| validation.required_with | field, otherField |

| Param name | Description                          |
|------------|--------------------------------------|
| field      | The field under validation.          |
| otherField | The other field to check its value.  |

<a id="requiredWithAll"></a>
## Required With All: `requiredWithAll:otherField,anotherField[,...]`
This rule check the field under validation exists if all given fields exist.

```go
rulesMap := validation.RulesMap{
	"username": {"requiredWithAll:phone,email"},
}
```

### Translation

| Key                          | Params             |
|------------------------------|--------------------|
| validation.required_with_all | field, otherFields |

| Param name  | Description                          |
|-------------|--------------------------------------|
| field       | The field under validation.          |
| otherFields | The other field to check its value.  |

<a id="requiredWithout"></a>
## Required Without: `requiredWithout:otherField[,anotherField,...]`
This rule check the field under validation exists if any of given fields doesn't exist.

```go
rulesMap := validation.RulesMap{
	"username": {"requiredWithout:phone,email"},
}
```

### Translation

| Key                         | Params            |
|-----------------------------|-------------------|
| validation.required_without | field, otherField |

| Param name | Description                          |
|------------|--------------------------------------|
| field      | The field under validation.          |
| otherField | The other field to check its value.  |

<a id="requiredWithoutAll"></a>
## Required Without All: `requiredWithoutAll:otherField,anotherField[,...]`
This rule check the field under validation exists if all given fields not exist.

```go
rulesMap := validation.RulesMap{
	"username": {"requiredWithoutAll:phone,email"},
}
```

### Translation

| Key                             | Params             |
|---------------------------------|--------------------|
| validation.required_without_all | field, otherFields |

| Param name  | Description                          |
|-------------|--------------------------------------|
| field       | The field under validation.          |
| otherFields | The other field to check its value.  |

<a id="sameAs"></a>
## Same As: `sameAs:otherField`
This rule check the field under validation has the value same as the other given field.

```go
rulesMap := validation.RulesMap{
	"passwordConfirmation": {"sameAs:password"},
}
```

### Translation

| Key                 | Params            |
|---------------------|-------------------|
| validation.same_as  | field, otherField |

| Param name | Description                          |
|------------|--------------------------------------|
| field      | The field under validation.          |
| otherField | The other field to check its value.  |

<a id="startsWith"></a>
## Starts With: `startsWith:prefix`
This rule check the field under validation starts with given sub string.

```go
rulesMap := validation.RulesMap{
    "functionName": {"startsWith:Set"}
}
```

> The `prefix` parameter is case-sensitive.

### Translation

| Key                     | Params       |
|-------------------------|--------------|
| validation.starts_with   | field, value |

| Param name | Description                   |
|------------|-------------------------------|
| field      | The field under validation.   |
| value      | The prefix value set in rule. |

## Timezone: `timezone`
This rule checks the field under validation is a valid IANA time zone name.

```go
rulesMap := validation.RulesMap{
    "tz": {"timezone"},
}
```

### Translation

| Key                 | Params |
|---------------------|--------|
| validation.timezone | field  |

| Param name | Description                 |
|------------|-----------------------------|
| field      | The field under validation. |

<a id="string"></a>
## String: `string`
This rule checks the field under validation is a string.

```go
rulesMap := validation.RulesMap{
    "title": {"string"},
}
```

### Translation

| Key               | Params |
|-------------------|--------|
| validation.string | field  |

| Param name | Description                 |
|------------|-----------------------------|
| field      | The field under validation. |

<a id="uppercase"></a>
## Uppercase: `uppercase`
This rule checks the field under validation is all uppercase.

```go
rulesMap := validation.RulesMap{
	"code": {"uppercase"},
}
```

### Translation

| Key                  | Params |
|----------------------|--------|
| validation.uppercase | field  |

| Param name | Description                 |
|------------|-----------------------------|
| field      | The field under validation. |

<a id="url"></a>
## URL: `url[:scheme]`
This rule checks the field under validation is a valid URL. By default, the scheme is optional; pass `scheme` to require it.

```go
rulesMap := validation.RulesMap{
    "site": {"url"},          // scheme optional
    "strict": {"url:scheme"}, // scheme required
}
```

### Translation

| Key            | Params |
|----------------|--------|
| validation.url | field  |

| Param name | Description                 |
|------------|-----------------------------|
| field      | The field under validation. |

<a id="uuid"></a>
## UUID: `uuid`
This rule checks field under validation is a valid uuid.

```go
rulesMap := validation.RulesMap{
	"id": {"uuid"},
}
```

### Translation

| Key             | Params |
|-----------------|--------|
| validation.uuid | field  |

| Param name | Description                 |
|------------|-----------------------------|
| field      | The field under validation. |
