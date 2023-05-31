package main

import "fmt"

func main() {
    names := []string{"Tom", "Polly", "Huck", "Becky"}
    i := 0
Loop:
    if i < len(names) {
        fmt.Println(names[i])
        i++
        goto Loop
    }
}
