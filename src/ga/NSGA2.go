package ga

import (
	"log"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"../constants"
	"../img"
)

func (population *Population) Evolve() {
	for i := 0; i < population.Size / 2; i++ {
		tournamentIndices := rand.Perm(population.Size)
		p1Tournament, p2Tournament := tournamentIndices[0:2], tournamentIndices[2:4]
		p1Index := int(math.Min(float64(p1Tournament[0]), float64(p1Tournament[1])))
		p2Index := int(math.Min(float64(p2Tournament[0]), float64(p2Tournament[1])))
		c1, c2 := population.Individuals[p1Index].Crossover(population.Individuals[p2Index])
		population.Individuals = append(population.Individuals, c1, c2)
	}

	population.NonDominatedSort()
	population.CrowdingDistance()
	sort.Sort(ByNSGA(population.Individuals))
	population.Individuals = population.Individuals[:population.Size]
	population.NonDominatedSort()
	population.CrowdingDistance()
	sort.Sort(ByNSGA(population.Individuals))

}

func (population *Population) NSGA2Run() {
	for i := 0; i < constants.NGenerations; i++ {
		population.Evolve()
		log.Println("Generation:", i, "Number of fronts:", len(population.Fronts), "Size of front 1:", len(population.Fronts[0]))
	}
	for i := range population.Fronts[0] {
		img.SaveEdgeDetectionImage(population.Fronts[0][i].SegmentIDMatrix, strconv.Itoa(len(population.Fronts[0][i].SegmentMap))+ "_" + strconv.Itoa(i))
	}
}

func (population *Population) NonDominatedSort() {
	fronts := make([][]*Individual, 0)
	fronts = append(fronts, make([]*Individual, 0))

	for i := range population.Individuals {
		individual := &population.Individuals[i]
		dominatedByIndividual := make([]*Individual, 0)
		numDominatingIndividual := 0
		for j := range population.Individuals {
			if i == j {
				continue
			} else if individual.IsDominating(&population.Individuals[i]) {
				dominatedByIndividual = append(dominatedByIndividual, &population.Individuals[j])
			} else if population.Individuals[j].IsDominating(individual) {
				numDominatingIndividual += 1
			}
		}
		individual.Dominates = dominatedByIndividual
		individual.DominatedByN = numDominatingIndividual

		if numDominatingIndividual == 0 {
			individual.Rank = 1
			fronts[0] = append(fronts[0], individual)
		}
	}
	frontCounter := 0
	for len(fronts[frontCounter]) > 0 {
		nextFront := make([]*Individual, 0)
		for i := range fronts[frontCounter] {
			individual := fronts[frontCounter][i]
			for j := range individual.Dominates {
				dominatedIndividual := individual.Dominates[j]
				dominatedIndividual.DominatedByN -= 1
				if dominatedIndividual.DominatedByN == 0 {
					dominatedIndividual.Rank = frontCounter + 2
					nextFront = append(nextFront, dominatedIndividual)
				}
			}
		}
		fronts = append(fronts, nextFront)
		frontCounter++
	}
	population.Fronts = fronts
}

func (population *Population) CrowdingDistance() {
	maxDeviation, minDeviation, maxEdgeValue, minEdgeValue := population.getMaxMinValues()
	deltaDeviation, deltaEdgeValue := maxDeviation - minDeviation, maxEdgeValue - minEdgeValue
	for i := range population.Fronts {
		if len(population.Fronts[i]) == 0 {
			continue
		}
		for j := range population.Fronts[i] {
			population.Fronts[i][j].CrowdingDistance = 0
		}
		sort.Sort(byDeviation(population.Fronts[i]))
		population.Fronts[i][0].CrowdingDistance = math.Inf(1)
		population.Fronts[i][len(population.Fronts[i]) - 1].CrowdingDistance = math.Inf(1)
		for j := range population.Fronts[i] {
			population.Fronts[i][j].CrowdingDistance = population.Fronts[i][j].CrowdingDistance + ((population.Fronts[i][j+1].Deviation - population.Fronts[i][j+1].Deviation) / deltaDeviation)
		}
		sort.Sort(byEdge(population.Fronts[i]))
		population.Fronts[i][0].CrowdingDistance = math.Inf(1)
		population.Fronts[i][len(population.Fronts) - 1].CrowdingDistance = math.Inf(1)
		for j := 0; j < len(population.Fronts) - 1; j++ {
			population.Fronts[i][j].CrowdingDistance = population.Fronts[i][j].CrowdingDistance + ((population.Fronts[i][j-1].EdgeValue - population.Fronts[i][j+1].EdgeValue) / deltaEdgeValue)
		}
	}
}

func (population *Population) getMaxMinValues() (float64, float64, float64, float64) {
	maxDeviation, maxEdgeValue := 0.0, 0.0
	minDeviation, minEgeValue := math.Inf(1), math.Inf(1)
	for i := range population.Individuals {
		if population.Individuals[i].EdgeValue > maxEdgeValue {
			maxEdgeValue = population.Individuals[i].EdgeValue
		}
		if population.Individuals[i].EdgeValue < minEgeValue {
			minEgeValue = population.Individuals[i].EdgeValue
		}
		if population.Individuals[i].Deviation > maxDeviation {
			maxDeviation = population.Individuals[i].Deviation
		}
		if population.Individuals[i].Deviation < minDeviation {
			minDeviation = population.Individuals[i].Deviation
		}
	}
	return maxDeviation, minDeviation, maxEdgeValue, minEgeValue
}
