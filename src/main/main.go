
package main

import (
	"fmt"
	"time"

	"../chromosome"
	"../image"
)

func main() {
	start := time.Now()
	path := "./test_images/"
	folderNumber := "3"

	myImage := image.ReadImageFromFile(path, folderNumber)
	imageGraph := myImage.ConvertToGraph()
	segments := imageGraph.GraphSegmentation(1000)
	c := chromosome.NewChromosome(segments, myImage)


	fmt.Println(c.CalcEdgeValue())
	fmt.Println("Total runtime:", time.Now().Sub(start))

	segmentedImage := image.ReconstructImage(c.Segments, myImage)
	image.SaveImageToFile(segmentedImage)



}

