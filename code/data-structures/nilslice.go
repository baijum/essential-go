package main

import "fmt"

func main() {
    var v []string
    fmt.Printf("%#v, %#v\n", v, v == nil)
    // Output: []string(nil), true
}
