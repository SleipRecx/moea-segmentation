package genetic

import (
	"../image"
	"../graph"
)

type Population struct {
	Individuals    []Chromosome
	PopulationSize int
}

func NewPopulation(size int, myImage image.Image) Population {
	individuals := make([]Chromosome, 0)
	imageGraph := myImage.ConvertToGraph()


	results := make(chan [][]graph.Vertex, size)

	for w := 1; w <= 10; w++ {
		go worker(imageGraph,results)
	}

	for a := 1; a <= 10; a++ {
		individuals = append(individuals , NewChromosome(<-results, myImage))
	}

	return Population{Individuals:individuals, PopulationSize:size}
}

func worker(imageGraph graph.Graph, results chan<- [][]graph.Vertex) {
	//k := rand.Intn(10000)
	seg := imageGraph.GraphSegmentation(2000)
	println(len(seg))
	results <- seg
}

/*
func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// Read image
	image := readTestImage(imageNumber)
	imagePixels = rgbaImageToPixelArray(image)
	imageWidth, imageHeight = len(imagePixels), len(imagePixels[0])

	//pop := Population{}
	//pop.InitialPopulation(nPopulation)
	//fmt.Println(len(pop.Individuals))
	//pop.Evolve()
	//fmt.Println(len(pop.Individuals))
	//for pop.Generation < nGenerations {
	//	pop.Evolve()
	//}
	//writeImage("img.png", pixelArrayToRgbaImage(SegmentToColorImage(seg)))
	// Create and initialize image graph
	start := time.Now()
	imageGraph := ImageGraph{}
	imageGraph.Init()

	// Create 50 workers to produce random solutions
	results := make(chan [][]Node, 100)
	for w := 1; w <= 50; w++ {
		go worker(imageGraph, w, results)
	}

	// Collect the results of the workers
	solutions := make([][][]Node, 0)
	for a := 1; a <= 50; a++ {
		solutions = append(solutions, <-results)
	}
	fmt.Println("Time:", time.Now().Sub(start))
	 //Write all segmentations to file
	for _, segmentation := range solutions {
		writeImage("outputs/img_"+strconv.Itoa(len(segmentation))+".png",
			pixelArrayToRgbaImage(SegmentToColorImage(segmentation)))
	}

}
*/



