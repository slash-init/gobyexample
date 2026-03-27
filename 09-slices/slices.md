# Go Slices

## Key Points

- **Dynamic Size:** Slices are dynamically-sized, unlike arrays.
- **Zero Value:** A nil slice has zero length and capacity.
- **Backed by Arrays:** Slices are views over arrays.
- **Capacity:** Use `cap()` to get the capacity of a slice.
- **Copy and Append:** Use `copy()` to duplicate slices and `append()` to grow them.

---

## Examples

### Declaring and Initializing Slices

```go
var s []string
fmt.Println("uninit:", s, s == nil, len(s) == 0)  // Output: uninit: [] true true

s = make([]string, 3)
fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))  // Output: emp: [  ] len: 3 cap: 3
```

### Setting and Getting Values

```go
s[0] = "a"
s[1] = "b"
s[2] = "c"
fmt.Println("set:", s)  // Output: set: [a b c]
fmt.Println("get:", s[2])  // Output: get: c
```

### Appending to Slices

```go
s = append(s, "d")
s = append(s, "e", "f")
fmt.Println("apd:", s)  // Output: apd: [a b c d e f]
```

### Copying Slices

```go
c := make([]string, len(s))
copy(c, s)
fmt.Println("cpy:", c)  // Output: cpy: [a b c d e f]
```

### Slicing

```go
l := s[2:5]
fmt.Println("sl1:", l)  // Output: sl1: [c d e]

l = s[:5]
fmt.Println("sl2:", l)  // Output: sl2: [a b c d e]

l = s[2:]
fmt.Println("sl3:", l)  // Output: sl3: [c d e f]
```

### Multi-dimensional Slices

```go
twoD := make([][]int, 3)
for i := 0; i < 3; i++ {
    innerLen := i + 1
    twoD[i] = make([]int, innerLen)
    for j := 0; j < innerLen; j++ {
        twoD[i][j] = i + j
    }
}
fmt.Println("2d:", twoD)  // Output: 2d: [[0] [1 2] [2 3 4]]
```

### Comparing Slices

```go
t := []string{"g", "h", "i"}
t2 := []string{"g", "h", "i"}
if slices.Equal(t, t2) {
    fmt.Println("t == t2")  // Output: t == t2
}
```

---

## Differences Between Arrays and Slices

| Feature             | Arrays                 | Slices                           |
| ------------------- | ---------------------- | -------------------------------- |
| **Size**            | Fixed at compile-time  | Dynamic, can grow with `append`  |
| **Zero Value**      | Zero-valued elements   | `nil` (zero length and capacity) |
| **Length**          | `len()` gives the size | `len()` gives current length     |
| **Capacity**        | Not applicable         | `cap()` gives total capacity     |
| **Initialization**  | `[5]int{}`             | `make([]int, size)`              |
| **Copy**            | Not applicable         | Use `copy()`                     |
| **Underlying Data** | Independent            | Backed by an array               |

---

## Summary

- Slices are more flexible than arrays and are the preferred way to work with collections in Go.
- Use `append()` to grow slices dynamically.
- Use `copy()` to duplicate slices.
- Slices are references to arrays, so modifying a slice affects the underlying array.

---

## Practice

Try these examples:

```go
// Create and append
nums := []int{1, 2, 3}
nums = append(nums, 4, 5)
fmt.Println(nums)

// Copy slices
src := []int{1, 2, 3}
dst := make([]int, len(src))
copy(dst, src)
fmt.Println(dst)

// Multi-dimensional slices
twoD := [][]int{
    {1, 2},
    {3, 4, 5},
}
fmt.Println(twoD)
```

### Slicing Explained

Slicing allows you to create a new slice by referencing a portion of an existing slice or array. The syntax is:

```go
slice[start:end]
```

- `start`: The index where the slice begins (inclusive).
- `end`: The index where the slice ends (exclusive).
- If `start` is omitted, it defaults to `0`.
- If `end` is omitted, it defaults to the length of the slice.

#### Examples

```go
s := []string{"a", "b", "c", "d", "e"}

// Slice from index 2 to 4 (exclusive)
l := s[2:5]
fmt.Println("sl1:", l)  // Output: sl1: [c d e]

// Slice from the start to index 5 (exclusive)
l = s[:5]
fmt.Println("sl2:", l)  // Output: sl2: [a b c d e]

// Slice from index 2 to the end
l = s[2:]
fmt.Println("sl3:", l)  // Output: sl3: [c d e]
```

#### Key Points

- **Slicing does not copy data:** The new slice references the same underlying array. Modifying the new slice affects the original slice/array.
- **Capacity:** The capacity of the new slice is determined by the length of the underlying array, starting from the `start` index.

#### Example Showing Shared Data

```go
s := []int{1, 2, 3, 4, 5}
l := s[1:4]
l[0] = 100
fmt.Println("s:", s)  // Output: s: [1 100 3 4 5]
fmt.Println("l:", l)  // Output: l: [100 3 4]
```

In this example, modifying `l` also modifies `s` because they share the same underlying array.
