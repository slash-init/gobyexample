# Go by Example: Range over Built-in Types

The `range` keyword in Go is a powerful tool for iterating over elements in various data structures. It simplifies the process of accessing elements in slices, arrays, maps, and strings. Below are some examples of how `range` can be used effectively:

### Iterating Over Slices and Arrays

You can use `range` to iterate over slices and arrays. It provides both the index and the value of each element. If you only need the value, you can ignore the index by using the blank identifier `_`.

Example:

```go
nums := []int{2, 3, 4}
sum := 0
for _, num := range nums {
    sum += num
}
fmt.Println("Sum:", sum)
```

If you need the index, you can use it as follows:

```go
for i, num := range nums {
    fmt.Println("Index:", i, "Value:", num)
}
```

### Iterating Over Maps

When iterating over a map, `range` provides both the key and value for each entry.

Example:

```go
kvs := map[string]string{"a": "apple", "b": "banana"}
for k, v := range kvs {
    fmt.Printf("%s -> %s\n", k, v)
}
```

You can also iterate over just the keys:

```go
for k := range kvs {
    fmt.Println("Key:", k)
}
```

### Iterating Over Strings

Using `range` on a string iterates over its Unicode code points. It provides the starting byte index of the rune and the rune itself.

Example:

```go
for i, c := range "go" {
    fmt.Println(i, c)
}
```

