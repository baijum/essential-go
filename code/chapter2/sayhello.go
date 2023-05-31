package main

import "fmt"

// SayHello returns wishing message based on input
func SayHello(name string) string {
	if name == "" { // check for empty string
		return "Hello, World!"
	} else {
		return "Hello, " + name + "!"
	}
}

func main() {
	w := SayHello("")
	fmt.Println(w)
	b := SayHello("Baiju")
	fmt.Println(b)
}
