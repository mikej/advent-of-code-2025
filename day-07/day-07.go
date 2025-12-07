package main

import (
	"fmt"

	"github.com/mikej/advent-of-code-2025/shared/input"
)

type Beam struct {
	x          int
	hasSplit   bool
	worldCount int
}

func (b *Beam) Split() {
	b.hasSplit = true
	b.worldCount = 0
}

type Splitter struct {
	x, y int
}

type Manifold struct {
	beams     []Beam
	splitters []Splitter
	rows      int
	splits    int
	worlds    int
}

func NewManifold(input []string) (*Manifold, error) {
	beams := []Beam{}
	splitters := []Splitter{}
	for y, line := range input {
		for x, char := range line {
			if char == 'S' {
				beams = append(beams, Beam{x, false, 1})
			} else if char == '^' {
				splitters = append(splitters, Splitter{x, y})
			} else if char != '.' {
				return nil, fmt.Errorf("invalid character %c at position (%d, %d)", char, x, y)
			}
		}
	}
	return &Manifold{beams, splitters, len(input), 0, 1}, nil
}

func (m *Manifold) Run() {
	for row := 0; row < m.rows; row++ {
		for i, beam := range m.beams {
			if beam.hasSplit {
				continue
			}
			for _, splitter := range m.splitters {
				if row == splitter.y {
					if beam.x == splitter.x {
						m.splits++
						m.worlds += beam.worldCount

						worldsBeforeSplit := beam.worldCount
						m.beams[i].Split()

						left := m.beamForColumn(beam.x - 1)
						left.hasSplit = false
						left.worldCount += worldsBeforeSplit

						right := m.beamForColumn(beam.x + 1)
						right.hasSplit = false
						right.worldCount += worldsBeforeSplit

						break
					}
				}
			}
		}
	}
}

func (m *Manifold) beamForColumn(col int) *Beam {
	for i, beam := range m.beams {
		if beam.x == col {
			return &m.beams[i]
		}
	}

	newBeam := Beam{col, false, 0}
	m.beams = append(m.beams, newBeam)
	return &m.beams[len(m.beams)-1]
}

func (m *Manifold) SplitCount() int {
	return m.splits
}

func (m *Manifold) TimelineCount() int {
	return m.worlds
}

func main() {
	lines, err := input.ReadFromFile("/Users/mike/Downloads/input-day-7.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	manifold, err := NewManifold(lines)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}
	manifold.Run()
	fmt.Println("Total number of splits is", manifold.SplitCount())
	fmt.Println("Total number of timelines is", manifold.TimelineCount())
}
