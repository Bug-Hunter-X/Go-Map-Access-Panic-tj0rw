```go
package main

import "fmt"

func main() {
    var m map[string]int
    value, ok := m["0"]
    if ok {
        fmt.Println("Value found:", value)
    } else {
        fmt.Println("Key not found")
    }
}
```