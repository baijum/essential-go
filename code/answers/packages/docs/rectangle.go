package shape

// Rectangle represents a rectangle shape
type Rectangle struct {
	Length float64
	Width float64
}

// Area return the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}
