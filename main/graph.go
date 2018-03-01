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
	From     Node
	To       Node
	Weight   float64
	Directed bool
}

type Graph struct {
	Edges    []Edge
	Vertices []Node
}

func (e Edge) String() string {
	return fmt.Sprintf("From: %d, To: %d, Weight: %f, Directed: %t", e.From, e.To, e.Weight, e.Directed)
}

func imageToGraph(myImage Image) Graph {
	var edges []Edge
	var verticies []Node
	pixels := myImage.pixels
	for i := range pixels {
		for j := range pixels[i] {
			verticies = append(verticies, Node{X: i, Y: j})
			from := pixels[i][j]
			if j+1 < len(pixels[i]) {
				to := pixels[i][j+1]
				edge := Edge{
					From: Node{X: i, Y: j},
					To: Node{X: i, Y: j + 1},
					Weight: from.distance(to),
					Directed: false}
				edges = append(edges, edge)
			}
			if i+1 < len(pixels) {
				to := pixels[i+1][j]
				edge := Edge{
					From: Node{X: i, Y: j},
					To: Node{X: i + 1, Y: j},
					Weight: from.distance(to),
					Directed: false}
				edges = append(edges, edge)
			}
		}
	}
	return Graph{Edges: edges, Vertices: verticies}
}

func (g Graph) minimalSpanningTree() []Edge {

	var edges = g.Edges[:]
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	F := map[Node]bool{}

	for _, node := range g.Vertices {
		F[node] = true
	}


	/*
	for _, edge := range edges {
		fmt.Println(edge)
	}
	*/

	return []Edge{}
}
