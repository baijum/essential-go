package main

import (
    "fmt"
)

func main() {
    fmt.Printf("Value: %#v, Type: %T\n", "Baiju", "Baiju")
    fmt.Printf("Value: %#v, Type: %T\n", 7, 7)
    fmt.Printf("Value: %#v, Type: %T\n", uint(7), uint(7))
    fmt.Printf("Value: %#v, Type: %T\n", int8(7), int8(7))
    fmt.Printf("Value: %#v, Type: %T\n", true, true)
    fmt.Printf("Value: %#v, Type: %T\n", 7.0, 7.0)
    fmt.Printf("Value: %#v, Type: %T\n", (1 + 6i), (1 + 6i))
}
