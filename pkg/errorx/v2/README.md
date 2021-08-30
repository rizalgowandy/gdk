# Errorx

Error package ideas comes and is a subset copy from [Upspin project](https://github.com/upspin/upspin).

## Error

Error is value in Go, and because error is a value, we need to check them. Go proverbs said:

> Don't just check errorx, handle them gracefully

and

> Log an error or return the error, never both - Dave Channey

### Why another error package

As `Dave Channey` said, we should log an error or just return the error, but never both. But how can we log a meaningful error in go and still can compare the error itself?

In order to do that, we need a modified implementation of error. Put more context into error and print the context when we need to log. That way we don't need to log and return an error at the same time, just to put more context into the log.

### errorx function

In order to create a meaningful error from this package, we need to use `errorx.E(args...)` function. Why `errorx.E()` instead of `errorx.New()`like `errorx` package from Go itself?

1. Following `upspin` convention to create the error.
2. Let the standard be a standard (`errors.New`), and the new one should have a new convention.

## Example

### Simple error creation

```go
import "github.com/rizalgowandy/gdk/pkg/errorx/v2"

func main() {
    err := errorx.E("this is error from library")
    // do something with the error
}

```

### Error with fields

Error with fields is useful to give context to error. For example `user_id` of user.

```go
import "github.com/rizalgowandy/gdk/pkg/errorx/v2"

func main() {
    err := errorx.E("this is error from library", errorx.Fields{"user_id": 1234})
    // do something with the error
}
```

### Error with operations

Sometimes we need to know what kind of operations we do in error, we want to know where exactly error happens.

```go
import "github.com/rizalgowandy/gdk/pkg/errorx/v2"

func main() {
    err := SomeFunction()
    // do something with the error
}

func SomeFunction() error {
    const op errorx.Op = "userService.FindUser"
    return errorx.E(op, "this is error from library")
}
```

### Real life example

This is an example where we need to call a function from handler and we need to know the error context

```go
import (
    "net/http"
    "strings"

    "github.com/rizalgowandy/gdk/pkg/errorx/v2"
)

// Error variables for error matching example
var (
    // Given string parameter on errorx.E() will directly converted to error message
    ErrParamIsEqual      = errorx.E("param1 is equal to param2")
    ErrMoreThanConstanta = errorx.E("param1 length is more than constanta")
)

func main() {
    http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
        param1 := r.FormValue("param1")
        param2 := r.FormValue("param2")

        err := BusinessLogic(param1, param2)
        // sample implementation of errorx.Match() to handle error regarding to error types
        if errorx.Match(err, ErrParamIsEqual) {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("Not OK"))
            return
        }

        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })

    http.ListenAndServe(":9090", nil)
}

func BusinessLogic(param1, param2 string) error {
    const op errorx.Op = "business/BusinessLogic"

    if strings.Compare(param1, param2) == 0 {
        return errorx.E(ErrParamIsEqual, errorx.Fields{
            "param1": param1,
            "param2": param2,
        }, op)
    }
    return ResourceLogic(param1)
}

const constVal string = "constanta"

func ResourceLogic(param1 string) error {
    const op errorx.Op = "resource/ResourceLogic"

    if len(param1) > len(constVal) {
        return errorx.E(ErrMoreThanConstanta, op, errorx.Fields{"param1": param1})
    }
    return nil
}

```