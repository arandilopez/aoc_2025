package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

func (aoc AdventOfCode) Day1Part1(input *string) (string, error) {
	var position int64 = 50
	var zeroCounter int64 = 0

	fmt.Printf("The dial starts by pointing at %d\n", position)

	for rotation := range strings.SplitSeq(*input, "\n") {
		direction := rune(rotation[0])
		steps, err := strconv.ParseInt(rotation[1:], 10, 64)
		if err != nil {
			return "", err
		}

		if steps >= 100 {
			steps = steps % 100
		}

		switch direction {
		case 'L':
			position -= steps
			if position < 0 {
				position += 100
			}
		case 'R':
			position = (position + steps) % 100
		default:
			return "", fmt.Errorf("invalid direction: %v", direction)
		}

		fmt.Printf("The dial is rotated %c%d to point at %d", direction, steps, position)

		if position == 0 {
			fmt.Printf(" (exactly zero)")

			zeroCounter++
		}

		fmt.Printf("\n")
	}

	return strconv.FormatInt(zeroCounter, 10), nil
}

func (aoc AdventOfCode) Day1Part2(input *string) (string, error) {
	zeroCounter := int64(0)
	position := int64(50)

	fmt.Printf("The dial starts by pointing at %d\n", position)

	for rotation := range strings.SplitSeq(*input, "\n") {
		timesMadeItToZero := int64(0)
		previousPosition := position
		direction := rune(rotation[0])
		steps, err := strconv.ParseInt(rotation[1:], 10, 64)
		if err != nil {
			return "", err
		}

		if steps >= 100 {
			fullLoops := steps / 100 // would cross zero this many times
			if fullLoops > 0 {
				timesMadeItToZero += fullLoops
				zeroCounter += fullLoops
			}
			steps = steps % 100
		}

		switch direction {
		case 'L':
			position -= steps
			if position < 0 {
				if previousPosition != 0 {
					zeroCounter++ // crossed zero once
					timesMadeItToZero++
				}
				position += 100
			}
		case 'R':
			position += steps
			if position > 100 && previousPosition != 0 {
				zeroCounter++ // crossed zero once
				timesMadeItToZero++
			}
			position = position % 100
		default:
			return "", fmt.Errorf("invalid direction: %v", direction)
		}

		fmt.Printf("The dial is rotated %c%d to point at %d", direction, steps, position)
		if position == 0 {
			fmt.Printf(" (exactly zero)")
		}
		if timesMadeItToZero > 0 {
			fmt.Printf("; during this rotation, it points at zero %d time(s)", timesMadeItToZero)
		}
		fmt.Printf("\n")

		if position == 0 {
			zeroCounter++
		}
	}

	return strconv.FormatInt(zeroCounter, 10), nil
}
