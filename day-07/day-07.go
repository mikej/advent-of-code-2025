package main

import (
	"fmt"

	"github.com/mikej/advent-of-code-2025/shared/input"
)

type Beam struct {
	x        int
	hasSplit bool
}

type Splitter struct {
	x, y int
}

type Manifold struct {
	beams     []Beam
	splitters []Splitter
	rows      int
	splits    int
}

func NewManifold(input []string) (*Manifold, error) {
	beams := []Beam{}
	splitters := []Splitter{}
	for y, line := range input {
		for x, char := range line {
			if char == 'S' {
				beams = append(beams, Beam{x, false})
			} else if char == '^' {
				splitters = append(splitters, Splitter{x, y})
			} else if char != '.' {
				return nil, fmt.Errorf("invalid character %c at position (%d, %d)", char, x, y)
			}
		}
	}
	return &Manifold{beams, splitters, len(input), 0}, nil
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
						
						m.beams[i].hasSplit = true

						left := m.beamForColumn(beam.x - 1)
						left.hasSplit = false

						right := m.beamForColumn(beam.x + 1)
						right.hasSplit = false

						break
					}
				}
			}
		}
	}
}

func (m *Manifold) isExistingBeamInColumn(col int) bool {
	for _, beam := range m.beams {
		if beam.x == col && !beam.hasSplit {
			return true
		}
	}
	return false
}

func (m *Manifold) beamForColumn(col int) *Beam {
	for i, beam := range m.beams {
		if beam.x == col {
			return &m.beams[i]
		}
	}

	newBeam := Beam{col, false}
	m.beams = append(m.beams, newBeam)
	return &m.beams[len(m.beams)-1]
}

func (m *Manifold) SplitCount() int {
	return m.splits
}

func (m *Manifold) WorldCount() int {
	return 0
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
}
