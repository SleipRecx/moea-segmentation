package main

import "fmt"

func main() {
	path := "./test_images/"
	folderNumber := "1"

	pixels := getPixelsFromImageFile(path, folderNumber)
	for i := range pixels {
		for j := range pixels[i] {
			fmt.Println(pixels[i][j])
		}
	}

	// createNewImageFromPixels(pixels)
}
