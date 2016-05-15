package parser

import (
	"errors"
	"regexp"
	"strconv"
)

type Plateau struct {
	X, Y int
}

func ParsePlateau(input string) (*Plateau, error) {
	regex := regexp.MustCompile(`(?i)^Plateau:(\d+),\s+(\d+)$`)

	results := regex.FindAllStringSubmatch(input, -1)

	if len(results) == 0 {
		return nil, errors.New("invalid input for Plateau")
	}

	x, _ := strconv.Atoi(results[0][1])
	y, _ := strconv.Atoi(results[0][2])

	return &Plateau{X: x, Y: y}, nil
}
