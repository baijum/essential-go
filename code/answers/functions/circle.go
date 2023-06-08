package main

import "fmt"

type Circle struct {
	Radius float64
}

// Area return the area of a circle for the given radius
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	c := Circle{5.0}
	fmt.Println(c.Area())
}
