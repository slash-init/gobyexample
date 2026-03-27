# Go Constants

## Overview

Go supports constants of character, string, boolean, and numeric values. This lesson covers:

- Declaring constants with `const`
- Constants at package and function level
- Type inference for constants
- Arbitrary precision arithmetic
- Type conversions

---

## Declaring Constants

The `const` keyword declares a constant value:

```go
const s string = "constant"
fmt.Println(s)  // Output: constant
```

**Key Points:**

- Constants must have a compile-time value
- Cannot be changed after declaration
- Can declare with explicit type

---

## Constants at Package vs Function Level

Constants can be declared at package level (outside functions):

```go
const s string = "constant"  // Package level
```

Or inside a function body:

```go
func main() {
    const n = 5000000000  // Function level
}
```

**Scope:**

- Package-level constants: accessible to entire package
- Function-level constants: accessible only within that function

---

## Type Inference

Constants can be declared without explicit type, and Go infers the type:

```go
const n = 5000000000
```

The type is inferred based on the value assigned.

---

## Arbitrary Precision Arithmetic

Constant expressions perform arithmetic with arbitrary precision:

```go
const d = 3e20 / n
fmt.Println(d)  // Output: 60000
```

**Key Points:**

- Constants can represent very large numbers
- Arithmetic is done at compile-time with arbitrary precision
- Scientific notation: `3e20` means 3 × 10^20

---

## Untyped Constants

A numeric constant has no type until it's given one.

**Ways to enforce a type:**

### 1. Explicit Conversion

```go
fmt.Println(int64(d))  // Convert to int64
```

### 2. Using in Typed Context

A number can be given a type by using it in a context that requires one:

```go
fmt.Println(math.Sin(n))
```

Here, `math.Sin` expects a `float64`, so `n` is automatically converted.

---

## Constants vs Variables

| Aspect    | Constants                  | Variables            |
| --------- | -------------------------- | -------------------- |
| Checked   | Compile-time               | Runtime              |
| Changed   | Cannot be changed          | Can be modified      |
| Type      | Optional (inferred)        | Required or inferred |
| Precision | Arbitrary                  | Limited to type      |
| Speed     | Fast (no runtime overhead) | Standard             |

---

## Common Constant Types

```go
const (
    s string = "hello"      // String constant
    n int = 42              // Integer constant
    f float64 = 3.14        // Float constant
    b bool = true           // Boolean constant
)
```

---

## Key Takeaways

1. Use `const` to declare immutable values
2. Can be declared at package or function level
3. Numeric constants have arbitrary precision
4. Type is inferred or can be explicitly set
5. Type is enforced when used in typed contexts
6. Constants provide compile-time safety

---

## Practice

Try these:

```go
const maxSize = 1000
const pi = 3.14159
const msg = "Hello, Constants!"
const x, y = 10, 20
```
