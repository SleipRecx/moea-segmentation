package graph

import (
	"fmt"
	"../img"
)

type Edge struct {
	U, V   img.Coordinate
	Weight float64
}

// For sorting by distance
type ByWeight []Edge

func (s ByWeight) Len() int {
	return len(s)
}

// Swaps the places of two edges in the ByWeight array
func (s ByWeight) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Boolean check to compare the cost of two edges
func (s ByWeight) Less(i, j int) bool {
	return s[i].Weight < s[j].Weight
}

func (e Edge) String() string {
	return fmt.Sprintf("%v <--> %v, Weight: %f", e.U, e.V, e.Weight)
}
