package gointerface

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Retangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (retangle Retangle) Perimeter() float64 {
	return (2 * retangle.Width) + (2 * retangle.Height)
}

func (retangle Retangle) Area() float64 {
	return retangle.Width * retangle.Height
}

func (triangle Triangle) Area() float64 {
	return 0.5 * triangle.Base * triangle.Height
}
