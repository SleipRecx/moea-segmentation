package main

import "fmt"

func calcOverallDeviation(segmentSet [][]Pixel) float64 {
	var deviation float64
	for i := 0; i < len(segmentSet); i++ {
		deviation += calcSegmentDeviation(segmentSet[i])
		fmt.Print("\nSegment number: ", i, " Deviation: ", 1 / deviation, "\n")

	}
	return 1 / deviation
}

func calcSegmentDeviation(segment []Pixel) float64 {
	deviation := 0.0
	centroid := calcCentroid(segment)

	for _, pixel := range segment {
		deviation += colorDistance(pixel, centroid)
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
	centroid := Pixel{    uint8(r / numPixels),
	   					  uint8(g / numPixels),
	   					  uint8(b / numPixels),
			              segment[0].A}
    fmt.Print("\nSegment centroid:", " R:", centroid.R, " G:", centroid.G, " B:", centroid.B)
	return Pixel{    centroid.R,
					 centroid.G,
					 centroid.B,
					 segment[0].A}
}
