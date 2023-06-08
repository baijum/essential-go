package main

import (
    "bufio"
    "fmt"
    "os"
    "os/signal"
    "strings"
    "time"
)

func watch(word, fp string) error {

    f, err := os.Open(fp)
    if err != nil {
        return err
    }
    r := bufio.NewReader(f)
    defer f.Close()
    for {
        line, err := r.ReadBytes('\n')
        if err != nil {
            if err.Error() == "EOF" {
                time.Sleep(2 * time.Second)
                continue
            }
            fmt.Printf("Error: %s\n%v\n", line, err)
        }
        if strings.Contains(string(line), word) {
            fmt.Printf("%s: Matched: %s\n", fp, line)
        }
        time.Sleep(2 * time.Second)
    }
}

func main() {
    word := os.Args[1]
    files := []string{}
    for _, f := range os.Args[2:len(os.Args)] {
        files = append(files, f)
        go watch(word, f)
    }
    sig := make(chan os.Signal, 1)
    done := make(chan bool)
    signal.Notify(sig, os.Interrupt)
    go func() {
        for _ = range sig {
            done <- true
        }
    }()
    <-done
}
