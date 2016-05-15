package exploration

import (
	"strconv"
)

type Rover struct {
	Name  string
	X     int
	Y     int
	Sense Sense
}

func NewRover(name string, initX int, initY int, initSense rune) *Rover {
	sense := NewSense(initSense)
	if sense == nil {
		return nil
	}

	rover := Rover{
		Name:  name,
		X:     initX,
		Y:     initY,
		Sense: *sense,
	}

	return &rover
}

func (ro *Rover) TurnCounterClockwise() {
	ro.Sense = (ro.Sense - 1) % 4

	if ro.Sense < 0 {
		ro.Sense = 3
	}
}

func (ro *Rover) TurnClockwise() {
	ro.Sense = (ro.Sense + 1) % 4
}

func (ro *Rover) MoveForward() {
	switch ro.Sense % 4 {
	case 0:
		ro.Y += 1
	case 1:
		ro.X += 1
	case 2:
		ro.Y -= 1
	case 3:
		ro.X -= 1
	}
}

func (ro *Rover) String() string {
	return strconv.Itoa(ro.X) + " " + strconv.Itoa(ro.Y) + " " + ro.Sense.String()
}
