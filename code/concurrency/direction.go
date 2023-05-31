package main

import "fmt"

func sendOnly(name chan<- string) {
    name <- "Hi"
}

func receiveOnly(name <-chan string) {
    fmt.Println(<-name)
}

func main() {
    n := make(chan string)

    go func() {
        fmt.Println(<-n)
    }()

    sendOnly(n)

    go func() {
        n <- "Hello"
    }()

    receiveOnly(n)
}
