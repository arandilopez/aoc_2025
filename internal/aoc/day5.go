package aoc

import (
	"fmt"
	"sort"
	"strings"

	"github.com/arandilopez/aoc_2025/internal/datastructures"
)

func (aoc AdventOfCode) Day5Part1(input *string) (string, error) {
	inputLines := strings.Split(*input, "\n")

	// first lines until empty line are ranges
	ranges := make([][2]int64, 0)
	i := 0
	for ; i < len(inputLines); i++ {
		line := strings.TrimSpace(inputLines[i])
		if line == "" {
			i++
			break
		}
		var start, end int64
		_, err := fmt.Sscanf(line, "%d-%d", &start, &end)
		if err != nil {
			return "", fmt.Errorf("Day5Part1: failed to parse range from line '%s': %v", line, err)
		}
		ranges = append(ranges, [2]int64{start, end})
	}

	// remaining lines are numbers to check
	numbers := make([]int64, 0)
	for ; i < len(inputLines); i++ {
		line := strings.TrimSpace(inputLines[i])
		if line == "" {
			continue
		}
		var num int64
		_, err := fmt.Sscanf(line, "%d", &num)
		if err != nil {
			return "", fmt.Errorf("Day5Part1: failed to parse number from line '%s': %v", line, err)
		}
		numbers = append(numbers, num)
	}

	okNumbers := datastructures.NewSet[int64]()

	// Check each number against the ranges
	for _, num := range numbers {
		for _, r := range ranges {
			if num >= r[0] && num <= r[1] {
				okNumbers.Add(num)
				break
			}
		}
	}

	count := okNumbers.Size()
	return fmt.Sprintf("%d", count), nil
}

func (aoc AdventOfCode) Day5Part2(input *string) (string, error) {
	inputLines := strings.Split(*input, "\n")
	ranges := make([][2]int64, 0)

	// first lines until empty line are ranges
	for _, line := range inputLines {
		if line == "" {
			break
		}
		var start, end int64
		_, err := fmt.Sscanf(line, "%d-%d", &start, &end)
		if err != nil {
			return "", fmt.Errorf("Day5Part1: failed to parse range from line '%s': %v", line, err)
		}
		ranges = append(ranges, [2]int64{start, end})
	}

	mergedRanges := make([][2]int64, 0)
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	// Merge overlapping ranges
	for _, r := range ranges {
		if len(mergedRanges) == 0 {
			mergedRanges = append(mergedRanges, r)
			continue
		}
		lastRange := mergedRanges[len(mergedRanges)-1]
		if r[0] <= lastRange[1] {
			// Overlapping ranges, merge them
			if r[1] > lastRange[1] {
				lastRange[1] = r[1]
			}
			mergedRanges[len(mergedRanges)-1] = lastRange
		} else {
			mergedRanges = append(mergedRanges, r)
		}
	}

	totalCount := int64(0)

	for _, r := range mergedRanges {
		totalCount += r[1] - r[0] + 1
	}

	return fmt.Sprintf("%d", totalCount), nil
}
