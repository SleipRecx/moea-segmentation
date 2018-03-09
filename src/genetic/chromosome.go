package genetic

import (


	"../graph"
	"../image"
)

type Coordinate = image.Coordinate

type Pixel = image.Pixel

type Chromosome struct {
	MyImage  image.Image
	Segments [][]Coordinate
}

func NewChromosome(partitions [][]graph.Vertex, myImage image.Image) Chromosome {
	var segments [][]Coordinate
	for i := range partitions {
		var coordinateList []Coordinate
		for j := range partitions[i] {
			coordinates := partitions[i][j].(Coordinate)
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
	for _, segment := range segments {
		edgeValue += calcSegmentEdgeValue(segment, myImage)
		//fmt.Println("Segment number: ", i, " Edge value: ", 1/edgeValue)
	}
	return -edgeValue
}

func calcSegmentEdgeValue(segment []image.Coordinate, myImage image.Image) float64 {
	var edgeValue float64
	pixelSegment := coordinatesToPixels(segment, myImage)

	for i, cord := range segment {
		x, y := cord.X, cord.Y
		if checkIfItemInSegment(segment, Coordinate{x + 1, y}) {
			edgeValue += pixelSegment[i].Distance(myImage.Pixels[x+1][y])
		}
		if checkIfItemInSegment(segment, Coordinate{x - 1, y}) {
			edgeValue += pixelSegment[i].Distance(myImage.Pixels[x-1][y])
		}
		if checkIfItemInSegment(segment, Coordinate{x, y + 1}) {
			edgeValue += pixelSegment[i].Distance(myImage.Pixels[x][y+1])
		}
		if checkIfItemInSegment(segment, Coordinate{x, y - 1}) {
			edgeValue += pixelSegment[i].Distance(myImage.Pixels[x][y-1])
		}
	}
	return edgeValue
}

func checkIfItemInSegment(segment []image.Coordinate, coordinate Coordinate) bool {
	for _, item := range segment {
		if item == coordinate {
			return true
		}
	}
	return false
}

func (c Chromosome) CalcDeviation() float64 {
	var deviation float64
	segments := c.Segments
	myImage := c.MyImage
	for i := range segments {
		pixelSegment := coordinatesToPixels(segments[i], myImage)
		deviation += calcSegmentDeviation(pixelSegment)
		//fmt.Println("Segment number: ", i, " Deviation: ", 1/deviation)
	}

	if deviation == 0.0 {
		return 0.0
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
