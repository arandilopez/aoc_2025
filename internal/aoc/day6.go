package aoc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/arandilopez/aoc_2025/internal/maths"
	"github.com/arandilopez/aoc_2025/internal/utils"
)

func (aoc AdventOfCode) Day6Part1(input *string) (string, error) {
	lines := strings.Split(*input, "\n")
	operators := strings.Fields(lines[len(lines)-1])
	colSize := len(operators)
	rowSize := len(lines) - 1

	numbers := make([][]int, 0)

	for i := range rowSize {
		numStrs := strings.Fields(lines[i])
		numRow := utils.MapTo(numStrs, func(s string) int {
			num, _ := strconv.Atoi(s)
			return num
		})
		numbers = append(numbers, numRow)
	}

	result := 0
	for col := range colSize {
		colNumber := make([]int, rowSize)
		for row := range rowSize {
			colNumber[row] = numbers[row][col]
		}
		switch operators[col] {
		case "*":
			result += maths.Prod(colNumber)
		case "+":
			result += maths.Sum(colNumber)
		default:
			return "", fmt.Errorf("Day6Part1: invalid operator %s at col %d", operators[col], col)
		}
	}

	return fmt.Sprintf("%d", result), nil
}

func (aoc AdventOfCode) Day6Part2(input *string) (string, error) {
	lines := strings.Split(*input, "\n")

	colSize := 0
	for i := range len(lines) - 1 {
		if len(lines[i]) > colSize {
			colSize = len(lines[i])
		}
	}

	result := 0

	startOperatorIndexes := make([]int, 0)
	operatorsLine := lines[len(lines)-1]
	for index, char := range operatorsLine {
		if char == '*' || char == '+' {
			startOperatorIndexes = append(startOperatorIndexes, index)
		}
	}

	for index, startIndex := range startOperatorIndexes {
		endIndex := colSize // Max column size

		if index < len(startOperatorIndexes)-1 {
			endIndex = startOperatorIndexes[index+1]
		}

		if index == len(startOperatorIndexes)-1 {
			endIndex = colSize
		}

		numStrs := make([]string, 0)
		for row := range len(lines) - 1 {
			line := lines[row]

			end := endIndex
			if startIndex >= len(line) {
				continue
			}

			if endIndex > len(line) {
				end = len(line)
			}

			numStrs = append(numStrs, line[startIndex:end]) // number can't be trimmed, spaces are significant
		}

		nums, err := cephalopodNumbersToInt(numStrs)
		if err != nil {
			return "", fmt.Errorf("Day6Part2: %v", err)
		}

		switch operatorsLine[startIndex] {
		case '*':
			result += maths.Prod(nums)
		case '+':
			result += maths.Sum(nums)
		default:
			return "", fmt.Errorf("Day6Part2: invalid operator %c at index %d", operatorsLine[startIndex], startIndex)
		}
	}

	return fmt.Sprintf("%d", result), nil
}

func cephalopodNumbersToInt(numStrs []string) ([]int, error) {
	maxSize := 0
	for _, s := range numStrs {
		if len(s) > maxSize {
			maxSize = len(s)
		}
	}

	paddedNums := make([]string, len(numStrs))
	for i, s := range numStrs {
		paddedNums[i] = s + strings.Repeat(" ", maxSize-len(s))
	}

	result := make([]int, 0)

	for col := range maxSize {
		numStr := ""
		for row := range len(paddedNums) {
			char := paddedNums[row][col]
			if char != ' ' {
				numStr = numStr + string(char)
			}
		}
		if len(numStr) == 0 {
			continue
		}
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, fmt.Errorf("cephalopodNumbersToInt: invalid number %s", numStr)
		}
		result = append(result, num)
	}

	return result, nil
}
