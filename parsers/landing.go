package parsers

import (
	"errors"
	"github.com/ofer987/mars-go/exploration"
	"regexp"
	"strconv"
)

func landingRegexp() *regexp.Regexp {
	return regexp.MustCompile(`(?i)^(\w+)\s+Landing\s*:\s*(\d+)\s+(\d+)\s+(\w)$`)
}

func ParseLanding(input string) (*exploration.Rover, error) {
	results := landingRegexp().FindAllStringSubmatch(input, -1)

	if results == nil || len(results) < 1 {
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
