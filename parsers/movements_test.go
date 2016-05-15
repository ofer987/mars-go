package parsers

import (
	"reflect"
	"testing"
)

func TestParseMovements(t *testing.T) {
	validCases := []struct {
		input    string
		expected Movements
	}{
		{
			"Rover1 Instructions: LMRLMLM",
			Movements{
				Name:  "Rover1",
				Moves: "LMRLMLM",
			},
		},
		{
			"MrSteel Instructions: LMRMLLLMMrmlM",
			Movements{
				Name:  "MrSteel",
				Moves: "LMRMLLLMMrmlM",
			},
		},
	}

	invalidCases := []struct {
		input string
	}{
		{"MrSteel Instructions:  "},
		{"Rover1 Instructions:MLRJTO"},
		{"Rover1 Instructions:000"},
		{"Rover1 Instructions:DEADBEEF"},
		{"Rover 1 Instructions:LMR"},
	}

	for index, c := range validCases {
		actual, err := ParseMovements(c.input)

		if err != nil {
			t.Fatalf("Test case %v: an error has occurred: %v", index, err)
		}
		if !reflect.DeepEqual(c.expected, *actual) {
			t.Errorf("Test case %v: movements should have been %+v instead of %+v", index, c.expected, *actual)
		}
	}

	for index, c := range invalidCases {
		_, err := ParseMovements(c.input)
		if err == nil {
			t.Errorf("Test case %v: The input (%v) should be invalid", index, c.input)
		}
	}
}
