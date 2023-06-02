package main

import (
    "fmt"
)

func value(a int) {
    fmt.Printf("%v\n", &a)
}

func pointer(a *int) {
    fmt.Printf("%v\n", a)
}

func main() {
    a := 4
    fmt.Printf("%v\n", &a)
    value(a)
    pointer(&a)
}
