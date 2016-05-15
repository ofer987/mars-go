package parsers

import (
	"errors"
	"regexp"
	"strconv"
)

type Plateau struct {
	X, Y int
}

func plateauRegexp() *regexp.Regexp {
	return regexp.MustCompile(`(?i)^Plateau:(\d+)\s+(\d+)$`)
}

func ParsePlateau(input string) (*Plateau, error) {
	results := plateauRegexp().FindAllStringSubmatch(input, -1)

	if results == nil || len(results) < 1 {
		return nil, errors.New("invalid input for Plateau")
	}

	x, _ := strconv.Atoi(results[0][1])
	y, _ := strconv.Atoi(results[0][2])

	return &Plateau{X: x, Y: y}, nil
}
