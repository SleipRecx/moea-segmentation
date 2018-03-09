package graph

import (
	"sort"
	"math"
	"../ds"
)

type Graph struct {
	Edges    []Edge
	Vertices []Vertex
}

func (g *Graph) GraphSegmentation(k int) [][]Vertex {
	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].Weight < g.Edges[j].Weight
	})

	disjointMap := make(map[Vertex]*ds.DisjointSet)
	maxInternal := make(map[*ds.DisjointSet]float64)
	sizeMap := make(map[*ds.DisjointSet]int)

	for _, vertex := range g.Vertices {
		element := ds.MakeSet(vertex)
		sizeMap[element] = 1
		disjointMap[vertex] = element
	}

	for _, edge := range g.Edges {
		from, to := ds.FindSet(disjointMap[edge.U]), ds.FindSet(disjointMap[edge.V])
		if from != to {
			minInt := math.Min(maxInternal[from]+float64(k/sizeMap[from]), maxInternal[to]+float64(k/sizeMap[to]))
			if edge.Weight <= minInt {
				size1, size2 := sizeMap[from], sizeMap[to]
				ds.Union(disjointMap[edge.U], disjointMap[edge.V])
				maxInternal[ds.FindSet(disjointMap[edge.U])] = edge.Weight
				sizeMap[from] = size1 + size2
			}
		}
	}
	segMap := make(map[*ds.DisjointSet][]Vertex)

	for _, vertex := range g.Vertices {
		rep := ds.FindSet(disjointMap[vertex])
		segMap[rep] = append(segMap[rep], vertex)
	}
	return extractMapValues(segMap)
}

func (g *Graph) MinimalSpanningTree() []Edge {
	tree := make([]Edge, 0)
	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].Weight < g.Edges[j].Weight
	})

	disjointMap := make(map[Vertex]*ds.DisjointSet)
	for _, vertex := range g.Vertices {
		element := ds.MakeSet(vertex)
		disjointMap[vertex] = element
	}
	for _, edge := range g.Edges {
		if ds.FindSet(disjointMap[edge.U]) != ds.FindSet(disjointMap[edge.V]) {
			ds.Union(disjointMap[edge.U], disjointMap[edge.V])
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
	queue := ds.Stack{}
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

func extractMapValues(hashMap map[*ds.DisjointSet][]Vertex) [][]Vertex {
	r := make([][]Vertex, len(hashMap))
	for _, v := range hashMap {
		r = append(r, v)
	}
	return r
}
