package ga

import (
	"fmt"

	"../graph"
	"../img"
	"math/rand"
	"time"
)

type Population struct {
	Individuals []Phenotype
	Size        int
}

func NewPopulation(size int, myImage img.Image) Population {
	individuals := make([]Phenotype, 0)
	imageGraph := myImage.ConvertToGraph()

	results := make(chan [][]img.Coordinate, size)
	jobs := make(chan int, size)
	nWorkers := 4

	for w := 1; w <= nWorkers; w++ {
		go graphSegmentWorker(imageGraph, w, jobs, results)
	}

	for j := 1; j <= size; j++ {
		jobs <- j
	}

	for r := 1; r <= size; r++ {
		result := <-results
		individuals = append(individuals, NewPhenotype(result, myImage))
	}

	return Population{Individuals: individuals, Size: size}
}

func graphSegmentWorker(imageGraph graph.Graph, id int, jobs <-chan int, results chan<- [][]img.Coordinate) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		rand.Seed(time.Now().UnixNano())
		k := rand.Intn(300-200) + 200
		segments := imageGraph.GraphSegmentation(k)
		fmt.Println("worker", id, "finished job", j)
		results <- mapVertexToCoordinate(segments)
	}
}

func mapVertexToCoordinate(partitions [][]graph.Vertex) [][]img.Coordinate {
	segments := make([][]img.Coordinate, 0, 0)
	for i := range partitions {
		seg := make([]img.Coordinate, 0)
		for j := range partitions[i] {
			seg = append(seg, partitions[i][j].(img.Coordinate))
		}
		segments = append(segments, seg)
	}
	return segments
}
