package main

import "fmt"

var _ = fmt.Println

func main() {
	path := "./test_images/"
	folderNumber := "1"
	myImage := readImageFromFile(path, folderNumber)
	graph := imageToGraph(myImage)
	graph.minimalSpanningTree()


}
