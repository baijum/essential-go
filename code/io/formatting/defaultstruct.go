package main

import (
    "fmt"
)

// Circle represents a circle
type Circle struct {
    radius float64
    color  string
}

func main() {
    c := Circle{radius: 76.45, color: "blue"}
    fmt.Printf("Value: %#v\n", c)
    fmt.Printf("Value with fields: %+v\n", c)
    fmt.Printf("Type: %T\n", c)
}
