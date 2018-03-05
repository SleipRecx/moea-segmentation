package chromosome

import (

	"../image"
	"../graph"
	"../deviation"
)


type Chromosome struct {
	MyImage image.Image
	Segments [][]image.Coordinate
}

func NewChromosome(mst graph.Tree, myImage image.Image) Chromosome {
	vertices := mst.RandomDisjointPartition(100)
	var segments [][]image.Coordinate
	for i := range vertices {
		var coordinateList []image.Coordinate
			for j := range vertices[i] {
				coordinates := vertices[i][j].(image.Coordinate)
				coordinateList = append(coordinateList, coordinates)
			}
			segments = append(segments, coordinateList)
	}
	return Chromosome{MyImage: myImage, Segments: segments}
}

func (c Chromosome) CalcDeviation () float64 {
	return deviation.CalcOverallDeviation(c.Segments, c.MyImage)
}
func (c Chromosome) CalcEdgeValue() float64 {
	return 2.0
}