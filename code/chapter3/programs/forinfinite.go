package main

import "fmt"

func main() {
    names := []string{"Tom", "Polly", "Huck", "Becky"}
    i := 0
    for {
        if i >= len(names) {
            break
        }
        fmt.Println(names[i])
        i++
    }
}
