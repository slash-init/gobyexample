# Go Conditional Statements: `if` and `else`

## Key Points

- **Basic Syntax:**
  ```go
  if condition {
      // code
  } else {
      // code
  }
  ```
- **No parentheses** around conditions, but braces `{}` are mandatory.
- **Logical Operators:** `&&` (AND), `||` (OR), `!` (NOT).
- **Variable Declaration:** You can declare variables in the `if` statement:
  ```go
  if num := 10; num > 5 {
      fmt.Println("Greater than 5")
  }
  ```
- **No Ternary Operator:** Use full `if-else` for conditional assignments.

---

## Examples

### Basic `if-else`

```go
if 7%2 == 0 {
    fmt.Println("Even")
} else {
    fmt.Println("Odd")
}
```

### `if` Without `else`

```go
if 8%4 == 0 {
    fmt.Println("Divisible by 4")
}
```

### Logical Operators

```go
if x > 0 && y > 0 {
    fmt.Println("Both positive")
}
```

### Variable Declaration

```go
if num := 9; num < 10 {
    fmt.Println("Single digit")
}
```

---

## Summary

- Use `if` for branching logic.
- Combine conditions with `&&`, `||`, and `!`.
- Declare variables inline for concise code.
- Always use braces `{}` for clarity.
