package navigator

import (
	"diego.pizza/equalexperts/mars/internal/command"
	"diego.pizza/equalexperts/mars/internal/position"
)

func Navigate(dir position.Position, cmd command.Command) position.Position {
	curFacingDirection := dir.Direction
	curX := dir.X
	curY := dir.Y
	for _, step := range cmd.Sequence {
		switch step {
		case 'F':
			switch curFacingDirection {
			case position.DirectionNorth:
				curY++
			case position.DirectionSouth:
				curY--
			case position.DirectionEast:
				curX++
			case position.DirectionWest:
				curX--
			}
		case 'B':
			switch curFacingDirection {
			case position.DirectionNorth:
				curY--
			case position.DirectionSouth:
				curY++
			case position.DirectionEast:
				curX--
			case position.DirectionWest:
				curX++
			}
		case 'L':
			switch curFacingDirection {
			case position.DirectionNorth:
				curFacingDirection = position.DirectionWest
			case position.DirectionSouth:
				curFacingDirection = position.DirectionEast
			case position.DirectionEast:
				curFacingDirection = position.DirectionNorth
			case position.DirectionWest:
				curFacingDirection = position.DirectionSouth
			}
		case 'R':
			switch curFacingDirection {
			case position.DirectionNorth:
				curFacingDirection = position.DirectionEast
			case position.DirectionSouth:
				curFacingDirection = position.DirectionWest
			case position.DirectionEast:
				curFacingDirection = position.DirectionSouth
			case position.DirectionWest:
				curFacingDirection = position.DirectionNorth
			}
		}
	}
	return position.Position{
		X:         curX,
		Y:         curY,
		Direction: curFacingDirection,
	}
}
