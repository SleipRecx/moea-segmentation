package graph

import "fmt"

type Edge struct {
	U, V   Vertex
	Weight float64
}

func (e Edge) String() string {
	return fmt.Sprintf("%v <--> %v, Weight: %f", e.U, e.V, e.Weight)
}
