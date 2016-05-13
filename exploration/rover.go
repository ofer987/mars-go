package exploration

import ()

type Rover struct {
	Name  string
	X     int
	Y     int
	Sense int
}

func NewRover(name string, initX int, initY int, sense int) *Rover {
	r := Rover{Name: name, X: initX, Y: initY, Sense: sense}

	return &r
}

func (ro *Rover) TurnCounterClockwise() {
	ro.Sense = (ro.Sense - 1) % 4

	if ro.Sense < 0 {
		ro.Sense = 4 + ro.Sense
	}
}

func (ro *Rover) TurnClockwise() {
	ro.Sense = (ro.Sense + 1) % 4
}

func (ro *Rover) MoveForward() {
	switch ro.Sense {
	case 0:
		ro.Y += 1
		break
	case 1:
		ro.X += 1
		break
	case 2:
		ro.Y -= 1
		break
	case 3:
		ro.X -= 1
		break
	}
}
