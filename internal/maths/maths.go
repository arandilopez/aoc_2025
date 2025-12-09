// Package maths provides mathematical utility functions.
package maths

import "slices"

type Numeric interface {
	~int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Sum[T Numeric](numbers []T) T {
	var total T
	for _, num := range numbers {
		total += num
	}
	return total
}

func Prod[T Numeric](numbers []T) T {
	var product T = 1
	for _, num := range numbers {
		product *= num
	}
	return product
}

func Avg[T Numeric](numbers []T) float64 {
	if len(numbers) == 0 {
		return 0
	}
	var total T
	for _, num := range numbers {
		total += num
	}
	return float64(total) / float64(len(numbers))
}

func Min[T Numeric](numbers []T) T {
	if len(numbers) == 0 {
		var zero T
		return zero
	}
	slices.Sort(numbers)
	return numbers[0]
}

func Max[T Numeric](numbers []T) T {
	if len(numbers) == 0 {
		var zero T
		return zero
	}
	slices.Sort(numbers)
	return numbers[len(numbers)-1]
}
