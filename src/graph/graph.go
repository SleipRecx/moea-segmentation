package graph

import (
	"sort"

	"../stack"
	"math"
)

type Graph struct {
	Edges    []Edge
	Vertices []Vertex
}

func (g *Graph) GraphSegmentation(k int) [][]Vertex {
	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].Weight < g.Edges[j].Weight
	})

	disjointMap := make(map[Vertex]*Element)
	maxInternal := make(map[*Element]float64)
	sizeMap := make(map[*Element]int)

	for _, vertex := range g.Vertices {
		element := MakeSet(vertex)
		sizeMap[element] = 1
		disjointMap[vertex] = element
	}

	for _, edge := range g.Edges {
		from, to := FindSet(disjointMap[edge.U]), FindSet(disjointMap[edge.V])
		if from != to {
			minInt := math.Min(maxInternal[from]+float64(k/sizeMap[from]), maxInternal[to]+float64(k/sizeMap[to]))
			if edge.Weight <= minInt {
				size1, size2 := sizeMap[from], sizeMap[to]
				Union(disjointMap[edge.U], disjointMap[edge.V])
				maxInternal[FindSet(disjointMap[edge.U])] = edge.Weight
				sizeMap[from] = size1 + size2
			}
		}
	}
	segMap := make(map[*Element][]Vertex)

	for _, vertex := range g.Vertices {
		rep := FindSet(disjointMap[vertex])
		segMap[rep] = append(segMap[rep], vertex)
	}
	return extractMapValues(segMap)
}

func (g *Graph) MinimalSpanningTree() []Edge {
	tree := make([]Edge,0)
	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].Weight < g.Edges[j].Weight
	})

	disjointMap := make(map[Vertex]*Element)
	for _, vertex := range g.Vertices {
		element := MakeSet(vertex)
		disjointMap[vertex] = element
	}
	for _, edge := range g.Edges {
		if FindSet(disjointMap[edge.U]) != FindSet(disjointMap[edge.V]) {
			Union(disjointMap[edge.U], disjointMap[edge.V])
			tree = append(tree, edge)
		}
	}
	return tree
}

func (g *Graph) DepthFirstSearch(root Vertex) {
	adjacentMap := make(map[Vertex][]Vertex)
	visitedOrderMap := make(map[Vertex]int)
	visitedCount := 0
	for _, edge := range g.Edges {
		adjacentMap[edge.U] = append(adjacentMap[edge.U], edge.V)
		adjacentMap[edge.V] = append(adjacentMap[edge.V], edge.U)
	}
	queue := stack.Stack{}
	queue.Push(root)
	for queue.Len() != 0 {
		v := queue.Pop()
		_, visitedBefore := visitedOrderMap[v]
		if !visitedBefore {
			visitedOrderMap[v] = visitedCount
			visitedCount++
			for _, adjacentVertex := range adjacentMap[v] {
				queue.Push(adjacentVertex)
			}
		}
	}
}

func extractMapValues(hashMap map[*Element][]Vertex) [][]Vertex {
	r := make([][]Vertex, len(hashMap))
	for _, v := range hashMap {
		r = append(r, v)
	}
	return r
}
