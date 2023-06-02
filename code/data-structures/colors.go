package main

import "fmt"

func main() {
    colors := [3]string{"Red", "Green", "Blue"}
    fmt.Println("Length:", len(colors))
    for i, v := range colors {
        fmt.Println(i, v)
    }
}
