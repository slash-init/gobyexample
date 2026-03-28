# Go Multiple Return Values

## Key Points

- **Definition:** Functions in Go can return multiple values, often used for returning results and errors.
- **Blank Identifier:** Use `_` to ignore specific return values.

---

## Examples

### Function with Multiple Returns

```go
func vals() (int, int) {
    return 3, 7
}
```

### Using Multiple Return Values

```go
a, b := vals()
fmt.Println(a)  // Output: 3
fmt.Println(b)  // Output: 7
```

### Ignoring a Return Value

```go
_, c := vals()
fmt.Println(c)  // Output: 7
```

---

## Summary

- Functions can return multiple values by specifying multiple types in the return signature.
- Use multiple assignment to capture all return values.
- Use `_` to ignore unwanted return values.

---

## Practice

Try these examples:

```go
// Write a function that returns a string and an int
func info() (string, int) {
    return "age", 30
}

// Call the function
name, age := info()
fmt.Println(name, age)  // Output: age 30
```
