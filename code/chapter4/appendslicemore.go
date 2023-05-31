package main

import (
    "fmt"
)

func main() {
    v := make([]string, 3)
    fmt.Printf("%v\n", len(v))
    v = append(v, "Yellow", "Black")
    fmt.Printf("%v\n", len(v))
}
