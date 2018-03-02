package graph

import (
	"fmt"
	"sort"
)

type Node = interface{}

type Edge struct {
	U, V   Node
	Weight float64
}

func (e Edge) String() string {
	return fmt.Sprintf("%v <--> %v, Weight: %f", e.U, e.V, e.Weight)
}

type Tree struct {
	Edges []Edge
}

type Graph struct {
	Edges    []Edge
	Vertices []Node
}

func (g Graph) MinimalSpanningTree() Tree {
	var tree Tree
	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].Weight < g.Edges[j].Weight
	})

	var disjointMap = make(map[Node]*Element)
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

func (g Graph) makeForest() []Graph {
	for _, edge := range g.Edges {
		fmt.Println(edge)
	}
	return []Graph{}
}
