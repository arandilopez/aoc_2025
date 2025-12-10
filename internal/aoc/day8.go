package aoc

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/arandilopez/aoc_2025/internal/datastructures"
	"github.com/arandilopez/aoc_2025/internal/maths"
	"github.com/arandilopez/aoc_2025/internal/utils"
)

type Point struct {
	X int
	Y int
	Z int
}

type Distance struct {
	P1       Point
	P2       Point
	Distance int
}

func (aoc AdventOfCode) Day8Part1(input *string) (string, error) {
	points := make([]Point, 0)

	for line := range strings.SplitSeq(*input, "\n") {
		point := Point{}
		_, err := fmt.Sscanf(line, "%d,%d,%d", &point.X, &point.Y, &point.Z)
		if err != nil {
			return "", err
		}
		points = append(points, point)
	}

	distances := make([]Distance, 0)

	for i, p1 := range points {
		for j, p2 := range points {
			if i == j {
				continue
			}

			// duplicate check
			if j < i {
				continue
			}

			distance := maths.EuclideanDistance3D(float64(p1.X), float64(p1.Y), float64(p1.Z), float64(p2.X), float64(p2.Y), float64(p2.Z))
			distances = append(distances, Distance{
				P1:       p1,
				P2:       p2,
				Distance: int(distance),
			})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	connections := []*datastructures.Set[Point]{}

	for _, distance := range distances[:10] {
		foundConnection := false

		for _, connection := range connections {
			if connection.Contains(distance.P1) || connection.Contains(distance.P2) {
				connection.Add(distance.P1)
				connection.Add(distance.P2)
				foundConnection = true
				break
			}
		}

		if !foundConnection {
			newConnection := datastructures.NewSet[Point]()
			newConnection.Add(distance.P1)
			newConnection.Add(distance.P2)
			connections = append(connections, newConnection)
		}
	}

	// Merge connections that share points
	merged := true
	for merged {
		merged = false
		for i := 0; i < len(connections); i++ {
			for j := i + 1; j < len(connections); j++ {
				contained := slices.ContainsFunc(connections[i].Elements(), func(point Point) bool {
					return connections[j].Contains(point)
				})

				if contained {
					for _, p := range connections[j].Elements() {
						connections[i].Add(p)
					}
					connections = append(connections[:j], connections[j+1:]...)
					merged = true
					break
				}
			}
		}
	}

	// Sort connections by size
	sort.Slice(connections, func(i, j int) bool {
		return connections[i].Size() > connections[j].Size()
	})

	sizes := utils.MapTo(connections, func(s *datastructures.Set[Point]) int {
		return s.Size()
	})

	total := maths.Prod(sizes[:3])

	return fmt.Sprintf("%d", total), nil
}

func (aoc AdventOfCode) Day8Part2(input *string) (string, error) {
	points := make([]Point, 0)

	for line := range strings.SplitSeq(*input, "\n") {
		point := Point{}
		_, err := fmt.Sscanf(line, "%d,%d,%d", &point.X, &point.Y, &point.Z)
		if err != nil {
			return "", err
		}
		points = append(points, point)
	}

	distances := make([]Distance, 0)

	for i, p1 := range points {
		for j, p2 := range points {
			if i == j {
				continue
			}

			// duplicate check
			if j < i {
				continue
			}

			distance := maths.EuclideanDistance3D(float64(p1.X), float64(p1.Y), float64(p1.Z), float64(p2.X), float64(p2.Y), float64(p2.Z))
			distances = append(distances, Distance{
				P1:       p1,
				P2:       p2,
				Distance: int(distance),
			})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	unionFind := datastructures.NewDisjointSet[Point]()

	for _, point := range points {
		unionFind.Add(point)
	}

	lastPairPoints := make([]Point, 2)

	for _, distance := range distances {
		if ok := unionFind.Union(distance.P1, distance.P2); ok {
			lastPairPoints[0] = distance.P1
			lastPairPoints[1] = distance.P2
		}
	}

	distanceFromWall := lastPairPoints[0].X * lastPairPoints[1].X

	return fmt.Sprintf("%d", distanceFromWall), nil
}
