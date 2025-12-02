package test

import (
	"testing"

	"github.com/arandilopez/aoc_2025/internal/aoc"
)

var inputDay2 = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

func TestDay2Part1(t *testing.T) {
	expected := "1227775554"
	aocInstance := aoc.AdventOfCode{}
	result, err := aocInstance.Day2Part1(&inputDay2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result != expected {
		t.Errorf("Day2Part1(): Expected %v, got %v", expected, result)
	}
}

func TestDay2Part2(t *testing.T) {
	expected := "4174379265"
	aocInstance := aoc.AdventOfCode{}
	result, err := aocInstance.Day2Part2(&inputDay2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result != expected {
		t.Errorf("Day2Part2: Expected %v, got %v", expected, result)
	}
}
