package parsers

import (
	"errors"
	"regexp"
)

type Movements struct {
	Name, Moves string
}

func movementsRegexp() *regexp.Regexp {
	return regexp.MustCompile(`(?i)^(\w+)\s+Instructions\s*:\s*([LMR]+)$`)
}

func ParseMovements(input string) (*Movements, error) {
	results := movementsRegexp().FindAllStringSubmatch(input, -1)

	if results == nil || len(results) < 1 {
		return nil, errors.New("Invalid Input for Rover Instructions")
	}

	movements := Movements{
		Name:  results[0][1],
		Moves: string(results[0][2]),
	}

	return &movements, nil
}
