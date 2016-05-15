package parser

import (
	"github.com/ofer987/mars-go/exploration"
	"reflect"
	"strings"
	"testing"
)

// func TestBar(t *testing.T) {
// 	inputs := []string{
// 		"Plateau:5 5",
// 		"Rover1 Landing:1 2 N",
// 		"Rover1 Instructions:LMLMLMLMM",
// 		"Rover2 Landing:3 3 E",
// 		"Rover2 Instructions:MMRMMRMRRM",
// 	}
//
// 	router := Router{}
//
// }

func TestSetRover(t *testing.T) {
	cases := []struct {
		inputs   string
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
		router := NewRouter()
		array := strings.Split(c.inputs, "\n")
		for _, input := range array {
			router.setRover(input)
		}

		actual := router.Rovers["Rover1"]
		if !reflect.DeepEqual(c.expected, actual) {
			t.Errorf("Test case %v: '%+v' should have created %+v instead of %+v", index, c.inputs, c.expected, actual)
		}
	}
}

func TestSetInstructions(t *testing.T) {
	cases := []struct {
		inputs   string
		expected string
	}{
		{
			"Rover1 Instructions:LMR",
			"LMR",
		},
		{
			"Rover1 Instructions:LMR\nRover1 Instructions: MMM",
			"LMRMMM",
		},
	}

	for index, c := range cases {
		router := NewRouter()
		array := strings.Split(c.inputs, "\n")
		for _, input := range array {
			router.setInstructions(input)
		}

		actual := router.Instructions["Rover1"]
		if !reflect.DeepEqual(c.expected, actual) {
			t.Errorf("Test case %v: '%+v' should have created %+v instead of %+v", index, c.inputs, c.expected, actual)
		}
	}
}
