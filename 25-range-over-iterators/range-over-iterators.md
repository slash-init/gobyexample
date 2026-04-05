# Go by Example: Range over Iterators

Starting with version 1.23, Go has added support for iterators, which lets us
range over pretty much anything.

```go
package main

import (
    "fmt"
    "iter"
    "slices"
    "strings"
)
```

Let’s look at the `List` type from the previous example again. In that example
we had an `AllElements` method that returned a slice of all elements in the
list. With Go iterators, we can do it better.

```go
type List[T any] struct {
    head, tail *element[T]
}

type element[T any] struct {
    next *element[T]
    val  T
}

func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}
```

`All` returns an iterator, which in Go is a function with a special signature.
`iter.Seq[T]` is essentially a function that receives a callback:
`func(yield func(T) bool)`.

The iterator produces values by calling `yield(value)`.
If `yield` returns `true`, iteration continues.
If `yield` returns `false`, iteration stops.

```go
func (lst *List[T]) All() iter.Seq[T] {
    return func(yield func(T) bool) {
        for e := lst.head; e != nil; e = e.next {
            if !yield(e.val) {
                return
            }
        }
    }
}
```

The iterator function takes another function as a parameter, called `yield` by
convention (the name itself can be anything). Here we walk the linked list and
call `yield` for each element. The `if !yield(...) { return }` check is what
allows early termination.

Iteration doesn’t require an underlying data structure, and doesn’t even have
to be finite. Here’s an iterator over Fibonacci numbers: it keeps generating
values until the caller signals stop.

```go
func genFib() iter.Seq[int] {
    return func(yield func(int) bool) {
        a, b := 0, 1

        for {
            if !yield(a) {
                return
            }
            a, b = b, a+b
        }
    }
}
```

```go
func main() {
    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)

    // Since List.All returns an iterator, we can use it in a regular range loop.
    for e := range lst.All() {
        fmt.Println(e)
    }

    // Packages like slices can collect iterator values into a slice.
    all := slices.Collect(lst.All())
    fmt.Println("all:", all)

    // Standard library packages now expose iterator helpers too.
    for part := range strings.SplitSeq("go-by-example", "-") {
        fmt.Printf("part: %s\n", part)
    }

    for n := range genFib() {
        if n >= 10 {
            break
        }
        fmt.Println(n)
    }
}
```

When a `for range` loop over an iterator hits `break` (or returns from the
surrounding function), the runtime makes the internal `yield` return `false`,
and the iterator function exits.

```bash
$ go run range-over-iterators.go
10
13
23
all: [10 13 23]
part: go
part: by
part: example
0
1
1
2
3
5
8
```
