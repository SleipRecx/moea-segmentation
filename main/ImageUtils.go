package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"os"
)

type Pixel struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

type Image struct {
	pixels [][]Pixel
}

func readImageFromFile(path string, folderNumber string) Image {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

	file, err := os.Open(path + folderNumber + "/Test image.jpg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()
	myImage, err := parseImageFile(file)

	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}
	return myImage
}

func parseImageFile(file io.Reader) (Image, error) {
	var myImage Image

	img, _, err := image.Decode(file)

	if err != nil {
		return myImage, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		myImage.pixels = append(myImage.pixels, row)
	}
	return myImage, nil
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
}

func saveImageToFile(myImage Image) {
	file, err := os.Create("image.jpg")
	defer file.Close()

	if err != nil {
		fmt.Print("Failed to create file, with error message", err)
		os.Exit(1)
	}

	height, width := len(myImage.pixels), len(myImage.pixels[0])
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			R := myImage.pixels[y][x].R
			G := myImage.pixels[y][x].G
			B := myImage.pixels[y][x].B
			A := myImage.pixels[y][x].A
			img.Set(x, y, color.RGBA{R, G, B, A})
		}
	}

	jpeg.Encode(file, img, nil)
}

func imageToGraph(myImage Image) Graph {
	var edges []Edge
	var verticies []interface{}
	pixels := myImage.pixels
	for i := range pixels {
		for j := range pixels[i] {
			verticies = append(verticies, Node{X: i, Y: j})
			from := pixels[i][j]
			if j+1 < len(pixels[i]) {
				to := pixels[i][j+1]
				edge := Edge{U: Node{X: i, Y: j}, V: Node{X: i, Y: j + 1}, Weight: from.distance(to)}
				edges = append(edges, edge)
			}
			if i+1 < len(pixels) {
				to := pixels[i+1][j]
				edge := Edge{U: Node{X: i, Y: j}, V: Node{X: i + 1, Y: j}, Weight: from.distance(to)}
				edges = append(edges, edge)
			}
		}
	}
	return Graph{Edges: edges, Vertices: verticies}
}
