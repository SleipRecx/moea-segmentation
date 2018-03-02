package main

import (
	"../image"
	"fmt"
)

var _ = fmt.Println

func main() {
	path := "./test_images/"
	folderNumber := "git1"
	myImage := image.ReadImageFromFile(path, folderNumber)
	graph := image.ImageToGraph(myImage)
	mst := graph.MinimalSpanningTree()
	fmt.Println(mst.CalculateTotalCost())
}
