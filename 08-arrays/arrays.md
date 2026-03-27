# Go Arrays

## Key Points

- **Fixed Size:** Arrays in Go have a fixed size defined at declaration.
- **Zero Values:** Elements are zero-valued by default (e.g., `0` for `int`).
- **Length:** Use `len()` to get the array length.
- **Multi-dimensional Arrays:** Supported with nested arrays.

---

## Examples

### Declaring and Accessing Arrays

```go
var a [5]int
fmt.Println("emp:", a)  // Output: emp: [0 0 0 0 0]

a[4] = 100
fmt.Println("set:", a)  // Output: set: [0 0 0 0 100]
fmt.Println("get:", a[4])  // Output: get: 100

fmt.Println("len:", len(a))  // Output: len: 5
```

### Array Initialization

```go
b := [5]int{1, 2, 3, 4, 5}
fmt.Println("dcl:", b)  // Output: dcl: [1 2 3 4 5]

b = [...]int{1, 2, 3, 4, 5}  // Compiler infers size
fmt.Println("dcl:", b)  // Output: dcl: [1 2 3 4 5]

b = [...]int{100, 3: 400, 500}  // Indexed initialization
fmt.Println("idx:", b)  // Output: idx: [100 0 0 400 500]
```

### Multi-dimensional Arrays

```go
var twoD [2][3]int
for i := 0; i < 2; i++ {
    for j := 0; j < 3; j++ {
        twoD[i][j] = i + j
    }
}
fmt.Println("2d:", twoD)  // Output: 2d: [[0 1 2] [1 2 3]]

twoD = [2][3]int{
    {1, 2, 3},
    {1, 2, 3},
}
fmt.Println("2d:", twoD)  // Output: 2d: [[1 2 3] [1 2 3]]
```

---

## Summary

- Arrays are fixed-size collections of elements of the same type.
- Use `[...]` to let the compiler infer the size.
- Multi-dimensional arrays are supported.
- Default values depend on the element type.

---

## Practice

Try these examples:

```go
// Declare and initialize
var nums [3]int
nums[0] = 10
nums[1] = 20
nums[2] = 30
fmt.Println(nums)

// Compiler-inferred size
letters := [...]string{"a", "b", "c"}
fmt.Println(letters)

// Multi-dimensional array
matrix := [2][2]int{
    {1, 2},
    {3, 4},
}
fmt.Println(matrix)
```
