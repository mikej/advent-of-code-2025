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

type Playground struct {
	boxes []Coord
}

func NewPlayground(lines []string) *Playground {
	coords := parseCoords(lines)
	return &Playground{coords}
}

func main() {
	lines, err := input.ReadFromFile("/Users/mike/Downloads/input-day-8.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	pg := NewPlayground(lines)

	fmt.Println("Part 1:", pg.SolvePart1(1000))
	fmt.Println("Part 2:", pg.SolvePart2())
}

func (pg *Playground) SolvePart1(circuitCount int) int {
	distances := pg.buildDistances(pg.boxes)
	circuits, _ := pg.buildCircuits(distances, circuitCount, 1)
	largesCircuits := pg.findLargest(circuits, 3)

	return largesCircuits[0] * largesCircuits[1] * largesCircuits[2]
}

func (pg *Playground) SolvePart2() int {
	distances := pg.buildDistances(pg.boxes)
	_, lastJunctionBoxAdded := pg.buildCircuits(distances, 0, 2)

	return lastJunctionBoxAdded.box1.x * lastJunctionBoxAdded.box2.x
}

func (pg *Playground) findLargest(circuits map[Coord]int, n int) []int {
	sizes := pg.buildCircuitSizes(circuits)

	slices.Sort(sizes)
	slices.Reverse(sizes)

	return sizes[:n]
}

func (pg *Playground) buildCircuitSizes(circuits map[Coord]int) []int {
	sizeMap := make(map[int]int)
	for _, id := range circuits {
		sizeMap[id]++
	}
	var sizes []int
	for _, size := range sizeMap {
		sizes = append(sizes, size)
	}
	return sizes
}

func (pg *Playground) buildCircuits(distances []JunctionBoxPair, n, part int) (map[Coord]int, JunctionBoxPair) {
	circuits := make(map[Coord]int)

	var currentCircuitId int
	var lastJunctionBoxPairAdded JunctionBoxPair

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

		lastJunctionBoxPairAdded = nextJoin

		if part == 1 && i >= n-1 {
			break
		}

		if part == 2 {
			sizes := pg.buildCircuitSizes(circuits)
			if len(sizes) == 1 && sizes[0] == len(pg.boxes) {
				break
			}
		}

	}
	return circuits, lastJunctionBoxPairAdded
}

func (pg *Playground) buildDistances(coords []Coord) []JunctionBoxPair {
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
