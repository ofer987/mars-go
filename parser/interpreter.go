package parser

import (
	"github.com/ofer987/mars-go/exploration"
	"regexp"
)

type Interpreter struct {
	BoundaryX, BoundaryY int
	Rovers               map[string]exploration.Rover
	Movements            map[string]string
}

func NewInterpreter() *Interpreter {
	initRovers := make(map[string]exploration.Rover)
	initMovements := make(map[string]string)

	return &Interpreter{
		Rovers:    initRovers,
		Movements: initMovements,
	}
}

func (i *Interpreter) Interpret(input string) {
	plateauRegex := regexp.MustCompile(`(?i)^Plateau:(\d+)\s+(\d+)$`)
	landingRegex := regexp.MustCompile(`(?i)^(\w+)\s+Landing\s*:\s*(\d+)\s+(\d+)\s+(\w)$`)
	instructionsRegex := regexp.MustCompile(`(?i)^(\w+)\s+Instructions\s*:\s*([LMR]+)$`)

	if plateauRegex.MatchString(input) {
		i.setPlateau(input)
	} else if landingRegex.MatchString(input) {
		i.setRover(input)
	} else if instructionsRegex.MatchString(input) {
		i.setMovements(input)
	}
}

func (i *Interpreter) setPlateau(input string) {
	plateau, err := ParsePlateau(input)

	if err == nil {
		i.BoundaryX = plateau.X
		i.BoundaryY = plateau.Y
	}
}

func (i *Interpreter) setRover(input string) {
	rover, err := ParseLanding(input)

	if err == nil {
		i.Rovers[rover.Name] = *rover
	}
}

func (i *Interpreter) setMovements(input string) {
	movements, err := ParseMovements(input)

	if err == nil {
		i.Movements[movements.Name] += movements.Moves
	}
}
