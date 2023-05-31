package main

import "fmt"

func main() {
    names := []string{"Tom", "Polly", "Huck", "Becky"}
    for i := 0; i < len(names); i++ {
        if names[i] == "Polly" {
            continue
        }
        fmt.Println(names[i])
    }
}
