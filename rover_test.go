package rover

import (
	"testing"
)

func TestTurnCounterClockwise(t *testing.T) {
	rover := New("Foo", 0, 0, 0)

	expectedSenses := [...]int{3, 2, 1, 0, 3}
	for _, expectedSense := range expectedSenses {
		rover.TurnCounterClockwise()

		if rover.Sense != expectedSense {
			t.Errorf("The sense of the rover is %v instead of %v.", rover.Sense, expectedSense)
		}
	}
}

func TestTurnClockwise(t *testing.T) {
	rover := New("Foo", 0, 0, 2)

	expectedSenses := [...]int{3, 0, 1, 2, 3, 0}
	for _, expectedSense := range expectedSenses {
		rover.TurnClockwise()

		if rover.Sense != expectedSense {
			t.Errorf("The sense of the rover is %v instead of %v.", rover.Sense, expectedSense)
		}
	}
}
