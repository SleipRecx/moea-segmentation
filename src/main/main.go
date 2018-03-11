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
	img.FolderNumber = "1"
	img.MyImage = img.ReadImageFromFile(img.Path, img.FolderNumber)
	img.ImageWidth, img.ImageHeight = len(img.MyImage.Pixels), len(img.MyImage.Pixels[0])
	img.MyImageGraph = img.MyImage.ConvertToGraph()






	population := ga.NewPopulation(1)
	phenotype := population.Individuals[0]


	img.SaveEdgeDetectionImage(phenotype.Segments, img.MyImage, phenotype.SegmentMap, "yo")
	//tester.CalculatePRI("/Users/markusandresen/Documents/moea/src/tester/run.py", "/Users/markusandresen/Documents/moea/test_images/1/", "/Users/markusandresen/Documents/moea/output/")
	fmt.Println("Total runtime:", time.Now().Sub(start))

}
