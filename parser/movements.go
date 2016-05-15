package parser

import (
	"errors"
	"regexp"
)

type Movements struct {
	Name, Moves string
}

func ParseMovements(input string) (*Movements, error) {
	regex := regexp.MustCompile(`(?i)^(\w+)\s+Instructions\s*:\s*([LMR]+)$`)

	results := regex.FindAllStringSubmatch(input, -1)

	if len(results) == 0 {
		return nil, errors.New("Invalid Input for Rover Instructions")
	}

	movements := Movements{
		Name:  results[0][1],
		Moves: string(results[0][2]),
	}

	return &movements, nil
}
