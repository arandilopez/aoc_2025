package test

import (
	"testing"

	"github.com/arandilopez/aoc_2025/internal/aoc"
)

var inputDay4 = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func TestDay4Part1(t *testing.T) {
	aocInstance := aoc.AdventOfCode{}
	expected := "13"
	result, err := aocInstance.Day4Part1(&inputDay4)
	if err != nil {
		t.Errorf("Day4Part1 returned an error: %v", err)
	}

	if result != expected {
		t.Errorf("Day4Part1 = %v; want %v", result, expected)
	}
}

func TestDay4Part2(t *testing.T) {
	aocInstance := aoc.AdventOfCode{}
	expected := "43"
	result, err := aocInstance.Day4Part2(&inputDay4)
	if err != nil {
		t.Errorf("Day4Part2 returned an error: %v", err)
	}

	if result != expected {
		t.Errorf("Day4Part2 = %v; want %v", result, expected)
	}
}
