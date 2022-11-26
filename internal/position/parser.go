package position

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrPositionInvalid       = errors.New("the position supplied is invalid: valid format is x,y,direction")
	ErrInvalidDirectionValue = errors.New("direction is invalid: valid format is one of the following: NORTH, EAST, SOUTH, WEST")
	ErrPositionCannotBeEmpty = errors.New("position cannot be empty")
)

func ParsePosition(pos string) (Position, error) {
	var position Position
	if len(pos) < 1 {
		return position, ErrPositionCannotBeEmpty
	}
	parts := strings.Split(pos, ",")
	if len(parts) != 3 {
		return position, ErrPositionInvalid
	}
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return position, fmt.Errorf("position coordinate X must be a valid integer: %v", err)
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return position, fmt.Errorf("position coordinate Y must be a valid integer: %v", err)
	}
	dir, err := ParseDirection(parts[2])
	if err != nil {
		return position, err
	}
	position.X = x
	position.Y = y
	position.Direction = dir

	return position, nil
}

func ParseDirection(dir string) (Direction, error) {
	switch dir {
	case string(DirectionNorth):
		return DirectionNorth, nil
	case string(DirectionEast):
		return DirectionEast, nil
	case string(DirectionSouth):
		return DirectionSouth, nil
	case string(DirectionWest):
		return DirectionWest, nil
	}

	return "", ErrInvalidDirectionValue
}
