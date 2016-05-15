package parser

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
		interpreter := NewInterpreter()
		commands := strings.Split(c.input, "\n")
		for _, command := range commands {
			interpreter.setRover(command)
		}

		actual := interpreter.Rovers["Rover1"]
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
		interpreter := NewInterpreter()
		commands := strings.Split(c.input, "\n")
		for _, command := range commands {
			interpreter.setMovements(command)
		}

		actual := interpreter.Movements["Rover1"]
		if !reflect.DeepEqual(c.expected, actual) {
			t.Errorf("Test case %v: '%+v' should have created %+v instead of %+v", index, c.input, c.expected, actual)
		}
	}
}

func TestInterpret(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			"Rover1 Instructions:MR",
			"MR",
		},
	}

	for index, c := range cases {
		interpreter := NewInterpreter()

		interpreter.Interpret(c.input)

		if interpreter.Movements["Rover1"] != "MR" {
			t.Errorf("Test Case %d: Failed to set the instructions", index)
		}
	}
}
