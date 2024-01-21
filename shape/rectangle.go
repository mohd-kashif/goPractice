package shape

type Rectangle struct {
	length float64
	width  float64
}

func NewRectangle(length float64, width float64) Rectangle {
	return Rectangle{
		length: length,
		width:  width,
	}
}
func (r *Rectangle) GetPerimeter() float64 {
	return 2 * (r.length + r.width)
}
