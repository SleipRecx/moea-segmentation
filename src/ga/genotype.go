package ga

import (
	"../ds"
	"../graph"
	"../img"
	"math"
	"sort"

)

type Graph = graph.Graph
type Edge = graph.Edge
type Vertex = graph.Vertex

type Genotype = []Direction

func NewGenotype(myImage img.Image) Genotype {
	width, height := len(myImage.Pixels), len(myImage.Pixels[0])
	return GraphSegmentation(myImage.ConvertToGraph(), 300, width, height)
}

func GraphSegmentation(g Graph, k int, imgWidth int, imgHeight int) Genotype {
	disjointMap := make(map[Vertex]*ds.DisjointSet)
	segMap := make(map[*ds.DisjointSet][]Vertex)
	maxInternal := make(map[*ds.DisjointSet]float64)
	sizeMap := make(map[*ds.DisjointSet]int)
	genotype := make(Genotype, 0)
	msf := make([]Edge, 0)

	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].Weight < g.Edges[j].Weight
	})

	for _, vertex := range g.Vertices {
		element := ds.MakeSet(vertex)
		sizeMap[element] = 1
		disjointMap[vertex] = element
		genotype = append(genotype, None)
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
	visitedOrderMap := make(map[img.Coordinate]bool)
	parentMap := make(map[img.Coordinate]img.Coordinate)
	directionMap := make(map[img.Coordinate]Direction)

	for _, edge := range msf {
		from, to := edge.U.(img.Coordinate), edge.U.(img.Coordinate)
		adjacentMap[from] = append(adjacentMap[from], to)
		adjacentMap[from] = append(adjacentMap[from], to)
	}

	for _, segment := range segments {
		queue := []img.Coordinate{segment[0]}
		for len(queue) != 0 {
			v := queue[0]
			queue = queue[1:]
			_, visitedBefore := visitedOrderMap[v]
			if !visitedBefore {
				visitedOrderMap[v] = true
				for _, child := range adjacentMap[v] {
					queue = append(queue, child)
					parentMap[child] = v
				}
			}
		}
	}

	for parent, child := range parentMap {
		directionMap[parent] = whichDirection(child, parent)

	}

	genotype = make([]Direction, 0)

	for row := 0; row < imgWidth; row++ {
		for col := 0; col < imgHeight; col++ {
			cord := img.Coordinate{X: row, Y: col}
			genotype = append(genotype, directionMap[cord])
		}
	}
	return genotype
}

func ConvertToPhenotype(genotype Genotype, imgWidth int, imgHeight int, myImage img.Image) Phenotype {
	directionMap := make(map[img.Coordinate]Direction)
	disjointMap := make(map[img.Coordinate]*ds.DisjointSet)
	counter := 0
	for i := 0; i < imgWidth; i++ {
		for j := 0; j < imgHeight; j++ {
			n := img.Coordinate{X:i, Y:j}
			disjointMap[n] = ds.MakeSet(n)
			directionMap[n] = genotype[counter]
			counter++
		}
	}
	for cord, direction := range directionMap {
		nCord := cordFromDirection(cord, direction, imgWidth, imgHeight)
		s1, s2 := ds.FindSet(disjointMap[cord]), ds.FindSet(disjointMap[nCord])
		ds.Union(s1, s2)
	}

	segmentMap := make(map[*ds.DisjointSet][]img.Coordinate)

	for cord := range directionMap {
		rep := ds.FindSet(disjointMap[cord])
		segmentMap[rep] = append(segmentMap[rep], cord)
	}
	segments := make([][]img.Coordinate, 0)
	for _,v := range segmentMap {
		segment := make([]img.Coordinate, 0)
		for _, cord := range v {
			segment = append(segment, cord)
		}
		segments = append(segments, segment)
 	}

	return NewPhenotype(segments, myImage)
}

func cordFromDirection(cord img.Coordinate, direction Direction, imgWidth int, imgHeight int) img.Coordinate {
	newCord := img.Coordinate{X:cord.X, Y:cord.Y}
	if direction == Right {
		newCord.X = cord.X + 1
	}
	if direction == Left{
		newCord.X = cord.X - 1
	}
	if direction == Up {
		newCord.Y = cord.Y + 1
	}
	if direction == Down {
		newCord.Y = cord.Y - 1
	}
	if newCord.X >= imgWidth || newCord.Y >= imgHeight {
		return cord
	}
	if newCord.X < 0 || newCord.Y < 0 {
		return cord
	}
	return newCord
}

func whichDirection(c1, c2 img.Coordinate) Direction {
	dx, dy := c2.X-c1.X, c2.Y-c1.Y

	if dx >= 1 {
		return Right
	}
	if dx <= -1 {
		return Left
	}
	if dy >= 1 {
		return Up
	}
	return Down
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
