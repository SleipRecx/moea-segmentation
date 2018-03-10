package main

import (
	"../ga"
	"../img"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	path := "./test_images/"
	folderNumber := "3"
	myImage := img.ReadImageFromFile(path, folderNumber)
	geno := ga.NewGenotype(myImage)
	imgWidth, imgHeight := len(myImage.Pixels), len(myImage.Pixels[0])
	pheno := ga.ConvertToPhenotype(geno, imgWidth, imgHeight, myImage)
	newImg := img.ReconstructImage(pheno.Segments, myImage)
	img.SaveImageToFile(newImg, "fuuck")
	fmt.Println(len(pheno.Segments))
	fmt.Println("Total runtime:", time.Now().Sub(start))

}
