package ga

import (
	"math/rand"
	"../img"
	"../graph"
	"../constants"
)

type Individual struct {
	DirectionMatrix                [][]img.Direction
	SegmentIDMatrix                [][]int
	SegmentMap                     map[int][]img.Coordinate
	EdgeValue, Deviation, Weighted float64
	Dominates                      []*Individual
	DominatedByN                   int
	Rank                           int
	CrowdingDistance               float64
}

func (individual *Individual) Init(imageGraph graph.Graph) {
	_, individual.DirectionMatrix = imageGraph.GraphSegmentation(rand.Intn(6000-500) + 500)
	individual.SegmentIDMatrix, individual.SegmentMap = graph.DirectionToSegment(individual.DirectionMatrix)
	individual.CalculateFitness()
}

func (individual *Individual) CalculateFitness() {
	individual.EdgeValue = individual.TotalEdgeValue()
	individual.Deviation = individual.TotalDeviation()
	individual.Weighted =  constants.EdgeWeight*individual.EdgeValue + constants.DeviationWeight*individual.Deviation
}

func (individual *Individual) TotalDeviation() float64 {
	total := 0.0
	for i := range individual.SegmentMap {
		centroid := img.CalcCentroid(individual.SegmentMap[i])
		for j := range individual.SegmentMap[i] {
			x, y := individual.SegmentMap[i][j].X, individual.SegmentMap[i][j].Y
			total += img.MyImage.Pixels[x][y].Distance(centroid)
		}
	}
	return total
}

func (individual *Individual) TotalEdgeValue() float64 {
	total := 0.0
	for segId := range individual.SegmentMap {
		for j := range individual.SegmentMap[segId] {
			x, y := individual.SegmentMap[segId][j].X, individual.SegmentMap[segId][j].Y
			neighbors := img.GetFourNeighbors(individual.SegmentMap[segId][j])
			for _, neighbor := range neighbors {
				if individual.SegmentIDMatrix[neighbor.X][neighbor.Y] != segId {
					nX, nY := neighbor.X, neighbor.Y
					total += img.MyImage.Pixels[x][y].Distance(img.MyImage.Pixels[nX][nY])
				}
			}
		}
	}
	return total
}

func (individual *Individual) SegmentMergeMutation() {
	segId1 := getRandomSegmentId(individual.SegmentMap)
	segId2 := segId1

	visited := make(map[img.Coordinate]bool)
	opened := make(map[img.Coordinate]bool)
	queue := []img.Coordinate{individual.SegmentMap[segId1][rand.Intn(len(individual.SegmentMap[segId1]))]}
	for len(queue) > 0 && segId1 == segId2 {
		cord := queue[0]
		visited[cord] = true
		queue = queue[1:]
		neighbors := img.GetFourNeighbors(cord)
		for _, neighbor := range neighbors {
			if _, visited := visited[neighbor]; visited {
				continue
			}
			if _, o := opened[neighbor]; !o {
				queue = append(queue, neighbor)
				opened[neighbor] = true
				newSegId := individual.SegmentIDMatrix[neighbor.X][neighbor.Y]
				if newSegId != segId1 {
					segId2 = segId1
					continue
				}
			}
		}
	}
	if segId1 == segId2 {
		return
	}
	for _, cord := range individual.SegmentMap[segId2] {
		individual.SegmentIDMatrix[cord.X][cord.Y] = segId1
	}
	individual.SegmentMap[segId1] = append(individual.SegmentMap[segId1], individual.SegmentMap[segId2]...)
	delete(individual.SegmentMap, segId2)
}

func (individual *Individual) IsDominating(individual2 *Individual) bool {
	if individual.EdgeValue == individual2.EdgeValue && individual.Deviation == individual2.Deviation {
		return false
	}
	return individual.EdgeValue >= individual2.EdgeValue && individual.Deviation <= individual2.Deviation
}

func (p1 Individual) Crossover(p2 Individual) (Individual, Individual) {
	c1, c2 := Individual{}, Individual{}
	c1.DirectionMatrix = p1.DirectionMatrix
	c2.DirectionMatrix = p2.DirectionMatrix

	if rand.Float64() < constants.CrossoverRate {
		for i := 0; i < rand.Intn(400 - 2) + 2; i++ {
			c1.SegmentMergeMutation()
			c2.SegmentMergeMutation()
		}
		c1.DirectionMatrix = graph.SegmentToDirection(c1.SegmentIDMatrix, c1.SegmentMap)
		c2.DirectionMatrix = graph.SegmentToDirection(c2.SegmentIDMatrix, c2.SegmentMap)
	}
	c1.CalculateFitness()
	c2.CalculateFitness()
	return c1, c2
}

// ----- SORTING -----
type byFitness []Individual

func (p byFitness) Len() int           { return len(p) }
func (p byFitness) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p byFitness) Less(i, j int) bool { return p[i].Weighted > p[j].Weighted }

type byEdge []*Individual

func (p byEdge) Len() int           { return len(p) }
func (p byEdge) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p byEdge) Less(i, j int) bool { return p[i].EdgeValue > p[j].EdgeValue }

type byDeviation []*Individual

func (p byDeviation) Len() int           { return len(p) }
func (p byDeviation) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p byDeviation) Less(i, j int) bool { return p[i].Deviation < p[j].Deviation }

type ByNSGA []Individual

func (p ByNSGA) Len() int      { return len(p) }
func (p ByNSGA) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p ByNSGA) Less(i, j int) bool {
	if p[i].Rank < p[j].Rank {
		return true
	} else if p[i].Rank == p[j].Rank && p[i].CrowdingDistance > p[j].CrowdingDistance {
		return true
	} else {
		return false
	}
}


func getRandomSegmentId(segmentMap map[int][]img.Coordinate) int {
	segId := rand.Intn(len(segmentMap))
	i := 0
	for j := range segmentMap {
		if i == segId {
			 segId = j
		}
		i++
	}
	return segId
}