package exploration

import ()

type Rover struct {
	Name  string
	X     int
	Y     int
	sense Sense
}

func NewRover(name string, initX int, initY int, initSense rune) *Rover {
	return &Rover{
		Name:  name,
		X:     initX,
		Y:     initY,
		sense: NewSense(initSense),
	}
}

func (ro *Rover) TurnCounterClockwise() {
	ro.sense = (ro.sense - 1) % 4

	if ro.sense < 0 {
		ro.sense = 3
	}
}

func (ro *Rover) TurnClockwise() {
	ro.sense = (ro.sense + 1) % 4
}

func (ro *Rover) MoveForward() {
	switch ro.sense % 4 {
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

// func (ro *Rover) String() string {
// 	strings.Join()
// 	ro.X.String() + ", " + ro.Y.String() + ", " + ro.sense.String()
// }
