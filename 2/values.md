# Go Values & Basic Operations

## Overview

Go supports various basic data types and operations. This lesson covers:

- String concatenation
- Arithmetic operations
- Boolean logic operations

---

## String Operations

### Concatenation

Strings can be concatenated using the `+` operator:

```go
fmt.Println("go" + "lang")  // Output: golang
```

**Key Points:**

- The `+` operator joins two strings together
- Result: `golang`

---

## Arithmetic Operations

### Integer Arithmetic

```go
fmt.Println("1+1 =", 1+1)  // Output: 1+1 = 2
```

**Supported Operations:**

- Addition: `+`
- Subtraction: `-`
- Multiplication: `*`
- Division: `/`
- Modulo: `%`

### Float Division

```go
fmt.Println("7.0/3.0 =", 7.0/3.0)  // Output: 7.0/3.0 = 2.3333333333333335
```

**Key Differences:**
| Operation | Result |
|-----------|--------|
| `7 / 3` (integer division) | `2` (rounds down) |
| `7.0 / 3.0` (float division) | `2.333...` (decimal result) |

---

## Boolean Operations

### Logical AND (`&&`)

```go
fmt.Println(true && false)  // Output: false
```

- Returns `true` only if BOTH operands are `true`
- Evaluation:
  - `true && true` → `true`
  - `true && false` → `false`
  - `false && false` → `false`

### Logical OR (`||`)

```go
fmt.Println(true || false)  // Output: true
```

- Returns `true` if AT LEAST ONE operand is `true`
- Evaluation:
  - `true || true` → `true`
  - `true || false` → `true`
  - `false || false` → `false`

### Logical NOT (`!`)

```go
fmt.Println(!true)  // Output: false
```

- Inverts the boolean value
- `!true` → `false`
- `!false` → `true`

---

## Truth Table

| A   | B   | A && B | A \|\| B | !A  |
| --- | --- | ------ | -------- | --- |
| T   | T   | T      | T        | F   |
| T   | F   | F      | T        | F   |
| F   | T   | F      | T        | T   |
| F   | F   | F      | F        | T   |

---

## Key Takeaways

1. **Strings** can be combined with `+`
2. **Integer Division** (`/`) truncates decimals
3. **Float Division** (`/`) preserves decimals
4. **Boolean Operations** (`&&`, `||`, `!`) are fundamental for conditional logic
5. Use floating-point numbers when precision matters

---

## Practice

Try modifying the program:

```go
// Experiment with different values
fmt.Println("2+3 =", 2+3)
fmt.Println("10/3 =", 10/3)
fmt.Println("10.0/3.0 =", 10.0/3.0)
fmt.Println(false && true)
fmt.Println(!false)
```
