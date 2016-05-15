package parser

import (
	"reflect"
	"testing"
)

func TestParseInstructions(t *testing.T) {
	validCases := []struct {
		input    string
		expected Instructions
	}{
		{
			"Rover1 Instructions: LMRLMLM",
			Instructions{
				Name:      "Rover1",
				Movements: "LMRLMLM",
			},
		},
		{
			"MrSteel Instructions: LMRMLLLMMrmlM",
			Instructions{
				Name:      "MrSteel",
				Movements: "LMRMLLLMMrmlM",
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
		actual, err := ParseInstructions(c.input)

		if err != nil {
			t.Fatalf("Test case %v: an error has occurred: %v", index, err)
		}
		if !reflect.DeepEqual(c.expected, *actual) {
			t.Errorf("Test case %v: instructions should have been %+v instead of %+v", index, c.expected, *actual)
		}
	}

	for index, c := range invalidCases {
		_, err := ParseInstructions(c.input)
		if err == nil {
			t.Errorf("Test case %v: The input (%v) should be invalid", index, c.input)
		}
	}
}
