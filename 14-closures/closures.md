# Go Closures

## Key Points

- **Definition:** A closure is a function that captures variables from its surrounding scope.
- **Anonymous Functions:** Functions without a name, often used inline.
- **Stateful Functions:** Closures can maintain state between calls by capturing variables.

---

## Examples

### Closure Example

```go
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
```

### Using a Closure

```go
func main() {
    nextInt := intSeq()

    fmt.Println(nextInt())  // Output: 1
    fmt.Println(nextInt())  // Output: 2
    fmt.Println(nextInt())  // Output: 3

    newInts := intSeq()
    fmt.Println(newInts())  // Output: 1
}
```

---

## Explanation

1. **How it Works:**
   - The `intSeq` function returns an anonymous function.
   - The anonymous function captures the variable `i` from `intSeq`'s scope.
   - Each call to the returned function updates and returns the captured `i`.

2. **Stateful Behavior:**
   - Each call to `intSeq` creates a new `i` variable.
   - The state of `i` is unique to each closure.

---

## Summary

- Closures allow functions to capture and use variables from their surrounding scope.
- Each closure maintains its own state.
- Anonymous functions are commonly used to define closures.

---

## Practice

Try these examples:

```go
// Create a closure that multiplies numbers
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

// Use the closure
double := multiplier(2)
fmt.Println(double(5))  // Output: 10

triple := multiplier(3)
fmt.Println(triple(5))  // Output: 15
```
