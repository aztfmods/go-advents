// Reference: https://adventofcode.com/2021/day/1

// Description:
// Analyze a sequence of sea floor depth measurements to determine how often the depth increases.
// Each line of input represents a depth measurement. The goal is to count the number of times
// a depth measurement increases from the previous one.
// The program processes a list of depth measurements and outputs the count of these increases.

package main

import (
	"errors"
	"fmt"
)

type SonarSweep struct {
	Depths []int
}

func NewSonarSweep(data []int) (SonarSweep, error) {
	if len(data) == 0 {
		return SonarSweep{}, errors.New("no depth data provided")
	}

	for _, depth := range data {
		if depth < 0 {
			return SonarSweep{}, errors.New("invalid depth data: negative values are not allowed")
		}
	}

	return SonarSweep{Depths: data}, nil
}

func (s SonarSweep) AnalyzeDepths() (int, int, int) {
	if len(s.Depths) == 0 {
		return 0, 0, 0
	}

	increaseCount, decreaseCount, totalDepth := 0, 0, 0
	var averageDepth int

	for i := 0; i < len(s.Depths); i++ {
		if i > 0 {
			depth := s.Depths[i]
			switch {
			case depth > s.Depths[i-1]:
				fmt.Printf("Depth increased from %d to %d\n", s.Depths[i-1], depth)
				increaseCount++
			case depth < s.Depths[i-1]:
				fmt.Printf("Depth decreased from %d to %d\n", s.Depths[i-1], depth)
				decreaseCount++
			default:
				fmt.Printf("Depth stayed the same at %d\n", depth)
			}
		}
		totalDepth += s.Depths[i]
	}

	if len(s.Depths) > 0 {
		averageDepth = totalDepth / len(s.Depths)
	}

	return increaseCount, decreaseCount, averageDepth
}

func (s SonarSweep) FindLargestFluctuation() (int, string, float64) {
	if len(s.Depths) < 2 {
		return 0, "No fluctuation", 0.0
	}

	largestFluctuation := 0
	fluctuationType := "no fluctuation"
	largestPercentage := 0.0

	for i := 1; i < len(s.Depths); i++ {
		currentFluctuation := s.Depths[i] - s.Depths[i-1]

		absFluctuation := currentFluctuation
		if absFluctuation < 0 {
			absFluctuation = -absFluctuation
		}

		if absFluctuation > largestFluctuation {
			largestFluctuation = absFluctuation

			switch {
			case currentFluctuation > 0:
				fluctuationType = "increase"
			case currentFluctuation < 0:
				fluctuationType = "decrease"
			}
			largestPercentage = (float64(absFluctuation) / float64(s.Depths[i-1])) * 100
		}
	}
	return largestFluctuation, fluctuationType, largestPercentage
}

func main() {
	sonarSweep, err := NewSonarSweep([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	increaseCount, decreaseCount, averageDepth := sonarSweep.AnalyzeDepths()
	largestFluctuation, fluctuationType, largestPercentage := sonarSweep.FindLargestFluctuation()

	fmt.Printf("\nDetected %d increases and %d decreases with an average depth of %d\n",
		increaseCount,
		decreaseCount,
		averageDepth,
	)

	fmt.Printf("Largest fluctuation was a %s of %d or %v in percentage\n",
		fluctuationType,
		largestFluctuation,
		largestPercentage,
	)
}
