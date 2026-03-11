package aoc

import (
	"fmt"
	"sort"
	"strings"

	"github.com/arandilopez/aoc_2025/internal/maths"
)

type Point2D struct {
	X int
	Y int
}

type Rectangle struct {
	MinX int
	MinY int
	MaxX int
	MaxY int
	Area int
}

func (aoc AdventOfCode) Day9Part1(input *string) (string, error) {
	points := make([]Point2D, 0)

	for line := range strings.SplitSeq(*input, "\n") {
		var x, y int
		_, err := fmt.Sscanf(line, "%d,%d", &x, &y)
		if err != nil {
			return "", err
		}
		points = append(points, Point2D{X: x, Y: y})
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i].X == points[j].X {
			return points[i].Y < points[j].Y
		}
		return points[i].X < points[j].X
	})

	maxArea := 0

	for i, p1 := range points {
		for j, p2 := range points {
			if i == j {
				continue
			}

			if j < i {
				continue
			}
			area := int(maths.ManhattanArea(float64(p1.X), float64(p1.Y), float64(p2.X), float64(p2.Y)))
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return fmt.Sprintf("%d", maxArea), nil
}

func (aoc AdventOfCode) Day9Part2(input *string) (string, error) {
	points := make([]Point2D, 0)

	for line := range strings.SplitSeq(*input, "\n") {
		var x, y int
		_, err := fmt.Sscanf(line, "%d,%d", &x, &y)
		if err != nil {
			return "", err
		}
		points = append(points, Point2D{X: x, Y: y})
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i].X == points[j].X {
			return points[i].Y < points[j].Y
		}
		return points[i].X < points[j].X
	})

	maxArea := 0

	for i, p1 := range points {
		for j, p2 := range points {
			if i == j {
				continue
			}

			if j < i {
				continue
			}
			rectangle := createRectangle(p1, p2)
			if rectangle.Area > maxArea && !intersects(rectangle, points) {
				maxArea = rectangle.Area
			}
		}
	}

	return fmt.Sprintf("%d", maxArea), nil
}

func intersects(rectangle Rectangle, points []Point2D) bool {
	for i := range len(points) {
		p1 := points[i]
		p2 := points[(i+1)%len(points)] // wrap around to the first point
		minX, maxX, minY, maxY := minMax(p1, p2)

		if maxY < rectangle.MinY || minY > rectangle.MaxY || maxX < rectangle.MinX || minX > rectangle.MaxX {
			continue
		}
		return true
	}

	return false
}

func createRectangle(p1, p2 Point2D) Rectangle {
	minX, maxX, minY, maxY := minMax(p1, p2)
	area := int(maths.ManhattanArea(float64(minX), float64(minY), float64(maxX), float64(maxY)))
	return Rectangle{
		MinX: minX,
		MinY: minY,
		MaxX: maxX,
		MaxY: maxY,
		Area: area,
	}
}

func minMax(p1, p2 Point2D) (int, int, int, int) {
	minX := min(p1.X, p2.X)
	maxX := max(p1.X, p2.X)
	minY := min(p1.Y, p2.Y)
	maxY := max(p1.Y, p2.Y)
	return minX, maxX, minY, maxY
}
