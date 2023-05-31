package main

import "fmt"

func main() {
    if money := 20000; money > 15000 {
        fmt.Println("I am going to buy a car.", money)
    } else {
        fmt.Println("I am going to buy a bike.", money)
    }
    // can't use the variable `money` here
}
