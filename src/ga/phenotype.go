package ga

import (
	"../img"
	"../ds"
)

type Phenotype struct {
	Segments   [][]img.Coordinate
	SegmentMap map[img.Coordinate]int
	Chromosome Chromosome
}

func NewPhenotype() Phenotype {
	chromosome := NewChromosome(300)
	return NewPhenotypeFromChromosome(chromosome)
}

func NewPhenotypeFromChromosome(chromosome Chromosome) Phenotype {
	directionMap := make(map[img.Coordinate]img.Direction)

	i := 0
	for x := 0; x < img.ImageWidth; x++ {
		for y := 0; y < img.ImageHeight; y++ {
			cord := img.Coordinate{X: x, Y: y}
			directionMap[cord] = chromosome.genes[i]
			i++
		}
	}
	disjointMap := make(map[img.Coordinate]*ds.DisjointSet)
	for k := range directionMap {
		set := ds.MakeSet(k)
		disjointMap[k] = set
	}

	for cord, direction := range directionMap {
		nCord := img.CordFromDirection(cord, direction)
		s1, s2 := ds.FindSet(disjointMap[cord]), ds.FindSet(disjointMap[nCord])
		ds.Union(s1, s2)
	}

	segmentMap := make(map[*ds.DisjointSet][]img.Coordinate)

	for k := range directionMap{
		segmentMap[ds.FindSet(disjointMap[k])] = append(segmentMap[ds.FindSet(disjointMap[k])], k)
	}
	segments := make([][]img.Coordinate, 0)
	for _, value := range segmentMap{
		segments = append(segments, value)
	}
	coordinateToSegmentMap := make(map[img.Coordinate]int)
	for i := range segments {
		for j := range segments[i] {
			coordinateToSegmentMap[segments[i][j]] = i
		}
	}

	return Phenotype{Segments:segments, Chromosome:chromosome, SegmentMap:coordinateToSegmentMap}
}

func (p Phenotype) CalcDeviation() float64 {
	var deviation float64
	segments := p.Segments
	for i := range segments {
		if len(segments[i]) > 0 {
			pixelSegment := coordinatesToPixels(segments[i])
			deviation += calcSegmentDeviation(pixelSegment)
			//fmt.Println("Segment number: ", i, " Deviation: ", 1/deviation)
		}
	}

	if deviation == 0.0 {
		return 0.0
	}
	return 1.0 / deviation

}

func (p Phenotype) CalcEdgeValue() float64 {
	var edgeValue float64
	segments := p.Segments
	segmentMap := p.SegmentMap
	for i := range segments {
		for _, cord := range segments[i] {
			x, y := cord.X, cord.Y
			right := img.Coordinate{X: x + 1, Y: y}
			left := img.Coordinate{X: x - 1, Y: y}
			up := img.Coordinate{X: x, Y: y + 1}
			down := img.Coordinate{X: x, Y: y - 1}

			neighbours := make([]img.Coordinate, 0)
			neighbours = append(neighbours, right, left, up, down)

			for _, neighbour := range neighbours {
				if segmentMap[neighbour] == segmentMap[cord] {
					neighX, neighY := neighbour.X, neighbour.Y
					edgeValue += img.MyImage.Pixels[x][y].Distance(img.MyImage.Pixels[neighX][neighY])
					//fmt.Println("Edge", img.Coordinate{x, y}, img.Coordinate{neighX, neighY})
				}
			}

		}
	}
	return -edgeValue
}

func checkIfItemInSegment(segment []img.Coordinate, coordinate img.Coordinate) bool {
	for _, item := range segment {
		if item == coordinate {
			return true
		}
	}
	return false
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

func coordinatesToPixels(segment []img.Coordinate) []img.Pixel {
	var pixels []img.Pixel
	for i := range segment {
		x, y := segment[i].X, segment[i].Y
		pixel := img.Pixel{
			R: img.MyImage.Pixels[x][y].R,
			G: img.MyImage.Pixels[x][y].G,
			B: img.MyImage.Pixels[x][y].B,
			A: img.MyImage.Pixels[x][y].A,
		}
		pixels = append(pixels, pixel)
	}
	return pixels

}
