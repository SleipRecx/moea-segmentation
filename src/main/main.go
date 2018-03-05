package main

import (
	"fmt"
	"time"

	"../image"
)

var _ = fmt.Println

func main() {
	start := time.Now()

	path := "./test_images/"
	folderNumber := "1"
	myImage := image.ReadImageFromFile(path, folderNumber)
	imageGraph := image.ImageToGraph(myImage)
	minimalSpanningTree := imageGraph.MinimalSpanningTree()
	segments := minimalSpanningTree.RandomDisjointPartition(10)
	segmentedImage := image.ReconstructImage(segments, myImage)
	image.SaveImageToFile(segmentedImage)

	fmt.Println("Total runtime:", time.Now().Sub(start))
}
