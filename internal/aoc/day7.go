package aoc

import (
	"fmt"
	"strings"

	"github.com/arandilopez/aoc_2025/internal/datastructures"
	"github.com/arandilopez/aoc_2025/internal/maths"
)

func (aoc AdventOfCode) Day7Part1(input *string) (string, error) {
	lines := strings.Split(*input, "\n")
	grid := datastructures.NewGrid[rune](len(lines), len(lines[0]))

	for _, line := range lines {
		row := make([]rune, 0)
		for _, char := range line {
			row = append(row, char)
		}
		e := grid.AppendRow(row)
		if e != nil {
			return "", e
		}
	}

	fmt.Println(grid.String())

	totalSplit := 0

	for row := range grid.Rows {
		for col := range grid.Cols {
			cell, _ := grid.Get(row, col)
			if row == 0 && cell != 'S' {
				continue
			}

			if cell == 'S' || cell == '|' {
				adjacentPositions := computeAdjacentPositions(row, col)
				downPos := adjacentPositions["down"]
				downCell, _ := grid.Get(downPos[0], downPos[1])
				if downCell == '^' {
					downAdjacentPositions := computeAdjacentPositions(downPos[0], downPos[1])
					leftPos := downAdjacentPositions["left"]
					leftCell, _ := grid.Get(leftPos[0], leftPos[1])

					totalSplit++

					if leftCell == '.' {
						grid.Set(leftPos[0], leftPos[1], '|')
					}

					rightPos := downAdjacentPositions["right"]
					rightCell, _ := grid.Get(rightPos[0], rightPos[1])
					if rightCell == '.' {
						grid.Set(rightPos[0], rightPos[1], '|')
					}
				}

				if downCell == '.' {
					grid.Set(downPos[0], downPos[1], '|')
				}
			}
		}
	}

	fmt.Println(grid.String())

	return fmt.Sprintf("%d", totalSplit), nil
}

func (aoc AdventOfCode) Day7Part2(input *string) (string, error) {
	lines := strings.Split(*input, "\n")
	grid := datastructures.NewGrid[rune](len(lines), len(lines[0]))

	for _, line := range lines {
		row := make([]rune, 0)
		for _, char := range line {
			row = append(row, char)
		}
		e := grid.AppendRow(row)
		if e != nil {
			return "", e
		}
	}

	possibleBeamPaths := make([]int, len(lines[0]))

	_, y, found := grid.Find(func(r rune) bool { return r == 'S' })
	if !found {
		return "", fmt.Errorf("start position 'S' not found in the grid")
	}

	possibleBeamPaths[y] = 1

	for row := range grid.Rows {
		for col := range grid.Cols {
			cell, _ := grid.Get(row, col)
			if row == 0 && cell != 'S' {
				continue
			}

			if cell == 'S' || cell == '|' {
				adjacentPositions := computeAdjacentPositions(row, col)
				downPos := adjacentPositions["down"]
				downCell, _ := grid.Get(downPos[0], downPos[1])
				if downCell == '^' && possibleBeamPaths[col] > 0 {

					// Increases the visited paths that the beam can take
					possibleBeamPaths[col-1] += possibleBeamPaths[col]
					possibleBeamPaths[col+1] += possibleBeamPaths[col]
					possibleBeamPaths[col] = 0

					downAdjacentPositions := computeAdjacentPositions(downPos[0], downPos[1])
					leftPos := downAdjacentPositions["left"]
					leftCell, _ := grid.Get(leftPos[0], leftPos[1])

					if leftCell == '.' {
						grid.Set(leftPos[0], leftPos[1], '|')
					}

					rightPos := downAdjacentPositions["right"]
					rightCell, _ := grid.Get(rightPos[0], rightPos[1])
					if rightCell == '.' {
						grid.Set(rightPos[0], rightPos[1], '|')
					}
				}

				if downCell == '.' {
					grid.Set(downPos[0], downPos[1], '|')
				}
			}
		}
	}

	pathSum := maths.Sum(possibleBeamPaths)

	return fmt.Sprintf("%d", pathSum), nil
}

func computeAdjacentPositions(x, y int) map[string][2]int {
	return map[string][2]int{
		"up":    {x - 1, y},
		"down":  {x + 1, y},
		"left":  {x, y - 1},
		"right": {x, y + 1},
	}
}
