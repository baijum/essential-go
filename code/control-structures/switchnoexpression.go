package main

import "fmt"

func main() {
    age := 10
    switch {
    case age < 10:
        fmt.Println("Junior", age)
    case age < 20:
        fmt.Println("Senior", age)
    default:
        fmt.Println("Other", age)
    }
}
