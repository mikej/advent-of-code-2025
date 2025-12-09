package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/mikej/advent-of-code-2025/shared/input"
)

func main() {
	lines, err := input.ReadFromFile("/Users/mike/Downloads/input-day-9.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	tiles, _ := parseTiles(lines)
	fmt.Println("Largest area is", largestArea(tiles))
}

type Tile struct {
	x, y int
}

type TilePair struct {
	tile1, tile2 Tile
	area         int
}

func parseTiles(lines []string) ([]Tile, error) {
	tiles := make([]Tile, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		tiles = append(tiles, Tile{x, y})
	}
	return tiles, nil
}

func largestArea(tiles []Tile) int {
	pairs := buildTilePairs(tiles)

	return pairs[0].area
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func buildTilePairs(tiles []Tile) []TilePair {
	var pairs []TilePair
	for i, a := range tiles {
		for j, b := range tiles {
			if i <= j {
				continue
			}
			width := absInt(a.x-b.x) + 1
			height := absInt(a.y-b.y) + 1
			pairs = append(pairs, TilePair{a, b, width * height})
		}
	}

	slices.SortFunc(pairs, func(a, b TilePair) int {
		return cmp.Compare(a.area, b.area)
	})
	slices.Reverse(pairs)

	return pairs
}
