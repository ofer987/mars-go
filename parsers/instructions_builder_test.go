package parsers

import (
	"github.com/ofer987/mars-go/exploration"
	"reflect"
	"strings"
	"testing"
)

func TestSetRover(t *testing.T) {
	cases := []struct {
		input    string
		expected exploration.Rover
	}{
		{
			"Rover1 Landing:1 2 N",
			*exploration.NewRover("Rover1", 1, 2, 'N'),
		},
		{
			"Rover1 Landing:10 2 N\nRover1 Landing:1 2 N",
			*exploration.NewRover("Rover1", 1, 2, 'N'),
		},
	}

	for index, c := range cases {
		instructions := NewInstructions()
		commands := strings.Split(c.input, "\n")
		for _, command := range commands {
			instructions.setRover(command)
		}

		actual := instructions.Rovers["Rover1"]
		if !reflect.DeepEqual(c.expected, actual) {
			t.Errorf("Test case %v: '%+v' should have created %+v instead of %+v", index, c.input, c.expected, actual)
		}
	}
}

func TestsetMovements(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			"Rover1 Instructions:LMR",
			"LMR",
		},
		{
			"Rover1 Instructions:LM",
			"LM",
		},
		{
			"Rover1 Instructions:LMR\nRover1 Instructions: MMM",
			"LMRMMM",
		},
		{
			"Rover1 Instructions:MMRMMRMRRM",
			"MMRMMRMRRM",
		},
	}

	for index, c := range cases {
		instructions := NewInstructions()
		commands := strings.Split(c.input, "\n")
		for _, command := range commands {
			instructions.setMovements(command)
		}

		actual := instructions.Movements["Rover1"]
		if !reflect.DeepEqual(c.expected, actual) {
			t.Errorf("Test case %v: '%+v' should have created %+v instead of %+v", index, c.input, c.expected, actual)
		}
	}
}

func TestSet(t *testing.T) {
	cases := []struct {
		input    string
		expected Instructions
	}{
		{
			"Plateau:1 2",
			Instructions{
				1, 2,
				map[string]exploration.Rover{},
				map[string]string{},
			},
		},
		{
			"Rover1 Landing:0 10 E",
			Instructions{
				0, 0,
				map[string]exploration.Rover{
					"Rover1": *exploration.NewRover("Rover1", 0, 10, 'E'),
				},
				map[string]string{},
			},
		},
		{
			"Rover1 Instructions:MR",
			Instructions{
				0, 0,
				map[string]exploration.Rover{},
				map[string]string{"Rover1": "MR"},
			},
		},
		{
			"Invalid instruction that will be ignored",
			Instructions{
				0, 0,
				map[string]exploration.Rover{},
				map[string]string{},
			},
		},
	}

	for index, c := range cases {
		instructions := NewInstructions()

		instructions.Set(c.input)

		if !reflect.DeepEqual(c.expected, *instructions) {
			t.Errorf("Test Case %d: Failed to set the instructions. They should be %+v instead of %+v", index, c.expected, *instructions)
		}
	}
}
