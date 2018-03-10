package ga

import (
	"../graph"
	"../img"
)


type Chromosome struct {
	MyImage  img.Image
	Segments [][]img.Coordinate
	SegmentMap map[img.Coordinate]int
}

func NewChromosome(partitions [][]graph.Vertex, myImage img.Image) Chromosome {
	segments := make([][]img.Coordinate, 0)
	segmentMap := make(map[img.Coordinate]int)
	for i := range partitions {
		var coordinateList []img.Coordinate
		for j := range partitions[i] {
			coordinates := partitions[i][j].(img.Coordinate)
			coordinateList = append(coordinateList, coordinates)
		}
		segments = append(segments, coordinateList)
		for _, cord := range coordinateList {
			segmentMap[cord] = i
		}
	}
	return Chromosome{MyImage: myImage, Segments: segments, SegmentMap: segmentMap}
}

func (c Chromosome) CalcEdgeValue() float64 {
	var edgeValue float64
	segments := c.Segments
	myImage := c.MyImage
	segmentMap := c.SegmentMap
	for i := range segments {
		for _, cord := range segments[i] {
			x, y := cord.X, cord.Y
			right := img.Coordinate{x + 1, y}
			left := img.Coordinate{x - 1, y}
			up := img.Coordinate{x, y + 1}
			down := img.Coordinate{x, y - 1}

			neighbours := make([]img.Coordinate, 0)
			neighbours = append(neighbours, right, left, up, down)

			for _, neighbour := range neighbours {
				if segmentMap[neighbour] == segmentMap[cord] {
					neighX, neighY := neighbour.X, neighbour.Y
					edgeValue += myImage.Pixels[x][y].Distance(myImage.Pixels[neighX][neighY])
					//fmt.Println("Edge", img.Coordinate{x, y}, img.Coordinate{neighX, neighY})
				}
			}

		}
	}
	return - edgeValue
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
		if len(segments[i]) > 0 {
			pixelSegment := coordinatesToPixels(segments[i], myImage)
			deviation += calcSegmentDeviation(pixelSegment)
			//fmt.Println("Segment number: ", i, " Deviation: ", 1/deviation)
		}
	}

	if deviation == 0.0 {
		return 0.0
	}
	return 1.0 / deviation


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
