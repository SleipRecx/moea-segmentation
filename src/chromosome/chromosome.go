package chromosome

import (
	//"fmt"

	"../image"
	"../graph"
)

type Coordinate = image.Coordinate

type Pixel = image.Pixel

type Chromosome struct {
	MyImage image.Image
	Segments [][]Coordinate
}

func NewChromosome(mst graph.Tree, myImage image.Image, initialSegments int) Chromosome {
	vertices := mst.RandomDisjointPartition(initialSegments)
	var segments [][]Coordinate
	for i := range vertices {
		var coordinateList []Coordinate
			for j := range vertices[i] {
				coordinates := vertices[i][j].(Coordinate)
				coordinateList = append(coordinateList, coordinates)
			}
			segments = append(segments, coordinateList)
	}
	return Chromosome{MyImage: myImage, Segments: segments}
}

func (c Chromosome) CalcEdgeValue() float64 {
	var edgeValue float64
	segments := c.Segments
	myImage := c.MyImage
	for i := range segments {
		pixelSegment := coordinatesToPixels(segments[i], myImage)
		for j := range segments[i] {
			x, y := segments[i][j].X, segments[i][j].Y
			if x + 1 < len(myImage.Pixels) {
				edgeValue += pixelSegment[j].Distance(myImage.Pixels[x+1][y])
			}
			if x - 1 >= 0 {
				edgeValue += pixelSegment[j].Distance(myImage.Pixels[x-1][y])
			}
			if y + 1 < len(myImage.Pixels[x]) {
				edgeValue += pixelSegment[j].Distance(myImage.Pixels[x][y+1])
			}
			if y - 1 >= 0 {
				edgeValue += pixelSegment[j].Distance(myImage.Pixels[x][y-1])
			}
		}
		//fmt.Println("Segment number: ", i, " Edge value: ", 1/edgeValue)
	}
	return  - edgeValue
}


func (c Chromosome) CalcDeviation () float64 {
	var deviation float64
	segments := c.Segments
	myImage := c.MyImage
	for i := range segments {
		pixelSegment := coordinatesToPixels(segments[i], myImage)
		deviation += calcSegmentDeviation(pixelSegment)
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

func coordinatesToPixels(segment []Coordinate, myImage image.Image) []Pixel {
	var pixels []Pixel
	for i := range segment {
		x, y := segment[i].X, segment[i].Y
		pixel := Pixel{
			R: myImage.Pixels[x][y].R,
			G: myImage.Pixels[x][y].G,
			B: myImage.Pixels[x][y].B,
			A: myImage.Pixels[x][y].A,
		}
		pixels = append(pixels, pixel)
	}
	return pixels

}