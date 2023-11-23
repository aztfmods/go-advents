package main

import (
	"reflect"
	"testing"
)

func TestNewSonarSweep(t *testing.T) {
	emptyInput := []int{}

	_, err := NewSonarSweep(emptyInput)
	if err == nil {
		t.Errorf("NewSonarSweep with empty input should return an error")
	}

	validInput := []int{100, 200, 300}
	expected := SonarSweep{Depths: validInput}
	result, err := NewSonarSweep(validInput)
	if err != nil {
		t.Errorf("NewSonarSweep with valid input returned an unexpected error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("NewSonarSweep with valid input = %v; want %v", result, expected)
	}
}

func TestAnalyzeDepths(t *testing.T) {
	testCases := []struct {
		name              string
		inputDepths       []int
		expectedIncreases int
		expectedDecreases int
		expectedAverage   int
	}{
		{
			name:              "Standard case",
			inputDepths:       []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			expectedIncreases: 7,
			expectedDecreases: 2,
			expectedAverage:   225,
		},
		{
			name:              "No fluctuation",
			inputDepths:       []int{100, 100, 100, 100},
			expectedIncreases: 0,
			expectedDecreases: 0,
			expectedAverage:   100,
		},
		{
			name:              "Empty input",
			inputDepths:       []int{},
			expectedIncreases: 0,
			expectedDecreases: 0,
			expectedAverage:   0,
		},
		{
			name:              "Invalid Data Negative Depths",
			inputDepths:       []int{-10, -20, -30},
			expectedIncreases: 0,
			expectedDecreases: 0,
			expectedAverage:   0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sonarSweep, _ := NewSonarSweep(tc.inputDepths)
			gotIncreases, gotDecreases, gotAverage := sonarSweep.AnalyzeDepths()

			if gotIncreases != tc.expectedIncreases {
				t.Errorf("AnalyzeDepths() = %v increases, want %v increases", gotIncreases, tc.expectedIncreases)
			}

			if gotDecreases != tc.expectedDecreases {
				t.Errorf("AnalyzeDepths() = %v decreases, want %v decreases", gotDecreases, tc.expectedDecreases)
			}

			if gotAverage != tc.expectedAverage {
				t.Errorf("AnalyzeDepths() = %v average depth, want %v average depth", gotAverage, tc.expectedAverage)
			}
		})
	}
}

func TestFindLargestFluctuation(t *testing.T) {
	testCases := []struct {
		name                string
		inputDepths         []int
		expectedFluctuation int
		expectedType        string
		expectedPercentage  float64
	}{
		{
			name:                "Standard case",
			inputDepths:         []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			expectedFluctuation: 33, // Assuming 207 to 240 is the largest fluctuation
			expectedType:        "increase",
			expectedPercentage:  15.942028985507244, // (33/207)*100
		},
		{
			name:                "Constant values",
			inputDepths:         []int{100, 100, 100},
			expectedFluctuation: 0,
			expectedType:        "no fluctuation",
			expectedPercentage:  0.0,
		},
		{
			name:                "Single decrease",
			inputDepths:         []int{200, 150},
			expectedFluctuation: 50,
			expectedType:        "decrease",
			expectedPercentage:  25, // (50/200)*100
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sonarSweep, _ := NewSonarSweep(tc.inputDepths)
			gotFluctuation, gotType, gotPercentage := sonarSweep.FindLargestFluctuation()

			if gotFluctuation != tc.expectedFluctuation {
				t.Errorf("FindLargestFluctuation() = %v, want %v", gotFluctuation, tc.expectedFluctuation)
			}
			if gotType != tc.expectedType {
				t.Errorf("FindLargestFluctuation() = %v, want %v", gotType, tc.expectedType)
			}
			if gotPercentage != tc.expectedPercentage {
				t.Errorf("FindLargestFluctuation() = %v%%, want %v%%", gotPercentage, tc.expectedPercentage)
			}
		})
	}
}
