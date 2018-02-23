package main

import (
	"math"
	"github.com/lucasb-eyer/go-colorful"
	"image/color"
	"math/rand"
)

func calcEuclideanDistance(p1 Pixel, p2 Pixel) float64 {
	num := rand.Int() % 2
	if num % 2 == 0 {
		return calcRGBEuclidean(p1, p2)
	}

	return calcLABEuclidean(p1, p2)
}


func calcRGBEuclidean(p1 Pixel, p2 Pixel) float64 {
	r := math.Pow(float64(p1.R - p2.R), 2)

	g := math.Pow(float64(p1.G - p2.G), 2)

	b := math.Pow(float64(p1.B - p2.B), 2)


	return math.Sqrt(r + g + b)
}

func calcLABEuclidean(p1 Pixel, p2 Pixel) float64 {
	c1 := colorful.MakeColor(color.RGBA{p1.R, p1.G, p1.B, p1.A})
	l1, a1, b1 := c1.Lab()

	c2 := colorful.MakeColor(color.RGBA{p2.R, p2.G, p2.B, p2.A})
	l2, a2, b2 := c2.Lab()

	l := math.Pow(float64(l1 - l2), 2)

	a := math.Pow(float64(a1 - a2), 2)

	b := math.Pow(float64(b1 - b2), 2)

	return math.Sqrt(l + a + b)

}
