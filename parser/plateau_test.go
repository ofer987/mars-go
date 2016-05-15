package parser

import (
	"testing"
)

func TestParsePlateau(t *testing.T) {
	validCases := []struct {
		input    string
		expected Plateau
	}{
		{
			"Plateau:5, 5",
			Plateau{X: 5, Y: 5},
		},
		{
			"Plateau:0, 10",
			Plateau{X: 0, Y: 10},
		},
		{
			"plateau:0, 10",
			Plateau{X: 0, Y: 10},
		},
	}

	invalidCases := []struct {
		input string
	}{
		{"Plateau:-10, 10"},
		{"Plateau:100, -10"},
		{"Plateau:-10, -10"},
	}

	for index, c := range validCases {
		actual, err := ParsePlateau(c.input)

		if err != nil {
			t.Fatalf("Test case %v: An error has occurred: %v", index, err)
		}
		if c.expected != *actual {
			t.Errorf("Test case %v: plateau should have been %+v instead of %+v", index, c.expected, *actual)
		}
	}

	for index, c := range invalidCases {
		_, err := ParsePlateau(c.input)
		if err == nil {
			t.Errorf("Test case %v: The input (%v) should be invalid", index, c.input)
		}
	}
}
