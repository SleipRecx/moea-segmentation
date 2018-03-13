package main

import (
	"../img"
	"fmt"
	"time"

	"io/ioutil"
	"strings"
	"../ga"
	"strconv"
	"../tester"
)

func main() {


	start := time.Now()
	//rand.Seed(time.Now().UnixNano())
	img.Path = "./test_images/"
	img.FolderNumber = "flower"
	img.MyImage = img.ReadImageFromFile(img.Path, img.FolderNumber)
	img.ImageWidth, img.ImageHeight = len(img.MyImage.Pixels), len(img.MyImage.Pixels[0])
	img.MyImageGraph = img.MyImage.ConvertToGraph()
	pop := ga.NewPopulation(4)
	for i, ind := range pop.Individuals {
		img.SaveEdgeDetectionImage(ind.Phenotype.Segments, ind.Phenotype.SegmentMap, "you" + strconv.Itoa(i))
	}
	

	file, _ := ioutil.ReadFile("src/main/path.txt")
	split := strings.Split(string(file), "\n")
	scriptPath := split[0]
	testPath := split[1] + img.FolderNumber
	outPath := split[2]


	fmt.Println("Calculating PRI")
	tester.CalculatePRI(scriptPath, testPath, outPath)
	fmt.Println("Total runtime:", time.Now().Sub(start))





}
