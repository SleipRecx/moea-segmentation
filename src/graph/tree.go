package graph

import (
	"math/rand"
	"sort"
	"time"

	"../stack"
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
	rand.Seed(time.Now().UnixNano())
	splitPoints := rand.Perm(len(adjacentMap))[0:n-1]
	partition := make([][]Vertex,n,len(adjacentMap))
	sort.Ints(splitPoints)
	currentSplit := 0

	visitedOrderMap := make(map[Vertex]int)
	visitedCount := 0
	queue := stack.Stack{}
	queue.Push(t.Edges[0].U)
	for queue.Len() != 0 {
		v := queue.Pop()
		_, visitedBefore := visitedOrderMap[v]
		if !visitedBefore {
			visitedOrderMap[v] = visitedCount
			partition[currentSplit] = append(partition[currentSplit], v)
			if currentSplit < len(splitPoints){
				if visitedCount == splitPoints[currentSplit] {
					currentSplit ++
				}
			}
			visitedCount++
			for _, adjacentVertex := range adjacentMap[v] {
				queue.Push(adjacentVertex)
			}
		}
	}

	return partition
}
