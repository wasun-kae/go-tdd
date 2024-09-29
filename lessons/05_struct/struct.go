package gostruct

type Retangle struct {
	Width  float64
	Height float64
}

func Perimeter(retangle Retangle) float64 {
	return (2 * retangle.Width) + (2 * retangle.Height)
}
