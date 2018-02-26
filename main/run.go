package main

import "fmt"

func main() {
	path := "./test_images/"
	folderNumber := "1"

	pixels := getPixelsFromImageFile(path, folderNumber)
	createNewImageFromPixels(pixels)
	fmt.Print("\n Total deviation: ", calcOverallDeviation(pixels))
}
