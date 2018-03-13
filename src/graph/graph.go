package graph

import (
	"math"
	"sort"
	"../img"
	"../constants"
	"../ds"
	"fmt"
)

type Graph struct {
	Edges    []Edge
	Coordinates []img.Coordinate
}

func (g *Graph) Init() {
	coordinates := make([]img.Coordinate, constants.ImageWidth*constants.ImageHeight)
	edges := make([]Edge, (constants.ImageWidth-1)*constants.ImageHeight+(constants.ImageHeight-1)*constants.ImageWidth)
	nodeI := 0
	edgeI := 0

	for x := 0; x < constants.ImageWidth; x++ {
		for y := 0; y < constants.ImageHeight; y++ {
			cord := img.Coordinate{x, y}
			coordinates[nodeI] = cord
			nodeI += 1
			right, rightErr, down, downErr := img.GetTwoNeighbors(cord)
			if rightErr == nil {
				edges[edgeI] = Edge{cord, right, edgeWeight(cord, right)}
				edgeI += 1

			}
			if downErr == nil {
				edges[edgeI] = Edge{cord, down, edgeWeight(cord, down)}
				edgeI += 1
			}

		}
	}
	sort.Sort(ByWeight(edges))
	g.Coordinates = coordinates
	g.Edges = edges
}

func (g *Graph) GraphSegmentation(k int) ([]img.Direction, [][]img.Direction) {
	coordinates := g.Coordinates
	edges := g.Edges
	tree := make([]Edge, 0)
	disjointMap := make(map[img.Coordinate]*ds.DisjointSet)
	maxInternal := make(map[*ds.DisjointSet]float64)
	sizeMap := make(map[*ds.DisjointSet]int)

	for i := 0; i < len(coordinates); i++ {
		disjointSet := ds.DisjointSet{Parent: nil, Rank: 0}
		sizeMap[&disjointSet] = 1
		disjointMap[coordinates[i]] = &disjointSet
	}

	for _, edge := range edges {
		u, v := ds.FindSet(disjointMap[edge.U]), ds.FindSet(disjointMap[edge.V])
		if u != v {
			minInt := math.Min(maxInternal[u]+float64(k/sizeMap[u]), maxInternal[v]+float64(k/sizeMap[v]))
			if edge.Weight <= minInt {
				tree = append(tree, edge)
				size1, size2 := sizeMap[u], sizeMap[v]
				ds.Union(disjointMap[edge.U], disjointMap[edge.V])
				maxInternal[ds.FindSet(disjointMap[edge.U])] = edge.Weight
				sizeMap[u] = size1 + size2
			}
		}
	}

	directions := make(map[img.Coordinate]img.Direction)
	parentMap := make(map[img.Coordinate]img.Coordinate)

	for _, cord := range g.Coordinates {
		parentMap[cord] = cord
	}

	adjacentMap := make(map[img.Coordinate][]img.Coordinate)

	for _, edge := range tree {
		adjacentMap[edge.U] = append(adjacentMap[edge.U], edge.V)
		adjacentMap[edge.V] = append(adjacentMap[edge.V], edge.U)
	}

	segmentSet := make(map[*ds.DisjointSet][]img.Coordinate)
	for _, k := range g.Coordinates {
		segmentSet[ds.FindSet(disjointMap[k])] = append(segmentSet[ds.FindSet(disjointMap[k])], k)
	}
	segments := make([][]img.Coordinate, 0)
	for _, item := range segmentSet {
		segments = append(segments, item)
	}

	//for _, segment := range segments {
	//	visited := make(map[Coordinate]bool)
	//	queue := []Coordinate{segment[0]}
	//	for len(queue)
	//}

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
	fmt.Println("Test")
	for cord, parent := range parentMap {
		directions[cord] = img.WhichDirection(cord, parent)
	}
	chromosome := make([]img.Direction, len(directions))
	directionMatrix := make([][]img.Direction, constants.ImageWidth)

	i := 0
	for x := 0; x < constants.ImageWidth; x++ {
		directionMatrix[x] = make([]img.Direction, constants.ImageHeight)
		for y := 0; y < constants.ImageHeight; i++ {
			chromosome[i] = directions[img.Coordinate{x, y}]
			directionMatrix[x][y] = directions[img.Coordinate{x, y}]
			i++
		}
	}
	return chromosome, directionMatrix
}

