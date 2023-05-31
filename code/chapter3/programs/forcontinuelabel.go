package main

import "fmt"

func main() {
    names := []string{"Tom", "Polly", "Huck", "Becky"}
Outer:
    for i := 0; i < len(names); i++ {
        for j := 0; j < len(names[i]); j++ {
            if names[i][j] == 'u' {
                continue Outer
            }
        }
        fmt.Println(names[i])
    }
}
