package test

import (
	"testing"

	"github.com/arandilopez/aoc_2025/internal/aoc"
)

var inputDay3 = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestDay3Part1(t *testing.T) {
	aocInstance := aoc.AdventOfCode{}
	expected := "357"

	result, err := aocInstance.Day3Part1(&inputDay3)
	if err != nil {
		t.Errorf("Day3Part1 returned an error: %v", err)
	}

	if result != expected {
		t.Errorf("Day3Part1 = %v; want %v", result, expected)
	}
}

func TestDay3Part2(t *testing.T) {
	aocInstance := aoc.AdventOfCode{}
	expected := "3121910778619"

	result, err := aocInstance.Day3Part2(&inputDay3)
	if err != nil {
		t.Errorf("Day3Part2 returned an error: %v", err)
	}

	if result != expected {
		t.Errorf("Day3Part2 = %v; want %v", result, expected)
	}
}
