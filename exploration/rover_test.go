package exploration

import (
	"testing"
)

func TestTurnCounterClockwise(t *testing.T) {
	rover := NewRover("Foo", 0, 0, 0)

	expectedSenses := [...]Sense{Sense(3), Sense(2), Sense(1), (0), Sense(3)}
	for _, expectedSense := range expectedSenses {
		rover.TurnCounterClockwise()

		if rover.sense != expectedSense {
			t.Errorf("The sense of the rover is %v instead of %v.", rover.sense, expectedSense)
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

		if expectedSense != rover.sense {
			t.Errorf("The sense of the rover is %v instead of %v.", rover.sense, expectedSense)
		}
	}
}
