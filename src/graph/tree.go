package graph

import (
	"math/rand"
	"sort"
)

type Tree struct {
	Edges []Edge
}

func (t Tree) CalculateTotalCost() float64 {
	cost := 0.0
	for _, edge := range t.Edges {
		cost += edge.Weight
	}
	return cost
}

// Todo: Refactor this (Corry, do not touch)
func (t Tree) RandomDisjointPartition(n int) [][]Vertex {

	adjacentMap := make(map[Vertex][]Vertex)
	for _, edge := range t.Edges {
		adjacentMap[edge.U] = append(adjacentMap[edge.U], edge.V)
		adjacentMap[edge.V] = append(adjacentMap[edge.V], edge.U)
	}
	rand.Seed(123)
	splitPoints := rand.Perm(len(adjacentMap))[0 : n-1]
	partition := make([][]Vertex, n, len(adjacentMap))
	sort.Ints(splitPoints)
	currentSplit := 0

	visitedOrderMap := make(map[Vertex]int)
	visitedCount := 0
	queue := make([]Vertex, 0)
	queue = append(queue, t.Edges[0].U)
	for len(queue) != 0 {
		v := queue[0]
		queue = queue[1:]
		_, visitedBefore := visitedOrderMap[v]
		if !visitedBefore {
			visitedOrderMap[v] = visitedCount
			partition[currentSplit] = append(partition[currentSplit], v)
			if currentSplit < len(splitPoints) {
				if visitedCount == splitPoints[currentSplit] {
					currentSplit++
				}
			}
			visitedCount++
			for _, adjacentVertex := range adjacentMap[v] {
				queue = append(queue, adjacentVertex)
			}
		}
	}
	return partition
}
