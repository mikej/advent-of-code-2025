package main

import "fmt"

type Beam struct {
	x        int
	hasSplit bool
}

type Splitter struct {
	x, y int
}

type Manifold struct {
	beams          []Beam
	splitters      []Splitter
	rows, currentY int
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

}

func (m *Manifold) SplitCount() int {
	return 0
}

func main() {
}
