package lol

// RGB is a simple holder for Red, Green and Blue
type RGB struct {
	R, G, B float64
}

// Sum up the three values
func (rgb RGB) Sum() float64 {
	return rgb.R + rgb.G + rgb.B
}
