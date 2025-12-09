package test

import (
	"testing"

	"github.com/arandilopez/aoc_2025/internal/aoc"
)

var inputDay7 = `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`

func TestDay7Part1(t *testing.T) {
	expected := "21"
	aocInstance := aoc.AdventOfCode{}
	result, err := aocInstance.Day7Part1(&inputDay7)
	if err != nil {
		t.Errorf("Day7Part1 returned an error: %v", err)
	}
	if result != expected {
		t.Errorf("Day7Part1 = %v; want %v", result, expected)
	}
}

func TestDay7Part2(t *testing.T) {
	expected := "40"
	aocInstance := aoc.AdventOfCode{}
	result, err := aocInstance.Day7Part2(&inputDay7)
	if err != nil {
		t.Errorf("Day7Part2 returned an error: %v", err)
	}
	if result != expected {
		t.Errorf("Day7Part2 = %v; want %v", result, expected)
	}
}
