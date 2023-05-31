package main

import "fmt"

func main() {
    var name string
    fmt.Printf("Enter your name: ")
    fmt.Scanf("%s", &name)
    fmt.Println("Your name:", name)
}
