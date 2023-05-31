package main

import (
    "fmt"
    "time"
)

func main() {
    v := "Becky"
    t := time.Now()
    switch v {
    case "Huck":
        if t.Hour() < 12 {
            fmt.Println("Good morning,", v)
            break
        }
        fmt.Println("Hello,", v)
    case "Becky":
        if t.Hour() < 12 {
            fmt.Println("Good morning,", v)
            break
        }
        fmt.Println("Hello,", v)
    default:
        fmt.Println("Hello")
    }
}
