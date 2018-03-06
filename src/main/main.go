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

	path := "./test_images/"
	folderNumber := "1"
	myImage := image.ReadImageFromFile(path, folderNumber)
	imageGraph := image.ImageToGraph(myImage)
	mst := imageGraph.MinimalSpanningTree()
	c := chromosome.NewChromosome(mst, myImage, 100)

	/*
	BUG: code below terminates.
	c := chromosome.NewChromosome(mst, myImage, 321)

	BUG: code below gives index out of range.
	c := chromosome.NewChromosome(mst, myImage, 322)
	*/

	fmt.Println("Total deviation", c.CalcDeviation())
	fmt.Println("Total edge value", c.CalcEdgeValue())
	fmt.Println("Total runtime:", time.Now().Sub(start))

	/*
	segments := minimalSpanningTree.RandomDisjointPartition(1000)
	segmentedImage := image.ReconstructImage(segments, myImage)
	image.SaveImageToFile(segmentedImage)
	*/
}
