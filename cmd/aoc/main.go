package main

import (
	"fmt"

	"github.com/arandilopez/aoc_2025/internal/aoc"
	"github.com/arandilopez/aoc_2025/internal/utils"
)

// Execute the requested day's solution.
func main() {
	// parse command line arguments to get the day number
	day, part, inputFile, err := utils.ArgsValidation()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Advent of Code 2025 - Command Line Interface")
	fmt.Println("--------------------------------------------")
	fmt.Println()

	fmt.Println("Executing Day", *day, "with input file", *inputFile)

	aocInstance := aoc.AdventOfCode{}
	result, err := utils.ExecDay(&aocInstance, day, part, inputFile)
	if err != nil {
		fmt.Printf("Error executing Day %d Part %d: %v\n", *day, *part, err)
		return
	}

	fmt.Printf("Result for Day %d Part %d:\n%d\n", *day, *part, result)
}
