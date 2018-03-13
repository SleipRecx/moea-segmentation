package img

import (
	"../graph"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"github.com/disintegration/imaging"
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
	img = imaging.Blur(img, .8)

	if err != nil {
		return myImage, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	for x := 0; x < width; x++ {
		var row []Pixel
		for y := 0; y < height; y++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		myImage.Pixels = append(myImage.Pixels, row)
	}
	return myImage, nil
}

func SaveImageToFile(saveImage Image, filename string) {
	file, err := os.Create(filename + ".png")
	defer file.Close()

	if err != nil {
		fmt.Print("Failed to create file, with error message", err)
		os.Exit(1)
	}

	width, height:= len(saveImage.Pixels), len(saveImage.Pixels[0])
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			R := saveImage.Pixels[x][y].R
			G := saveImage.Pixels[x][y].G
			B := saveImage.Pixels[x][y].B
			A := saveImage.Pixels[x][y].A
			img.Set(x, y, color.RGBA{R, G, B, A})
		}
	}
	encoder := png.Encoder{CompressionLevel: -1}
	encoder.Encode(file, img)
}

func ReconstructImage(segments [][]Coordinate) Image {
	for _, segment := range segments {
		r := 1.0
		g := 1.0
		b := 1.0
		for _, cord := range segment {
			r += float64(MyImage.Pixels[cord.X][cord.Y].R)
			g += float64(MyImage.Pixels[cord.X][cord.Y].G)
			b += float64(MyImage.Pixels[cord.X][cord.Y].B)
		}
		r = r / float64(len(segment))
		g = g / float64(len(segment))
		b = b / float64(len(segment))
		for _, cord := range segment {
			MyImage.Pixels[cord.X][cord.Y].R = uint8(r)
			MyImage.Pixels[cord.X][cord.Y].G = uint8(g)
			MyImage.Pixels[cord.X][cord.Y].B = uint8(b)
		}
	}
	return MyImage
}

func SaveEdgeDetectionImage(segments [][]Coordinate, segmentMap map[Coordinate]int, filename string) {
	newImageBlackAndWhite := image.NewRGBA(image.Rect(0, 0, ImageWidth, ImageHeight))
	newImageGreen := image.NewRGBA(image.Rect(0, 0, ImageWidth, ImageHeight))
	for i := range segments {
		for _, cord := range segments[i] {
			x, y := cord.X, cord.Y
			right := Coordinate{x + 1, y}
			down := Coordinate{x, y - 1}

			neighbours := make([]Coordinate, 0)
			neighbours = append(neighbours, right, down)

			for _, neighbour := range neighbours {
				if inImage(neighbour, MyImage) {
					i := MyImage.Pixels[x][y]
					r, g, b, a := i.R, i.G, i.B, i.A
					newImageBlackAndWhite.Set(x, y, color.RGBA{255, 255, 255, 255})
					newImageGreen.Set(x, y, color.RGBA{r, g, b, a})
					if segmentMap[neighbour] != segmentMap[cord] {
						newImageBlackAndWhite.Set(x, y, color.RGBA{0, 0, 0, 255})
						newImageGreen.Set(x, y, color.RGBA{0, 255, 0, 255})
						break
					}
				}
			}
		}
	}
	fBlack, _ := os.OpenFile("output/edge/"+filename+ ".png", os.O_WRONLY|os.O_CREATE, 0600)
	defer fBlack.Close()
	png.Encode(fBlack, newImageBlackAndWhite)

	fGreen, _ := os.OpenFile("output/green/" + filename + ".png", os.O_WRONLY|os.O_CREATE, 0600)
	defer fGreen.Close()
	png.Encode(fGreen, newImageGreen)
}

func inImage(cord Coordinate, myImage Image) bool {
	if cord.X < len(myImage.Pixels) && cord.Y < len(myImage.Pixels[0]) {
		return true
	}
	return false
}
