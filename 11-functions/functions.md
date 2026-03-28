# Go Functions

## Key Points

- **Definition:** Functions are reusable blocks of code that take inputs, perform operations, and return outputs.
- **Explicit Returns:** Go requires explicit `return` statements; it does not automatically return the last expression.
- **Parameter Types:** Consecutive parameters of the same type can share the type declaration.

---

## Examples

### Basic Function

```go
func plus(a int, b int) int {
    return a + b
}
```

### Function with Shared Parameter Types

```go
func plusPlus(a, b, c int) int {
    return a + b + c
}
```

### Calling Functions

```go
res := plus(1, 2)
fmt.Println("1+2 =", res)  // Output: 1+2 = 3

res = plusPlus(1, 2, 3)
fmt.Println("1+2+3 =", res)  // Output: 1+2+3 = 6
```

---

## Summary

- Functions are defined using the `func` keyword.
- Parameters and return types must be explicitly declared.
- Use functions to encapsulate reusable logic.

---

## Practice

Try these examples:

```go
// Write a function to multiply two numbers
func multiply(a, b int) int {
    return a * b
}

// Call the function
fmt.Println(multiply(3, 4))  // Output: 12
```
