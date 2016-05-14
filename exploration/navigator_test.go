package exploration

import (
	"reflect"
	"testing"
)

func Testmove(t *testing.T) {
	cases := []struct {
		startRover    Rover
		movement      rune
		expectedRover Rover
	}{
		{
			*NewRover("", 0, 0, 'N'),
			'M',
			*NewRover("", 0, 1, 'N'),
		},
		{
			*NewRover("", 0, 0, 'N'),
			'R',
			*NewRover("", 0, 0, 'E'),
		},
		{
			*NewRover("", 1, 2, 'N'),
			'L',
			*NewRover("", 1, 2, 'W'),
		},
		{
			*NewRover("", 1, 2, 'N'),
			'l',
			*NewRover("", 1, 2, 'W'),
		},
		{
			*NewRover("", 1, 2, 'W'),
			'm',
			*NewRover("", 1, 1, 'W'),
		},
	}

	for _, c := range cases {
		navigator := NewNavigator("", 999, 999, &c.startRover)
		navigator.move(c.movement)

		if reflect.DeepEqual(c.expectedRover, navigator.rover) {
			t.Errorf("The rover should have ended at (x, y, sense) = (%v) rather than at (%v)", c.expectedRover, navigator.rover)
		}
	}
}

func TestBoundaries(t *testing.T) {
	cases := []struct {
		boundaryX, boundaryY int
		X, Y                 int
		isWithinBounds       bool
	}{
		{
			1, 2,
			1, 2,
			true,
		},
		{
			1, 2,
			1, 3,
			false,
		},
	}

	for _, c := range cases {
		rover := NewRover("Foo", c.X, c.Y, 0)
		navigator := NewNavigator("", c.boundaryX, c.boundaryY, rover)

		if navigator.validateBoundary() != c.isWithinBounds {
			t.Errorf("Navigator.validateBoundary() failed")
		}
	}
}

func TestNavigate(t *testing.T) {
	cases := []struct {
		boundaryX, boundaryY int
		movements            string
		startRover           Rover
		expectedRover        Rover
	}{
		{
			1, 2,
			"LMRLMLM",
			*NewRover("", 1, 2, 'N'),
			*NewRover("", 0, 1, 'S'),
		},
		{
			2, 2,
			"LMRMRMM",
			*NewRover("", 1, 1, 'N'),
			*NewRover("", 2, 2, 'E'),
		},
		{
			2, 2,
			"lmrmrmm",
			*NewRover("", 1, 1, 'N'),
			*NewRover("", 2, 2, 'E'),
		},
	}

	for index, c := range cases {
		navigator := NewNavigator(c.movements, c.boundaryX, c.boundaryY, &c.startRover)

		navigator.Navigate()

		if !reflect.DeepEqual(*navigator.rover, c.expectedRover) {
			t.Errorf("Test case %v, The rover should have ended at (x, y, sense) = (%v) rather than at (%v)", index, c.expectedRover, *navigator.rover)
		}
	}
}
