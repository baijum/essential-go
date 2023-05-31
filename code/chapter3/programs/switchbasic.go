package main

import "fmt"

func main() {
    v := 2
    switch v {
    case 0:
        fmt.Println("zero")
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    default:
        fmt.Println("unknown")
    }
}
