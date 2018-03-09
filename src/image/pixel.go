package image

import (
	"math"
)

type Pixel struct {
	R, G, B, A uint8
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
}

func (pixel Pixel) Distance(pixel2 Pixel) float64 {
	return ColorDistance(pixel, pixel2)
}

func ColorDistance(p1 Pixel, p2 Pixel) float64 {
	r := math.Pow(float64(p1.R)-float64(p2.R), 2)
	g := math.Pow(float64(p1.G)-float64(p2.G), 2)
	b := math.Pow(float64(p1.B)-float64(p2.B), 2)
	a := math.Pow(float64(p1.A)-float64(p2.A), 2)
	return math.Sqrt(r + g + b + a)
}
