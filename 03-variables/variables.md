# Go Variables

## Overview

In Go, variables are explicitly declared and used by the compiler to check type-correctness of function calls. This lesson covers:

- Variable declaration with `var`
- Type inference
- Multiple variable declarations
- Zero values
- Shorthand syntax with `:=`

---

## Variable Declaration

### Basic Declaration with `var`

The `var` keyword declares one or more variables.

**Key Points:**

- Go infers the type automatically from the assigned value
- Variables hold values of their declared type

### Type Inference

Go automatically infers types from initial values:

**What Go Infers:**

- `"text"` → `string`
- `42` → `int`
- `3.14` → `float64`
- `true` / `false` → `bool`

---

## Multiple Variable Declarations

You can declare multiple variables at once:

```go
var b, c int = 1, 2
```

**Syntax:**

```go
var name1, name2 type = value1, value2
```

**Key Points:**

- All variables must be of the same type
- Assign corresponding values in order

---

## Zero Values

Variables declared without initialization receive a **zero value** (default value for that type).

**Zero Values by Type:**
| Type | Zero Value |
|------|------------|
| `int`, `float64` | `0` |
| `string` | `""` (empty string) |
| `bool` | `false` |
| pointers | `nil` |

---

## Shorthand Declaration (`:=`)

The `:=` syntax is shorthand for declaring and initializing a variable. For example, `f := "apple"` is shorthand for `var f string = "apple"`.

**Important:**

- `:=` can **only** be used inside functions
- Cannot be used at package level
- Type is always inferred

---

## Comparison

| Syntax                | Usage              | Type Inference | Scope               |
| --------------------- | ------------------ | -------------- | ------------------- |
| `var x = value`       | Full declaration   | Yes            | Package or function |
| `var x int`           | With explicit type | No             | Package or function |
| `var x, y int = 1, 2` | Multiple variables | No             | Package or function |
| `x := value`          | Shorthand          | Yes            | Functions only      |

---

## Key Takeaways

1. Use `var` for explicit variable declarations
2. Go infers types from assigned values
3. Variables without initialization get zero values
4. Use `:=` inside functions for concise syntax
5. Every variable has a type in Go (strictly typed language)

---

## Practice

Try these variations:

```go
var name = "John"           // String inference
var age int = 25             // Explicit type
var x, y = 10, 20            // Multiple with inference
z := 3.14                    // Shorthand
var empty string             // Zero value
```
