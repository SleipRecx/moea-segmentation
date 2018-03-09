package main

import (
	"fmt"
	"time"

	ga "../genetic"
	"../image"
)

func main() {
	start := time.Now()
	path := "./test_images/"
	folderNumber := "9"

	myImage := image.ReadImageFromFile(path, folderNumber)



	pop := ga.NewPopulation(10, myImage)
	c := pop.Individuals[0]
	img := image.ReconstructImage(c.Segments, myImage)
	image.SaveImageToFile(img)
	fmt.Println("Total runtime:", time.Now().Sub(start))

}
