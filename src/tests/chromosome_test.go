package tests

/*
import (
	"testing"

	"../ga"
	"../img"
	"fmt"
	"image"
)

func initTest() ga.Chromosome {
	pixels := make([][]Pixel, 3, 3)
	for i := range pixels {
		fmt.Println(i)
		tmp := []img.Pixel{{R: 255, G: 0, B: 0, A: 255}, {R: 255, G: 0, B: 0, A: 255}, {R: 255, G: 0, B: 0, A: 255}}
		pixels[i] = tmp
	}
	img := image.Image{Pixels: pixels}
	imageGraph := image.ImageToGraph(img)
	mst := imageGraph.MinimalSpanningTree()
	c := NewChromosome(mst, img, 1)
	return c
}

func TestEdgeValue(t *testing.T) {
	c := initTest()
	edgeValue := c.CalcEdgeValue()
	expectedEdgeValue := 0.0
	if edgeValue != expectedEdgeValue {
		t.Error("Expected", expectedEdgeValue, "but got", edgeValue)
	}
}

func TestDeviation(t *testing.T) {
	c := initTest()
	deviation := c.CalcDeviation()
	expectedDeviation := 0.0
	if expectedDeviation != deviation {
		t.Error("Expected", expectedDeviation, "but got", deviation)
	}
}
*/
