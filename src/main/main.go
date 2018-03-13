package main

import (
	"fmt"
	"time"
	"io/ioutil"
	"../ga"
	"../tester"
	"../img"
	"../graph"
	"../constants"
	"strings"
)

func main() {


	start := time.Now()
	//rand.Seed(time.Now().UnixNano())
	constants.Path = "./test_images/"
	constants.FolderNumber = "flower"
	constants.UseNSGA2 = false
	constants.DeviationWeight = -2
	constants.EdgeWeight = 0
	constants.NGenerations = 10
	constants.NPopulation = 20
	constants.CrossoverRate = 0.2
	constants.MutationRate = 1.0
	constants.NElites = 1
	img.MyImage = img.ReadImageFromFile(constants.Path, constants.FolderNumber)
	constants.ImageWidth, constants.ImageHeight = len(img.MyImage.Pixels), len(img.MyImage.Pixels[0])
	myImageGraph := graph.Graph{}
	myImageGraph.Init()

	population := ga.Population{}
	population.Init(constants.NPopulation, myImageGraph)

	if constants.UseNSGA2 {
		population.NSGA2Run()
	} else {
		population.WeightedSumRun()
	}
	

	file, _ := ioutil.ReadFile("src/main/path.txt")
	split := strings.Split(string(file), "\n")
	scriptPath := split[0]
	testPath := split[1] + constants.FolderNumber
	outPath := split[2]


	fmt.Println("Calculating PRI")
	tester.CalculatePRI(scriptPath, testPath, outPath)
	fmt.Println("Total runtime:", time.Now().Sub(start))





}
