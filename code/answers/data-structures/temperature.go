package main

import "fmt"

type Temperature float64

func (t Temperature) Freezing() bool {
	if t < Temperature(0.0) {
		return true
	}
	return false
}

func main() {

	temperatures := map[string]Temperature{
		"New York":  9.3,
		"London":    13.5,
		"New Delhi": 31.5,
		"Montreal":  -9.0,
	}

	location := "New Delhi"
	fmt.Println(location, temperatures[location].Freezing())

	location = "Montreal"
	fmt.Println(location, temperatures[location].Freezing())

}
