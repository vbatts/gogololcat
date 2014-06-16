package lol

import (
	"math"
)

type distanceMatch struct {
	Distance float64
	Pos      int
}

type distanceSorter []distanceMatch

func (ds distanceSorter) Len() int {
	return len(ds)
}
func (ds distanceSorter) Swap(i, j int) {
	ds[i], ds[j] = ds[j], ds[i]
}
func (ds distanceSorter) Less(i, j int) bool {
	return ds[i].Distance < ds[j].Distance
}

func Distance(rgb1, rgb2 RGB) float64 {
	var (
		sum  float64
		a, b float64
	)
	for i := 0; i < 3; i++ {
		switch i {
		case 0:
			a = rgb1.R
			b = rgb2.R
		case 1:
			a = rgb1.G
			b = rgb2.G
		case 2:
			a = rgb1.B
			b = rgb2.B
		}
		sum += math.Pow((a - b), 2)
	}
	return sum
}
