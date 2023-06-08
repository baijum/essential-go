package shape

// Triangle represents a rectangle shape
type Triangle struct {
	Breadth float64
	Height float64
}

// Area return the area of a triangle
func (t Triangle) Area() float64 {
	return (t.Breadth * t.Height)/2
}
