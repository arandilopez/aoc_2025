// Package utils provides utility functions for executing Advent of Code solutions.
package utils

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/arandilopez/aoc_2025/internal/aoc"
)

// ExecSolution executes the specified day's part function from the AdventOfCode interface.
func ExecSolution(aoc aoc.AdventOfCode, args *Args) (string, error) {
	functionName := "Day" + string(rune('0'+args.Day)) + "Part" + string(rune('0'+args.Part))
	function := reflect.ValueOf(aoc).MethodByName(functionName)

	if !function.IsValid() {
		return "", fmt.Errorf("function %s not found", functionName)
	}

	results := function.Call([]reflect.Value{reflect.ValueOf(&args.InputFile)})

	result, err := results[0], results[1]

	if !err.IsNil() {
		return "", err.Interface().(error)
	}

	if result.Kind() != reflect.String {
		return "", errors.New("unexpected return type")
	}

	return result.String(), nil
}
