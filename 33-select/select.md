# Go by Example: Select

Go's `select` lets you wait on multiple channel operations. Combining
goroutines and channels with `select` is a powerful feature of Go.

```go
package main

import (
    "fmt"
    "time"
)
```

For our example we'll select across two channels.

```go
func main() {
    c1 := make(chan string)
    c2 := make(chan string)
```

Each channel will receive a value after some amount of time, to simulate
blocking operations executing in concurrent goroutines.

```go
    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()
```

We'll use `select` to await both of these values simultaneously, printing each
one as it arrives.

```go
    for range 2 {
        select {
        case msg1 := <-c1:
            fmt.Println("recieved", msg1)
        case msg2 := <-c2:
            fmt.Println("recieved", msg2)
        }
    }
}
```

We receive the values from `c1` and `c2` as they complete, roughly one second
apart.

```bash
$ time go run select.go
recieved one
recieved two

real    0m2.xxxs
```

The total execution time is only about 2 seconds because the sleeps execute
concurrently.
