package exploration

import (
	"reflect"
	"testing"
)

func TestIsRoverCrash(t *testing.T) {
	firstRover := NewRover("Bumped", 5, 5, 'N')
	secondRover := NewRover("Bumper", 5, 6, 'S')

	na := NewNavigator("MM", 10, 10, secondRover)
	safetyNa := SafetyNavigator{
		*na,
		[]*Rover{firstRover},
	}

	safetyNa.move('M')

	if !safetyNa.isRoverCrash() {
		t.Errorf("The two rovers should have crashed into one another\n")
	}
}

func TestIsNotRoverCrash(t *testing.T) {
	firstRover := NewRover("Bumped", 5, 5, 'N')
	secondRover := NewRover("Bumper", 5, 6, 'N')

	na := NewNavigator("MM", 10, 10, secondRover)
	safetyNa := SafetyNavigator{
		*na,
		[]*Rover{firstRover},
	}

	safetyNa.move('M')

	if safetyNa.isRoverCrash() {
		t.Errorf("The two rovers should NOT have crashed into one another\n")
	}
}

func TestSafetyNavigate(t *testing.T) {
	otherRovers := []*Rover{
		NewRover("1", 5, 5, 'N'),
		NewRover("2", 3, 3, 'N'),
		NewRover("3", 5, 0, 'N'),
	}
	boundaryX := 10
	boundaryY := 10
	cases := []struct {
		movements     string
		expectedRover Rover
	}{
		{
			"MMMLM",
			*NewRover("MyRover", 6, 4, 'S'),
		},
		{
			"RMMLMM",
			*NewRover("MyRover", 5, 7, 'W'),
		},
		{
			"LMMMLMMMMM",
			*NewRover("MyRover", 10, 2, 'E'),
		},
	}

	for index, c := range cases {
		myRover := NewRover("MyRover", 7, 5, 'W')
		navigator := NewNavigator(c.movements, boundaryX, boundaryY, myRover)
		safetyNa := SafetyNavigator{*navigator, otherRovers}

		safetyNa.Navigate()

		if !reflect.DeepEqual(*myRover, c.expectedRover) {
			t.Errorf("Test case %v, The rover should have ended at (x, y, sense) = (%v) rather than at (%+v)", index, c.expectedRover, *myRover)
		}
	}
}
