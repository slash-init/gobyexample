# Go by Example: Goroutines

A goroutine is a lightweight thread of execution.

```go
package main

import (
    "fmt"
    "time"
)
```

```go
func f(from string) {
    for i := range 3 {
        fmt.Println(from, ":", i)
    }
}
```

Suppose we have a function call `f(s)`. Here is how we would call it in the
usual way, running it synchronously.

```go
func main() {
    f("direct")
```

To invoke this function in a goroutine, use `go f(s)`. This new goroutine will
execute concurrently with the calling one.

```go
    go f("goroutine")
```

You can also start a goroutine for an anonymous function call.

```go
    go func(msg string) {
        fmt.Println(msg)
    }("going")
```

Our two function calls are running asynchronously in separate goroutines now.
Wait for them to finish. For a more robust approach, use a `sync.WaitGroup`.

```go
    time.Sleep(time.Second)
    fmt.Println("done")
}
```

When we run this program, we see the output of the blocking call first, then
output from the two goroutines. The goroutines' output may be interleaved,
because goroutines are executed concurrently by the Go runtime.

```bash
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done
```
