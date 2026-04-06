# Go by Example: Channels

Channels are the pipes that connect concurrent goroutines. You can send values
into channels from one goroutine and receive those values in another goroutine.

```go
package main

import "fmt"
```

Create a new channel with `make(chan val-type)`. Channels are typed by the
values they convey.

```go
func main() {
    messages := make(chan string)
```

Send a value into a channel using the `channel <-` syntax. Here we send
`"ping"` to the `messages` channel from a new goroutine.

```go
    go func() { messages <- "ping" }()
```

The `<-channel` syntax receives a value from the channel. Here we receive the
`"ping"` message and print it.

```go
    msg := <-messages
    fmt.Println(msg)
}
```

When we run the program, the `"ping"` message is passed from one goroutine to
another via our channel.

```bash
$ go run channels.go
ping
```

By default, sends and receives block until both the sender and receiver are
ready. This property lets us wait for `"ping"` without using other
synchronization primitives.
