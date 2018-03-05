package main

import (
	"fmt"
	"time"

	"../image"
)

var _ = fmt.Println

func main() {
	start := time.Now()

	path := "./test_images/"
	folderNumber := "1"
	myImage := image.ReadImageFromFile(path, folderNumber)
	graph := image.ImageToGraph(myImage)
	fmt.Println(len(graph.Vertices))
	mst := graph.MinimalSpanningTree()
	segments := mst.RandomDisjointPartition(10)
	segmentedImage := image.ReconstructImage(segments, myImage)
	image.SaveImageToFile(segmentedImage)
	end := time.Now()
	fmt.Println(end.Sub(start))
}
