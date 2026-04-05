# Go by Example: Errors

In Go, it is idiomatic to communicate errors via an explicit, separate return
value. This contrasts with the exceptions used in languages like Java, Python,
and Ruby. Go's approach makes it easy to see which functions return errors and
to handle them using the same language constructs used for other tasks.

```go
package main

import (
    "errors"
    "fmt"
)
```

By convention, errors are the last return value and have type `error`, a
built-in interface.

```go
func f(arg int) (int, error) {
    if arg == 42 {
        return -1, errors.New("can't work with 42")
    }
    return arg + 3, nil
}
```

`errors.New` constructs a basic error value with the given error message. A
`nil` value in the error position indicates success.

Sentinel errors are predeclared variables used to signal specific error
conditions.

```go
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")
```

We can wrap errors with higher-level context using the `%w` verb in
`fmt.Errorf`. Wrapped errors create an error chain that can be inspected with
functions like `errors.Is` and `errors.As`.

```go
func makeTea(arg int) error {
    if arg == 2 {
        return ErrOutOfTea
    } else if arg == 4 {
        return fmt.Errorf("making tea: %w", ErrPower)
    }
    return nil
}
```

```go
func main() {
    for _, i := range []int{7, 42} {
        if r, e := f(i); e != nil {
            fmt.Println("f failed:", e)
        } else {
            fmt.Println("f worked:", r)
        }
    }

    for i := range 5 {
        if err := makeTea(i); err != nil {
            if errors.Is(err, ErrOutOfTea) {
                fmt.Println("We should buy new tea!")
            } else if errors.Is(err, ErrPower) {
                fmt.Println("Now it is dark.")
            } else {
                fmt.Printf("unknown error: %s\n", err)
            }
            continue
        }
        fmt.Println("Tea is ready")
    }
}
```

It is common to use inline error checks in `if` statements, as shown above.
`errors.Is` checks whether an error (or any error in its chain) matches a
specific sentinel error value.

```bash
$ go run errors.go
f worked: 10
f failed: can't work with 42
Tea is ready
Tea is ready
We should buy new tea!
Tea is ready
Now it is dark.
```
