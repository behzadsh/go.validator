# Go Validator
![Coverage](https://img.shields.io/badge/Coverage-99.7%25-brightgreen)

The `go.validator` package provides a simple and convenient way to validate your data.

## Installation

to install `go.validator` package, run the following command

```
go get -u github.com/behzadsh/go.validator
```

## How to use
You can use this package to validate the http request data in any form.

### Validating map data

For validating a map data you could use the function `ValidateMap`. The `ValidateMap` accept two parameters, the input data
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
            "birthDate": {"datetime"},
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
* `ValidateMapSlice`: For validating a slice of maps. The rule will be repeated for each of maps in the slice
* `ValidaeStruct`: For validating a struct data, just like the map.
* `ValidateStructSlice`: Same as `ValidateMapSlice`.

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
		// the validation errors stores in res.Errors under `variable` key when using `validation.Validate()`.
		for _, value := range res.Errors["variable"] {
            fmt.Println(value)
		}
		os.Exit(1)
	}
}
```

## Available Rules

The complete list of available rules can be found [here](https://github.com/behzadsh/go.validator/tree/main/rules.md).
