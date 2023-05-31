package main

import (
    "fmt"
)

var c = make(chan int)
var msg string

func setMessage() {
    msg = "Hello, World!"
    c <- 0
}

func main() {
    go setMessage()
    <-c
    fmt.Println(msg)
}
