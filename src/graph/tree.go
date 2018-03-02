package graph

type Tree struct {
	Edges []Edge
}

func (t Tree) CalculateTotalCost() float64 {
	cost := 0.0
	for _, edge := range t.Edges {
		cost += edge.Weight
	}
	return cost
}
