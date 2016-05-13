package exploration

import (
	"testing"
)

func Testmove(t *testing.T) {
	cases := []struct {
		startX, startY, startSense          int
		movements                           rune
		expectedX, expectedY, expectedSense int
	}{
		{
			0, 0, 0,
			'M',
			0, 1, 0,
		},
		{
			0, 0, 0,
			'R',
			0, 0, 1,
		},
		{
			1, 2, 0,
			'L',
			1, 2, 3,
		},
		{
			1, 2, 0,
			'l',
			1, 2, 3,
		},
		{
			1, 2, 3,
			'm',
			1, 1, 3,
		},
	}

	for _, c := range cases {
		rover := NewRover("Foo", c.startX, c.startY, c.startSense)
		navigator := NewNavigator("", 999, 999, rover)

		navigator.move(c.movements)

		if rover.X != c.expectedX || rover.Y != c.expectedY {
			t.Errorf("The rover should have ended at (x, y, sense) = (%v, %v, %v) rather than at (%v, %v, %v)", c.expectedX, c.expectedY, c.expectedSense, rover.X, rover.Y, rover.Sense)
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
		boundaryX, boundaryY                int
		startX, startY, startSense          int
		movements                           string
		expectedX, expectedY, expectedSense int
	}{
		{
			1, 2,
			1, 2, 0,
			"LMRLMLM",
			0, 1, 2,
		},
		{
			2, 2,
			1, 1, 0,
			"LMRMRMM",
			2, 2, 1,
		},
		{
			2, 2,
			1, 1, 0,
			"lmrmrmm",
			2, 2, 1,
		},
	}

	for _, c := range cases {
		rover := NewRover("Foo", c.startX, c.startY, c.startSense)
		navigator := NewNavigator(c.movements, c.boundaryX, c.boundaryY, rover)

		navigator.Navigate()

		if navigator.rover.X != c.expectedX || navigator.rover.Y != c.expectedY || navigator.rover.Sense != c.expectedSense {
			t.Errorf("The rover should have ended at (x, y, sense) = (%v, %v, %v) rather than at (%v, %v, %v)", c.expectedX, c.expectedY, c.expectedSense, navigator.rover.X, navigator.rover.Y, navigator.rover.Sense)
		}
	}
}
