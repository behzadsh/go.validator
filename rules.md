# Available Rules

## After: `after:otherField[,timeZone]`
This rule checks the field under validation is a datetime value after the given datetime. The `after` rule accepts two parameters, `otherField` and `timeZone`. The `otherField` is a mandatory parameter that specifies the field to compare the current field value.

> Note that when using this rule, the field under validation must pass the `datetime` rule, and the `otherField` must be present and also pass the `datetime` rule.

```go
ruleMap := validation.RuleMap{
	"end": {"after:start"} 
}
```

In this example, then `end` and `start` must pass the `datetime` rule, and also `start` must be present.

> The `after` rule does not imply the `required` rule on the field under validation.

### Timezone parameter
If you want to compare the time values in a specified time zone, you can pass your desired time zone string as the second parameter.

```go
ruleMap := validation.RuleMap{
	"end": {"after:start,America/New_York"} 
}
```

> The default value for the time zone parameter is **UTC**.

## After Or Equal: `afterOrEqual:otherField[,timeZone]`
This rule checks the field under validation is a datetime value after or equal to the given datetime. The `afterOrEqual` rule accepts two parameters, `otherField` and `timeZone`. The `otherField` is a mandatory parameter that specifies the field to compare the current field value.

> Note that when using this rule, the field under validation must pass the `datetime` rule, and the `otherField` must be present and also pass the `datetime` rule.

```go
ruleMap := validation.RuleMap{
	"end": {"afterOrEqual:start"} 
}
```

In this example, then `end` and `start` must pass the `datetime` rule, and also `start` must be present.

> The `afterOrEqual` rule does not imply the `required` rule on the field under validation.

### Timezone parameter
If you want to compare the time values in a specified time zone, you can pass your desired time zone string as the second parameter.

```go
ruleMap := validation.RuleMap{
	"end": {"afterOrEqual:start,America/New_York"} 
}
```

> The default value for the time zone parameter is **UTC**.

## Alpha: `alpha`
This rule checks the field under validation is entirely alphabetic characters.

```go
ruleMap := validation.RuleMap{
	"name": {"alpha"} 
}
```

## Alpha Dash: `alphaDash`
This rule checks the field under validation has alphabetic characters as well as dashes and underscores.

```go
ruleMap := validation.RuleMap{
	"filename": {"alphaDash"} 
}
```

## Alpha Num: `alphaNum`
This rule checks the field under validation has alphanumeric characters.

```go
ruleMap := validation.RuleMap{
	"title": {"alphaNum"} 
}
```

## Alpha Space: `alphaSpace`
This rule checks the field under validation has alphabetic characters as well as spaces.

```go
ruleMap := validation.RuleMap{
	"title": {"alphaSpace"} 
}
```

## Array: `array`
This rule checks the field under validation is an array or slice.

```go
ruleMap := validation.RuleMap{
	"title": {"array"} 
}
```

## Before: `before:otherField[,timeZone]`
This rule checks the field under validation is a datetime value before the given datetime. The `before` rule accepts two parameters, `otherField` and `timeZone`. The `otherField` is a mandatory parameter that specifies the field to compare the current field value.

> Note that when using this rule, the field under validation must pass the `datetime` rule, and the `otherField` must be present and also pass the `datetime` rule.

```go
ruleMap := validation.RuleMap{
	"start": {"before:end"} 
}
```

In this example, then `start` and `end` must pass the `datetime` rule, and also `end` must be present.

> The `before` rule does not imply the `required` rule on the field under validation.

### Timezone parameter
If you want to compare the time values in a specified time zone, you can pass your desired time zone string as the second parameter.

```go
ruleMap := validation.RuleMap{
	"end": {"before:start,America/New_York"} 
}
```

> The default value for the time zone parameter is **UTC**.

## Before Or Equal: `beforeOrEqual:otherField[,timeZone]`
This rule checks the field under validation is a datetime value before the given datetime. The `beforeOrEqual` rule accepts two parameters, `otherField` and `timeZone`. The `otherField` is a mandatory parameter that specifies the field to compare the current field value.

> Note that when using this rule, the field under validation must pass the `datetime` rule, and the `otherField` must be present and also pass the `datetime` rule.

```go
ruleMap := validation.RuleMap{
	"start": {"beforeOrEqual:end"} 
}
```

In this example, then `start` and `end` must pass the `datetime` rule, and also `end` must be present.

> The `beforeOrEqual` rule does not imply the `required` rule on the field under validation.

### Timezone parameter
If you want to compare the time values in a specified time zone, you can pass your desired time zone string as the second parameter.

```go
ruleMap := validation.RuleMap{
	"end": {"beforeOrEqual:start,America/New_York"} 
}
```

> The default value for the time zone parameter is **UTC**.

## Between: `between:min,max`
This rule acts a little differently based on the type of field value.

 * If the value is numeric, the rule will check the value is between the given min and max.
 * If the value is a string, slice, array, or map, the rule will check the length of the value is between the given min and max.
 * Otherwise, the rule checks nothing.

```go
ruleMap := validation.RuleMap{
	"age": {"between:18,30"},
	"score": {"between:8.5,10"},
	"title": {"between:5,50},
	"skills": {"between:2,5"},
}
```

> The `between` rule only checks the condition when the field under validation is present.

## Boolean: `boolean`
This rule checks the field under validation has a boolean value.

```go
ruleMap := validation.RuleMap{
	"accept": {"boolean"} 
}
```

## Datetime: `datetime`
This rule checks the field under validation is a valid datetime string.

```go
ruleMap := validation.RuleMap{
	"start": {"datetime"} 
}
```

## Datetime Format: `datetimeFormat:format`
This rule checks the field under validation matches the given datetime layout format.

```go
ruleMap := validation.RuleMap{
	"start": {"datetimeFormat:2006-01-02T15:04:05Z07:00"}
}
```

> The `datetimeFormat` rule only checks the condition when the field under validation is present

## Different: `different:otherField`
This rule checks the field under validation has a different value than the given field.

```go
ruleMap := validation.RuleMap{
	"newPassword": {"different:oldPassword"} 
}
```

## Digits: `digits:count`
This rule checks the field under validation is all digits and has exact digits as given count.

```go
ruleMap := validation.RuleMap{
	"code": {"digits:6"} 
}
```

## Digits Between: `digitsBetween:min,max`
This rule checks the field under validation is all digits and has digits between the given min and max.

```go
ruleMap := validation.RuleMap{
	"code": {"digits:4,6"} 
}
```