func SegmentToDirection(segmentMatrix [][]int, segmentMap map[int][]img.Coordinate) [][]img.Direction {
	directionMatrix := make([][]img.Direction, constants.ImageWidth)

	for x := 0; x < constants.ImageWidth; x++ {
		directionMatrix[x] = make([]img.Direction, constants.ImageHeight)
		for y := 0; y < constants.ImageHeight; y++ {
			directionMatrix[x][y] = img.None
		}
	}
	for segId := range segmentMap {
		visited := make(map[img.Coordinate]bool)
		opened := make(map[img.Coordinate]bool)
		queue := []img.Coordinate{segmentMap[segId][0]}
		for len(queue) > 0 {
			cord := queue[0]
			visited[cord] = true
			queue = queue[1:]
			neighbors := img.GetFourNeighbors(cord)
			for _, neighbor := range neighbors {
				if _, visited := visited[neighbor]; visited {
					continue
				}
				if _, o := opened[neighbor]; !o {
					if segmentMatrix[neighbor.X][neighbor.Y] == segmentMatrix[cord.X][cord.Y] {
						queue = append(queue, neighbor)
						directionMatrix[neighbor.X][neighbor.Y] = img.WhichDirection(neighbor, cord)
						opened[neighbor] = true
					}
				}
			}

		}
	}
	return directionMatrix
}

func DirectionToSegment(directionMatrix [][]img.Direction) ([][]int, map[int][]img.Coordinate) {
	segmentMatrix := make([][]int, constants.ImageWidth)

	dsSets := make(map[img.Coordinate]*ds.DisjointSet)
	for x := range segmentMatrix {
		segmentMatrix[x] = make([]int, len(segmentMatrix[x]))
		for y := range segmentMatrix[x] {
			set := ds.DisjointSet{Parent: nil, Rank: 0}
			dsSets[img.Coordinate{x, y}] = &set
		}
	}
	for x := range directionMatrix {
		for y := range directionMatrix {
			cord := img.Coordinate{x, y}
			neighbor := img.CordFromDirection(cord, directionMatrix[x][y])
			set1, set2 := ds.FindSet(dsSets[cord]), ds.FindSet(dsSets[neighbor])
			ds.Union(set1, set2)
		}
	}
	setToSegment := make(map[*ds.DisjointSet][]img.Coordinate)
	for x := range directionMatrix {
		for y := range directionMatrix[x] {
			cord := img.Coordinate{x, y}
			set := ds.FindSet(dsSets[cord])
			setToSegment[set] = append(setToSegment[set], cord)
		}
	}
	segId := 0
	segmentMap := make(map[int][]img.Coordinate)
	for set := range setToSegment {
		for i := range setToSegment[set] {
			cord := setToSegment[set][i]
			segmentMatrix[cord.X][cord.Y] = segId
		}
		segmentMap[segId] = setToSegment[set]
		segId++
	}
	return segmentMatrix, segmentMap
}

func edgeWeight(u, v img.Coordinate) float64 {
	return img.MyImage.Pixels[u.X][u.Y].Distance(img.MyImage.Pixels[v.X][v.Y])
}

//func (g *Graph) MinimalSpanningTree() []Edge {
//	tree := make([]Edge, 0)
//	sort.Slice(g.Edges, func(i, j int) bool {
//		return g.Edges[i].Weight < g.Edges[j].Weight
//	})
//
//	disjointMap := make(map[Vertex]*ds.DisjointSet)
//	for _, vertex := range g.Vertices {
//		element := ds.MakeSet(vertex)
//		disjointMap[vertex] = element
//	}
//	for _, edge := range g.Edges {
//		if ds.FindSet(disjointMap[edge.U]) != ds.FindSet(disjointMap[edge.V]) {
//			ds.Union(disjointMap[edge.U], disjointMap[edge.V])
//			tree = append(tree, edge)
//		}
//	}
//	return tree
//}

//func (g *Graph) DepthFirstSearch(root Vertex) {
//	adjacentMap := make(map[Vertex][]Vertex)
//	visitedOrderMap := make(map[Vertex]int)
//	visitedCount := 0
//	for _, edge := range g.Edges {
//		adjacentMap[edge.U] = append(adjacentMap[edge.U], edge.V)
//		adjacentMap[edge.V] = append(adjacentMap[edge.V], edge.U)
//	}
//	queue := ds.Stack{}
//	queue.Push(root)
//	for queue.Len() != 0 {
//		v := queue.Pop()
//		_, visitedBefore := visitedOrderMap[v]
//		if !visitedBefore {
//			visitedOrderMap[v] = visitedCount
//			visitedCount++
//			for _, adjacentVertex := range adjacentMap[v] {
//				queue.Push(adjacentVertex)
//			}
//		}
//	}
//}
//
//func extractMapValues(hashMap map[*ds.DisjointSet][]Vertex) [][]Vertex {
//	r := make([][]Vertex, len(hashMap))
//	for _, v := range hashMap {
//		r = append(r, v)
//	}
//	return r
//}
