package datastructures

import (
	"fmt"
)

type Grid[T any] struct {
	Cells      [][]T
	Rows, Cols int
}

func NewGrid[T any](rows, cols int) *Grid[T] {
	cells := make([][]T, 0)
	return &Grid[T]{
		Cells: cells,
		Rows:  rows,
		Cols:  cols,
	}
}

func (g *Grid[T]) AppendRow(row []T) error {
	if len(row) != g.Cols {
		return fmt.Errorf("Grid#AppendRow: invalid row, got %d cols, expected %d", len(row), g.Cols)
	}

	if len(g.Cells) >= g.Rows {
		return fmt.Errorf("Grid#AppendRow: cannot append more rows, grid already has %d rows", len(g.Cells))
	}
	g.Cells = append(g.Cells, row)

	return nil
}

func (g *Grid[T]) Get(row, col int) (T, bool) {
	if row < 0 || row >= g.Rows || col < 0 || col >= g.Cols {
		// out of bounds
		var zero T
		return zero, false
	}
	return g.Cells[row][col], true
}

func (g *Grid[T]) Set(row, col int, value T) bool {
	if row < 0 || row >= g.Rows || col < 0 || col >= g.Cols {
		return false
	}
	g.Cells[row][col] = value
	return true
}

func (g *Grid[T]) String() string {
	result := ""
	for _, row := range g.Cells {
		result += fmt.Sprintln(row)
	}
	return result
}
