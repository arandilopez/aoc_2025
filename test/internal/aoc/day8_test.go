package test

import (
	"testing"

	"github.com/arandilopez/aoc_2025/internal/aoc"
)

var inputDay8 = `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

func TestDay8Part1(t *testing.T) {
	expected := "40"
	aocInstance := aoc.AdventOfCode{}
	result, err := aocInstance.Day8Part1(&inputDay8)
	if err != nil {
		t.Errorf("Day8Part1 returned an error: %v", err)
	}
	if result != expected {
		t.Errorf("Day8Part1 = %v; want %v", result, expected)
	}
}

func TestDay8Part2(t *testing.T) {
	expected := "25272"
	aocInstance := aoc.AdventOfCode{}
	result, err := aocInstance.Day8Part2(&inputDay8)
	if err != nil {
		t.Errorf("Day8Part2 returned an error: %v", err)
	}
	if result != expected {
		t.Errorf("Day8Part2 = %v; want %v", result, expected)
	}
}
