package aoc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/arandilopez/aoc_2025/internal/datastructures"
	"github.com/arandilopez/aoc_2025/internal/utils"
)

func (aoc AdventOfCode) Day3Part1(input *string) (string, error) {
	banks := strings.SplitSeq(*input, "\n")

	totalMaxJoltage := int64(0)

	for bank := range banks {
		bankStrs := strings.Split(bank, "")
		batteryBank := utils.MapTo(bankStrs, func(s string) int64 {
			n, _ := strconv.ParseInt(s, 10, 64)
			return n
		})
		maxJoltage, err := findMaxJoltage(&batteryBank, 2)
		if err != nil {
			return "", err
		}

		fmt.Printf("Max joltage for bank %s: %d\n", bank, maxJoltage)

		totalMaxJoltage += maxJoltage
	}

	fmt.Printf("Total max joltage: %d\n", totalMaxJoltage)

	return fmt.Sprintf("%d", totalMaxJoltage), nil
}

func (aoc AdventOfCode) Day3Part2(input *string) (string, error) {
	banks := strings.SplitSeq(*input, "\n")

	totalMaxJoltage := int64(0)

	for bank := range banks {
		bankStrs := strings.Split(bank, "")
		batteryBank := utils.MapTo(bankStrs, func(s string) int64 {
			n, _ := strconv.ParseInt(s, 10, 64)
			return n
		})
		maxJoltage, err := findMaxJoltage(&batteryBank, 12)
		if err != nil {
			return "", err
		}

		fmt.Printf("Max joltage for bank %s: %d\n", bank, maxJoltage)

		totalMaxJoltage += maxJoltage
	}

	fmt.Printf("Total max joltage: %d\n", totalMaxJoltage)

	return fmt.Sprintf("%d", totalMaxJoltage), nil
}

// Find the largest two battery joltages combined in the order they appear
func findMaxJoltage(bank *[]int64, openBatteries int) (int64, error) {
	batteriesStack := datastructures.NewStack[int64]()
	allowedRemovals := len(*bank) - openBatteries

	for _, battery := range *bank {
		for !batteriesStack.IsEmpty() && allowedRemovals > 0 && batteriesStack.Peek() < battery {
			_, err := batteriesStack.Pop()
			if err != nil {
				return 0, err
			}
			allowedRemovals--
		}
		batteriesStack.Push(battery)
	}

	for batteriesStack.Size() > openBatteries {
		_, err := batteriesStack.Pop()
		if err != nil {
			return 0, err
		}
	}

	finalBatteries := batteriesStack.Elements()
	joltage := strings.Join(utils.MapTo(finalBatteries, func(n int64) string {
		return fmt.Sprintf("%d", n)
	}), "")

	return strconv.ParseInt(joltage, 10, 64)
}
