# Go by Example: Custom Errors

It is possible to define custom error types by implementing the `Error()`
method on them. Here is a variant on the previous example that uses a custom
type to explicitly represent an argument error.

```go
package main

import (
    "errors"
    "fmt"
)
```

Custom error types often include the suffix `Error`.

```go
type argError struct {
    arg     int
    message string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.message)
}
```

Adding the `Error()` method makes `argError` implement the `error` interface.

```go
func f(arg int) (int, error) {
    if arg == 42 {
        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}
```

In `main`, we use `errors.As` to check whether an error in the chain has type
`*argError`.

```go
func main() {
    _, err := f(42)
    var ae *argError
    if errors.As(err, &ae) {
        fmt.Println(ae.arg)
        fmt.Println(ae.message)
    } else {
        fmt.Println("err doesn't match argError")
    }
}
```

```bash
$ go run custom-errors.go
42
can't work with it
```
