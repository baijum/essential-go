package main

import (
    "fmt"
)

func main() {
    v := make([]string, 3)
    fmt.Printf("%v\n", len(v))
    a := []string{"Yellow", "Black"}
    v = append(v, a...)
    fmt.Printf("%v\n", len(v))
}
