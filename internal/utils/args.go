package utils

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type Args struct {
	Day       int
	Part      int
	InputFile string
}

func ArgsValidation() (*Args, error) {
	day := flag.Int("day", 1, "Day of the Advent of Code to execute")
	part := flag.Int("part", 1, "Part of the day's challenge to execute (1 or 2)")
	inputFile := flag.String("input", "input.in", "Path to the input file")
	flag.Parse()

	if *day < 1 || *day > 25 {
		return nil, errors.New("day must be between 1 and 25")
	}
	if *part != 1 && *part != 2 {
		return nil, errors.New("part must be either 1 or 2")
	}

	if *inputFile == "" || len(*inputFile) == 0 {
		return nil, errors.New("input file path cannot be empty")
	}

	_, err := os.Stat(*inputFile)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("input file '%s' does not exist", *inputFile)
	}

	return &Args{
		Day:       *day,
		Part:      *part,
		InputFile: *inputFile,
	}, nil
}
