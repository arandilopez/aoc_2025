package test

import (
	"testing"

	"github.com/arandilopez/aoc_2025/internal/aoc"
)

var inputDay1 = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func TestDay1Part1(t *testing.T) {
	expected := "3"

	aocInstance := aoc.AdventOfCode{}
	result, err := aocInstance.Day1Part1(&inputDay1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result != expected {
		t.Errorf("Day1Part1(): Expected %v, got %v", expected, result)
	}
}

func TestDay1Part2(t *testing.T) {
	expected := "6"

	aocInstance := aoc.AdventOfCode{}
	result, err := aocInstance.Day1Part2(&inputDay1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result != expected {
		t.Errorf("Day1Part2: Expected %v, got %v", expected, result)
	}
}
