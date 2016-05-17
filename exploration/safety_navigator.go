package exploration

import (
	"strings"
)

type SafetyNavigator struct {
	Navigator
	otherRovers []*Rover
}

func (sn *SafetyNavigator) isRoverCrash() bool {
	for _, otherRover := range sn.otherRovers {
		if otherRover.X == sn.rover.X && otherRover.Y == sn.rover.Y {
			return true
		}
	}

	return false
}

func (sn *SafetyNavigator) Navigate() {
	for _, movement := range strings.ToUpper(sn.movements) {
		// Copy rover coordinates
		originalX := sn.rover.X
		originalY := sn.rover.Y
		originalSense := sn.rover.Sense

		sn.move(movement)

		if !sn.validateBoundary() || sn.isRoverCrash() {
			sn.rover.X = originalX
			sn.rover.Y = originalY
			sn.rover.Sense = originalSense
		}
	}
}
