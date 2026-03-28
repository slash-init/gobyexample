# Go Recursion

## Key Points

- **Definition:** Recursion is when a function calls itself to solve smaller instances of a problem.
- **Base Case:** Every recursive function must have a base case to stop the recursion.
- **Anonymous Recursive Functions:** Anonymous functions can also be recursive, but require a `var` declaration to reference the function.

---

## Examples

### Factorial Function (Classic Recursion)

```go
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
    fmt.Println(fact(7))  // Output: 5040
}
```

### Fibonacci Sequence (Anonymous Recursive Function)

```go
func main() {
    var fib func(n int) int

    fib = func(n int) int {
        if n < 2 {
            return n
        }
        return fib(n-1) + fib(n-2)
    }

    fmt.Println(fib(7))  // Output: 13
}
```

---

## Explanation

1. **Factorial Function:**
   - The `fact` function calculates the factorial of a number.
   - It multiplies `n` by the result of `fact(n-1)` until `n` reaches 0 (base case).

2. **Anonymous Recursive Function:**
   - The `fib` function calculates the Fibonacci sequence.
   - It uses a `var` declaration to reference itself within its definition.
   - The base case handles `n < 2`, and the recursive case sums `fib(n-1)` and `fib(n-2)`.

---

## Summary

- Recursive functions solve problems by breaking them into smaller subproblems.
- Always include a base case to prevent infinite recursion.
- Anonymous functions can be recursive with a `var` declaration.

---

## Practice

Try these examples:

```go
// Write a recursive function to calculate the sum of numbers from 1 to n
func sum(n int) int {
    if n == 0 {
        return 0
    }
    return n + sum(n-1)
}

func main() {
    fmt.Println(sum(10))  // Output: 55
}
```
