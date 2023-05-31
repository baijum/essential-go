package main

import (
    "fmt"
    "sync"
)

var msg string
var wg sync.WaitGroup

func setMessage() {
    msg = "Hello, World!"
    wg.Done()
}

func main() {
    wg.Add(1)
    go setMessage()
    wg.Wait()
    fmt.Println(msg)
}
