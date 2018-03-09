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

	//segmentedImage := image.ReconstructImage(segments, myImage)
	//image.SaveImageToFile(segmentedImage)
	fmt.Println("Total runtime:", time.Now().Sub(start))

}
