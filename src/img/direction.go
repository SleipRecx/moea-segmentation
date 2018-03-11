package img

type Direction int

func DirectionFactory(n int) Direction{
	switch n {
	case 0:
		return None
	case 1:
		return Up
	case 2:
		return Down
	case 3:
		return Left
	default:
		return Right
	}
}

const (
	None  Direction = iota
	Up    Direction = iota
	Down  Direction = iota
	Left  Direction = iota
	Right Direction = iota
)
