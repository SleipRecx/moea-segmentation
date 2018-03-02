package graph

import (
	"sort"
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
