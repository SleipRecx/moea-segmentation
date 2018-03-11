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

	parentMap := make(map[img.Coordinate]img.Coordinate)
	directionMap := make(map[img.Coordinate]Direction)

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
		directionMap[parent] = whichDirection(parent, child)
	}

	genotype = make([]Direction, 0)
	myCords := make([]img.Coordinate, 0)

	for x := 0; x < imgWidth; x++ {
		for y := 0; y < imgHeight; y++ {
			cord := img.Coordinate{X: x, Y: y}
			genotype = append(genotype, directionMap[cord])
			myCords = append(myCords, cord)
		}
	}
	return genotype
}


func ConvertToPhenotype(genotype Genotype, imgWidth int, imgHeight int, myImage img.Image) Phenotype {
	directionMap := make(map[img.Coordinate]Direction)

	i := 0
	for x := 0; x < imgWidth; x++ {
		for y := 0; y < imgHeight; y++ {
			cord := img.Coordinate{X: x, Y: y}
			directionMap[cord] = genotype[i]
			i++
		}
	}
	disjointMap := make(map[img.Coordinate]*ds.DisjointSet)
	for k := range directionMap {
		set := ds.MakeSet(k)
		disjointMap[k] = set
	}

	for cord, direction := range directionMap {
		nCord := cordFromDirection(cord, direction, imgWidth, imgHeight)
		s1, s2 := ds.FindSet(disjointMap[cord]), ds.FindSet(disjointMap[nCord])
		ds.Union(s1, s2)
	}

	segmentMap := make(map[*ds.DisjointSet][]img.Coordinate)

	for k := range directionMap{
		segmentMap[ds.FindSet(disjointMap[k])] = append(segmentMap[ds.FindSet(disjointMap[k])], k)
	}
	segments := make([][]img.Coordinate, 0)
	for _, value := range segmentMap{
		segments = append(segments, value)
	}
	return NewPhenotype(segments, myImage)
}

/*
func nodeAndDirectionToNode(node Node, direction Direction) Node {
	var n Node
	switch direction {
	case Up:
		n = Node{node.X, node.Y - 1}
	case Down:
		n = Node{node.X, node.Y + 1}
	case Right:
		n = Node{node.X + 1, node.Y}
	case Left:
		n = Node{node.X - 1, node.Y}
	default:
		return Node{node.X, node.Y}
	}
	if n.isInRange() {
		return n
	} else {
		return node
	}
}
 */

func cordFromDirection(cord img.Coordinate, direction Direction, imgWidth int, imgHeight int) img.Coordinate {
	newCord := img.Coordinate{X: cord.X, Y: cord.Y}
	switch direction {
	case Up:
		newCord.Y -= 1
	case Down:
		newCord.Y += 1
	case Right:
		newCord.X += 1
	case Left:
		newCord.X -= 1
	}
	if newCord.X >= imgWidth || newCord.Y >= imgHeight{
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
		return Down
	}
	if dy <= -1 {
		return Up
	}
	return None
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
