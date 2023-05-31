package main

import "fmt"

type Circle struct {
    radius float64
}

func (c *Circle) Area() float64 {
    return 3.14 * c.radius * c.radius
}

func main() {
    c1 := Circle{3.4}
    a1 := c1.Area()
    fmt.Println(a1)

    c2 := &Circle{3.4}
    a2 := c2.Area()
    fmt.Println(a2)

}
