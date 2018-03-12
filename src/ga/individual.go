package ga


type Individual struct {
	Genotype  Genotype
	Phenotype Phenotype
	Fitness   float64
}

func NewIndividual(k int) Individual {
	genotype := NewGenotype(k)
	phenotype := NewPhenotype(genotype)
	return Individual{Genotype:genotype, Phenotype:phenotype}
}

