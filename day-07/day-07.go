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
	return &Manifold{beams, splitters, len(input)}, nil
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
						m.beams[i].hasSplit = true
						if !m.isExistingBeamInColumn(beam.x - 1) {
							m.beams = append(m.beams, Beam{splitter.x - 1, false})
						}
						if !m.isExistingBeamInColumn(beam.x + 1) {
							m.beams = append(m.beams, Beam{splitter.x + 1, false})
						}
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

func (m *Manifold) SplitCount() int {
	count := 0
	for _, beam := range m.beams {
		if beam.hasSplit {
			count++
		}
	}
	return count
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
