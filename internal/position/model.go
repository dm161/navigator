package position

import "fmt"

type Direction string

const (
	DirectionNorth Direction = "NORTH"
	DirectionSouth Direction = "SOUTH"
	DirectionEast  Direction = "EAST"
	DirectionWest  Direction = "WEST"
)

type Position struct {
	X         int
	Y         int
	Direction Direction
}

func (p Position) ToString() string {
	return fmt.Sprintf("(%d, %d, %s)", p.X, p.Y, p.Direction)
}
