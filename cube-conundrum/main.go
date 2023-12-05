// Reference: https://adventofcode.com/2023/day/2

// Description:
// Develop a program to solve the 'Cube Conundrum' puzzle, involving analysis of games with colored cubes.
// Each game, identified by an ID, contains sets of red, green, and blue cubes drawn from a bag.
// Determine which games are feasible with a bag containing 12 red, 13 green, and 14 blue cubes.
// Calculate and output the sum of the IDs of all feasible games.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type CubeCount struct {
	Red   int
	Green int
	Blue  int
}

func (c CubeCount) CanDraw(set CubeCount) bool {
	return set.Red <= c.Red && set.Green <= c.Green && set.Blue <= c.Blue
}

func ProcessGame(gameData string, maxCubes CubeCount) bool {
	sets := strings.Split(gameData, ";")

	for _, set := range sets {
		var redCount, greenCount, blueCount int

		colors := strings.Split(set, ",")

		for _, color := range colors {
			color = strings.TrimSpace(color)
			parts := strings.Split(color, " ")

			if len(parts) < 2 {
				continue
			}

			count, err := strconv.Atoi(parts[0])
			if err != nil {
				continue
			}

			switch parts[1] {
			case "red":
				redCount = count
			case "green":
				greenCount = count
			case "blue":
				blueCount = count
			}
		}

		currentSet := CubeCount{Red: redCount, Green: greenCount, Blue: blueCount}

		// Check if the current set can be drawn from maxCubes.
		if !maxCubes.CanDraw(currentSet) {
			return false
		}
	}
	return true
}

func main() {
	maxCubes := CubeCount{Red: 12, Green: 13, Blue: 14}

	gameData := []string{
		"3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 16 blue",
	}

	for i, data := range gameData {
		if ProcessGame(data, maxCubes) {
			fmt.Printf("Game %d is possible\n", i+1)
		} else {
			fmt.Printf("Game %d is impossible\n", i+1)
		}
	}
}
