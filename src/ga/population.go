package ga

import (
	"fmt"
	"math/rand"
	"time"
)

type Population struct {
	Individuals []Phenotype
	Size        int
}

func NewPopulation(size int) Population {
	individuals := make([]Phenotype, 0)
	results := make(chan Phenotype, size)
	jobs := make(chan int, size)
	nWorkers := 4

	for w := 1; w <= nWorkers; w++ {
		go initPopulationWorker(w, jobs, results)
	}

	for j := 1; j <= size; j++ {
		jobs <- j
	}

	for r := 1; r <= size; r++ {
		individual := <-results
		individuals = append(individuals, individual)
	}

	return Population{Individuals: individuals, Size: size}
}



func initPopulationWorker(id int, jobs <-chan int, results chan<- Phenotype) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		rand.Seed(time.Now().UnixNano())
		k := rand.Intn(10000 - 200) + 200
		individual := NewPhenotype(k)
		fmt.Println("worker", id, "finished job", j)
		results <- individual
	}
}