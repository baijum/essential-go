package main

import "fmt"

func main() {
    o := "=="
    switch o {
    case "+", "-", "*", "/", "%", "&", "|", "^", "&^", "<<", ">>":
        fmt.Println("Arithmetic operator")
    case "==", "!=", "<", "<=", ">", ">=":
        fmt.Println("Comparison operators")
    case "&&", "||", "!":
        fmt.Println("Logical operators")
    default:
        fmt.Println("Unknown operator")
    }
}
