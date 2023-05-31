package main

import (
    "fmt"
    "os"
    "strconv"
)

type Number int

func (num Number) Even() bool {
    if num%2 == 0 {
        return true
    } else {
        return false
    }

}

func main() {
    i := os.Args[1]
    n, err := strconv.Atoi(i)
    if err != nil {
        fmt.Println("Not a number:", i)
        os.Exit(1)
    }
    num := Number(n)
    fmt.Println(num.Even())
}
