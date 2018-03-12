package ga

import (
	"../ds"
	"../graph"
	"../img"
	"math"
	"sort"
	"math/rand"
)

type Graph = graph.Graph
type Edge = graph.Edge
type Vertex = graph.Vertex

type Genotype struct {
	genes []img.Direction
}

func (g *Genotype) Mutate() {
	size := len(g.genes) / 4
	toMutate := rand.Perm(len(g.genes))[:size]
	for index := range toMutate {
		g.genes[index] = img.DirectionFactory(rand.Intn(4))
	}
}

func NewGenotype(k int) Genotype {
	g := img.MyImageGraph
	disjointMap := make(map[Vertex]*ds.DisjointSet)
	segMap := make(map[*ds.DisjointSet][]Vertex)
	maxInternal := make(map[*ds.DisjointSet]float64)
	sizeMap := make(map[*ds.DisjointSet]int)
	genes := make([]img.Direction, 0)
	msf := make([]Edge, 0)

	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].Weight < g.Edges[j].Weight
	})

	for _, vertex := range g.Vertices {
		element := ds.MakeSet(vertex)
		sizeMap[element] = 1
		disjointMap[vertex] = element
		genes = append(genes, img.None)
	}

	for _, edge := range g.Edges {
		from, to := ds.FindSet(disjointMap[edge.U]), ds.FindSet(disjointMap[edge.V])
		if from != to {
			minInt := math.Min(maxInternal[from]+float64(k/sizeMap[from]), maxInternal[to]+float64(k/sizeMap[to]))
			if edge.Weight <= minInt {
				size1, size2 := sizeMap[from], sizeMap[to]
				ds.Union(disjointMap[edge.U], disjointMap[edge.V])
				maxInternal[ds.FindSet(disjointMap[edge.U])] = edge.Weight
				sizeMap[from] = size1 + size2
				msf = append(msf, edge)
			}
		}
	}

	for _, vertex := range g.Vertices {
		rep := ds.FindSet(disjointMap[vertex])
		segMap[rep] = append(segMap[rep], vertex)
	}

	segments := extractMapValues(segMap)
	adjacentMap := make(map[img.Coordinate][]img.Coordinate)
	parentMap := make(map[img.Coordinate]img.Coordinate)
	directionMap := make(map[img.Coordinate]img.Direction)

	for _, edge := range msf {
		from, to := edge.U.(img.Coordinate), edge.V.(img.Coordinate)
		adjacentMap[from] = append(adjacentMap[from], to)
		adjacentMap[to] = append(adjacentMap[to], from)
	}
	for _, segment := range segments {
		queue := []img.Coordinate{segment[0]}
		visitedOrderMap := make(map[img.Coordinate]bool)
		for len(queue) > 0 {
			v := queue[0]
			visitedOrderMap[v] = true
			queue = queue[1:]
			for _, child := range adjacentMap[v] {
				_, visitedBefore := visitedOrderMap[child]
				if !visitedBefore {
					queue = append(queue, child)
					parentMap[child] = v
				}

			}
		}
	}

	for parent, child := range parentMap {
		directionMap[parent] = img.WhichDirection(parent, child)
	}

	genes = make([]img.Direction, 0)
	myCords := make([]img.Coordinate, 0)

	for x := 0; x < img.ImageWidth; x++ {
		for y := 0; y < img.ImageHeight; y++ {
			cord := img.Coordinate{X: x, Y: y}
			genes = append(genes, directionMap[cord])
			myCords = append(myCords, cord)
		}
	}
	return Genotype{genes: genes}
}

func extractMapValues(hashMap map[*ds.DisjointSet][]Vertex) [][]img.Coordinate {
	r := make([][]img.Coordinate, 0)
	for _, v := range hashMap {
		tmp := make([]img.Coordinate, 0)
		for _, cord := range v {
			tmp = append(tmp, cord.(img.Coordinate))
		}
		r = append(r, tmp)
	}
	return r
}
