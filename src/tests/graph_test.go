package tests

import (
	"../graph"
	"testing"
)

func TestMinimalSpanningTree(t *testing.T) {
	vertices := []graph.Vertex{0, 1, 2, 3, 4, 5, 6, 7, 8}
	edges := []graph.Edge{
		{U: 7, V: 6, Weight: 1},
		{U: 8, V: 2, Weight: 2},
		{U: 6, V: 5, Weight: 2},
		{U: 0, V: 1, Weight: 4},
		{U: 2, V: 5, Weight: 4},
		{U: 8, V: 6, Weight: 6},
		{U: 2, V: 3, Weight: 7},
		{U: 7, V: 8, Weight: 7},
		{U: 0, V: 7, Weight: 8},
		{U: 1, V: 2, Weight: 8},
		{U: 3, V: 4, Weight: 9},
		{U: 5, V: 4, Weight: 10},
		{U: 1, V: 7, Weight: 11},
		{U: 3, V: 5, Weight: 14}}

	g := graph.Graph{Vertices: vertices, Edges: edges}
	mst := g.MinimalSpanningTree()

	expected := 37.0
	actual := mst.CalculateTotalCost()
	if expected != actual {
		t.Errorf("Expected: %f Was: %f", expected, actual)
	}
}
