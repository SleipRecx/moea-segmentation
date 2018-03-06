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
	segmentedImage := image.ReconstructImage(c.Segments, myImage)
	image.SaveImageToFile(segmentedImage)

	fmt.Println("Total deviation", c.CalcDeviation())
	fmt.Println("Total edge value", c.CalcEdgeValue())
	fmt.Println("Total runtime:", time.Now().Sub(start))
}
