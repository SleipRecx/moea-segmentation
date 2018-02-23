package main

import "fmt"

func calcOverallDeviation(segmentSet [][]Pixel) float64 {
	var overallDeviation float64
	fmt.Print(len(segmentSet))
	for i := 0; i < len(segmentSet); i++ {
		overallDeviation += calcSegmentDeviation(segmentSet[i])
		fmt.Print("\n", i, 1 / overallDeviation)

	}
	return 1 / overallDeviation
}

func calcSegmentDeviation(segment []Pixel) float64 {
	euclideanDistance := 0.0
	var segmentDeviation float64
	for i, mainPixel := range segment {
		if euclideanDistance > 0.0 {
			centroid := euclideanDistance / float64(len(segment))
			deviation :=  euclideanDistance * centroid
			segmentDeviation += deviation
		}
		euclideanDistance = 0.0
		tempSegment := segment
		tempSegment = append(tempSegment[:i], tempSegment[i+1:]...) // mulig dette gir out of bounds exception
		for _, pixel := range tempSegment {
			euclideanDistance += calcEuclideanDistance(mainPixel, pixel)
		}
	}

	return segmentDeviation
}
