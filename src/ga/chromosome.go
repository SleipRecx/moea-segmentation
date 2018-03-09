package ga

import (
	"../graph"
	"../img"
)


type Chromosome struct {
	MyImage  img.Image
	Segments [][]img.Coordinate
}

func NewChromosome(partitions [][]graph.Vertex, myImage img.Image) Chromosome {
	segments := make([][]img.Coordinate, 0)
	for i := range partitions {
		var coordinateList []img.Coordinate
		for j := range partitions[i] {
			coordinates := partitions[i][j].(img.Coordinate)
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

func calcSegmentEdgeValue(segment []img.Coordinate, myImage img.Image) float64 {
	var edgeValue float64
	pixelSegment := coordinatesToPixels(segment, myImage)

	for i, cord := range segment {
		x, y := cord.X, cord.Y
		if checkIfItemInSegment(segment, img.Coordinate{x + 1, y}) {
			edgeValue += pixelSegment[i].Distance(myImage.Pixels[x+1][y])
		}
		if checkIfItemInSegment(segment, img.Coordinate{x - 1, y}) {
			edgeValue += pixelSegment[i].Distance(myImage.Pixels[x-1][y])
		}
		if checkIfItemInSegment(segment, img.Coordinate{x, y + 1}) {
			edgeValue += pixelSegment[i].Distance(myImage.Pixels[x][y+1])
		}
		if checkIfItemInSegment(segment, img.Coordinate{x, y - 1}) {
			edgeValue += pixelSegment[i].Distance(myImage.Pixels[x][y-1])
		}
	}
	return edgeValue
}

func checkIfItemInSegment(segment []img.Coordinate, coordinate img.Coordinate) bool {
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

func calcSegmentDeviation(segment []img.Pixel) float64 {
	deviation := 0.0
	centroid := calcCentroid(segment)

	for _, pixel := range segment {
		deviation += pixel.Distance(centroid)
	}
	return deviation
}

func calcCentroid(segment []img.Pixel) img.Pixel {
	r, g, b := 0, 0, 0
	for _, pixel := range segment {
		r += int(pixel.R)
		g += int(pixel.G)
		b += int(pixel.B)
	}
	numPixels := len(segment)
	centroid := img.Pixel{
		R: uint8(r / numPixels),
		G: uint8(g / numPixels),
		B: uint8(b / numPixels),
		A: segment[0].A}

	return img.Pixel{
		R: centroid.R,
		G: centroid.G,
		B: centroid.B,
		A: segment[0].A}
}

func coordinatesToPixels(segment []img.Coordinate, myImage img.Image) []img.Pixel {
	var pixels []img.Pixel
	for i := range segment {
		x, y := segment[i].X, segment[i].Y
		pixel := img.Pixel{
			R: myImage.Pixels[x][y].R,
			G: myImage.Pixels[x][y].G,
			B: myImage.Pixels[x][y].B,
			A: myImage.Pixels[x][y].A,
		}
		pixels = append(pixels, pixel)
	}
	return pixels

}
