package parser

import (
	"errors"
	"regexp"
)

type Instructions struct {
	Name, Movements string
}

func ParseInstructions(input string) (*Instructions, error) {
	regex := regexp.MustCompile(`(?i)^(\w+)\s+Instructions\s*:\s*([LMR]+)$`)

	results := regex.FindAllStringSubmatch(input, -1)

	if len(results) == 0 {
		return nil, errors.New("Invalid Input for Rover Instructions")
	}

	instructions := Instructions{
		Name:      results[0][1],
		Movements: string(results[0][2]),
	}

	return &instructions, nil
}
