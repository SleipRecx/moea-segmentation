package ga

import (
	"../graph"
)

type Population struct {
	Generation int
	Individuals []Individual
	Fronts [][]*Individual
	Size        int
}

func (population Population) Init(size int, graph graph.Graph) {
	population.Size = size
	for j := 0; j < size; j++ {
		individual := Individual{}
		individual.Init(graph)
		population.Individuals = append(population.Individuals, individual)
	}
}

/*
func (p *Population) UniformCrossover(parent1, parent2 *Individual) (Individual, Individual) {
	numGenes := len(parent1.Genotype.genes)
	crossoverPoints := rand.Perm(numGenes)[0:int(numGenes/2)]
	sort.Ints(crossoverPoints)
	genes1, genes2 := make([]img.Direction, 0), make([]img.Direction, 0)

	prev := 0
	parents := []Individual{*parent1, *parent2}
	for i, n := range crossoverPoints {
		if i == len(crossoverPoints) - 1 {
			n = numGenes
		}
		genes1 = append(genes1, parents[0].Genotype.genes[prev:n]...)
		genes2 = append(genes2, parents[1].Genotype.genes[prev:n]...)
		parents[0], parents[1] = parents[1], parents[0]
		prev = n
	}

	geno1, geno2 := Genotype{genes:genes1}, Genotype{genes:genes2}
	individual1 := Individual{Genotype:geno1, Phenotype:NewPhenotype(geno1)}
	individual2 := Individual{Genotype:geno2, Phenotype:NewPhenotype(geno2)}
	return individual1, individual2
}


func NewPopulation(size int) Population {
	individuals := make([]Individual, 0)
	results := make(chan Genotype, size)
	jobs := make(chan int, size)
	nWorkers := 4

	for w := 1; w <= nWorkers; w++ {
		go initPopulationWorker(w, jobs, results)
	}

	for j := 1; j <= size; j++ {
		jobs <- j
	}

	for r := 1; r <= size; r++ {
		genotype := <-results
		individual := Individual{Genotype:genotype, Phenotype:NewPhenotype(genotype)}
		individuals = append(individuals, individual)
	}

	return Population{Individuals: individuals, Size: size}
}

func initPopulationWorker(id int, jobs <-chan int, results chan<- Genotype) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		k := rand.Intn(8000-200) + 200
		genotype := NewGenotype(k)

		fmt.Println("worker", id, "finished job", j)
		results <- genotype
	}
}
*/