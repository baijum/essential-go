package main

import (
	"fmt"
	"os"
)

func main() {
	fd, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("Cannot write file:", err)
		os.Exit(1)
	}
	fd.Write([]byte("Hello, World!\n"))
	fd.Close()
}
