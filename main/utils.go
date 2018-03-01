package main

import (
	"math"
)

func colorDistance(p1 Pixel, p2 Pixel) float64 {
	r := math.Pow(float64(p1.R-p2.R), 2)
	g := math.Pow(float64(p1.G-p2.G), 2)
	b := math.Pow(float64(p1.B-p2.B), 2)
	a := math.Pow(float64(p1.A-p2.A), 2)

	return math.Sqrt(r + g + b + a)
}

func (p1 Pixel) distance(p2 Pixel) float64 {
	return colorDistance(p1, p2)
}
