package main

import (
	"fmt"

	"github.com/arandilopez/aoc_2025/internal/aoc"
	"github.com/arandilopez/aoc_2025/internal/utils"
)

// Execute the requested day's solution.
func main() {
	// parse command line arguments to get the day number
	args, err := utils.ArgsValidation()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Advent of Code 2025 - Command Line Interface")
	fmt.Println("--------------------------------------------")
	fmt.Println()

	fmt.Println("Executing Day", args.Day, "Part", args.Part, "with input file", args.InputFile)

	aocInstance := aoc.AdventOfCode{}
	result, err := utils.ExecSolution(aocInstance, args)
	if err != nil {
		fmt.Printf("Error executing Day %d Part %d: %v\n", args.Day, args.Part, err)
		return
	}

	fmt.Printf("Result for Day %d Part %d:\n", args.Day, args.Part)
	fmt.Println(result)
	err = utils.CopyToClipboard(result)
	if err != nil {
		fmt.Println("Warning: Failed to copy result to clipboard:", err)
	} else {
		fmt.Println("Result copied to clipboard!")
	}
}
