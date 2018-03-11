package main

import (
	"../img"
	"../tester"
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

	tester.CalculatePRI(ScriptPath, TestPath + img.FolderNumber, OutPath)

	fmt.Println("Total runtime:", time.Now().Sub(start))

}
