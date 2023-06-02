package main

import "fmt"

var age int = 10

func main() {
    if age < 10 {
        fmt.Println("Junior")
    } else if age < 20 {
        fmt.Println("Senior")
    } else {
        fmt.Println("Other")
    }
}
