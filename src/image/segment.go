package image

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


