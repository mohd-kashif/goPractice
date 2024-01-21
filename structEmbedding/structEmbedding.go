package structembedding

import (
	"math"
)

// Shape interface represents a generic shape
type Shape interface {
	Area() float64
}

// 2DShape is a base struct for 2D shapes
type _2DShape struct {
	name string
}

// Area calculates the area of the 2D shape
func (_2d *_2DShape) PArea() float64 {
	return 0
}

// Square is a 2D shape that extends _2DShape
type Square struct {
	_2DShape
	sideLength float64
}

// Area calculates the area of the square
func (s *Square) Area() float64 {
	return s.PArea()
	// return s.sideLength * s.sideLength
}

// 3DShape is a base struct for 3D shapes
type _3DShape struct {
	name string
}

// Volume calculates the volume of the 3D shape
func (_3d *_3DShape) Volume() float64 {
	return 0
}

// Cube is a 3D shape that extends _3DShape
type Cube struct {
	_3DShape
	sideLength float64
}

// Volume calculates the volume of the cube
func (c *Cube) Volume() float64 {
	return math.Pow(c.sideLength, 3)
}
