package aoc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/arandilopez/aoc_2025/internal/utils"
)

func (aoc AdventOfCode) Day2Part1(input *string) (string, error) {
	var sum int64 = 0

	for r := range strings.SplitSeq(*input, ",") {
		boundsStr := strings.Split(r, "-")
		boundsInt := utils.MapTo(boundsStr, func(s string) int64 {
			n, _ := strconv.ParseInt(s, 10, 64)
			return n
		})
		invalidIDs := make([]int64, 0)

		for id := boundsInt[0]; id <= boundsInt[1]; id++ {
			invalidID := detectInvalidID(id)
			sum += invalidID

			if invalidID != 0 {
				invalidIDs = append(invalidIDs, invalidID)
			}
		}

		if len(invalidIDs) > 0 {
			fmt.Printf("%s has %d invalid IDs, %v\n", r, len(invalidIDs), invalidIDs)
		}
	}
	return fmt.Sprintf("%d", sum), nil
}

func (aoc AdventOfCode) Day2Part2(input *string) (string, error) {
	var sum int64 = 0

	for r := range strings.SplitSeq(*input, ",") {
		boundsStr := strings.Split(r, "-")
		boundsInt := utils.MapTo(boundsStr, func(s string) int64 {
			n, _ := strconv.ParseInt(s, 10, 64)
			return n
		})
		invalidIDs := make([]int64, 0)

		for id := boundsInt[0]; id <= boundsInt[1]; id++ {
			invalidID := detectMoreInvalidID(id)
			sum += invalidID

			if invalidID != 0 {
				invalidIDs = append(invalidIDs, invalidID)
			}
		}

		if len(invalidIDs) > 0 {
			fmt.Printf("%s has %d invalid IDs, %v\n", r, len(invalidIDs), invalidIDs)
		}
	}
	return fmt.Sprintf("%d", sum), nil
}

func detectMoreInvalidID(id int64) int64 {
	idStr := fmt.Sprintf("%d", id)
	for size := 1; size <= len(idStr)/2; size++ {
		if len(idStr)%size != 0 {
			continue
		}

		pattern := idStr[:size]
		matchesPattern := true

		for i := 0; i < len(idStr); i += size {
			if idStr[i:i+size] != pattern {
				matchesPattern = false
				break
			}
		}

		if matchesPattern {
			return id
		}
	}

	return 0
}

func detectInvalidID(id int64) int64 {
	idStr := fmt.Sprintf("%d", id)

	if len(idStr)%2 != 0 {
		return 0
	}

	halfID := idStr[:len(idStr)/2]
	restID := idStr[len(idStr)/2:]

	if halfID == restID {
		return id
	}

	return 0
}
