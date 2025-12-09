package test

import (
	"testing"

	"github.com/arandilopez/aoc_2025/internal/aoc"
)

var day6Input = `123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +`

func TestDay6Part1(t *testing.T) {
	expected := "4277556"
	aocInstance := aoc.AdventOfCode{}
	result, err := aocInstance.Day6Part1(&day6Input)
	if err != nil {
		t.Errorf("Day6Part1 returned an error: %v", err)
	}
	if result != expected {
		t.Errorf("Day6Part1 = %v; want %v", result, expected)
	}
}

func TestDay6Part2(t *testing.T) {
	expected := "3263827"
	aocInstance := aoc.AdventOfCode{}
	result, err := aocInstance.Day6Part2(&day6Input)
	if err != nil {
		t.Errorf("Day6Part2 returned an error: %v", err)
	}
	if result != expected {
		t.Errorf("Day6Part2 = %v; want %v", result, expected)
	}
}
