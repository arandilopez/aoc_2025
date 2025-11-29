package aoc

import (
	"testing"

	"github.com/arandilopez/aoc_2025/internal/aoc"
)

func TestDay1Part1(t *testing.T) {
	input := "1 2 3"
	expected := "42" // Replace with the expected result for the given input

	aocInstance := aoc.AdventOfCode{}
	result, err := aocInstance.Day1Part1(&input)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result != expected {
		t.Errorf("Day1Part1(%v): Expected %v, got %v", input, expected, result)
	}
}

func TestDay1Part2(t *testing.T) {
	t.Skip("Not implemented yet")
	input := "4 5 6"
	expected := "84" // Replace with the expected result for the given input

	aocInstance := aoc.AdventOfCode{}
	result, err := aocInstance.Day1Part2(&input)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result != expected {
		t.Errorf("Day1Part2(%v): Expected %v, got %v", input, expected, result)
	}
}
