package main

import (
	"fmt"
	"sort"
)

type Node struct {
	X int
	Y int
}

type Edge struct {
	U      interface{}
	V      interface{}
	Weight float64
}

func (e Edge) String() string {
	return fmt.Sprintf("%v <--> %v, Weight: %f", e.U, e.V, e.Weight)
}

type Graph struct {
	Edges    []Edge
	Vertices []interface{}
}

func (g Graph) minimalSpanningTree() []Edge {
	var tree []Edge
	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].Weight < g.Edges[j].Weight
	})

	var myMap = make(map[interface{}]*Element)
	for _, vertex := range g.Vertices {
		element := makeSet(vertex)
		myMap[vertex] = element
	}

	for _, edge := range g.Edges {
		if findSet(myMap[edge.U]) != findSet(myMap[edge.V]) {
			union(myMap[edge.U], myMap[edge.V])
			tree = append(tree, edge)
		}
	}
	return tree
}
