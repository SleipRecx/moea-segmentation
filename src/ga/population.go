package ga

import (
	"fmt"

	"../graph"
	"../img"
	"math/rand"
	"time"
)

type Population struct {
	Individuals []Chromosome
	Size        int
}

func NewPopulation(size int, myImage img.Image) Population {
	individuals := make([]Chromosome, 0)
	imageGraph := myImage.ConvertToGraph()

	results := make(chan [][]graph.Vertex, size)
	jobs := make(chan int, size)
	nWorkers := 4

	for w := 1; w <= nWorkers; w++ {
		go graphSegmentWorker(imageGraph, w, jobs, results)
	}

	for j:= 1; j <= size; j++ {
		jobs <- j
	}

	for r := 1; r <= size; r++ {
		result := <- results
		individuals = append(individuals, NewChromosome(result, myImage))
	}

	return Population{Individuals: individuals, Size: size}
}

func graphSegmentWorker(imageGraph graph.Graph, id int, jobs <-chan int, results chan<- [][]graph.Vertex) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		rand.Seed(time.Now().UnixNano())
		k := rand.Intn(300 - 200) + 200
		segments := imageGraph.GraphSegmentation(k)
		fmt.Println("worker", id, "finished job", j)
		results <- segments
	}
}