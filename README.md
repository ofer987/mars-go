# MaRS

## Instructions

Source code is in [Go-lang](https://golang.org).
1. Install the [Go development environment](https://golang.org/doc/install)
2. Set up your [GOPATH](https://golang.org/doc/code.html#GOPATH)
3. Compile `go build .`
4. Run the program `./mars-go`

*Note*: Not feeling like installing go? No problem! Use the included binary **main.go** (compiled for darwin OS on amd64 architecture)

### Example files

Try the program using one of the example files, e.g.,

`cat examples/commands.txt | ./mars-go`

## Background

A squad of robotic rovers are to be landed by NASA on a plateau on Mars.

This plateau, which is curiously rectangular, must be navigated by the rovers so that their on board cameras can get a complete view of the surrounding terrain to send back to Earth.

A rover's position is represented by a combination of an x and y co-ordinates and a letter representing one of the four cardinal compass points. The plateau is divided up into a grid to simplify navigation. An example position might be 0, 0, N, which means the rover is in the bottom left corner and facing North.

In order to control a rover, NASA sends a simple string of letters. The possible letters are 'L', 'R' and 'M'. 'L' and 'R' makes the rover spin 90 degrees left or right respectively, without moving from its current spot.

'M' means move forward one grid point, and maintain the same heading.

Assume that the square directly North from (x, y) is (x, y+1).

## Input

Configuration Input: The first line of input is the upper-right coordinates of the plateau, the lower-left coordinates are assumed to be 0,0.
Per Rover.Test Input:8 4
Input 1: Landing co-ordinates for the Rover The position is made up of two integers and a letter separated by spaces, corresponding to the x and y co-ordinates and the rover's orientation. Test Input:1 2 N
Input 2: Navigation instructions i.e a string containing ('L', 'R', 'M'). Test Input:LMLMLMLMM

### Example

```
Plateau:5 5
Rover1 Landing:1 2 N
Rover1 Instructions:LMLMLMLMM
Rover2 Landing:3 3 E
Rover2 Instructions:MMRMMRMRRM

Expected Output:
Rover1:1 3 N
Rover2:5 1 E
```
