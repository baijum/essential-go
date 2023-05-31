package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    i := os.Args[1]
    n, err := strconv.Atoi(i)
    if err != nil {
        fmt.Println("Not a number:", i)
        os.Exit(1)
    }
    if n%2 == 0 {
        fmt.Printf("%d is even\n", n)
    } else {
        fmt.Printf("%d is odd\n", n)
    }
}
