# Available Rules

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

## Before Or Equal: `beforeOrEqual:otherField[,timeZone]`
This rule checks the field under validation is a datetime value before the given datetime. The `beforeOrEqual` rule accepts two parameters, `otherField` and `timeZone`. The `otherField` is a mandatory parameter that specifies the field to compare the current field value.

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
	"end": {"beforeOrEqual:start,America/New_York"} 
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

## Digits Between: `digitsBetween:min,max`
This rule checks the field under validation is all digits and has digits between the given min and max.

```go
rulesMap := validation.RulesMap{
	"code": {"digits:4,6"} 
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

## In: `in:value1,value2[,value3,...]`
This rule check the field under validation exists in one of the given options.

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

# Max: `max:value`
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

# Min: `min:value`
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

## Not Equal: `neq:value`
This rule checks the field under validation is not equal to given value.

```go
rulesMap := validation.RulesMap{
	"username": {"neq:admin"}
}
```

### Translation

| Key                  | Params       |
|----------------------|--------------|
| validation.not_equal | field, value |

| Param name | Description                     |
|------------|---------------------------------|
| field      | The field under validation.     |
| value      | The maximum length set in rule. |

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

## Start With: `startWith:prefix`
This rule check the field under validation starts with given sub string.

```go
rulesMap := validation.RulesMap{
	"functionName": {"startWith:Set"}
}
```

> The `prefix` parameter is case-sensitive.

### Translation

| Key                  | Params       |
|----------------------|--------------|
| validation.ends_with | field, value |

| Param name | Description                   |
|------------|-------------------------------|
| field      | The field under validation.   |
| value      | The prefix value set in rule. |

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
