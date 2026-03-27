# Go Maps

## Key Points

- **Definition:** Maps are Go’s built-in associative data type (similar to dictionaries in Python).
- **Zero Value:** An uninitialized map is `nil` and has no keys.
- **Dynamic Size:** Maps grow dynamically as you add key-value pairs.
- **Key Operations:** Add, retrieve, delete, and check existence of keys.

---

## Examples

### Creating and Initializing Maps

```go
// Create an empty map
m := make(map[string]int)

// Declare and initialize a map
n := map[string]int{"foo": 1, "bar": 2}
fmt.Println("map:", n)  // Output: map: map[bar:2 foo:1]
```

### Adding and Retrieving Values

```go
m["k1"] = 7
m["k2"] = 13
fmt.Println("map:", m)  // Output: map: map[k1:7 k2:13]

v1 := m["k1"]
fmt.Println("v1:", v1)  // Output: v1: 7

// Accessing a non-existent key returns the zero value
v3 := m["k3"]
fmt.Println("v3:", v3)  // Output: v3: 0
```

### Checking Key Existence

```go
_, prs := m["k2"]
fmt.Println("prs:", prs)  // Output: prs: true

_, prs = m["k3"]
fmt.Println("prs:", prs)  // Output: prs: false
```

### Deleting Keys

```go
delete(m, "k2")
fmt.Println("map:", m)  // Output: map: map[k1:7]
```

### Clearing a Map

```go
clear(m)
fmt.Println("map:", m)  // Output: map: map[]
```

### Comparing Maps

```go
n2 := map[string]int{"foo": 1, "bar": 2}
if maps.Equal(n, n2) {
    fmt.Println("n == n2")  // Output: n == n2
}
```

---

## Summary

- Use `make` to create an empty map.
- Use `map[key] = value` to add or update key-value pairs.
- Use `delete(map, key)` to remove a key.
- Use `clear(map)` to remove all keys.
- Use `_, ok := map[key]` to check if a key exists.
- Maps are not safe for concurrent use; use synchronization for shared maps.

---

## Practice

Try these examples:

```go
// Create and manipulate a map
m := make(map[string]int)
m["a"] = 10
m["b"] = 20
fmt.Println(m)

delete(m, "a")
fmt.Println(m)

// Check existence
_, ok := m["b"]
fmt.Println("b exists:", ok)

// Compare maps
m1 := map[string]int{"x": 1, "y": 2}
m2 := map[string]int{"x": 1, "y": 2}
fmt.Println(maps.Equal(m1, m2))
```
