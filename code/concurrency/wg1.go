package main

import (
    "fmt"
    "sync"
    "time"
)

func someWork(i int) {
    time.Sleep(time.Millisecond * 10)
    fmt.Println(i)
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(j int) {
            defer wg.Done()
            someWork(j)
        }(i)
    }
    wg.Wait()
}
