package test

import (
	"testing"

	"github.com/arandilopez/aoc_2025/internal/aoc"
)

var inputDay9 = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

func TestDay9Part1(t *testing.T) {
	expected := "50"
	aoc := aoc.AdventOfCode{}
	result, err := aoc.Day9Part1(&inputDay9)
	if err != nil {
		t.Errorf("Day9Part1 returned an error: %v", err)
	}
	if result != expected {
		t.Errorf("Day9Part1 = %v; want %v", result, expected)
	}
}

func TestDay9Part2(t *testing.T) {
	expected := "24"
	aoc := aoc.AdventOfCode{}
	result, err := aoc.Day9Part2(&inputDay9)
	if err != nil {
		t.Errorf("Day9Part2 returned an error: %v", err)
	}
	if result != expected {
		t.Errorf("Day9Part2 = %v; want %v", result, expected)
	}
}
