# Go Conditional Statements: `switch`

## Key Points

- **Basic Syntax:**
  ```go
  switch value {
  case condition1:
      // code
  case condition2:
      // code
  default:
      // code
  }
  ```
- **Multiple Cases:** Use commas to group cases.
- **Expressionless Switch:** Acts like `if-else`.
- **Type Switch:** Matches types of interface values.

---

## Examples

### Basic `switch`

```go
i := 2
switch i {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
case 3:
    fmt.Println("three")
}
```

### Multiple Cases

```go
switch time.Now().Weekday() {
case time.Saturday, time.Sunday:
    fmt.Println("It's the weekend")
default:
    fmt.Println("It's a weekday")
}
```

### Expressionless `switch`

```go
t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("It's before noon")
default:
    fmt.Println("It's after noon")
}
```

### Type Switch

```go
whatAmI := func(i interface{}) {
    switch t := i.(type) {
    case bool:
        fmt.Println("I'm a bool")
    case int:
        fmt.Println("I'm an int")
    default:
        fmt.Printf("Don't know type %T\n", t)
    }
}
whatAmI(true)
whatAmI(1)
whatAmI("hey")
```

---

## Summary

- Use `switch` for cleaner multi-branch logic.
- Combine cases with commas.
- Use expressionless `switch` for flexible conditions.
- Type switches are useful for interface type checks.
