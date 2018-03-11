package main

import (
	"../ga"
	"../img"
	"fmt"
	"time"

)

func main() {
	start := time.Now()
	img.Path = "./test_images/"
	img.FolderNumber = "3"
	img.MyImage = img.ReadImageFromFile(img.Path, img.FolderNumber)
	img.ImageWidth, img.ImageHeight = len(img.MyImage.Pixels), len(img.MyImage.Pixels[0])
	img.MyImageGraph = img.MyImage.ConvertToGraph()

	population := ga.NewPopulation(1)
	myImg := img.ReconstructImage(population.Individuals[0].Segments)
	img.SaveImageToFile(myImg, "Before")
	a := population.Individuals[0].Mutate()
	myImg2 := img.ReconstructImage(a.Segments)
	img.SaveImageToFile(myImg2, "After")


	fmt.Println("Total runtime:", time.Now().Sub(start))

}
