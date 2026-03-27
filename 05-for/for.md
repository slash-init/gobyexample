# Go Loops: `for`

## Overview

In Go, the `for` loop is the only looping construct. It can be used in various ways to achieve different looping behaviors. This lesson covers:

- Basic `for` loops
- Conditional loops
- Infinite loops
- Using `range` with `for`
- Controlling loops with `break` and `continue`

---

## Basic `for` Loop

The most common form of a `for` loop includes an initializer, condition, and post statement:

```go
for j := 0; j < 3; j++ {
    fmt.Println(j)
}
```

**Key Points:**

- `j := 0` initializes the loop variable
- `j < 3` is the condition to continue looping
- `j++` increments the loop variable after each iteration
- Output: `0`, `1`, `2`

---

## Conditional Loops

A `for` loop can also act like a `while` loop by omitting the initializer and post statement:

```go
i := 1
for i <= 3 {
    fmt.Println(i)
    i++
}
```

**Key Points:**

- Only the condition (`i <= 3`) is specified
- Output: `1`, `2`, `3`

---

## Infinite Loops

An infinite loop is created by omitting all three components:

```go
for {
    fmt.Println("loop")
    break
}
```

**Key Points:**

- Runs indefinitely unless explicitly stopped
- Use `break` to exit the loop
- Output: `loop`

---

## Using `range` with `for`

The `range` keyword is used to iterate over elements in a collection (e.g., slices, maps, strings):

```go
for i := range []int{0, 1, 2} {
    fmt.Println("range", i)
}
```

**Key Points:**

- `range` provides the index of each element
- Output: `range 0`, `range 1`, `range 2`

---

## Controlling Loops

### `break`

Exits the loop immediately:

```go
for {
    fmt.Println("loop")
    break
}
```

### `continue`

Skips the current iteration and moves to the next:

```go
for n := range []int{0, 1, 2, 3, 4, 5} {
    if n%2 == 0 {
        continue
    }
    fmt.Println(n)
}
```

**Key Points:**

- `continue` skips even numbers
- Output: `1`, `3`, `5`

---

## Summary

| Loop Type         | Example Code                | Description                    |
| ----------------- | --------------------------- | ------------------------------ |
| Basic `for`       | `for j := 0; j < 3; j++ {}` | Standard loop with initializer |
| Conditional `for` | `for i <= 3 {}`             | Acts like a `while` loop       |
| Infinite `for`    | `for {}`                    | Runs indefinitely              |
| `range` loop      | `for i := range []int{}`    | Iterates over collections      |
| `break`           | `break` inside `for`        | Exits the loop                 |
| `continue`        | `continue` inside `for`     | Skips current iteration        |

---

## Practice

Try these examples:

```go
// Basic loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// Conditional loop
x := 1
for x <= 5 {
    fmt.Println(x)
    x++
}

// Infinite loop
for {
    fmt.Println("Running forever")
    break
}

// Using range
for idx, val := range []string{"a", "b", "c"} {
    fmt.Println(idx, val)
}

// Skipping even numbers
for n := 0; n < 10; n++ {
    if n%2 == 0 {
        continue
    }
    fmt.Println(n)
}
```
