package graph

import (
	"fmt"
	"sort"
)

type Node = interface{}

type Edge struct {
	U      Node
	V      Node
	Weight float64
}

func (e Edge) String() string {
	return fmt.Sprintf("%v <--> %v, Weight: %f", e.U, e.V, e.Weight)
}

type Graph struct {
	Edges    []Edge
	Vertices []Node
}

func (g Graph) MinimalSpanningTree() []Edge {
	var tree []Edge
	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].Weight < g.Edges[j].Weight
	})

	var myMap = make(map[Node]*Element)
	for _, vertex := range g.Vertices {
		element := MakeSet(vertex)
		myMap[vertex] = element
	}

	for _, edge := range g.Edges {
		if FindSet(myMap[edge.U]) != FindSet(myMap[edge.V]) {
			Union(myMap[edge.U], myMap[edge.V])
			tree = append(tree, edge)
		}
	}
	return tree
}

func (g Graph) makeForest() []Graph {
	for _, edge := range g.Edges {
		fmt.Println(edge)
	}
	return []Graph{}
}
