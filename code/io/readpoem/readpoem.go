package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fd, err := os.Open("poem.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	for {
		chars := make([]byte, 50)
		n, err := fd.Read(chars)
		if n == 0 && err == io.EOF {
			break
		}
		fmt.Print(string(chars))
	}
	fd.Close()
}
