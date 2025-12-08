package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/mikej/advent-of-code-2025/shared/input"
)

type Coord struct {
	x, y, z int
}

type JunctionBoxPair struct {
	box1, box2 Coord
	distance   int
}

func main() {
	lines, err := input.ReadFromFile("/Users/mike/Downloads/input-day-8.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := SolvePart1(lines, 1000)

	fmt.Println("Result:", result)
}

func SolvePart1(lines []string, circuitCount int) int {
	coords := parseCoords(lines)
	distances := buildDistances(coords)
	circuits := buildCircuits(distances, circuitCount)
	largesCircuits := findLargest(circuits, 3)

	return largesCircuits[0] * largesCircuits[1] * largesCircuits[2]
}

func findLargest(circuits map[Coord]int, n int) []int {
	sizeMap := make(map[int]int)
	for _, id := range circuits {
		sizeMap[id]++
	}
	var sizes []int
	for _, size := range sizeMap {
		sizes = append(sizes, size)
	}

	slices.Sort(sizes)
	slices.Reverse(sizes)

	return sizes[:n]
}

func buildCircuits(distances []JunctionBoxPair, n int) map[Coord]int {
	circuits := make(map[Coord]int)
	var currentCircuitId int
	for i, nextJoin := range distances {
		circuitId1, inCircuit1 := circuits[nextJoin.box1]
		circuitId2, inCircuit2 := circuits[nextJoin.box2]
		if inCircuit1 && inCircuit2 {
			// both boxes are already in a circuit, join them
			for coord, circuitId := range circuits {
				if circuitId == circuitId2 {
					circuits[coord] = circuitId1
				}
			}
		} else if inCircuit1 {
			// put box 2 in the same circuit as box 1
			circuits[nextJoin.box2] = circuitId1
		} else if inCircuit2 {
			// put box 1 in the same circuit as box 2
			circuits[nextJoin.box1] = circuitId2
		} else {
			// put the 2 together in a new circuit
			circuits[nextJoin.box1] = currentCircuitId
			circuits[nextJoin.box2] = currentCircuitId
			currentCircuitId++
		}

		if i >= n-1 {
			break
		}

	}
	return circuits
}

func buildDistances(coords []Coord) []JunctionBoxPair {
	var distances []JunctionBoxPair
	for i, a := range coords {
		for j, b := range coords {
			if i <= j {
				continue
			}
			xdistance := a.x - b.x
			ydistance := a.y - b.y
			zdistance := a.z - b.z
			distance := xdistance*xdistance + ydistance*ydistance + zdistance*zdistance
			distances = append(distances, JunctionBoxPair{a, b, distance})
		}
	}

	slices.SortFunc(distances, func(a, b JunctionBoxPair) int {
		return cmp.Compare(a.distance, b.distance)
	})

	return distances
}

func parseCoords(lines []string) []Coord {
	coords := make([]Coord, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		coords = append(coords, Coord{x, y, z})
	}
	return coords
}
