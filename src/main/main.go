package main

import (
	"../image"
	"fmt"
)

var _ = fmt.Println

func main() {
	path := "./test_images/"
	folderNumber := "1"
	myImage := image.ReadImageFromFile(path, folderNumber)
	graph := image.ImageToGraph(myImage)
	mst := graph.MinimalSpanningTree()
	fmt.Println(mst.CalculateTotalCost())
}
