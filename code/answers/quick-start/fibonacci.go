package main

import "fmt"

func Fib(n int) {
	for i, j := 0, 1; i < n; i, j = i+j, i {
		fmt.Println(i)
	}

}

func main() {
	Fib(200)
}
