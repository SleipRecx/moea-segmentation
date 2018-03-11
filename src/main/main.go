package main

import (
	"../ga"
	"../img"
	"fmt"
	"time"
	"../tester"

)

func main() {
	start := time.Now()
	img.Path = "./test_images/"
	img.FolderNumber = "outback"
	img.MyImage = img.ReadImageFromFile(img.Path, img.FolderNumber)
	img.ImageWidth, img.ImageHeight = len(img.MyImage.Pixels), len(img.MyImage.Pixels[0])
	img.MyImageGraph = img.MyImage.ConvertToGraph()


	phenotype := ga.NewPhenotype()
	img.SaveEdgeDetectionImage(phenotype.Segments, img.MyImage, phenotype.SegmentMap, img.FolderNumber)
	tester.CalculatePRI(ScriptPath, TestPath + img.FolderNumber, OutPath)
	rImage := img.ReconstructImage(phenotype.Segments)
	img.SaveImageToFile(rImage, "fuuck")
	fmt.Println("Total runtime:", time.Now().Sub(start))

}
