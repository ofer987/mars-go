package parser

import (
	"github.com/ofer987/mars-go/exploration"
	"regexp"
)

type Router struct {
	BoundaryX, BoundaryY int
	Rovers               map[string]exploration.Rover
	Instructions         map[string]string
}

func NewRouter() *Router {
	initRovers := make(map[string]exploration.Rover)
	initInstructions := make(map[string]string)

	return &Router{
		Rovers:       initRovers,
		Instructions: initInstructions,
	}
}

func (ro *Router) SetValues(inputs []string) {
	plateauRegex := regexp.MustCompile(`(?i)^(\w+)\s+Instructions\s*:\s*([LMR]+)$`)
	landingRegex := regexp.MustCompile(`(?i)^(\w+)\s+Landing\s*:\s*(\d+)\s+(\d+)\s+(\w)$`)
	instructionsRegex := regexp.MustCompile(`(?i)^(\w+)\s+Instructions\s*:\s*([LMR]+)$`)

	for _, input := range inputs {
		if plateauRegex.MatchString(input) {
			ro.setPlateau(input)
		} else if landingRegex.MatchString(input) {
			ro.setRover(input)
		} else if instructionsRegex.MatchString(input) {
			ro.setInstructions(input)
		}
	}
}

func (ro *Router) setPlateau(input string) {
	plateau, err := ParsePlateau(input)

	if err == nil {
		ro.BoundaryX = plateau.X
		ro.BoundaryY = plateau.Y
	}
}

func (ro *Router) setRover(input string) {
	rover, err := ParseLanding(input)

	if err == nil {
		ro.Rovers[rover.Name] = *rover
	}
}

func (ro *Router) setInstructions(input string) {
	instructions, err := ParseInstructions(input)

	if err == nil {
		ro.Instructions[instructions.Name] += instructions.Movements
	}
}
