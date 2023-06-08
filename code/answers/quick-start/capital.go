package main

import "fmt"

func StartsCapital(s string) bool {
	for _, v := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		if string(s[0]) == string(v) {
			return true
		}
	}
	return false
}

func main() {
	h := StartsCapital("Hello")
	fmt.Println(h)
	w := StartsCapital("world")
	fmt.Println(w)
}
