package main

import (
	"bufio"
	"fmt"
	"github.com/ofer987/mars-go/exploration"
	"github.com/ofer987/mars-go/parser"
	"os"
)

func main() {
	interpreter := parser.NewInterpreter()
	stdinScanner := bufio.NewScanner(os.Stdin)

	for {
		stdinScanner.Scan()
		input := string(stdinScanner.Bytes())

		if input == "" {
			break
		}

		interpreter.Interpret(input)
	}

	for name, rover := range interpreter.Rovers {
		movements := interpreter.Instructions[name]
		navigator := exploration.NewNavigator(
			movements,
			interpreter.BoundaryX,
			interpreter.BoundaryY,
			&rover,
		)

		navigator.Navigate()

		fmt.Printf("%v: %s\n", rover.Name, rover.String())
	}
}
