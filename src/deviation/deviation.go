package deviation

import (
	//"fmt"

	"../image"
)

type Pixel = image.Pixel


func CalcOverallDeviation(segments [][]image.Coordinate, myImage image.Image) float64 {
	var deviation float64
	for i := range segments {
		var pixels []Pixel
		for j := range segments[i] {
			x, y := segments[i][j].X, segments[i][j].Y
			pixel := Pixel{
				myImage.Pixels[x][y].R,
				myImage.Pixels[x][y].G,
				myImage.Pixels[x][y].B,
				myImage.Pixels[x][y].A,
			}
			pixels = append(pixels, pixel)
		}
		deviation += calcSegmentDeviation(pixels)
		//fmt.Println("Segment number: ", i, " Deviation: ", 1/deviation)

	}
	return 1 / deviation
}

func calcSegmentDeviation(segment []Pixel) float64 {
	deviation := 0.0
	centroid := calcCentroid(segment)

	for _, pixel := range segment {
		deviation += pixel.Distance(centroid)
	}
	return deviation
}

func calcCentroid(segment []Pixel) Pixel {
	r, g, b := 0, 0, 0
	for _, pixel := range segment {
		r += int(pixel.R)
		g += int(pixel.G)
		b += int(pixel.B)
	}
	numPixels := len(segment)
	centroid := Pixel{
		R: uint8(r / numPixels),
		G: uint8(g / numPixels),
		B: uint8(b / numPixels),
		A: segment[0].A}

	return Pixel{
		R: centroid.R,
		G: centroid.G,
		B: centroid.B,
		A: segment[0].A}
}
