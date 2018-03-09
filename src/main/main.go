
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
	segments := imageGraph.GraphSegmentation(100000)
	c := chromosome.NewChromosome(segments, myImage)

	segmentedImage := image.ReconstructImage(c.Segments, myImage)
	image.SaveImageToFile(segmentedImage)
	fmt.Println("Total runtime:", time.Now().Sub(start))
	fmt.Println(c.CalcEdgeValue())
	fmt.Println("Total runtime:", time.Now().Sub(start))







}

