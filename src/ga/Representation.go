package ga

import (
	"../img"
	"math"
)

type Representation struct {
	SegmentMatrix [][]int
	DirectionMatrix [][]img.Direction
	SegmentMap map[img.Coordinate]int
	OverallDeviation float64
	EdgeValue float64
	Rank int
	CrowdImgDistance float64
}

func NewRepresentation(genotype Genotype) Representation {
	representation := Representation{genotype.SegmentMatrix, genotype.Genes, genotype.GenesToSegmentMap, math.MaxFloat64, math.MaxFloat64, 0, math.MaxFloat64}
	return representation
}
