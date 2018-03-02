package graph

type Segment struct {
	Vertices []Vertex
}

/*
func updateRandom() {
	rand.Seed(time.Now().UnixNano())
}

func (t Tree) DisjointPartition(n int) []Segment {
	var se []Segment
	forest = append(forest, t)
	for len(forest) < n {
		updateRandom()
		i := rand.Intn(len(forest))
		if len(forest[i].Edges) > 1 {
			t1, t2 := forest[i].RandomSplit()
			if len(t1.Edges) > 0 && len(t2.Edges) > 0 {
				forest = forest[:i+copy(forest[i:], forest[i+1:])]
				forest = append(forest, t1, t2)
			}
		}
	}
	return forest
}

func (t Tree) RandomSplit() (Tree, Tree) {
	updateRandom()
	randomIndex := rand.Intn(len(t.Edges))
	t1 := Tree{Edges: t.Edges[0:randomIndex]}
	t2 := Tree{Edges: t.Edges[randomIndex:]}
	return t1, t2
}
func createImage(segments []graph.Tree, myImage image.Image) image.Image {
	var myMap = make(map[image.Coordinate]image.Pixel)
	for _, tree := range segments {
		var cords [] image.Coordinate
		r := 1.0
		g := 1.0
		b := 1.0
		for _,edge := range tree.Edges {
			cord1 := edge.U.(image.Coordinate)
			cord2 := edge.V.(image.Coordinate)
			cords = append(cords, cord1, cord2)
			r += float64(myImage.Pixels[cord1.X][cord1.Y].R)
			g += float64(myImage.Pixels[cord1.X][cord1.Y].G)
			b += float64(myImage.Pixels[cord1.X][cord1.Y].B)
			r += float64(myImage.Pixels[cord2.X][cord2.Y].R)
			g += float64(myImage.Pixels[cord2.X][cord2.Y].G)
			b += float64(myImage.Pixels[cord2.X][cord2.Y].B)
		}
		r = r / float64(len(tree.Edges)) * 2
		g = g / float64(len(tree.Edges)) * 2
		b = b / float64(len(tree.Edges)) * 2
		p := image.Pixel{R:uint8(r), G:uint8(g), B:uint8(b), A:255}
		for _, v := range cords {
			_, ok := myMap[v]
			fmt.Println(ok)
			myMap[v] = p
		}
	}

	for i := range myImage.Pixels {
		for j := range myImage.Pixels[i] {
			pixel,_ := myMap[image.Coordinate{X: i, Y: j}]
			myImage.Pixels[i][j] = pixel
		}
	}
	return myImage
}

*/
