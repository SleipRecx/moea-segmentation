package graph

import (
	"../ds"
	"math"
	"sort"
	"../img"
)

type Graph struct {
	Edges    []Edge
	Coordinates []img.Coordinate
}

func (g *Graph) NewGraph() {
	coordinates := g.Coordinates
	edges := g.Edges
	cordI := 0
	edgeI := 0

	for x := 0; x < img.ImageWidth; x++ {
		for y := 0; y < img.ImageWidth; y++ {
			cord := img.Coordinate{x, y}
			coordinates[cordI] = cord
			cordI++
			right, rightError, down, downError := img.GetTwoNeighbors(cord)
			if rightError == nil {
				edges[edgeI] = Edge{cord, right, edgeWeight(cord, right)}
				edgeI++
			}
			if downError == nil {
				edges[edgeI] = Edge{cord, down, edgeWeight(cord, down)}
				edgeI++
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
		from, to := ds.FindSet(disjointMap[edge.U]), ds.FindSet(disjointMap[edge.V])
		if from != to {
			minInt := math.Min(maxInternal[from]+float64(k/sizeMap[from]), maxInternal[to]+float64(k/sizeMap[to]))
			if edge.Weight <= minInt {
				tree = append(tree, edge)
				size1, size2 := sizeMap[from], sizeMap[to]
				ds.Union(disjointMap[edge.U], disjointMap[edge.V])
				maxInternal[ds.FindSet(disjointMap[edge.U])] = edge.Weight
				sizeMap[from] = size1 + size2
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
	for cord, parent, := range parentMap {
		directions[cord] = img.WhichDirection(cord, parent)
	}
	chromosome := make([]img.Direction, len(directions))
	directionMatrix := make([][]img.Direction, img.ImageWidth)

	i := 0
	for x := 0; x < img.ImageWidth; x++ {
		directionMatrix[x] = make([]img.Direction, img.ImageHeight)
		for y := 0; y < img.ImageHeight; i++ {
			chromosome[i] = directions[img.Coordinate{x, y}]
			directionMatrix[x][y] = directions[img.Coordinate{x, y}]
			i++
		}
	}
	return chromosome, directionMatrix
}

func segmentToDirection(segmentMatrix [][]int, segmentMap map[int][]img.Coordinate) [][]img.Direction {
	directionMatrix := make([][]img.Direction, img.ImageWidth)

	for x := 0; x < img.ImageWidth; x++ {
		directionMatrix[x] = make([]img.Direction, img.ImageHeight)
		for y := 0; y < img.ImageHeight; y++ {
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

func directionToSegment(directionMatrix [][]img.Direction) ([][]int, map[int][]img.Coordinate) {
	segmentMatrix := make([][]int, img.ImageWidth)

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
	return img.ColorDistance(img.MyImage.Pixels[u.X][u.Y], img.MyImage.Pixels[v.X][v.Y])
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
