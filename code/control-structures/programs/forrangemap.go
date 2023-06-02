package main

import "fmt"

func main() {
    var characters = map[string]int{
        "Tom":   8,
        "Polly": 51,
        "Huck":  9,
        "Becky": 8,
    }
    for name, age := range characters {
        fmt.Println(name, age)
    }
}
