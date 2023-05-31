package main

import "fmt"

type Complex struct {
	Real      float64
	Imaginary float64
}

func (c Complex) String() string {
	return fmt.Sprintf("%.02f + %.02fi", c.Real, c.Imaginary)
}

func main() {
	c1 := Complex{Real: 2.3, Imaginary: 5}
	fmt.Println(c1)
}
