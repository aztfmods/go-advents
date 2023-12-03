// Reference: https://adventofcode.com/2023/day/1

// Description:
// Your task involves using a calibration document, now obscured by a young Elf's art, to recover calibration values.
// Each line in the document contains a calibration value formed by combining the first and last digits into a two-digit number.
// Sum up these calibration values from the entire document to complete the puzzle.

package main

import (
	"fmt"
	"strconv"
)

type CalibrationDocument struct {
	lines []string
}

func (cd *CalibrationDocument) ReadDocument(lines []string) {
	cd.lines = lines
}

func (cd *CalibrationDocument) ExtractCalibrationValue() int {
	var totalCalibrationValue int

	for _, line := range cd.lines {
		firstPart, lastPart := "", ""
		for _, char := range line {
			if char >= '0' && char <= '9' {
				if firstPart == "" {
					firstPart = string(char) // Assign the first digit
				}
				lastPart = string(char) // Continuously assign last digit until end of line
			}
		}
		if firstPart != "" && lastPart != "" {
			combinedValue := firstPart + lastPart
			calibrationValue, _ := strconv.Atoi(combinedValue)
			totalCalibrationValue += calibrationValue

		}
	}
	return totalCalibrationValue
}

func main() {
	calibrationDocument := CalibrationDocument{}
	calibrationDocument.ReadDocument([]string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"})

	fmt.Printf("Calibration value is: %d\n", calibrationDocument.ExtractCalibrationValue())
}
