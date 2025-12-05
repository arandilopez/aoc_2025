package aoc

import (
	"fmt"
	"strings"

	"github.com/arandilopez/aoc_2025/internal/datastructures"
)

func (aoc AdventOfCode) Day4Part1(input *string) (string, error) {
	count := 0
	grid, err := initGridFromInput(input)
	if err != nil {
		return "", err
	}

	fmt.Println(grid.String())

	for row := 0; row < grid.Rows; row++ {
		for col := 0; col < grid.Cols; col++ {
			cell, _ := grid.Get(row, col)
			if cell != "@" {
				continue
			}
			adjacentCells := getAdjacentCells(grid, row, col)
			if countAdjacentSymbols(adjacentCells, "@") <= 3 {
				count++
			}
		}
	}

	return fmt.Sprintf("%d", count), nil
}

func (aoc AdventOfCode) Day4Part2(input *string) (string, error) {
	grid, err := initGridFromInput(input)
	if err != nil {
		return "", err
	}

	fmt.Println(grid.String())

	count := 0

	for ok := true; ok; {
		ok = false
		// Will stop when no more cells can be marked
		for row := 0; row < grid.Rows; row++ {
			for col := 0; col < grid.Cols; col++ {
				cell, _ := grid.Get(row, col)
				if cell != "@" {
					continue
				}
				adjacentCells := getAdjacentCells(grid, row, col)
				if countAdjacentSymbols(adjacentCells, "@") <= 3 {
					count++
					ok = true
					grid.Set(row, col, "x")
				}
			}
		}
	}

	return fmt.Sprintf("%d", count), nil
}

func initGridFromInput(input *string) (*datastructures.Grid[string], error) {
	rowsData := strings.Split(*input, "\n")
	if len(rowsData) == 0 {
		return nil, fmt.Errorf("empty input")
	}

	rows := len(rowsData)
	cols := len(rowsData[0])
	fmt.Printf("Initializing grid with %d rows and %d cols\n", rows, cols)
	grid := datastructures.NewGrid[string](rows, cols)

	for _, rowStr := range rowsData {
		row := strings.Split(rowStr, "")
		if err := grid.AppendRow(row); err != nil {
			return nil, err
		}
	}

	return grid, nil
}

func getAdjacentCells(grid *datastructures.Grid[string], row, col int) []string {
	directions := [][2]int{
		{-1, 0}, // Up
		{1, 0},  // Down
		{0, -1}, // Left
		{0, 1},  // Right
		// Diagonals
		{-1, -1}, // Up-Left
		{-1, 1},  // Up-Right
		{1, -1},  // Down-Left
		{1, 1},   // Down-Right
	}

	var adjacentCells []string
	for _, dir := range directions {
		newRow, newCol := row+dir[0], col+dir[1]
		if cell, ok := grid.Get(newRow, newCol); ok {
			adjacentCells = append(adjacentCells, cell)
		}
	}

	return adjacentCells
}

func countAdjacentSymbols(haystack []string, needle string) int {
	count := 0
	for _, cell := range haystack {
		if cell == needle {
			count++
		}
	}
	return count
}
