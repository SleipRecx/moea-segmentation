package image

import (
	"../graph"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"math"
	"os"
)

type Coordinate struct {
	X, Y int
}

type Pixel struct {
	R, G, B, A uint8
}

func ColorDistance(p1 Pixel, p2 Pixel) float64 {
	r := math.Pow(float64(p1.R-p2.R), 2)
	g := math.Pow(float64(p1.G-p2.G), 2)
	b := math.Pow(float64(p1.B-p2.B), 2)
	a := math.Pow(float64(p1.A-p2.A), 2)

	return math.Sqrt(r + g + b + a)
}

func (pixel Pixel) Distance(pixel2 Pixel) float64 {
	return ColorDistance(pixel, pixel2)
}

type Image struct {
	Pixels [][]Pixel
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

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
}

func SaveImageToFile(myImage Image) {
	file, err := os.Create("image.jpg")
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

	jpeg.Encode(file, img, &jpeg.Options{Quality: 100})
}

func ImageToGraph(myImage Image) graph.Graph {
	var edges []graph.Edge
	var vertices []graph.Vertex
	pixels := myImage.Pixels
	for i := range pixels {
		for j := range pixels[i] {
			fromCord := Coordinate{X: i, Y: j}
			vertices = append(vertices, fromCord)
			if j + 1 < len(pixels[i]) {
				toCord := Coordinate{X: i, Y: j + 1}
				edge := graph.Edge{U: fromCord, V: toCord, Weight: ColorDistance(pixels[i][j], pixels[i][j+1])}
				edges = append(edges, edge)
			}
			if i + 1 < len(pixels) {
				toCord := Coordinate{X: i + 1, Y: j}
				edge := graph.Edge{U: fromCord, V: toCord, Weight: ColorDistance(pixels[i][j], pixels[i+1][j])}
				edges = append(edges, edge)
			}
		}
	}
	return graph.Graph{Edges: edges, Vertices: vertices}
}


func ReconstructImage(segments [][]interface{}, myImage Image) Image {
	for _,segment := range segments {
		r := 1.0
		g := 1.0
		b := 1.0
		for _, c := range segment {
			cord := c.(Coordinate)
			r += float64(myImage.Pixels[cord.X][cord.Y].R)
			g += float64(myImage.Pixels[cord.X][cord.Y].G)
			b += float64(myImage.Pixels[cord.X][cord.Y].B)
		}
		r = r / float64(len(segment))
		g = g / float64(len(segment))
		b = b / float64(len(segment))
		for _, c := range segment {
			cord := c.(Coordinate)
			myImage.Pixels[cord.X][cord.Y].R = uint8(r)
			myImage.Pixels[cord.X][cord.Y].G = uint8(g)
			myImage.Pixels[cord.X][cord.Y].B = uint8(b)
		}
	}
	return myImage
}