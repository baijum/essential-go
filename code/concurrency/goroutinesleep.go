package main

import (
    "fmt"
    "time"
)

var msg string

func setMessage() {
    msg = "Hello, World!"
}

func main() {
    go setMessage()
    time.Sleep(1 * time.Millisecond)
    fmt.Println(msg)
}
