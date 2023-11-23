package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FruitData struct {
	Data string
}

func (fd *FruitData) CountFruits() int {
	trimmedData := strings.TrimSpace(fd.Data)
	if trimmedData == "" {
		return 0
	}
	return len(strings.Split(trimmedData, "\n"))
}

func (fd *FruitData) FilterFruits(minLength int) *FruitData {
	var filteredFruits []string
	fruits := strings.Split(fd.Data, "\n")

	for _, fruit := range fruits {
		trimmedFruit := strings.TrimSpace(fruit)
		if len(trimmedFruit) >= minLength {
			filteredFruits = append(filteredFruits, trimmedFruit)
		}
	}
	fd.Data = strings.Join(filteredFruits, "\n")
	return fd
}

func (fd *FruitData) ReadFromConsole() error {
	fmt.Println("Enter fruit names (one per line). Press Ctrl+D when finished:")
	scanner := bufio.NewScanner(os.Stdin)
	var fruits []string

	for scanner.Scan() {
		line := scanner.Text()
		fruits = append(fruits, strings.TrimSpace(line))
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input: %w", err)
	}

	fd.Data = strings.Join(fruits, "\n")
	return nil
}

func main() {
	var fruitdata FruitData
	if err := fruitdata.ReadFromConsole(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fruitdata.FilterFruits(5)
	fmt.Printf("\n%s\nwhich is a total of: %d\n", fruitdata.Data, fruitdata.CountFruits())
}
