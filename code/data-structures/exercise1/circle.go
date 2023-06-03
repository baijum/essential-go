package main

import "fmt"

type Circle float64

func (r Circle) Area() float64 {
	return float64(3.14 * r * r)
}

func (r Circle) Perimeter() float64 {
	return float64(2 * 3.14 * r)
}

func main() {
	c := Circle(4.0)
	fmt.Println(c.Area())
	fmt.Println(c.Perimeter())
}
