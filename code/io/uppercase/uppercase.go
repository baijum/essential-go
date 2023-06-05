package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	str := string(stdin)
	newStr := strings.TrimSuffix(str, "\n")
	upper := strings.ToUpper(newStr)
	fmt.Println(upper)
}
