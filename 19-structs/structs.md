# Go by Example: Structs

Structs in Go are a way to group together data fields. They are particularly useful for creating records or representing entities in your program. Structs are typed collections of fields, and each field has a name and a type.

To define a struct, you use the `type` keyword followed by the struct name and its fields. Here’s an example:

```go
type person struct {
    name string
    age  int
}
```

This defines a `person` struct with two fields: `name` (a string) and `age` (an integer).

## Creating Structs

You can create a new struct by specifying its field values. For example:

```go
fmt.Println(person{"Bob", 20})
```

This creates a `person` struct with the name "Bob" and age 20. Alternatively, you can name the fields explicitly:

```go
fmt.Println(person{name: "Alice", age: 30})
```

If you omit a field, it will be set to its zero value. For instance:

```go
fmt.Println(person{name: "Fred"})
```

Here, the `age` field is omitted, so it defaults to `0`.

## Struct Pointers

You can create a pointer to a struct using the `&` operator. For example:

```go
fmt.Println(&person{name: "Ann", age: 40})
```

This creates a pointer to a `person` struct. Go automatically dereferences struct pointers when accessing fields, so you can use the dot operator as if it were a regular struct.

## Constructor Functions

It’s common to encapsulate struct creation in a constructor function. Here’s an example:

```go
func newPerson(name string) *person {
    p := person{name: name}
    p.age = 42
    return &p
}
```

This function creates a new `person` struct, sets its `age` to 42, and returns a pointer to it. You can use it like this:

```go
fmt.Println(newPerson("Jon"))
```

## Accessing and Modifying Fields

You can access struct fields using the dot operator. For example:

```go
s := person{name: "Sean", age: 50}
fmt.Println(s.name)
```

You can also modify fields directly:

```go
s.age = 51
fmt.Println(s.age)
```

## Anonymous Structs

If you only need a struct for a single use, you can define an anonymous struct. This is often used in table-driven tests. For example:

```go
dog := struct {
    name   string
    isGood bool
}{
    "Rex",
    true,
}
fmt.Println(dog)
```

This creates a struct with fields `name` and `isGood`, initializes it, and prints it.

## Full Example

Here’s the complete program:

```go
package main

import "fmt"

type person struct {
    name string
    age  int
}

func newPerson(name string) *person {
    p := person{name: name}
    p.age = 42
    return &p
}

func main() {
    fmt.Println(person{"Bob", 20})
    fmt.Println(person{name: "Alice", age: 30})
    fmt.Println(person{name: "Fred"})
    fmt.Println(&person{name: "Ann", age: 40})
    fmt.Println(newPerson("Jon"))

    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    sp := &s
    fmt.Println(sp.age)

    sp.age = 51
    fmt.Println(sp.age)

    dog := struct {
        name   string
        isGood bool
    }{
        "Rex",
        true,
    }
    fmt.Println(dog)
}
```

When you run this program, the output will be:

```
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 40}
&{Jon 42}
Sean
50
51
{Rex true}
```
