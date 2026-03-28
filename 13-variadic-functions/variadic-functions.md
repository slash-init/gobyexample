# Go Variadic Functions

## Key Points

- **Definition:** Variadic functions accept any number of trailing arguments.
- **Common Example:** `fmt.Println` is a variadic function.
- **Argument Type:** Inside the function, variadic arguments are treated as a slice.
- **Usage:** Use `func(slice...)` to pass a slice to a variadic function.

---

## Examples

### Variadic Function Definition

```go
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}
```

### Calling a Variadic Function

```go
sum(1, 2)          // Output: [1 2] 3
sum(1, 2, 3)       // Output: [1 2 3] 6
```

### Passing a Slice to a Variadic Function

```go
nums := []int{1, 2, 3, 4}
sum(nums...)       // Output: [1 2 3 4] 10
```

---

## Summary

- Variadic functions use `...` to accept any number of arguments.
- Arguments are treated as a slice within the function.
- Use `slice...` to pass a slice to a variadic function.

---

## Practice

Try these examples:

```go
// Write a variadic function to calculate the product of numbers
func product(nums ...int) int {
    result := 1
    for _, num := range nums {
        result *= num
    }
    return result
}

// Call the function
fmt.Println(product(2, 3, 4))  // Output: 24
```
