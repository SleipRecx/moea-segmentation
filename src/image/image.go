package image

import (
	"../graph"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

type Image struct {
	Pixels [][]Pixel
}

func (i *Image) ConvertToGraph() graph.Graph {
	var edges []graph.Edge
	var vertices []graph.Vertex
	pixels := i.Pixels
	for i := range pixels {
		for j := range pixels[i] {
			fromCord := Coordinate{X: i, Y: j}
			vertices = append(vertices, fromCord)
			if j+1 < len(pixels[i]) {
				toCord := Coordinate{X: i, Y: j + 1}

				edge := graph.Edge{U: fromCord, V: toCord, Weight: ColorDistance(pixels[i][j], pixels[i][j+1])}
				edges = append(edges, edge)
			}
			if i+1 < len(pixels) {
				toCord := Coordinate{X: i + 1, Y: j}
				edge := graph.Edge{U: fromCord, V: toCord, Weight: ColorDistance(pixels[i][j], pixels[i+1][j])}
				edges = append(edges, edge)
			}
		}
	}
	return graph.Graph{Edges: edges, Vertices: vertices}
}

func ReadImageFromFile(path string, folderNumber string) Image {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

	file, err := os.Open(path + folderNumber + "/Test image.jpg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()
	myImage, err := ParseImageFile(file)

	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}
	return myImage
}

func ParseImageFile(file io.Reader) (Image, error) {
	var myImage Image

	img, _, err := image.Decode(file)
	img = imaging.Blur(img, .4)

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
		myImage.Pixels = append(myImage.Pixels, row)
	}
	return myImage, nil
}

func SaveImageToFile(myImage Image) {
	file, err := os.Create("image.png")
	defer file.Close()

	if err != nil {
		fmt.Print("Failed to create file, with error message", err)
		os.Exit(1)
	}

	height, width := len(myImage.Pixels), len(myImage.Pixels[0])
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			R := myImage.Pixels[y][x].R
			G := myImage.Pixels[y][x].G
			B := myImage.Pixels[y][x].B
			A := myImage.Pixels[y][x].A
			img.Set(x, y, color.RGBA{R, G, B, A})
		}
	}
	encoder := png.Encoder{CompressionLevel: -1}
	encoder.Encode(file, img)
}

func ReconstructImage(segments [][]Coordinate, myImage Image) Image {
	for _, segment := range segments {
		r := 1.0
		g := 1.0
		b := 1.0
		for _, cord := range segment {
			r += float64(myImage.Pixels[cord.X][cord.Y].R)
			g += float64(myImage.Pixels[cord.X][cord.Y].G)
			b += float64(myImage.Pixels[cord.X][cord.Y].B)
		}
		r = r / float64(len(segment))
		g = g / float64(len(segment))
		b = b / float64(len(segment))
		for _, cord := range segment {
			myImage.Pixels[cord.X][cord.Y].R = uint8(r)
			myImage.Pixels[cord.X][cord.Y].G = uint8(g)
			myImage.Pixels[cord.X][cord.Y].B = uint8(b)
		}
	}
	return myImage
}
