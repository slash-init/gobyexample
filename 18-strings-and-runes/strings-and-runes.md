# Go by Example: Strings and Runes

A Go string is a read-only slice of bytes. The language and the standard library treat strings specially as containers of text encoded in UTF-8. In other languages, strings are made of "characters." In Go, the concept of a character is called a **rune**—an integer that represents a Unicode code point. This [Go blog post](https://blog.golang.org/strings) is a good introduction to the topic.

---

### Example Code

```go
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    // A string assigned a literal value representing the word "hello" in Thai.
    const s = "สวัสดี"

    // Since strings are equivalent to []byte, this will produce the length of the raw bytes stored within.
    fmt.Println("Len:", len(s))

    // Indexing into a string produces the raw byte values at each index.
    // This loop generates the hex values of all the bytes that constitute the code points in `s`.
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
    fmt.Println()

    // To count how many runes are in a string, we can use the utf8 package.
    // Note: RuneCountInString depends on the size of the string, as it decodes each UTF-8 rune sequentially.
    fmt.Println("Rune count:", utf8.RuneCountInString(s))

    // A range loop handles strings specially and decodes each rune along with its offset in the string.
    for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }
}
```

---

### Explanation

1. **UTF-8 Encoding**: Go strings are UTF-8 encoded by default, making them compatible with a wide range of text representations.
2. **Raw Byte Length**: The `len` function returns the number of bytes, not runes, in the string.
3. **Hexadecimal Representation**: Indexing a string gives the raw byte values, which can be printed in hexadecimal format.
4. **Rune Count**: The `utf8.RuneCountInString` function counts the number of runes, which may differ from the byte length for multi-byte characters.
5. **Range Loop**: Using a `range` loop over a string decodes each rune and provides its Unicode representation and starting byte offset.

---

### Key Takeaways

- Strings in Go are immutable and represented as a sequence of bytes.
- Runes are Go's way of representing Unicode code points.
- Use the `utf8` package for working with multi-byte characters.
- The `range` loop is a powerful tool for iterating over runes in a string.
