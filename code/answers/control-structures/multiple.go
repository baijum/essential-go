package main

import (
	"fmt"
)

func main() {
	n := 7
	found := false
	for _, j := range []int{2, 3, 5} {
		if n%j == 0 {
			fmt.Printf("%d is a multiple of %d\n", n, j)
			found = true
		}
	}
	if !found {
		fmt.Printf("%d is not a multiple of 2, 3, or 5\n", n)
	}
}
