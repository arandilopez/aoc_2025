// Package aoc provides solutions for Advent of Code challenges.
package aoc

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/arandilopez/aoc_2025/internal/utils"
)

// AdventOfCode struct serves as a receiver for all day's part functions.
type AdventOfCode struct{}

var ErrNotImplemented = errors.New("not implemented")

// ExecSolution executes the specified day's part function from the AdventOfCode interface using the provided arguments.
func ExecSolution(aoc AdventOfCode, args *utils.Args) (string, error) {
	// Infer the function name based on day and part, e.g., Day1Part1
	functionName := "Day" + string(rune('0'+*args.Day)) + "Part" + string(rune('0'+*args.Part))
	function := reflect.ValueOf(aoc).MethodByName(functionName)

	if !function.IsValid() {
		return "", fmt.Errorf("function %s not found", functionName)
	}

	// Read the input file contents
	fileContents, e := utils.ReadFileContents(args.InputFile)
	if e != nil {
		return "", e
	}

	// Call the function with the file contents (input)
	results := function.Call([]reflect.Value{reflect.ValueOf(fileContents)})

	result, err := results[0], results[1]

	if !err.IsNil() {
		return "", err.Interface().(error)
	}

	if result.Kind() != reflect.String {
		return "", errors.New("unexpected return type")
	}

	return result.String(), nil
}
