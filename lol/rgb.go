package lol

type RGB struct {
	R, G, B float64
}

func (rgb RGB) Sum() float64 {
	return rgb.R + rgb.G + rgb.B
}
