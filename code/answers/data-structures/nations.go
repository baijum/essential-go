package main

import "fmt"

type Country struct {
	Capital    string
	Currency   string
	Popolation int
}

func main() {
	countries := map[string]Country{}
	countries["India"] = Country{Capital: "New Delhi",
		Currency: "Indian Rupee", Popolation: 1428600000}
	fmt.Printf("%#v\n", countries)
}
