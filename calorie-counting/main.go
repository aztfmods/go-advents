// Reference: https://adventofcode.com/2022/day/1

// Description:
// Determine which Elf in an expedition group is carrying the most and least calories.
// Each Elf's food calorie count is entered line by line, separated by blank lines for different Elves.
// The program should calculate and display the maximum total calories carried by any Elf.

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type ElfGroup struct {
	groupCalories []int
}

// constructor function
func NewElfGroup(elfGroups []string) ElfGroup {
	var groupCalories []int
	for _, groupData := range elfGroups {
		total := 0
		for _, line := range strings.Split(groupData, "\n") {
			if val, err := strconv.Atoi(strings.TrimSpace(line)); err == nil {
				total += val
			}
		}
		groupCalories = append(groupCalories, total)
	}
	return ElfGroup{groupCalories}
}

func (e ElfGroup) TotalCalories() int {
	total := 0
	for _, cal := range e.groupCalories {
		total += cal
	}
	return total
}

func (e ElfGroup) FindMaxCalories() (int, int) {
	maxCalories := math.MinInt64
	elfIndex := -1

	for i, total := range e.groupCalories {
		if total > maxCalories {
			maxCalories, elfIndex = total, i+1
		}
	}

	if elfIndex == -1 {
		return 0, 0
	}
	return elfIndex, maxCalories
}

func (e ElfGroup) FindMinCalories() (int, int) {
	minCalories := math.MaxInt64
	elfIndex := -1

	for i, total := range e.groupCalories {
		if total != 0 && total < minCalories {
			minCalories, elfIndex = total, i+1
		}
	}

	if elfIndex == -1 {
		return 0, 0
	}
	return elfIndex, minCalories
}

func readElfData(scanner *bufio.Scanner) string {
	var inputBuilder strings.Builder
	elfNumber := 1

	for {
		fmt.Printf("Enter calories for Elf %d (press Enter after each number, blank line to finish):\n", elfNumber)
		var elfData strings.Builder
		for scanner.Scan() {
			line := scanner.Text()
			if line == "done" || line == "" {
				break
			}
			elfData.WriteString(line + "\n")
		}

		if scanner.Err() != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", scanner.Err())
			return ""
		}

		if elfData.Len() > 0 {
			inputBuilder.WriteString(elfData.String() + "\n\n")
			elfNumber++
		} else {
			break
		}
	}

	return inputBuilder.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter calorie data for each Elf. Enter a blank line to move to the next Elf. Type 'done' when finished:")

	elfDataString := readElfData(scanner)
	elfGroups := strings.Split(elfDataString, "\n\n")
	elfGroup := NewElfGroup(elfGroups)

	elfMax, maxCalories := elfGroup.FindMaxCalories()
	fmt.Printf("Elf %d carries the most calories: %d\n", elfMax, maxCalories)

	elfMin, minCalories := elfGroup.FindMinCalories()
	fmt.Printf("Elf %d carries the least calories: %d\n", elfMin, minCalories)
}
