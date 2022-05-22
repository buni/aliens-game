package city

import (
	"errors"
	"strings"
)

type Direction int

func ParseDirection(direction string) (Direction, error) {
	direction = strings.ToLower(strings.TrimSpace(direction))
	switch direction {
	case "north":
		return North, nil
	case "east":
		return East, nil
	case "south":
		return South, nil
	case "west":
		return West, nil
	default:
		return 0, errors.New("bad direction")
	}
}

func (d Direction) String() string {
	switch d {
	case 1:
		return "north"
	case 2:
		return "east"
	case 3:
		return "south"
	case 4:
		return "west"
	default:
		return ""
	}
}
