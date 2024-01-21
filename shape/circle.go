package shape

type Circle struct {
	radius float64
	pie    float64
}

func NewCircle(raduis float64) Circle {
	return Circle{
		radius: raduis,
		pie:    3.142,
	}
}

func (c *Circle) SetRaduis(raduis float64) {
	c.radius = raduis
}

func (c *Circle) GetRaduis() float64 {
	return c.radius
}

func (c *Circle) GetPerimeter() float64 {
	return 2 * c.pie * c.radius
}

func (c *Circle) GetArea() float64 {
	return c.pie * (c.radius * c.radius)
}
