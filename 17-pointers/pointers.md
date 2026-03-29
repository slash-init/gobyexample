# Go by Example: Pointers

Go supports pointers, allowing you to pass references to values and modify them directly within your program. This can be useful when you want a function to modify the value of a variable outside its own scope.

### Example Code

```go
package main

import "fmt"

// zeroval has an int parameter, so arguments will be passed to it by value.
// It will get a copy of ival, and changes made here won't affect the original variable.
func zeroval(ival int) {
    ival = 0
}

// zeroptr has a *int parameter, meaning it takes a pointer to an int.
// This allows it to modify the value at the memory address the pointer refers to.
func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("after zeroval:", i) // i is still 1

    zeroptr(&i)
    fmt.Println("after zeroptr:", i) // i is now 0
}
```

### Explanation

- **`zeroval` (Pass by Value):**
  - The `zeroval` function takes an `int` as its parameter. When you pass a variable to it, the function receives a copy of the value, not the original variable. Any changes made to the parameter inside the function do not affect the original variable.

- **`zeroptr` (Pass by Reference):**
  - The `zeroptr` function takes a pointer to an `int` as its parameter. By dereferencing the pointer (using `*iptr`), the function can modify the value at the memory address the pointer refers to. This allows the function to change the original variable.

Pointers are a powerful feature in Go that enable you to work with memory addresses directly, providing more control over how data is passed and modified in your programs.
