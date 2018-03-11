package img


type Coordinate struct {
	X, Y int
}

func CordFromDirection(cord Coordinate, direction Direction) Coordinate {
	newCord := Coordinate{X: cord.X, Y: cord.Y}
	switch direction {
	case Up:
		newCord.Y -= 1
	case Down:
		newCord.Y += 1
	case Right:
		newCord.X += 1
	case Left:
		newCord.X -= 1
	}
	if newCord.X >= ImageWidth|| newCord.Y >= ImageHeight{
		return cord
	}
	if newCord.X < 0 || newCord.Y < 0 {
		return cord
	}
	return newCord
}

func WhichDirection(c1, c2 Coordinate) Direction {
	dx, dy := c2.X-c1.X, c2.Y-c1.Y
	if dx >= 1 {
		return Right
	}
	if dx <= -1 {
		return Left
	}
	if dy >= 1 {
		return Down
	}
	if dy <= -1 {
		return Up
	}
	return None
}

