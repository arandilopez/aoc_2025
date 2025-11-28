package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/arandilopez/aoc_2025/internal/aoc"
	"github.com/arandilopez/aoc_2025/internal/utils"
)

// Execute the requested day's solution.
func main() {
	// parse command line arguments to get the day number
	day := flag.Int("day", 1, "Day of the Advent of Code to execute")
	part := flag.Int("part", 1, "Part of the day's challenge to execute (1 or 2)")
	inputFile := flag.String("input", "input.in", "Path to the input file")
	flag.Parse()

	if *day < 1 || *day > 25 {
		fmt.Println("Error: Day must be between 1 and 25")
		return
	}
	if *part != 1 && *part != 2 {
		fmt.Println("Error: Part must be either 1 or 2")
		return
	}

	if *inputFile == "" || len(*inputFile) == 0 {
		fmt.Println("Error: Input file path cannot be empty")
		return
	}

	_, err := os.Stat(*inputFile)
	if os.IsNotExist(err) {
		fmt.Printf("Error: Input file '%s' does not exist\n", *inputFile)
		return
	}

	fmt.Println("Executing Day", *day, "with input file", *inputFile)

	aocInstance := aoc.AdventOfCode{}
	result, err := utils.ExecDay(aocInstance, *day, *part, *inputFile)
	if err != nil {
		fmt.Printf("Error executing Day %d Part %d: %v\n", *day, *part, err)
		return
	}

	fmt.Printf("Result for Day %d Part %d: %d\n", *day, *part, result)
}
