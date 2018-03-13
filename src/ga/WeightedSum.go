package ga

import (
	"../constants"
	"../img"
	"log"
	"math/rand"
	"sort"
)

func (population *Population) WeightedSumEvolve() {
	sort.Sort(byFitness(population.Individuals))
	population.Generation += 1
	newPopulation := make([]Individual, 0)
	newPopulation = append(newPopulation, population.Individuals[0:constants.NElites]...)

	for j := 0; j < population.Size / 2; j++ {
		tournamentIndices := rand.Perm(population.Size)[0:4]
		tournament := make([]Individual, 0)
		for _, index := range tournamentIndices {
			tournament = append(tournament, population.Individuals[index])
		}
		sort.Sort(byFitness(tournament))
		c1, c2 := tournament[0].Crossover(tournament[1])
		newPopulation = append(newPopulation, c1, c2)
	}
	population.Individuals = newPopulation[:population.Size]
	sort.Sort(byFitness(population.Individuals))
}

func (population *Population) WeightedSumRun() {
	for i := 0; i < constants.NGenerations; i++ {
		population.WeightedSumEvolve()
		log.Println("Generation", i, "Best weighted sum:", population.Individuals[0].Weighted, "Worst:", population.Individuals[len(population.Individuals)-1].Weighted)
		log.Println("Num Segments:",len(population.Individuals[0].SegmentMap))
		img.SaveEdgeDetectionImage(population.Individuals[0].SegmentIDMatrix, constants.FolderNumber)
	}
}