package method_and_receiver

type Retangle struct {
	Width  float64
	Height float64
}

func (retangle Retangle) Perimeter() float64 {
	return (2 * retangle.Width) + (2 * retangle.Height)
}

func (retangle Retangle) Area() float64 {
	return retangle.Width * retangle.Height
}
