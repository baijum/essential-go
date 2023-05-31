package main

import (
    "fmt"
)

func main() {
    v := make([]string, 0)
    for i := 0; i < 9000000; i++ {
        v = append(v, "Yellow")
    }
    fmt.Printf("%v\n", len(v))
}
