package shape

// Circle represents a circle shape
type Circle struct {
	Radius float64
}

// Area return the area of a circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
