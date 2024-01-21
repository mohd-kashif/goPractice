package coordinates

import "math"

type Point struct {
	X float64
	Y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(math.Pow(p.X, 2) + math.Pow(p.Y, 2))
}
