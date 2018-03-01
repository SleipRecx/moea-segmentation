package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"os"
)

func getPixelsFromImageFile(path string, folderNumber string) [][]Pixel {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

	file, err := os.Open(path + folderNumber + "/Test image.jpg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()
	pixels, err := getPixels(file)

	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}
	return pixels
}

func createNewImageFromPixels(pixels [][]Pixel) {
	file, err := os.Create("image.jpg")
	defer file.Close()

	if err != nil {
		fmt.Print("Failed to create file, with error message", err)
		os.Exit(1)
	}

	height, width := len(pixels), len(pixels[0])
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			R, G, B, A := pixels[y][x].R, pixels[y][x].G, pixels[y][x].B, pixels[y][x].A
			img.Set(x, y, color.RGBA{R, G, B, A})
		}
	}
	jpeg.Encode(file, img, nil)

}

func getPixels(file io.Reader) ([][]Pixel, error) {
	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	var pixels [][]Pixel

	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}
	return pixels, nil
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
}

type Pixel struct {
	R uint8
	G uint8
	B uint8
	A uint8
}
