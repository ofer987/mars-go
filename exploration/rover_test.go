package exploration

import (
	"reflect"
	"testing"
)

func TestTurnCounterClockwise(t *testing.T) {
	rover := NewRover("Foo", 0, 0, 'N')

	expectedSenses := [...]Sense{Sense(3), Sense(2), Sense(1), (0), Sense(3)}
	for _, expectedSense := range expectedSenses {
		rover.TurnCounterClockwise()

		if rover.Sense != expectedSense {
			t.Errorf("The sense of the rover is %v instead of %v.", rover.Sense, expectedSense)
		}
	}
}

func TestTurnClockwise(t *testing.T) {
	rover := NewRover("Foo", 0, 0, 'S')

	expectedSenses := [...]Sense{
		Sense(3),
		Sense(0),
		Sense(1),
		Sense(2),
		Sense(3),
		Sense(0),
	}
	for _, expectedSense := range expectedSenses {
		rover.TurnClockwise()

		if expectedSense != rover.Sense {
			t.Errorf("The sense of the rover is %v instead of %v.", rover.Sense, expectedSense)
		}
	}
}

func TestMoveForward(t *testing.T) {
	rover := NewRover("Foo", 0, 0, 'N')

	expectedY := [...]int{
		1,
		2,
		3,
		4,
		5,
		6,
	}
	for _, y := range expectedY {
		rover.MoveForward()

		if y != rover.Y {
			t.Errorf("The Y coordinate of the rover is %v instead of %v.", rover.Y, y)
		}
	}
}

func TestString(t *testing.T) {
	rover := NewRover("Foo", 10, 50, 'E')
	expected := "10 50 E"

	if !reflect.DeepEqual(expected, rover.String()) {
		t.Errorf("The string representation of rover: %+v should be %s", rover, expected)
	}
}
