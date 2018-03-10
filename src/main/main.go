package main

import (
	"fmt"
	"time"
	"strconv"

	"../ga"
	"../img"
)

func main() {
	start := time.Now()
	path := "./test_images/"
	folderNumber := "1"

	myImage := img.ReadImageFromFile(path, folderNumber)
	population := ga.NewPopulation(10, myImage)
	for i := range population.Individuals {
		segImage := img.ReconstructImage(population.Individuals[i].Segments, myImage)
		img.SaveImageToFile(segImage, "output/img" + strconv.Itoa(i))
	}
	fmt.Println("Total runtime:", time.Now().Sub(start))

}
