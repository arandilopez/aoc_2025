package utils

import (
	"errors"
	"reflect"

	"github.com/arandilopez/aoc_2025/internal/aoc"
)

func ExecDay(aoc aoc.AdventOfCode, day int, part int, inputFile string) (int, error) {
	function := reflect.ValueOf(aoc).MethodByName(
		"Day" + string(rune('0'+day)) +
			"Part" +
			string(rune('0'+part)),
	)

	if !function.IsValid() {
		return 0, errors.New("function not implemented for the specified day and part")
	}

	results := function.Call([]reflect.Value{reflect.ValueOf(inputFile)})

	if len(results) != 1 {
		return 0, errors.New("unexpected number of return values")
	}

	if results[0].Kind() != reflect.Int {
		return 0, errors.New("unexpected return type")
	}

	return int(results[0].Int()), nil
}
