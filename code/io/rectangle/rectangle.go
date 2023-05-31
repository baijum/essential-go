package main

import (
    "flag"
    "fmt"
    "log"
    "os"
)

// Rectangle represents a rectangle shape
type Rectangle struct {
    Length float64
    Width  float64
}

// Area return the area of a rectangle
func (r Rectangle) Area() float64 {
    return r.Length * r.Width
}

var length = flag.Float64("length", 0, "length of rectangle")
var width = flag.Float64("width", 0, "width of rectangle")

func main() {
    flag.Parse()
    if *length <= 0 {
        log.Println("Invalid length")
        os.Exit(1)
    }
    if *width <= 0 {
        log.Println("Invalid width")
        os.Exit(1)
    }
    r := Rectangle{Length: *length, Width: *width}
    a := r.Area()
    fmt.Println("Area: ", a)
}
