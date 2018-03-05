package main

import (
	"fmt"
	"time"

	"../image"
	"../chromosome"
)

var _ = fmt.Println

func main() {
	start := time.Now()
	var population []chromosome.Chromosome

	path := "./test_images/"
	folderNumber := "1"
	myImage := image.ReadImageFromFile(path, folderNumber)
	imageGraph := image.ImageToGraph(myImage)
	minimalSpanningTree := imageGraph.MinimalSpanningTree()
	c := chromosome.NewChromosome(minimalSpanningTree, myImage)
	population = append(population, c)
	segments := minimalSpanningTree.RandomDisjointPartition(10)
	segmentedImage := image.ReconstructImage(segments, myImage)
	image.SaveImageToFile(segmentedImage)

	var deviation float64
	for _, chromosome := range population {
		deviation += chromosome.CalcDeviation()
	}
	fmt.Println("Total deviation", deviation)
	fmt.Println("Total runtime:", time.Now().Sub(start))
}
