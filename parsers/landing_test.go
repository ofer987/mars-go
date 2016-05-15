package parsers

import (
	"github.com/ofer987/mars-go/exploration"
	"testing"
)

func TestParseLanding(t *testing.T) {
	validCases := []struct {
		input    string
		expected exploration.Rover
	}{
		{
			"Rover1 Landing:1 2 N",
			*exploration.NewRover("Rover1", 1, 2, 'N'),
		},
		{
			"MrSteel Landing: 100 200 E",
			*exploration.NewRover("MrSteel", 100, 200, 'E'),
		},
		{
			"MrSteel landing: 100 200 E",
			*exploration.NewRover("MrSteel", 100, 200, 'E'),
		},
	}

	invalidCases := []struct {
		input string
	}{
		{"MrSteel Landing: 100 200 T"},
		{"Rover1 Landing:-100, 0"},
		{"Rover1 Landing:0, -100"},
		{"Rover1 Landing:-10, -100"},
		{"Rover 1 Landing: 10, 100"},
		{"Rover Dover Landing: 10, 100"},
	}

	for index, c := range validCases {
		actual, err := ParseLanding(c.input)

		if err != nil {
			t.Fatalf("Test case %v: an error has occurred %v", index, err)
		}
		if c.expected != *actual {
			t.Errorf("Test case %v: rover should have been %+v instead of %+v", index, c.expected, *actual)
		}
	}

	for index, c := range invalidCases {
		_, err := ParseLanding(c.input)
		if err == nil {
			t.Errorf("Test case %v: The input (%v) should be invalid", index, c.input)
		}
	}
}
