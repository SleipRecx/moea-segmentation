package main

func main() {
	path := "./test_images/"
	folderNumber := "1"

	pixels := getPixelsFromImageFile(path, folderNumber)
	createNewImageFromPixels(pixels)
}
