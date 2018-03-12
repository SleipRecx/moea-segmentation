package main

import (
	"../img"
	"fmt"
	"time"
	"../ga"
	"strconv"
	"../tester"
)

func main() {
	const TestPath = "/Users/markusandresen/Documents/moea/test_images/"
	const OutPath = "/Users/markusandresen/Documents/moea/output/edge/"
	const ScriptPath = "/Users/markusandresen/Documents/moea/src/tester/run.py"


	start := time.Now()
	//rand.Seed(time.Now().UnixNano())
	img.Path = "./test_images/"
	img.FolderNumber = "7"
	img.MyImage = img.ReadImageFromFile(img.Path, img.FolderNumber)
	img.ImageWidth, img.ImageHeight = len(img.MyImage.Pixels), len(img.MyImage.Pixels[0])
	img.MyImageGraph = img.MyImage.ConvertToGraph()
	pop := ga.NewPopulation(4)
	for i, ind := range pop.Individuals {
		img.SaveEdgeDetectionImage(ind.Phenotype.Segments, ind.Phenotype.SegmentMap, "you" + strconv.Itoa(i))
	}
	

	fmt.Println("Total runtime:", time.Now().Sub(start))

	fmt.Println("Calculating PRI")
	tester.CalculatePRI(ScriptPath, TestPath + img.FolderNumber, OutPath)



}
