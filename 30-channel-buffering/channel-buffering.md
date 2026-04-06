# Go by Example: Channel Buffering

By default, channels are unbuffered, meaning that they only accept sends
(`chan <-`) when there is a corresponding receive (`<-chan`) ready to receive
the sent value. Buffered channels accept a limited number of values without a
corresponding receiver.

```go
package main

import "fmt"
```

Here we make a channel of strings buffering up to 2 values.

```go
func main() {
    messages := make(chan string, 2)
```

Because this channel is buffered, we can send these values into the channel
without a corresponding concurrent receive.

```go
    messages <- "buffered"
    messages <- "channel"
```

Later we can receive these two values as usual.

```go
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
```

```bash
$ go run channel-buffering.go
buffered
channel
```
