package main

import (
	"../image"
	"fmt"
	"time"
)

var _ = fmt.Println

func main() {
	path := "./test_images/"
	folderNumber := "1"
	myImage := image.ReadImageFromFile(path, folderNumber)
	graph := image.ImageToGraph(myImage)
	start := time.Now()
	mst := graph.MinimalSpanningTree()
	fmt.Println(time.Now().Sub(start))
	fmt.Println(len(mst))
}
