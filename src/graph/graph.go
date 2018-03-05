package graph

import (
	"sort"

	"../stack"
)

type Graph struct {
	Edges    []Edge
	Vertices []Vertex
}


func (g Graph) MinimalSpanningTree() Tree {
	var tree Tree
	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].Weight < g.Edges[j].Weight
	})

	var disjointMap = make(map[Vertex]*Element)
	for _, vertex := range g.Vertices {
		element := MakeSet(vertex)
		disjointMap[vertex] = element
	}

	for _, edge := range g.Edges {
		if FindSet(disjointMap[edge.U]) != FindSet(disjointMap[edge.V]) {
			Union(disjointMap[edge.U], disjointMap[edge.V])
			tree.Edges = append(tree.Edges, edge)
		}
	}
	return tree
}

func (g Graph) DepthFirstSearch(root Vertex) {
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
