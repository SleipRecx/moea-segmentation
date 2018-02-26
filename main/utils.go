package main

import (
	"math"
)

func calcEuclideanDistance(p1 Pixel, p2 Pixel) float64 {
	r := math.Pow(float64(p1.R - p2.R), 2)

	g := math.Pow(float64(p1.G - p2.G), 2)

	b := math.Pow(float64(p1.B - p2.B), 2)


	return math.Sqrt(r + g + b)
}
