package main

import "fmt"

func main() {
    num := 1
    goto Marker
    num = 2
Marker:
    fmt.Println("Value of num:", num)
}
