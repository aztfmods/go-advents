// Reference: https://adventofcode.com/2021/day/2

// Description:
// Calculate a submarine's final position from movement commands.
// Commands include "forward" (increases horizontal position), "down" (increases depth),
// and "up" (decreases depth). Each command is followed by an integer value.
// The program determines the submarine's final horizontal position and depth,
// and computes their product.

package main

import (
	"fmt"
	"strings"
)

type Submarine struct {
	HorizontalPosition int
	Depth              int
}

func (s *Submarine) Forward(x int) {
	s.HorizontalPosition += x
}

func (s *Submarine) Down(x int) {
	s.Depth += x
}

func (s *Submarine) Up(x int) {
	s.Depth -= x
}

func main() {
	sub := Submarine{0, 0}
	commands := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	for _, cmd := range commands {
		parts := strings.Fields(cmd)
		var value int
		fmt.Sscanf(parts[1], "%d", &value)

		switch parts[0] {
		case "forward":
			sub.Forward(value)
		case "down":
			sub.Down(value)
		case "up":
			sub.Up(value)
		}
	}

	fmt.Printf("horizontal position is: %d, Depth: %d\n", sub.HorizontalPosition, sub.Depth)
	fmt.Printf("multiplication of position and depth is: %d\n", sub.HorizontalPosition*sub.Depth)
}
