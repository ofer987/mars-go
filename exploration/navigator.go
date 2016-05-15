package exploration

import (
	"strings"
)

type Navigator struct {
	movements            string
	boundaryX, boundaryY int
	rover                *Rover
}

func NewNavigator(
	movements string,
	boundaryX, boundaryY int,
	rover *Rover,
) *Navigator {
	return &Navigator{
		movements: movements,
		boundaryX: boundaryX,
		boundaryY: boundaryY,
		rover:     rover,
	}
}

func (na *Navigator) Rover() *Rover {
	return na.rover
}

func (na *Navigator) move(movement rune) {
	switch movement {
	case 'L':
		na.rover.TurnCounterClockwise()
	case 'R':
		na.rover.TurnClockwise()
	case 'M':
		na.rover.MoveForward()
	default:
	}
}

func (na *Navigator) validateBoundary() bool {
	return na.rover.X >= 0 &&
		na.rover.X <= na.boundaryX &&
		na.rover.Y >= 0 &&
		na.rover.Y <= na.boundaryY
}

func (na *Navigator) Navigate() {
	for _, movement := range strings.ToUpper(na.movements) {
		// Copy rover coordinates
		originalX := na.rover.X
		originalY := na.rover.Y
		originalSense := na.rover.Sense

		na.move(movement)

		if !na.validateBoundary() {
			na.rover.X = originalX
			na.rover.Y = originalY
			na.rover.Sense = originalSense
		}
	}
}
