package main

import (
	"../img"
	//"../tester"
	"fmt"
	"time"

	"io/ioutil"
	"strings"
	"../tester"
)

func main() {
	start := time.Now()
	img.Path = "./test_images/"
	img.FolderNumber = "dessert"
	img.MyImage = img.ReadImageFromFile(img.Path, img.FolderNumber)
	img.ImageWidth, img.ImageHeight = len(img.MyImage.Pixels), len(img.MyImage.Pixels[0])
	img.MyImageGraph = img.MyImage.ConvertToGraph()

	file, _ := ioutil.ReadFile("src/main/path.txt")
	split := strings.Split(string(file), "\n")
	scriptPath := split[0]
	testPath := split[1] + img.FolderNumber
	outPath := split[2]

	fmt.Println(scriptPath)
	fmt.Println(testPath)
	fmt.Println(outPath)


	tester.CalculatePRI(scriptPath, testPath, outPath)

	fmt.Println("Total runtime:", time.Now().Sub(start))

}
