package main

import (
	"bufio"
	"fmt"
	"github.com/ofer987/mars-go/exploration"
	"github.com/ofer987/mars-go/parser"
	"os"
)

func main() {
	instructions := parser.NewInstructions()
	stdinScanner := bufio.NewScanner(os.Stdin)

	for {
		stdinScanner.Scan()
		input := string(stdinScanner.Bytes())

		if input == "" {
			break
		}

		instructions.Set(input)
	}

	for name, rover := range instructions.Rovers {
		movements := instructions.Movements[name]
		navigator := exploration.NewNavigator(
			movements,
			instructions.BoundaryX,
			instructions.BoundaryY,
			&rover,
		)

		navigator.Navigate()

		fmt.Printf("%v: %s\n", rover.Name, rover.String())
	}
}
