package parser

import (
	"errors"
	"github.com/ofer987/mars-go/exploration"
	"regexp"
	"strconv"
)

func ParseLanding(input string) (*exploration.Rover, error) {
	regex := regexp.MustCompile(`(?i)^(\w+)\s+Landing\s*:\s*(\d+)\s+(\d+)\s+(\w)$`)

	results := regex.FindAllStringSubmatch(input, -1)

	if len(results) == 0 {
		return nil, errors.New("Invalid Input for Rover Landing")
	}

	name := results[0][1]
	x, _ := strconv.Atoi(results[0][2])
	y, _ := strconv.Atoi(results[0][3])
	sense := string(results[0][4])

	rover := exploration.NewRover(name, x, y, rune(sense[0]))
	if rover == nil {
		return nil, errors.New("Rover could not be initialized")
	}

	return rover, nil
}
