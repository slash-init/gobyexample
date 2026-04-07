# Go by Example: Channel Synchronization

We can use channels to synchronize execution across goroutines. Here is an
example of using a blocking receive to wait for a goroutine to finish. When
waiting for multiple goroutines to finish, you may prefer to use a
`sync.WaitGroup`.

```go
package main

import (
    "fmt"
    "time"
)
```

This is the function we'll run in a goroutine. The `done` channel will be
used to notify another goroutine that this function's work is done.

```go
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}
```

Start a worker goroutine, giving it the channel to notify on. Then block until
we receive a notification from the worker on the channel.

```go
func main() {
    done := make(chan bool, 1)
    go worker(done)

    <-done
}
```

If you removed the `<-done` line from this program, the program could exit
before the worker finished its work.

```bash
$ go run channel-synchronization.go
working...done
```
