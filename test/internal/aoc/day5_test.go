package test

import (
	"testing"

	"github.com/arandilopez/aoc_2025/internal/aoc"
)

var inputDay5 = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func TestDay5Part1(t *testing.T) {
	expected := "3"
	aocInstance := aoc.AdventOfCode{}
	result, err := aocInstance.Day5Part1(&inputDay5)
	if err != nil {
		t.Errorf("Day5Part1 returned an error: %v", err)
	}
	if result != expected {
		t.Errorf("Day5Part1 = %v; want %v", result, expected)
	}
}

func TestDay5Part2(t *testing.T) {
	expected := "14"
	aocInstance := aoc.AdventOfCode{}
	result, err := aocInstance.Day5Part2(&inputDay5)
	if err != nil {
		t.Errorf("Day5Part2 returned an error: %v", err)
	}
	if result != expected {
		t.Errorf("Day5Part2 = %v; want %v", result, expected)
	}
}
